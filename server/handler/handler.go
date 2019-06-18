package handler

import (
	"net/http"
)

// Session wraps the HTTP handler with session related operations
func Session(
	responseWriter http.ResponseWriter,
	httpRequst *http.Request,
) {
	var endpoint, action, routeError = routeGetRouteInfo(
		httpRequst,
	)
	var sessionID = sessionRegister(
		endpoint,
		requestGetLoginID(
			httpRequst,
		),
		requestGetCorrelationID(
			httpRequst,
		),
		requestGetAllowedLogType(
			httpRequst,
		),
		httpRequst,
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
			httpRequst.Method,
		)
		responseError(
			sessionID,
			routeError,
			responseWriter,
		)
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequst.Method,
		)
	} else {
		loggerAPIEnter(
			sessionID,
			"handler",
			endpoint,
			httpRequst.Method,
		)
		action(
			responseWriter,
			httpRequst,
			sessionID,
		)
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequst.Method,
		)
	}
}
