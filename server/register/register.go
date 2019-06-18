package register

import (
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

// Routes is the function to be exposed to the consumers to customize routes registration
var Routes func() []model.Route

// Statics is the function to be exposed to the consumers to customize static contents registration
var Statics func() []model.Static

func doParameterReplacement(
	originalPath string,
	parameterName string,
	parameterType model.ParameterType,
	parameterReplacementsMap map[model.ParameterType]string,
) string {
	var parameterReplacement, found = parameterReplacementsMap[parameterType]
	if !found {
		loggerAppRoot(
			"register",
			"doParameterReplacement",
			"Path parameter [%v] in path [%v] has no type specification; fallback to default.",
			parameterName,
			originalPath,
		)
		return originalPath
	}
	var oldParameter = fmtSprintf(
		"{%v}",
		parameterName,
	)
	var newParameter = fmtSprintf(
		"{%v:%v}",
		parameterName,
		parameterReplacement,
	)
	return stringsReplace(
		originalPath,
		oldParameter,
		newParameter,
		-1,
	)
}

func evaluatePathWithParameters(
	path string,
	parameters map[string]model.Parameter,
	parameterReplacementsMap map[model.ParameterType]string,
) string {
	var updatedPath = path
	for _, parameter := range parameters {
		updatedPath = doParameterReplacementFunc(
			updatedPath,
			parameter.Name,
			parameter.Type,
			parameterReplacementsMap,
		)
	}
	return updatedPath
}

func registerRoutes(
	router *mux.Router,
) {
	if Routes == nil {
		loggerAppRoot(
			"register",
			"registerRoutes",
			"Routes function not set: no routes registered!",
		)
		return
	}
	var configuredRoutes = Routes()
	if configuredRoutes == nil ||
		len(configuredRoutes) == 0 {
		loggerAppRoot(
			"register",
			"registerRoutes",
			"Routes function empty: no routes returned!",
		)
		return
	}
	for _, configuredRoute := range configuredRoutes {
		var evaluatedPath = evaluatePathWithParametersFunc(
			configuredRoute.Path,
			configuredRoute.Parameters,
			model.ParameterTypeMap,
		)
		routeHandleFunc(
			router,
			configuredRoute.Endpoint,
			configuredRoute.Method,
			evaluatedPath,
			handlerSession,
			configuredRoute.ActionFunc,
		)
	}
}

func registerStatics(
	router *mux.Router,
) {
	if Statics == nil {
		loggerAppRoot(
			"register",
			"registerStatics",
			"Statics function not set: no static content registered!",
		)
		return
	}
	var statics = Statics()
	if statics == nil ||
		len(statics) == 0 {
		loggerAppRoot(
			"register",
			"registerStatics",
			"Statics function empty: no static content returned!",
		)
		return
	}
	for _, static := range statics {
		routeHostStatic(
			router,
			static.Name,
			static.PathPrefix,
			static.Handler,
		)
	}
}

// Instantiate instantiates and registers the given routes according to custom specification
func Instantiate() (*mux.Router, error) {
	var router = routeCreateRouter()
	registerRoutesFunc(
		router,
	)
	registerStaticsFunc(
		router,
	)
	var routerError = routeWalkRegisteredRoutes(
		router,
	)
	if routerError != nil {
		return nil,
			apperrorWrapSimpleError(
				routerError,
				"Failed to instantiate routes",
			)
	}
	return router, nil
}
