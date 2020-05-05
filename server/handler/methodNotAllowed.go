package handler

import (
	"net/http"
)

// MethodNotAllowedHandler handles the route not found error, returning error code with corresponding logging
type MethodNotAllowedHandler struct{}

func (handler *MethodNotAllowedHandler) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	var requestString = jsonutilMarshalIgnoreError(
		httpRequest,
	)
	loggerAppRoot(
		"RouteError",
		"MethodNotAllowed",
		"%v",
		requestString,
	)
	httpError(
		responseWriter,
		"405 - resource URI action not allowed",
		http.StatusMethodNotAllowed,
	)
}
