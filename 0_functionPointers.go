package main

import (
	"fmt"

	"github.com/zhongjie-cai/WebServiceTemplate/application"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
)

// func pointers for injection / testing: config.go
var (
	fmtPrintf        = fmt.Printf
	swaggerHandler   = swagger.Handler
	applicationStart = application.Start
)
