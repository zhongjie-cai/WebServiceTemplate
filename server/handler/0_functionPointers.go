package handler

import (
	"net/http"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/request"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

// func pointers for injection / testing: handler.go
var (
	routeGetRouteInfo             = route.GetRouteInfo
	sessionRegister               = session.Register
	panicHandle                   = panic.Handle
	responseWrite                 = response.Write
	loggerAPIEnter                = logger.APIEnter
	loggerAPIExit                 = logger.APIExit
	apperrorGetInvalidOperation   = apperror.GetInvalidOperation
	timeutilGetTimeNowUTC         = timeutil.GetTimeNowUTC
	timeSince                     = time.Since
	executeCustomizedFunctionFunc = executeCustomizedFunction
)

// func pointers for injection / testing: methodNotAllowed.go
var (
	requestFullDump = request.FullDump
	loggerAppRoot   = logger.AppRoot
	httpError       = http.Error
)
