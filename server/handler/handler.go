package handler

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func executeCustomizedFunction(
	session sessionModel.Session,
	customFunc func(sessionModel.Session) error,
) error {
	if customFunc == nil {
		return nil
	}
	return customFunc(
		session,
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
		httpRequest,
		responseWriter,
	)
	var startTime = timeutilGetTimeNowUTC()
	loggerAPIEnter(
		session,
		endpoint,
		httpRequest.Method,
		"",
	)
	defer func() {
		panicHandle(
			session,
			recover(),
		)
		loggerAPIExit(
			session,
			endpoint,
			httpRequest.Method,
			"%s",
			timeSince(startTime),
		)
	}()
	if routeError != nil {
		responseWrite(
			session,
			nil,
			apperrorGetInvalidOperation(
				routeError,
			),
		)
	} else {
		var preActionError = executeCustomizedFunctionFunc(
			session,
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
				session,
			)
			var postActionError = executeCustomizedFunctionFunc(
				session,
				customization.PostActionFunc,
			)
			if postActionError != nil {
				if responseError != nil {
					loggerAPIExit(
						session,
						endpoint,
						httpRequest.Method,
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
	}
}
