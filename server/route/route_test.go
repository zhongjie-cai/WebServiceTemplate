package route

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestGetName_Undefined(t *testing.T) {
	// arrange
	var dummyRouter = mux.NewRouter()

	// mock
	createMock(t)

	// SUT
	var dummyRoute = dummyRouter.NewRoute()

	// act
	var result = getName(
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
	var result = getName(
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
	var result, err = getPathTemplate(
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
	var result, err = getPathTemplate(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "/foo/{bar}", result)
	assert.NoError(t, err)

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
	var result, err = getPathRegexp(
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
	var result, err = getPathRegexp(
		dummyRoute,
	)

	// assert
	assert.Equal(t, "^/foo/(?P<v0>[^/]+)$", result)
	assert.NoError(t, err)

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
	var result = getQueriesTemplates(
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
	var result = getQueriesTemplates(
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
	var result = getQueriesRegexp(
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
	var result = getQueriesRegexp(
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
	var result = getMethods(
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
	var result = getMethods(
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
	var err = printRegisteredRouteDetails(
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
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
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
	var err = printRegisteredRouteDetails(
		dummyRoute,
		dummyRouter,
		dummyAncestors,
	)

	// assert
	assert.NoError(t, err)

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
	var err = WalkRegisteredRoutes(
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
	var err = WalkRegisteredRoutes(
		dummyRouter,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestHostStatic(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyPath = "/foo/"
	var dummyHandler = dummyHandler{t}

	// mock
	createMock(t)

	// expect
	muxNewRouterExpected = 1
	muxNewRouter = func() *mux.Router {
		muxNewRouterCalled++
		return mux.NewRouter()
	}

	// SUT
	var router = CreateRouter()

	// act
	var route = HostStatic(
		router,
		dummyName,
		dummyPath,
		dummyHandler,
	)
	var name = route.GetName()
	var pathTemplate, _ = route.GetPathTemplate()
	var handler = route.GetHandler()

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
	var dummyHTTPRequest, _ = http.NewRequest(
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
	var dummyActionFuncExpected = 0
	var dummyActionFuncCalled = 0
	var dummyActionFunc = func(uuid.UUID, string, map[string]string) {
		dummyActionFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	muxNewRouterExpected = 1
	muxNewRouter = func() *mux.Router {
		muxNewRouterCalled++
		return mux.NewRouter()
	}

	// SUT
	var router = CreateRouter()

	// act
	var route = HandleFunc(
		router,
		dummyEndpoint,
		dummyMethod,
		dummyPath,
		dummyHandlerFunc,
		dummyActionFunc,
	)
	var name = route.GetName()
	var methods, _ = route.GetMethods()
	var pathTemplate, _ = route.GetPathTemplate()
	route.GetHandler().ServeHTTP(dummyResponseWriter, dummyHTTPRequest)

	// assert
	assert.Equal(t, dummyMethod+":"+dummyEndpoint, name)
	assert.Equal(t, 1, len(methods))
	assert.Equal(t, dummyMethod, methods[0])
	assert.Equal(t, dummyPath, pathTemplate)
	assert.Equal(t, dummyHandlerFuncExpected, dummyHandlerFuncCalled)
	assert.Equal(t, dummyActionFuncExpected, dummyActionFuncCalled)

	// verify
	verifyAll(t)
}

func TestDefaultActionFunc(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequestBody = "some request body"
	var dummyParameters = map[string]string{"foo": "bar"}
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorGetNotImplementedErrorExpected = 1
	apperrorGetNotImplementedError = func(innerError error) apperror.AppError {
		apperrorGetNotImplementedErrorCalled++
		assert.NoError(t, innerError)
		return dummyAppError
	}
	responseErrorExpected = 1
	responseError = func(sessionID uuid.UUID, err error) {
		responseErrorCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyAppError, err)
	}

	// SUT + act
	defaultActionFunc(
		dummySessionID,
		dummyRequestBody,
		dummyParameters,
	)

	// verify
	verifyAll(t)
}

func TestGetActionByName_NotFound(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyAction func(sessionID uuid.UUID, requestBody string, parameters map[string]string)
	var dummyOtherName = "some other name"
	var expectedActionPointer = fmt.Sprintf("%v", reflect.ValueOf(defaultActionFunc))

	// stub
	registeredRouteActionFuncs = map[string]func(uuid.UUID, string, map[string]string){
		dummyName: dummyAction,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getActionByName(
		dummyOtherName,
	)

	// assert
	assert.Equal(t, expectedActionPointer, fmt.Sprintf("%v", reflect.ValueOf(result)))

	// verify
	verifyAll(t)
}

func TestGetActionByName_Found(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyActionExpected = 0
	var dummyActionCalled = 0
	var dummyAction = func(sessionID uuid.UUID, requestBody string, parameters map[string]string) {
		dummyActionCalled++
	}
	var expectedActionPointer = fmt.Sprintf("%v", reflect.ValueOf(dummyAction))

	// stub
	registeredRouteActionFuncs = map[string]func(uuid.UUID, string, map[string]string){
		dummyName: dummyAction,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getActionByName(
		dummyName,
	)

	// assert
	assert.Equal(t, expectedActionPointer, fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyActionExpected, dummyActionCalled, "Unexpected number of calls to dummyAction")

	// verify
	verifyAll(t)
}

func TestGetRouteInfo_NilRoute(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyRoute *mux.Route
	var dummyMessageFormat = "Failed to retrieve route info for request - no route found"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	muxCurrentRouteExpected = 1
	muxCurrentRoute = func(httpRequest *http.Request) *mux.Route {
		muxCurrentRouteCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRoute
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var name, action, err = GetRouteInfo(
		dummyHTTPRequest,
	)

	// assert
	assert.Zero(t, name)
	assert.Nil(t, action)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetRouteInfo_ValidRoute(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyRoute = &mux.Route{}
	var dummyName = "some name"
	var dummyActionExpected = 0
	var dummyActionCalled = 0
	var dummyAction = func(sessionID uuid.UUID, requestBody string, parameters map[string]string) {
		dummyActionCalled++
	}
	var dummyActionPointer = fmt.Sprintf("%v", reflect.ValueOf(dummyAction))

	// mock
	createMock(t)

	// expect
	muxCurrentRouteExpected = 1
	muxCurrentRoute = func(httpRequest *http.Request) *mux.Route {
		muxCurrentRouteCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRoute
	}
	getNameFuncExpected = 1
	getNameFunc = func(route *mux.Route) string {
		getNameFuncCalled++
		assert.Equal(t, dummyRoute, route)
		return dummyName
	}
	getActionByNameFuncExpected = 1
	getActionByNameFunc = func(name string) func(uuid.UUID, string, map[string]string) {
		getActionByNameFuncCalled++
		assert.Equal(t, dummyName, name)
		return dummyAction
	}

	// SUT + act
	var name, action, err = GetRouteInfo(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyName, name)
	assert.Equal(t, dummyActionPointer, fmt.Sprintf("%v", reflect.ValueOf(action)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyActionExpected, dummyActionCalled, "Unexpected number of calls to dummyAction")
}
