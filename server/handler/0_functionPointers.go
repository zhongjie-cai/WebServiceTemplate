package handler

import (
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// func pointers for injection / testing: common.go
var (
	muxVars                     = mux.Vars
	routeGetRouteInfo           = route.GetRouteInfo
	sessionRegister             = session.Register
	sessionUnregister           = session.Unregister
	panicHandle                 = panic.Handle
	requestGetLoginID           = request.GetLoginID
	requestGetCorrelationID     = request.GetCorrelationID
	requestGetAllowedLogType    = request.GetAllowedLogType
	requestGetRequestBody       = request.GetRequestBody
	responseWrite               = response.Write
	loggerAPIEnter              = logger.APIEnter
	loggerAPIExit               = logger.APIExit
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
)
