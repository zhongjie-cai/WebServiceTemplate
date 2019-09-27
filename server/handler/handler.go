package handler

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func verifyAuthorization(
	httpRequest *http.Request,
) error {
	if customization.AuthorizationFunc == nil {
		return nil
	}
	return customization.AuthorizationFunc(
		httpRequest,
	)
}

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
		responseWrite(
			sessionID,
			nil,
			apperrorGetInvalidOperation(
				routeError,
			),
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
		var authorizationError = verifyAuthorizationFunc(
			httpRequest,
		)
		if authorizationError == nil {
			var responseObject, responseError = action(
				sessionID,
			)
			responseWrite(
				sessionID,
				responseObject,
				responseError,
			)
		} else {
			responseWrite(
				sessionID,
				nil,
				authorizationError,
			)
		}
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
	}
}
