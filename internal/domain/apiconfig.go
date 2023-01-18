package domain

import (
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

// APIConfig which stores the api configuration
type APIConfig struct {
	ConcurrentConnections int            // number of the worker
	Duration              int            // time for the run request
	TimeOut               int            // timeout for request
	FinalStatus           chan *Status   // this chanel rerun final status to the print state
	Interrupt             int32          // wear going to interrupt  loading the if user send the signal
	Params                *RequestParams // all parameter is going to send to the request
}

// create new api config
func NewAPIConfig(goroutines, duration, timeOut int, finalStatusChan chan *Status, params *RequestParams) *APIConfig {
	a := &APIConfig{
		ConcurrentConnections: goroutines,
		Duration:              duration,
		TimeOut:               timeOut,
		FinalStatus:           finalStatusChan,
		Params:                params,
	}
	return a
}

// send stop signal to the go routin
func (conf *APIConfig) Stop() {
	atomic.StoreInt32(&conf.Interrupt, 1)
}

// som prepration for the send request
func (conf *APIConfig) PrepareRequest() error {
	val := conf.Params
	var buffer io.Reader
	req, err := http.NewRequest(val.Method, val.URL, buffer)
	if err != nil {
		fmt.Println("[Info] An error occurred while creating a new worker Request", err)
		return err
	}
	for headerKey, headerValue := range val.Header {
		req.Header.Add(headerKey, headerValue)
	}

	return nil
}
