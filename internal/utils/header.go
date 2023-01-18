package utils

import (
	"load-test/internal/domain"
	"net/http"
	"strings"
)

func HeaderSize(headers http.Header) (result int64) {
	result = 0
	for k, v := range headers {
		result += int64(len(k) + len(": \r\n"))
		for _, s := range v {
			result += int64(len(s))
		}
	}
	result += int64(len("\r\n"))
	return result
}

func RequestHeader(headerValues string) map[string]string {
	requestHeader := make(map[string]string)

	if headerValues != "" {
		hv := strings.Split(headerValues, ";")
		for _, hd := range hv {
			header := strings.SplitN(hd, ":", 2)

			requestHeader[header[0]] = header[1]
		}
	}
	return requestHeader
}

func GetRequestParams(requestURL, requestMethod string, requestHeader map[string]string, requestParams []*domain.RequestParams) []*domain.RequestParams {
	reqParam := new(domain.RequestParams)
	reqParam.URL = requestURL
	reqParam.Method = requestMethod
	reqParam.Header = requestHeader
	requestParams = append(requestParams, reqParam)
	return requestParams
}
