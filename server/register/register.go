package register

import (
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

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
	parameters map[string]model.ParameterType,
	parameterReplacementsMap map[model.ParameterType]string,
) string {
	var updatedPath = path
	for parameterName, parameterType := range parameters {
		updatedPath = doParameterReplacementFunc(
			updatedPath,
			parameterName,
			parameterType,
			parameterReplacementsMap,
		)
	}
	return updatedPath
}

func registerRoutes(
	router *mux.Router,
) {
	if customization.Routes == nil {
		loggerAppRoot(
			"register",
			"registerRoutes",
			"customization.Routes function not set: no routes registered!",
		)
		return
	}
	var configuredRoutes = customization.Routes()
	if configuredRoutes == nil ||
		len(configuredRoutes) == 0 {
		loggerAppRoot(
			"register",
			"registerRoutes",
			"customization.Routes function empty: no routes returned!",
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
	if customization.Statics == nil {
		loggerAppRoot(
			"register",
			"registerStatics",
			"customization.Statics function not set: no static content registered!",
		)
		return
	}
	var statics = customization.Statics()
	if statics == nil ||
		len(statics) == 0 {
		loggerAppRoot(
			"register",
			"registerStatics",
			"customization.Statics function empty: no static content returned!",
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