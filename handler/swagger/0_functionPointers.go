package swagger

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

// func pointers for injection / testing: swagger.go
var (
	configAppPath       = config.AppPath
	httpRedirect        = http.Redirect
	httpStripPrefix     = http.StripPrefix
	httpFileServer      = http.FileServer
	routeHandleFunc     = route.HandleFunc
	routeHostStatic     = route.HostStatic
	redirectHandlerFunc = redirectHandler
	contentHandlerFunc  = contentHandler
)
