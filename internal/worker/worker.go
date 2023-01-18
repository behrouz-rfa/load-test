package worker

import (
	"github.com/valyala/fasthttp"
	"load-test/internal/domain"
	"load-test/internal/ports"
	utils2 "load-test/internal/utils"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Worker struct {
	conf    *domain.APIConfig
	client  *fasthttp.Client
	counter *domain.Counter
}

// New Worker
func NewWorker(conf *domain.APIConfig, client *fasthttp.Client, counter *domain.Counter) *Worker {
	return &Worker{conf: conf, client: client, counter: counter}
}

var _ ports.WorkerRepo = (*Worker)(nil)

func (r Worker) AsyncHTTP() {

	//create api status
	status := &domain.Status{
		TotalDuration: time.Millisecond,
		MinTime:       time.Hour,
		MaxTime:       time.Millisecond,
	}

	// for calculation of send request
	start := time.Now()

	//prettier the request if we have something to send with request
	err := r.conf.PrepareRequest()
	if err != nil {
		log.Fatal(err)
	}

	//atomic interrupt is when we send signal from clt + c
	// after som duration It's going to interrupt the
	for time.Since(start).Seconds() <= float64(r.conf.Duration) && atomic.LoadInt32(&r.conf.Interrupt) == 0 {
		val := r.conf.Params
		reqDuration, respSize := r.Run(val.URL, val.Method)
		if respSize > 0 {
			//increment the success Request-counter
			status.RequestsCounter++
			// add response size
			status.TotalResponseSize += int64(respSize)
			// increment the request duration
			status.TotalDuration += reqDuration
			//calculate the min and max time
			status.MaxTime = utils2.FindMaxRequestTime(reqDuration, status.MaxTime)
			status.MinTime = utils2.FindMinRequestTime(reqDuration, status.MinTime)
		} else {
			status.ErrorCount++
		}

	}
	// pass status to the chanel to show on final result
	r.conf.FinalStatus <- status
}

// Run sends an HTTP request and returns an HTTP response, following
func (r Worker) Run(url string, method string) (requestDuration time.Duration, responseSize int) {

	// for calculate the request duration
	requestDuration = -1

	// calculate the response size
	responseSize = -1

	//start for  calculate the time for each request
	start := time.Now()

	//create request and response for fasthttp client
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	resp := fasthttp.AcquireResponse()

	// call the request
	if err := r.client.Do(req, resp); err != nil {
		r.counter.Add("-1xx", 1)
		return
	}

	//get body after we
	body := resp.Body()

	// check request code in general
	if resp.StatusCode() == http.StatusOK || resp.StatusCode() == http.StatusCreated {

		//calculate the duration base on start and time since
		requestDuration = time.Since(start)

		//for calculate the size of  header
		size := 0

		//calculate the size of header
		resp.Header.VisitAll(func(key, value []byte) {
			size += len(key) + len(value) + 2 // 2 for the \r\n that separates the headers.
		})
		//response size is sum up the header and body size
		responseSize = len(body) + size
		//increment the size if 2xx status code by 1
		r.counter.Add("2xx", 1)
	} else if resp.StatusCode() == http.StatusContinue || resp.StatusCode() == http.StatusSwitchingProtocols ||
		resp.StatusCode() == http.StatusProcessing {
		// all error from 1xx set to map
		r.counter.Add("1xx", 1)
	} else if resp.StatusCode() == http.StatusMultipleChoices || resp.StatusCode() == http.StatusMovedPermanently ||
		resp.StatusCode() == http.StatusFound || resp.StatusCode() == http.StatusSeeOther ||
		resp.StatusCode() == http.StatusNotModified {
		// all error from 3xx set to map
		r.counter.Add("3xx", 1)
	} else if resp.StatusCode() == http.StatusBadRequest || resp.StatusCode() == http.StatusUnauthorized ||
		resp.StatusCode() == http.StatusPaymentRequired || resp.StatusCode() == http.StatusForbidden ||
		resp.StatusCode() == http.StatusNotFound || resp.StatusCode() == http.StatusMethodNotAllowed ||
		resp.StatusCode() == http.StatusNotAcceptable || resp.StatusCode() == http.StatusProxyAuthRequired ||
		resp.StatusCode() == http.StatusRequestTimeout || resp.StatusCode() == http.StatusContinue {
		// all error from 4xx set to map
		r.counter.Add("4xx", 1)
	} else if resp.StatusCode() == http.StatusInternalServerError || resp.StatusCode() == http.StatusNotImplemented ||
		resp.StatusCode() == http.StatusBadGateway || resp.StatusCode() == http.StatusServiceUnavailable ||
		resp.StatusCode() == http.StatusGatewayTimeout || resp.StatusCode() == http.StatusHTTPVersionNotSupported {
		// all error from 5xx set to map
		r.counter.Add("5xx", 1)
	} else {
		// unknown Status code
		//maybe we dont get response
		r.counter.Add("0xx", 1)
	}
	return
}
