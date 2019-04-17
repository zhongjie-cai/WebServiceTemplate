package health

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// func pointers for injection / testing: health.go
var (
	httpHandleFunc              = http.HandleFunc
	configAppVersion            = config.AppVersion
	responseOk                  = response.Ok
	responseError               = response.Error
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
	commonHandleInSession       = common.HandleInSession
	handleHealthLogicFunc       = handleHealthLogic
	handlerFunc                 = handler
)
