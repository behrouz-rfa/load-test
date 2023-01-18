package integration_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"load-test/internal/application"
	"load-test/internal/domain"
	"load-test/internal/logger"
	utils2 "load-test/internal/utils"
	"load-test/internal/worker"
	"os"
	"testing"
)

const (
	NUM_PARALLEL             = 50  // Number of concurrent connection
	requestDurationInSeconds = 5   //second
	requestTimeOut           = 200 //milisec
	requestURL               = "https://jsonplaceholder.typicode.com/todos/1"
	requestURL2              = "https://wrongdomain.local"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestFastHttp(t *testing.T) {
	client := &fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://jsonplaceholder.typicode.com/todo/1")
	req.Header.SetMethod("GET")
	resp := fasthttp.AcquireResponse()

	if err := client.Do(req, resp); err != nil {
		fmt.Println(err)
		return
	}
	status := resp.StatusCode()

	fmt.Println(status)
}

func TestSuccessStatusRequest(t *testing.T) {
	requestHeader := utils2.RequestHeader("")

	requestParams := domain.CreateRequestParams(requestURL, "GET", requestHeader)

	statusChan := make(chan *domain.Status, NUM_PARALLEL)

	config := domain.NewAPIConfig(
		NUM_PARALLEL,
		requestDurationInSeconds,
		requestTimeOut,
		statusChan,
		requestParams,
	)

	//create new http client forom utils
	client := utils2.NewFastClient(config.TimeOut)
	//counter i
	counter := &domain.Counter{}

	request := worker.NewRequestRepository(config, client, counter)

	app := application.New(request)

	for i := 0; i < NUM_PARALLEL; i++ {
		go func() {
			app.CreateRequest()
		}()
	}

	responseCounter := 0
	status := domain.NewStatus()

	for responseCounter < NUM_PARALLEL {
		select {

		case s := <-statusChan:
			status.RequestsCounter += s.RequestsCounter
			status.ErrorCount += s.ErrorCount
			status.TotalDuration += s.TotalDuration
			status.MaxTime = s.MaxTime
			status.MinTime = s.MinTime
			status.TotalResponseSize = s.TotalResponseSize
			responseCounter++
		}
	}
	// check how many request we send if ==  0 show the message

	assert.Greater(t, status.RequestsCounter, 0)

	//finish the result when program exit
	defer func() {
		logger.PrintResult(status, counter)
	}()
}
func TestZeroStatusReturn(t *testing.T) {
	requestHeader := utils2.RequestHeader("")

	requestParams := domain.CreateRequestParams(requestURL2, "GET", requestHeader)

	statusChan := make(chan *domain.Status, NUM_PARALLEL)

	config := domain.NewAPIConfig(
		NUM_PARALLEL,
		requestDurationInSeconds,
		requestTimeOut,
		statusChan,
		requestParams,
	)

	//create new http client forom utils
	client := utils2.NewFastClient(config.TimeOut)
	//counter i
	counter := &domain.Counter{}

	request := worker.NewRequestRepository(config, client, counter)

	app := application.New(request)

	for i := 0; i < NUM_PARALLEL; i++ {
		go func() {
			app.CreateRequest()
		}()
	}

	responseCounter := 0
	status := domain.NewStatus()

	for responseCounter < NUM_PARALLEL {
		select {

		case s := <-statusChan:
			status.RequestsCounter += s.RequestsCounter
			status.ErrorCount += s.ErrorCount
			status.TotalDuration += s.TotalDuration
			status.MaxTime = s.MaxTime
			status.MinTime = s.MinTime
			status.TotalResponseSize = s.TotalResponseSize
			responseCounter++
		}
	}
	// check how many request we send if ==  0 show the message

	assert.Equal(t, status.RequestsCounter, 0)

	//finish the result when program exit
	defer func() {
		logger.PrintResult(status, counter)
	}()
}
