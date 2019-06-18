package health

import (
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// func pointers for injection / testing: health.go
var (
	configAppVersion = config.AppVersion
	responseOk       = response.Ok
)
