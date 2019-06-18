package register

import (
	"fmt"
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/handler"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

// func pointers for injection / testing: panic.go
var (
	stringsReplace                 = strings.Replace
	fmtSprintf                     = fmt.Sprintf
	loggerAppRoot                  = logger.AppRoot
	routeHandleFunc                = route.HandleFunc
	routeHostStatic                = route.HostStatic
	routeCreateRouter              = route.CreateRouter
	routeWalkRegisteredRoutes      = route.WalkRegisteredRoutes
	apperrorWrapSimpleError        = apperror.WrapSimpleError
	handlerSession                 = handler.Session
	doParameterReplacementFunc     = doParameterReplacement
	evaluatePathWithParametersFunc = evaluatePathWithParameters
	registerRoutesFunc             = registerRoutes
	registerStaticsFunc            = registerStatics
)
