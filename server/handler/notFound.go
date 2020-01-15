package handler

import (
	"net/http"
)

// NotFoundHandler handles the route not found error, returning error code with corresponding logging
type NotFoundHandler struct{}

func (handler *NotFoundHandler) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	loggerAppRoot(
		"RouteError",
		"NotFound",
		"%v",
		httpRequest,
	)
	httpError(
		responseWriter,
		"404 - resource URI not found",
		http.StatusNotFound,
	)
}
