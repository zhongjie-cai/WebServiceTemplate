package favicon

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

// func pointers for injection / testing: favicon.go
var (
	httpServeFile = http.ServeFile
	configAppPath = config.AppPath
)
