package handler

import (
	"net/http"
)

// NotFoundHandler handles the route not found error, returning error code with corresponding logging
type NotFoundHandler struct{}

func (handler *NotFoundHandler) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	var requestString = jsonutilMarshalIgnoreError(
		httpRequest,
	)
	loggerAppRoot(
		"RouteError",
		"NotFound",
		"%v",
		requestString,
	)
	httpError(
		responseWriter,
		"404 - resource URI not found",
		http.StatusNotFound,
	)
}
