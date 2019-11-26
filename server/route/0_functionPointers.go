package route

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// func pointers for injection / testing: server.go
var (
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	apperrorGetNotImplementedError  = apperror.GetNotImplementedError
	apperrorGetCustomError          = apperror.GetCustomError
	stringsJoin                     = strings.Join
	fmtSprintf                      = fmt.Sprintf
	loggerAppRoot                   = logger.AppRoot
	muxNewRouter                    = mux.NewRouter
	muxCurrentRoute                 = mux.CurrentRoute
	responseWrite                   = response.Write
	getNameFunc                     = getName
	getPathTemplateFunc             = getPathTemplate
	getPathRegexpFunc               = getPathRegexp
	getQueriesTemplatesFunc         = getQueriesTemplates
	getQueriesRegexpFunc            = getQueriesRegexp
	getMethodsFunc                  = getMethods
	getActionByNameFunc             = getActionByName
	printRegisteredRouteDetailsFunc = printRegisteredRouteDetails
)
