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
	var session = sessionRegister(
		endpoint,
		requestGetAllowedLogType(
			httpRequest,
		),
		requestGetAllowedLogLevel(
			httpRequest,
		),
		httpRequest,
		responseWriter,
	)
	var sessionID = session.GetID()
	defer func() {
		panicHandle(
			session,
			recover(),
		)
		sessionUnregister(
			session,
		)
	}()
	if routeError != nil {
		loggerAPIEnter(
			session,
			httpRequest.Method,
			endpoint,
			"",
		)
		responseWrite(
			session,
			nil,
			apperrorGetInvalidOperation(
				routeError,
			),
		)
		loggerAPIExit(
			session,
			httpRequest.Method,
			endpoint,
			"",
		)
	} else {
		loggerAPIEnter(
			session,
			httpRequest.Method,
			endpoint,
			"",
		)
		var preActionError = executeCustomizedFunctionFunc(
			sessionID,
			customization.PreActionFunc,
		)
		if preActionError != nil {
			responseWrite(
				session,
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
						session,
						httpRequest.Method,
						endpoint,
						"Post-action error: %v",
						postActionError,
					)
					responseWrite(
						session,
						nil,
						responseError,
					)
				} else {
					responseWrite(
						session,
						nil,
						postActionError,
					)
				}
			} else {
				responseWrite(
					session,
					responseObject,
					responseError,
				)
			}
		}
		loggerAPIExit(
			session,
			httpRequest.Method,
			endpoint,
			"",
		)
	}
}
