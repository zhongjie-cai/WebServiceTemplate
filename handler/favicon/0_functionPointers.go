package favicon

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

// func pointers for injection / testing: favicon.go
var (
	routeHandleFunc       = route.HandleFunc
	httpServeFile         = http.ServeFile
	configAppPath         = config.AppPath
	commonHandleInSession = common.HandleInSession
	handleGetFaviconFunc  = handleGetFavicon
)
