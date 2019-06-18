package swagger

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

// func pointers for injection / testing: swagger.go
var (
	configAppPath   = config.AppPath
	httpRedirect    = http.Redirect
	httpStripPrefix = http.StripPrefix
	httpFileServer  = http.FileServer
)
