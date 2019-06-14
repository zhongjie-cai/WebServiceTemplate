package common

import (
	"github.com/zhongjie-cai/WebServiceTemplate/handler/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// func pointers for injection / testing: common.go
var (
	routeGetEndpointName     = route.GetEndpointName
	sessionRegister          = session.Register
	sessionUnregister        = session.Unregister
	panicHandle              = panic.Handle
	requestGetLoginID        = request.GetLoginID
	requestGetCorrelationID  = request.GetCorrelationID
	requestGetAllowedLogType = request.GetAllowedLogType
	responseError            = response.Error
	loggerAPIEnter           = logger.APIEnter
	loggerAPIExit            = logger.APIExit
)
