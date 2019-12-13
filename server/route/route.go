package route

import (
	"net/http"

	"github.com/gorilla/mux"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

const (
	stringSeparator string = "|"
)

var registeredRouteActionFuncs map[string]model.ActionFunc

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
	return stringsJoin(queriesTemplates, stringSeparator)
}

func getQueriesRegexp(route *mux.Route) string {
	var queriesRegexps, _ = route.GetQueriesRegexp()
	return stringsJoin(queriesRegexps, stringSeparator)
}

func getMethods(route *mux.Route) string {
	var methods, _ = route.GetMethods()
	return stringsJoin(methods, stringSeparator)
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
	var consolidatedError = apperrorWrapSimpleError(
		[]error{
			pathTemplateError,
			pathRegexpError,
		},
		"Failed to register service route for name [%v]",
		name,
	)
	if consolidatedError != nil {
		return consolidatedError
	}
	loggerAppRoot(
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

// WalkRegisteredRoutes examines the registered router for errors
func WalkRegisteredRoutes(router *mux.Router) error {
	var walkError = router.Walk(
		printRegisteredRouteDetailsFunc,
	)
	if walkError != nil {
		return apperrorWrapSimpleError(
			[]error{
				walkError,
			},
			"Failed to walk through registered routes",
		)
	}
	return nil
}

// CreateRouter initializes a router for route registrations
func CreateRouter() *mux.Router {
	registeredRouteActionFuncs = map[string]model.ActionFunc{}
	return muxNewRouter()
}

// HandleFunc wraps the mux route handler
func HandleFunc(
	router *mux.Router,
	endpoint string,
	method string,
	path string,
	queries []string,
	handleFunc func(http.ResponseWriter, *http.Request),
	actionFunc model.ActionFunc,
) *mux.Route {
	var name = fmtSprintf(
		"%v:%v",
		endpoint,
		method,
	)
	var route = router.HandleFunc(
		path,
		handleFunc,
	).Methods(
		method,
	).Queries(
		queries...,
	).Name(
		name,
	)
	registeredRouteActionFuncs[name] = actionFunc
	return route
}

// HostStatic wraps the mux static content handler
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

// AddMiddleware wraps the mux middleware addition function
func AddMiddleware(
	router *mux.Router,
	middleware model.MiddlewareFunc,
) {
	router.Use((func(http.Handler) http.Handler)(middleware))
}

func defaultActionFunc(session sessionModel.Session) (interface{}, error) {
	return nil, apperrorGetNotImplementedError()
}

func getActionByName(name string) model.ActionFunc {
	var actionFunc, found = registeredRouteActionFuncs[name]
	if !found {
		return defaultActionFunc
	}
	return actionFunc
}

// GetRouteInfo retrieves the registered name and action for the given route
func GetRouteInfo(httpRequest *http.Request) (string, model.ActionFunc, error) {
	var route = muxCurrentRoute(httpRequest)
	if route == nil {
		return "",
			nil,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"Failed to retrieve route info for request - no route found",
			)
	}
	var name = getNameFunc(route)
	var action = getActionByNameFunc(name)
	return name, action, nil
}
