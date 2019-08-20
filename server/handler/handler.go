package handler

import (
	"net/http"
)

// Session wraps the HTTP handler with session related operations
func Session(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
) {
	var endpoint, action, routeError = routeGetRouteInfo(
		httpRequest,
	)
	var sessionID = sessionRegister(
		endpoint,
		requestGetLoginID(
			httpRequest,
		),
		requestGetCorrelationID(
			httpRequest,
		),
		requestGetAllowedLogType(
			httpRequest,
		),
		httpRequest,
		responseWriter,
	)
	defer func() {
		panicHandle(
			endpoint,
			sessionID,
			recover(),
			responseWriter,
		)
		sessionUnregister(
			sessionID,
		)
	}()
	if routeError != nil {
		loggerAPIEnter(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
		responseError(
			sessionID,
			routeError,
		)
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
	} else {
		loggerAPIEnter(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
		var requestBody = requestGetRequestBody(
			sessionID,
			httpRequest,
		)
		action(
			sessionID,
			requestBody,
		)
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
	}
}
