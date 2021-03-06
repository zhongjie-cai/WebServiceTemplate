package model

import "net/http"

// NetworkRequest is an interface for easy operating on network requests and responses
type NetworkRequest interface {
	// EnableRetry sets up automatic retry upon error of specific HTTP status codes; each entry maps an HTTP status code to how many times retry should happen if code matches
	EnableRetry(connectivityRetryCount int, httpStatusRetryCount map[int]int)
	// Process sends the network request over the wire, retrieves and serialize the response to dataTemplate, and provides status code, header and error if applicable
	Process(dataTemplate interface{}) (statusCode int, responseHeader http.Header, responseError error)
	// ProcessRaw sends the network request over the wire, retrieves the response, and returns that response and error if applicable
	ProcessRaw() (responseObject *http.Response, responseError error)
}
