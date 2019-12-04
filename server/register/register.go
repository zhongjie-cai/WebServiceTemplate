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
) string {
	if parameterType == "" {
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
		parameterType,
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
) string {
	var updatedPath = path
	for parameterName, parameterType := range parameters {
		updatedPath = doParameterReplacementFunc(
			updatedPath,
			parameterName,
			parameterType,
		)
	}
	return updatedPath
}

func evaluateQueries(
	queries map[string]model.ParameterType,
) []string {
	var evaluatedQueries = []string{}
	for key, value := range queries {
		var queryParameter string
		if value == "" {
			queryParameter = fmtSprintf(
				"{%v}",
				key,
			)
		} else {
			queryParameter = fmtSprintf(
				"{%v:%v}",
				key,
				value,
			)
		}
		evaluatedQueries = append(
			evaluatedQueries,
			key,
			queryParameter,
		)
	}
	return evaluatedQueries
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
		)
		var queries = evaluateQueriesFunc(
			configuredRoute.Queries,
		)
		routeHandleFunc(
			router,
			configuredRoute.Endpoint,
			configuredRoute.Method,
			evaluatedPath,
			queries,
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

func registerMiddlewares(
	router *mux.Router,
) {
	if customization.Middlewares == nil {
		loggerAppRoot(
			"register",
			"registerMiddlewares",
			"customization.Middlewares function not set: no middleware registered!",
		)
		return
	}
	var middlewares = customization.Middlewares()
	if middlewares == nil ||
		len(middlewares) == 0 {
		loggerAppRoot(
			"register",
			"registerMiddlewares",
			"customization.Middlewares function empty: no middleware returned!",
		)
		return
	}
	for _, middleware := range middlewares {
		routeAddMiddleware(
			router,
			middleware,
		)
	}
}

func instrumentRouter(router *mux.Router) *mux.Router {
	if customization.InstrumentRouter == nil {
		return router
	}
	return customization.InstrumentRouter(router)
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
	registerMiddlewaresFunc(
		router,
	)
	var routerError = routeWalkRegisteredRoutes(
		router,
	)
	if routerError != nil {
		return nil,
			apperrorWrapSimpleError(
				[]error{routerError},
				"Failed to instantiate routes",
			)
	}
	return instrumentRouterFunc(router), nil
}
