package utils

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func NewHTTPClient(timeOut int) (*http.Client, error) {

	httpTransport := &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		DisableCompression:    true,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(timeOut),
	}
	client := &http.Client{Transport: httpTransport}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("url redirection not allowed")
	}
	return client, nil
}

func NewRestyClient(timeOut int) *resty.Client {
	// Create a Resty Client
	client := resty.New()

	// Set client timeout as per your need
	client.SetTimeout(time.Millisecond * time.Duration(timeOut))
	return client
}

func NewFastClient(timeOut int) *fasthttp.Client {
	// Create a Resty Client
	client := &fasthttp.Client{}
	client.ReadTimeout = time.Millisecond * time.Duration(timeOut)
	client.MaxIdleConnDuration = 5 * time.Second
	// Set client timeout as per your need
	return client
}

