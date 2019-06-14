package health

import (
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

// func pointers for injection / testing: health.go
var (
	routeHandleFunc           = route.HandleFunc
	configAppVersion          = config.AppVersion
	responseOk                = response.Ok
	commonHandleInSession     = common.HandleInSession
	handleGetHealthFunc       = handleGetHealth
	handleGetHealthReportFunc = handleGetHealthReport
)
