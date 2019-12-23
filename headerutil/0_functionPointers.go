package headerutil

import (
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
)

// func pointers for injection / testing: headerutil.go
var (
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	stringsJoin                = strings.Join
	loggerAPIRequest           = logger.APIRequest
	getHeaderLogStyleFunc      = getHeaderLogStyle
	logCombinedHTTPHeaderFunc  = logCombinedHTTPHeader
	logPerNameHTTPHeaderFunc   = logPerNameHTTPHeader
	logPerValueHTTPHeaderFunc  = logPerValueHTTPHeader
	logHTTPHeaderFunc          = LogHTTPHeader
)
