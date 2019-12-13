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
	defer func() {
		panicHandle(
			session,
			recover(),
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
