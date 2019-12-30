package headerutil

import (
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
)

// func pointers for injection / testing: headerutil.go
var (
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	stringsJoin                = strings.Join
	getHeaderLogStyleFunc      = getHeaderLogStyle
	logCombinedHTTPHeaderFunc  = logCombinedHTTPHeader
	logPerNameHTTPHeaderFunc   = logPerNameHTTPHeader
	logPerValueHTTPHeaderFunc  = logPerValueHTTPHeader
	logHTTPHeaderFunc          = LogHTTPHeader
)
