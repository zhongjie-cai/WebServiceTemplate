package handler

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// func pointers for injection / testing: common.go
var (
	routeGetRouteInfo             = route.GetRouteInfo
	sessionRegister               = session.Register
	panicHandle                   = panic.Handle
	responseWrite                 = response.Write
	loggerAPIEnter                = logger.APIEnter
	loggerAPIExit                 = logger.APIExit
	apperrorGetInvalidOperation   = apperror.GetInvalidOperation
	executeCustomizedFunctionFunc = executeCustomizedFunction
)
