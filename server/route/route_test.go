package route

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// mock struct
type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.WriteHeader")
}

func TestGetName_Undefined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result := getName(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetName_Defined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Name(
		"test",
	)

	// act
	result := getName(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "test", result)

	// verify
	verifyAll(t)
}

func TestGetPathTemplate_Error(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result, err := getPathTemplate(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "mux: route doesn't have a path", err.Error())

	// verify
	verifyAll(t)
}

func TestGetPathTemplate_Success(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Path(
		"/foo/{bar}",
	)

	// act
	result, err := getPathTemplate(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "/foo/{bar}", result)
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestGetPathRegexp_Error(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result, err := getPathRegexp(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "mux: route does not have a path", err.Error())

	// verify
	verifyAll(t)
}

func TestGetPathRegexp_Success(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Path(
		"/foo/{bar}",
	)

	// act
	result, err := getPathRegexp(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "^/foo/(?P<v0>[^/]+)$", result)
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestGetQueriesTemplate_Undefined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result := getQueriesTemplates(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetQueriesTemplate_Defined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Queries(
		"abc",
		"{def}",
		"xyz",
		"{zyx}",
	)

	// act
	result := getQueriesTemplates(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "abc={def},xyz={zyx}", result)

	// verify
	verifyAll(t)
}

func TestGetQueriesRegexp_Undefined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result := getQueriesRegexp(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetQueriesRegexp_Defined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Queries(
		"abc",
		"{def}",
		"xyz",
		"{zyx}",
	)

	// act
	result := getQueriesRegexp(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "^abc=(?P<v0>.*)$,^xyz=(?P<v0>.*)$", result)

	// verify
	verifyAll(t)
}

func TestGetMethods_Undefined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	result := getMethods(
		dummyRoute,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetMethods_Defined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var dummyRoute = dummyRouter.NewRoute().Methods(
		"GET",
		"PUT",
	)

	// act
	result := getMethods(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "GET,PUT", result)

	// verify
	verifyAll(t)
}

func TestPrintRegisteredRouteDetails_ErrorConsolidated(t *testing.T) {
	// arrange
	var dummyRoute = &mux.Route{}
	var dummyRouter = &mux.Router{}
	var dummyAncestors = []*mux.Route{}
	var dummyName = "some name"
	var dummyPathTemplate string
	var dummyPathRegexp string
	var dummyQueriesTemplates string
	var dummyQueriesRegexps string
	var dummyMethods string
	var dummyPathTemplateError = errors.New("some path template error")
	var dummyPathRegexpError = errors.New("some path regexp error")
	var dummyMessageFormat = "Failed to register service route for name [%v]"
	var dummyBaseErrorMessage = "Failed to register service route for name [" + dummyName + "]"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	getNameFuncExpected = 1
	getNameFunc = func(route *mux.Route) string {
		getNameFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyName
	}
	getPathTemplateFuncExpected = 1
	getPathTemplateFunc = func(route *mux.Route) (string, error) {
		getPathTemplateFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyPathTemplate, dummyPathTemplateError
	}
	getPathRegexpFuncExpected = 1
	getPathRegexpFunc = func(route *mux.Route) (string, error) {
		getPathRegexpFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyPathRegexp, dummyPathRegexpError
	}
	getQueriesTemplatesFuncExpected = 1
	getQueriesTemplatesFunc = func(route *mux.Route) string {
		getQueriesTemplatesFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyQueriesTemplates
	}
	getQueriesRegexpFuncExpected = 1
	getQueriesRegexpFunc = func(route *mux.Route) string {
		getQueriesRegexpFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyQueriesRegexps
	}
	getMethodsFuncExpected = 1
	getMethodsFunc = func(route *mux.Route) string {
		getMethodsFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyMethods
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyBaseErrorMessage
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, dummyBaseErrorMessage, baseErrorMessage)
		assert.Equal(t, 2, len(allErrors))
		assert.Equal(t, dummyPathTemplateError, allErrors[0])
		assert.Equal(t, dummyPathRegexpError, allErrors[1])
		return dummyAppError
	}

	// SUT + act
	err := printRegisteredRouteDetails(
		dummyRoute,
		dummyRouter,
		dummyAncestors,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestPrintRegisteredRouteDetails_Success(t *testing.T) {
	// arrange
	var dummyRoute = &mux.Route{}
	var dummyRouter = &mux.Router{}
	var dummyAncestors = []*mux.Route{}
	var dummyName = "some name"
	var dummyPathTemplate = "some path template"
	var dummyPathRegexp = "some path regexp"
	var dummyQueriesTemplates = "some queries templates"
	var dummyQueriesRegexps = "some queries regexps"
	var dummyMethods = "some methods"
	var dummyPathTemplateError error
	var dummyPathRegexpError error
	var dummyMessageFormat = "Failed to register service route for name [%v]"
	var dummyBaseErrorMessage = "Failed to register service route for name [" + dummyName + "]"
	var dummyLoggerMessageFormat = "Route registered for name [%v]\nPath template:%v\nPath regexp:%v\nQueries templates:%v\nQueries regexps:%v\nMethods:%v"

	// mock
	createMock(t)

	// expect
	getNameFuncExpected = 1
	getNameFunc = func(route *mux.Route) string {
		getNameFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyName
	}
	getPathTemplateFuncExpected = 1
	getPathTemplateFunc = func(route *mux.Route) (string, error) {
		getPathTemplateFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyPathTemplate, dummyPathTemplateError
	}
	getPathRegexpFuncExpected = 1
	getPathRegexpFunc = func(route *mux.Route) (string, error) {
		getPathRegexpFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyPathRegexp, dummyPathRegexpError
	}
	getQueriesTemplatesFuncExpected = 1
	getQueriesTemplatesFunc = func(route *mux.Route) string {
		getQueriesTemplatesFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyQueriesTemplates
	}
	getQueriesRegexpFuncExpected = 1
	getQueriesRegexpFunc = func(route *mux.Route) string {
		getQueriesRegexpFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyQueriesRegexps
	}
	getMethodsFuncExpected = 1
	getMethodsFunc = func(route *mux.Route) string {
		getMethodsFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyMethods
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyBaseErrorMessage
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, dummyBaseErrorMessage, baseErrorMessage)
		assert.Equal(t, 2, len(allErrors))
		assert.Equal(t, dummyPathTemplateError, allErrors[0])
		assert.Equal(t, dummyPathRegexpError, allErrors[1])
		return nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "route", category)
		assert.Equal(t, "printRegisteredRouteDetails", subcategory)
		assert.Equal(t, dummyLoggerMessageFormat, messageFormat)
		assert.Equal(t, 6, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyPathTemplate, parameters[1])
		assert.Equal(t, dummyPathRegexp, parameters[2])
		assert.Equal(t, dummyQueriesTemplates, parameters[3])
		assert.Equal(t, dummyQueriesRegexps, parameters[4])
		assert.Equal(t, dummyMethods, parameters[5])
	}

	// SUT + act
	err := printRegisteredRouteDetails(
		dummyRoute,
		dummyRouter,
		dummyAncestors,
	)

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestWalkRegisteredRoutes_Error(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to walk through registered routes"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	dummyRouter.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})

	// mock
	createMock(t)

	// expect
	printRegisteredRouteDetailsFuncExpected = 1
	printRegisteredRouteDetailsFunc = func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		printRegisteredRouteDetailsFuncCalled++
		return dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	err := walkRegisteredRoutes(
		dummyRouter,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestWalkRegisteredRoutes_Success(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}

	// stub
	dummyRouter.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})

	// mock
	createMock(t)

	// expect
	printRegisteredRouteDetailsFuncExpected = 1
	printRegisteredRouteDetailsFunc = func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		printRegisteredRouteDetailsFuncCalled++
		return nil
	}

	// SUT + act
	err := walkRegisteredRoutes(
		dummyRouter,
	)

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestRegisterEntries_NilEntries(t *testing.T) {
	// arrange
	var dummyMessageFormat = "No host entries found"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var router, err = RegisterEntries()

	// assert
	assert.Nil(t, router)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestRegisterEntries_EmptyEntries(t *testing.T) {
	// arrange
	var dummyMessageFormat = "No host entries found"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	var dummyEntryFuncs = []func(*mux.Router){}

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var router, err = RegisterEntries(
		dummyEntryFuncs...,
	)

	// assert
	assert.Nil(t, router)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestRegisterEntries_ErrorRegister(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var entryFuncsCalled = 0
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to register routes"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	var dummyEntryFuncs = []func(router *mux.Router){
		func(router *mux.Router) { entryFuncsCalled++ },
		func(router *mux.Router) { entryFuncsCalled++ },
		func(router *mux.Router) { entryFuncsCalled++ },
	}

	// mock
	createMock(t)

	// expect
	muxNewRouterExpected = 1
	muxNewRouter = func() *mux.Router {
		muxNewRouterCalled++
		return dummyRouter
	}
	walkRegisteredRoutesFuncExpected = 1
	walkRegisteredRoutesFunc = func(router *mux.Router) error {
		walkRegisteredRoutesFuncCalled++
		assert.Equal(t, dummyRouter, router)
		return dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var router, err = RegisterEntries(
		dummyEntryFuncs...,
	)

	// assert
	assert.Nil(t, router)
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, len(dummyEntryFuncs), entryFuncsCalled)

	// verify
	verifyAll(t)
}

func TestRegisterEntries_Success(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var entryFuncsCalled = 0

	// stub
	var dummyEntryFuncs = []func(router *mux.Router){
		func(router *mux.Router) { entryFuncsCalled++ },
		func(router *mux.Router) { entryFuncsCalled++ },
		func(router *mux.Router) { entryFuncsCalled++ },
	}

	// mock
	createMock(t)

	// expect
	muxNewRouterExpected = 1
	muxNewRouter = func() *mux.Router {
		muxNewRouterCalled++
		return dummyRouter
	}
	walkRegisteredRoutesFuncExpected = 1
	walkRegisteredRoutesFunc = func(router *mux.Router) error {
		walkRegisteredRoutesFuncCalled++
		assert.Equal(t, dummyRouter, router)
		return nil
	}

	// SUT + act
	var router, err = RegisterEntries(
		dummyEntryFuncs...,
	)

	// assert
	assert.Equal(t, dummyRouter, router)
	assert.Nil(t, err)
	assert.Equal(t, len(dummyEntryFuncs), entryFuncsCalled)

	// verify
	verifyAll(t)
}

func TestHostStatic(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyPath = "/foo/"
	var dummyHandler http.Handler

	// mock
	createMock(t)

	// SUT
	var router = mux.NewRouter()

	// act
	route := HostStatic(
		router,
		dummyName,
		dummyPath,
		dummyHandler,
	)
	name := route.GetName()
	pathTemplate, _ := route.GetPathTemplate()
	handler := route.GetHandler()

	// assert
	assert.Equal(t, dummyName, name)
	assert.Equal(t, dummyPath, pathTemplate)
	assert.Equal(t, dummyHandler, handler)

	// verify
	verifyAll(t)
}

func TestHandleFunc(t *testing.T) {
	// arrange
	var dummyEndpoint = "some endpoint"
	var dummyMethod = "PUT"
	var dummyPath = "/foo/{bar}"
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)

	// stub
	var dummyHandlerFuncExpected = 1
	var dummyHandlerFuncCalled = 0
	var dummyHandlerFunc = func(http.ResponseWriter, *http.Request) {
		dummyHandlerFuncCalled++
	}

	// mock
	createMock(t)

	// SUT
	var router = mux.NewRouter()

	// act
	route := HandleFunc(
		router,
		dummyEndpoint,
		dummyMethod,
		dummyPath,
		dummyHandlerFunc,
	)
	name := route.GetName()
	methods, _ := route.GetMethods()
	pathTemplate, _ := route.GetPathTemplate()
	route.GetHandler().ServeHTTP(dummyResponseWriter, dummyRequest)

	// assert
	assert.Equal(t, dummyEndpoint, name)
	assert.Equal(t, 1, len(methods))
	assert.Equal(t, dummyMethod, methods[0])
	assert.Equal(t, dummyPath, pathTemplate)
	assert.Equal(t, dummyHandlerFuncExpected, dummyHandlerFuncCalled)

	// verify
	verifyAll(t)
}

func TestGetEndpointName_NilRoute(t *testing.T) {
	// arrange
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyRoute *mux.Route

	// mock
	createMock(t)

	// expect
	muxCurrentRouteExpected = 1
	muxCurrentRoute = func(request *http.Request) *mux.Route {
		muxCurrentRouteCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyRoute
	}

	// SUT + act
	result := GetEndpointName(
		dummyRequest,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetEndpointName_ValidRoute(t *testing.T) {
	// arrange
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyName = "some name"
	var dummyRoute = mux.NewRouter().NewRoute().Name(dummyName)

	// mock
	createMock(t)

	// expect
	muxCurrentRouteExpected = 1
	muxCurrentRoute = func(request *http.Request) *mux.Route {
		muxCurrentRouteCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyRoute
	}

	// SUT + act
	result := GetEndpointName(
		dummyRequest,
	)

	// assert
	assert.Equal(t, dummyName, result)

	// verify
	verifyAll(t)
}
