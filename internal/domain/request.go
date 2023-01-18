package domain

// RequestParams for worker Request
type RequestParams struct {
	URL    string            // Url from the user
	Method string            // Method for sending request
	Header map[string]string // Header for the specific request
}

// Create Request Base Params
func CreateRequestParams(requestURL, requestMethod string, requestHeader map[string]string) *RequestParams {
	reqParam := new(RequestParams)
	reqParam.URL = requestURL
	reqParam.Method = requestMethod
	reqParam.Header = requestHeader
	return reqParam

}
