package route

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
)

// func pointers for injection / testing: server.go
var (
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	apperrorConsolidateAllErrors    = apperror.ConsolidateAllErrors
	stringsJoin                     = strings.Join
	fmtSprintf                      = fmt.Sprintf
	loggerAppRoot                   = logger.AppRoot
	muxNewRouter                    = mux.NewRouter
	muxCurrentRoute                 = mux.CurrentRoute
	getNameFunc                     = getName
	getPathTemplateFunc             = getPathTemplate
	getPathRegexpFunc               = getPathRegexp
	getQueriesTemplatesFunc         = getQueriesTemplates
	getQueriesRegexpFunc            = getQueriesRegexp
	getMethodsFunc                  = getMethods
	printRegisteredRouteDetailsFunc = printRegisteredRouteDetails
	walkRegisteredRoutesFunc        = walkRegisteredRoutes
)
