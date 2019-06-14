package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getName(route *mux.Route) string {
	return route.GetName()
}

func getPathTemplate(route *mux.Route) (string, error) {
	return route.GetPathTemplate()
}

func getPathRegexp(route *mux.Route) (string, error) {
	return route.GetPathRegexp()
}

func getQueriesTemplates(route *mux.Route) string {
	var queriesTemplates, _ = route.GetQueriesTemplates()
	return stringsJoin(queriesTemplates, ",")
}

func getQueriesRegexp(route *mux.Route) string {
	var queriesRegexps, _ = route.GetQueriesRegexp()
	return stringsJoin(queriesRegexps, ",")
}

func getMethods(route *mux.Route) string {
	var methods, _ = route.GetMethods()
	return stringsJoin(methods, ",")
}

func printRegisteredRouteDetails(
	route *mux.Route,
	router *mux.Router,
	ancestors []*mux.Route,
) error {
	var (
		name                            = getNameFunc(route)
		pathTemplate, pathTemplateError = getPathTemplateFunc(route)
		pathRegexp, pathRegexpError     = getPathRegexpFunc(route)
		queriesTemplates                = getQueriesTemplatesFunc(route)
		queriesRegexps                  = getQueriesRegexpFunc(route)
		methods                         = getMethodsFunc(route)
	)
	var consolidatedError = apperrorConsolidateAllErrors(
		fmtSprintf(
			"Failed to register service route for name [%v]",
			name,
		),
		pathTemplateError,
		pathRegexpError,
	)
	if consolidatedError != nil {
		return consolidatedError
	}
	loggerAppRoot(
		uuid.Nil,
		"route",
		"printRegisteredRouteDetails",
		"Route registered for name [%v]\nPath template:%v\nPath regexp:%v\nQueries templates:%v\nQueries regexps:%v\nMethods:%v",
		name,
		pathTemplate,
		pathRegexp,
		queriesTemplates,
		queriesRegexps,
		methods,
	)
	return nil
}

func walkRegisteredRoutes(router *mux.Router) error {
	var walkError = router.Walk(
		printRegisteredRouteDetailsFunc,
	)
	if walkError != nil {
		return apperrorWrapSimpleError(
			walkError,
			"Failed to walk through registered routes",
		)
	}
	return nil
}

// RegisterEntries registeres the listed entry functions and returns the mux router for server hosting
func RegisterEntries(entryFuncs ...func(*mux.Router)) (*mux.Router, error) {
	if entryFuncs == nil ||
		len(entryFuncs) == 0 {
		return nil,
			apperrorWrapSimpleError(
				nil,
				"No host entries found",
			)
	}
	var router = muxNewRouter()
	for _, entryFunc := range entryFuncs {
		entryFunc(router)
	}
	var routerError = walkRegisteredRoutesFunc(
		router,
	)
	if routerError != nil {
		return nil,
			apperrorWrapSimpleError(
				routerError,
				"Failed to register routes",
			)
	}
	return router, nil
}

const handleMethod = "GET"

// HostStatic registers a given path URI with the given handler interface (for static content hosting)
func HostStatic(
	router *mux.Router,
	name string,
	path string,
	handler http.Handler,
) *mux.Route {
	return router.PathPrefix(
		path,
	).Handler(
		handler,
	).Name(
		name,
	)
}

// HandleFunc registers a given method/path URI with the given handler function
func HandleFunc(
	router *mux.Router,
	endpoint string,
	method string,
	path string,
	handlerFunc func(http.ResponseWriter, *http.Request),
) *mux.Route {
	return router.HandleFunc(
		path,
		handlerFunc,
	).Methods(
		method,
	).Name(
		endpoint,
	)
}

// GetEndpointName retrieves the name of the registered route by the current request
func GetEndpointName(request *http.Request) string {
	var route = muxCurrentRoute(request)
	if route == nil {
		return ""
	}
	return route.GetName()
}
