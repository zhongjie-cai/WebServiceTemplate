package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func executeCustomizedFunction(
	sessionID uuid.UUID,
	customFunc func(uuid.UUID) error,
) error {
	if customFunc == nil {
		return nil
	}
	return customFunc(
		sessionID,
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
		var preActionError = executeCustomizedFunctionFunc(
			sessionID,
			customization.PreActionFunc,
		)
		if preActionError != nil {
			responseWrite(
				sessionID,
				nil,
				preActionError,
			)
		} else {
			var responseObject, responseError = action(
				sessionID,
			)
			var postActionError = executeCustomizedFunctionFunc(
				sessionID,
				customization.PostActionFunc,
			)
			if postActionError != nil {
				if responseError != nil {
					loggerAPIExit(
						sessionID,
						"handler",
						endpoint,
						"Post-action error: %v",
						postActionError,
					)
					responseWrite(
						sessionID,
						nil,
						responseError,
					)
				} else {
					responseWrite(
						sessionID,
						nil,
						postActionError,
					)
				}
			} else {
				responseWrite(
					sessionID,
					responseObject,
					responseError,
				)
			}
		}
		loggerAPIExit(
			sessionID,
			"handler",
			endpoint,
			httpRequest.Method,
		)
	}
}
