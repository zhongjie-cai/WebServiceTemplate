package favicon

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// func pointers for injection / testing: favicon.go
var (
	httpHandleFunc              = http.HandleFunc
	httpServeFile               = http.ServeFile
	responseError               = response.Error
	configAppPath               = config.AppPath
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
	commonHandleInSession       = common.HandleInSession
	handleFaviconLogicFunc      = handleFaviconLogic
	handlerFunc                 = handler
)
