package register

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
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

func TestDoParameterReplacement_NilReplacementsMap(t *testing.T) {
	// arrange
	var dummyParameterName = "some name"
	var dummyOriginalPath = "/some/original/path/with/{" + dummyParameterName + "}/in/it"
	var dummyParameterType = model.ParameterType("some type")
	var dummyParameterReplacementsMap map[model.ParameterType]string

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "doParameterReplacement", subcategory)
		assert.Equal(t, "Path parameter [%v] in path [%v] has no type specification; fallback to default.", messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyParameterName, parameters[0])
		assert.Equal(t, dummyOriginalPath, parameters[1])
	}

	// SUT + act
	result := doParameterReplacement(
		dummyOriginalPath,
		dummyParameterName,
		dummyParameterType,
		dummyParameterReplacementsMap,
	)

	// assert
	assert.Equal(t, dummyOriginalPath, result)

	// verify
	verifyAll(t)
}

func TestDoParameterReplacement_NoReplacementFound(t *testing.T) {
	// arrange
	var dummyParameterName = "some name"
	var dummyOriginalPath = "/some/original/path/with/{" + dummyParameterName + "}/in/it"
	var dummyParameterType = model.ParameterType("some type")
	var dummyParameterReplacementsMap = map[model.ParameterType]string{
		model.ParameterType("foo"): "bar",
	}

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "doParameterReplacement", subcategory)
		assert.Equal(t, "Path parameter [%v] in path [%v] has no type specification; fallback to default.", messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyParameterName, parameters[0])
		assert.Equal(t, dummyOriginalPath, parameters[1])
	}

	// SUT + act
	result := doParameterReplacement(
		dummyOriginalPath,
		dummyParameterName,
		dummyParameterType,
		dummyParameterReplacementsMap,
	)

	// assert
	assert.Equal(t, dummyOriginalPath, result)

	// verify
	verifyAll(t)
}

func TestDoParameterReplacement_ValidReplacementFound(t *testing.T) {
	// arrange
	var dummyParameterName = "some name"
	var dummyOriginalPath = "/some/original/path/with/{" + dummyParameterName + "}/in/it"
	var dummyParameterType = model.ParameterType("some type")
	var dummyReplacement = "some replacement"
	var dummyParameterReplacementsMap = map[model.ParameterType]string{
		model.ParameterType("foo"): "bar",
		dummyParameterType:         dummyReplacement,
	}
	var dummyResult = "/some/original/path/with/{" + dummyParameterName + ":" + dummyReplacement + "}/in/it"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		if fmtSprintfCalled == 1 {
			assert.Equal(t, "{%v}", format)
			assert.Equal(t, 1, len(a))
			assert.Equal(t, dummyParameterName, a[0])
		} else if fmtSprintfCalled == 2 {
			assert.Equal(t, "{%v:%v}", format)
			assert.Equal(t, 2, len(a))
			assert.Equal(t, dummyParameterName, a[0])
			assert.Equal(t, dummyReplacement, a[1])
		}
		return fmt.Sprintf(format, a...)
	}
	stringsReplaceExpected = 1
	stringsReplace = func(s, old, new string, n int) string {
		stringsReplaceCalled++
		assert.Equal(t, dummyOriginalPath, s)
		assert.Equal(t, "{"+dummyParameterName+"}", old)
		assert.Equal(t, "{"+dummyParameterName+":"+dummyReplacement+"}", new)
		assert.Equal(t, -1, n)
		return strings.Replace(s, old, new, n)
	}

	// SUT + act
	result := doParameterReplacement(
		dummyOriginalPath,
		dummyParameterName,
		dummyParameterType,
		dummyParameterReplacementsMap,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestEvaluatePathWithParameters(t *testing.T) {
	// arrange
	var dummyOriginalPath = "some original path"
	var dummyParameterName1 = "some parameter name 1"
	var dummyParameterType1 = model.ParameterType("some paramter type 1")
	var dummyParameterName2 = "some parameter name 2"
	var dummyParameterType2 = model.ParameterType("some paramter type 2")
	var dummyParameterName3 = "some parameter name 3"
	var dummyParameterType3 = model.ParameterType("some paramter type 3")
	var dummyParameters = map[string]model.Parameter{
		"parameter1": model.Parameter{
			Name: dummyParameterName1,
			Type: dummyParameterType1,
		},
		"parameter2": model.Parameter{
			Name: dummyParameterName2,
			Type: dummyParameterType2,
		},
		"parameter3": model.Parameter{
			Name: dummyParameterName3,
			Type: dummyParameterType3,
		},
	}
	var dummyParameterReplacementsMap = map[model.ParameterType]string{
		model.ParameterType("foo"): "bar",
	}
	var dummyUpdatedPath1 = "some updated path 1"
	var dummyUpdatedPath2 = "some updated path 2"
	var dummyUpdatedPath3 = "some updated path 3"

	// mock
	createMock(t)

	// expect
	doParameterReplacementFuncExpected = 3
	doParameterReplacementFunc = func(originalPath string, parameterName string, parameterType model.ParameterType, parameterReplacementsMap map[model.ParameterType]string) string {
		doParameterReplacementFuncCalled++
		assert.Equal(t, dummyParameterReplacementsMap, parameterReplacementsMap)
		if dummyParameterName1 == parameterName {
			assert.Equal(t, dummyOriginalPath, originalPath)
			assert.Equal(t, dummyParameterType1, parameterType)
			return dummyUpdatedPath1
		} else if dummyParameterName2 == parameterName {
			assert.Equal(t, dummyUpdatedPath1, originalPath)
			assert.Equal(t, dummyParameterType2, parameterType)
			return dummyUpdatedPath2
		} else if dummyParameterName3 == parameterName {
			assert.Equal(t, dummyUpdatedPath2, originalPath)
			assert.Equal(t, dummyParameterType3, parameterType)
			return dummyUpdatedPath3
		}
		return ""
	}

	// SUT + act
	result := evaluatePathWithParameters(
		dummyOriginalPath,
		dummyParameters,
		dummyParameterReplacementsMap,
	)

	// assert
	assert.Equal(t, dummyUpdatedPath3, result)

	// verify
	verifyAll(t)
}

func TestRegisterRoutes_NilRoutesFunc(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}

	// stub
	Routes = nil

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "registerRoutes", subcategory)
		assert.Equal(t, "Routes function not set: no routes registered!", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	registerRoutes(
		dummyRouter,
	)

	// verify
	verifyAll(t)
}

func TestRegisterRoutes_EmptyRoutes(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var routesExpected int
	var routesCalled int
	var dummyRoutes []model.Route

	// mock
	createMock(t)

	// expect
	routesExpected = 1
	Routes = func() []model.Route {
		routesCalled++
		return dummyRoutes
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "registerRoutes", subcategory)
		assert.Equal(t, "Routes function empty: no routes returned!", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	registerRoutes(
		dummyRouter,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, routesExpected, routesCalled, "Unexpected number of calls to Routes")
}

func TestRegisterRoutes_ValidRoutes(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var routesExpected int
	var routesCalled int
	var dummyEndpoint1 = "some endpoint 1"
	var dummyMethod1 = "some method 1"
	var dummyPath1 = "some path 1"
	var dummyParameters1 = map[string]model.Parameter{
		"foo1": model.Parameter{Name: "bar1"},
	}
	var dummyActionFunc1 = func(http.ResponseWriter, *http.Request, uuid.UUID) {}
	var dummyActionFunc1Pointer = fmt.Sprintf("%v", reflect.ValueOf(dummyActionFunc1))
	var dummyEndpoint2 = "some endpoint 2"
	var dummyMethod2 = "some method 2"
	var dummyPath2 = "some path 2"
	var dummyParameters2 = map[string]model.Parameter{
		"foo2": model.Parameter{Name: "bar2"},
	}
	var dummyActionFunc2 = func(http.ResponseWriter, *http.Request, uuid.UUID) {}
	var dummyActionFunc2Pointer = fmt.Sprintf("%v", reflect.ValueOf(dummyActionFunc2))
	var dummyRoutes = []model.Route{
		model.Route{
			Endpoint:   dummyEndpoint1,
			Method:     dummyMethod1,
			Path:       dummyPath1,
			Parameters: dummyParameters1,
			ActionFunc: dummyActionFunc1,
		},
		model.Route{
			Endpoint:   dummyEndpoint2,
			Method:     dummyMethod2,
			Path:       dummyPath2,
			Parameters: dummyParameters2,
			ActionFunc: dummyActionFunc2,
		},
	}
	var dummyEvaluatedPath1 = "some evaluated path 1"
	var dummyEvaluatedPath2 = "some evaluated path 2"

	// mock
	createMock(t)

	// expect
	routesExpected = 1
	Routes = func() []model.Route {
		routesCalled++
		return dummyRoutes
	}
	evaluatePathWithParametersFuncExpected = 2
	evaluatePathWithParametersFunc = func(path string, parameters map[string]model.Parameter, replacementsMap map[model.ParameterType]string) string {
		evaluatePathWithParametersFuncCalled++
		assert.Equal(t, model.ParameterTypeMap, replacementsMap)
		if evaluatePathWithParametersFuncCalled == 1 {
			assert.Equal(t, dummyPath1, path)
			assert.Equal(t, dummyParameters1, parameters)
			return dummyEvaluatedPath1
		} else if evaluatePathWithParametersFuncCalled == 2 {
			assert.Equal(t, dummyPath2, path)
			assert.Equal(t, dummyParameters2, parameters)
			return dummyEvaluatedPath2
		}
		return ""
	}
	routeHandleFuncExpected = 2
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handlerFunc func(http.ResponseWriter, *http.Request), actionFunc func(http.ResponseWriter, *http.Request, uuid.UUID)) *mux.Route {
		routeHandleFuncCalled++
		assert.Equal(t, dummyRouter, router)
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(handlerSession)), fmt.Sprintf("%v", reflect.ValueOf(handlerFunc)))
		if routeHandleFuncCalled == 1 {
			assert.Equal(t, dummyEndpoint1, endpoint)
			assert.Equal(t, dummyMethod1, method)
			assert.Equal(t, dummyEvaluatedPath1, path)
			assert.Equal(t, dummyActionFunc1Pointer, fmt.Sprintf("%v", reflect.ValueOf(actionFunc)))
		} else if routeHandleFuncCalled == 2 {
			assert.Equal(t, dummyEndpoint2, endpoint)
			assert.Equal(t, dummyMethod2, method)
			assert.Equal(t, dummyEvaluatedPath2, path)
			assert.Equal(t, dummyActionFunc2Pointer, fmt.Sprintf("%v", reflect.ValueOf(actionFunc)))
		}
		return nil
	}

	// SUT + act
	registerRoutes(
		dummyRouter,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, routesExpected, routesCalled, "Unexpected number of calls to Routes")
}

func TestRegisterStatics_NilStaticsFunc(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}

	// stub
	Statics = nil

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "registerStatics", subcategory)
		assert.Equal(t, "Statics function not set: no static content registered!", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	registerStatics(
		dummyRouter,
	)

	// verify
	verifyAll(t)
}

func TestRegisterStatics_EmptyStatics(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var routesExpected int
	var routesCalled int
	var dummyStatics []model.Static

	// mock
	createMock(t)

	// expect
	routesExpected = 1
	Statics = func() []model.Static {
		routesCalled++
		return dummyStatics
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "register", category)
		assert.Equal(t, "registerStatics", subcategory)
		assert.Equal(t, "Statics function empty: no static content returned!", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	registerStatics(
		dummyRouter,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, routesExpected, routesCalled, "Unexpected number of calls to Statics")
}

func TestRegisterStatics_ValidStatics(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var staticsExpected int
	var staticsCalled int
	var dummyName1 = "some name 1"
	var dummyPathPrefix1 = "some path prefix 1"
	var dummyHandler1 = dummyHandler{t}
	var dummyName2 = "some name 2"
	var dummyPathPrefix2 = "some path prefix 2"
	var dummyHandler2 = dummyHandler{t}
	var dummyStatics = []model.Static{
		model.Static{
			Name:       dummyName1,
			PathPrefix: dummyPathPrefix1,
			Handler:    dummyHandler1,
		},
		model.Static{
			Name:       dummyName2,
			PathPrefix: dummyPathPrefix2,
			Handler:    dummyHandler2,
		},
	}

	// mock
	createMock(t)

	// expect
	staticsExpected = 1
	Statics = func() []model.Static {
		staticsCalled++
		return dummyStatics
	}
	routeHostStaticExpected = 2
	routeHostStatic = func(router *mux.Router, name string, path string, handler http.Handler) *mux.Route {
		routeHostStaticCalled++
		assert.Equal(t, dummyRouter, router)
		if routeHostStaticCalled == 1 {
			assert.Equal(t, dummyName1, name)
			assert.Equal(t, dummyPathPrefix1, path)
			assert.Equal(t, dummyHandler1, handler)
		} else if routeHostStaticCalled == 2 {
			assert.Equal(t, dummyName2, name)
			assert.Equal(t, dummyPathPrefix2, path)
			assert.Equal(t, dummyHandler2, handler)
		}
		return nil
	}

	// SUT + act
	registerStatics(
		dummyRouter,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, staticsExpected, staticsCalled, "Unexpected number of calls to Statics")
}

func TestInstantiate_RouterError(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}
	var dummyRouteError = errors.New("some route error")
	var dummyMessageFormat = "Failed to instantiate routes"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	routeCreateRouterExpected = 1
	routeCreateRouter = func() *mux.Router {
		routeCreateRouterCalled++
		return dummyRouter
	}
	registerRoutesFuncExpected = 1
	registerRoutesFunc = func(router *mux.Router) {
		registerRoutesFuncCalled++
		assert.Equal(t, dummyRouter, router)
	}
	registerStaticsFuncExpected = 1
	registerStaticsFunc = func(router *mux.Router) {
		registerStaticsFuncCalled++
		assert.Equal(t, dummyRouter, router)
	}
	routeWalkRegisteredRoutesExpected = 1
	routeWalkRegisteredRoutes = func(router *mux.Router) error {
		routeWalkRegisteredRoutesCalled++
		assert.Equal(t, dummyRouter, router)
		return dummyRouteError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyRouteError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	result, err := Instantiate()

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInstantiate_Success(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}

	// mock
	createMock(t)

	// expect
	routeCreateRouterExpected = 1
	routeCreateRouter = func() *mux.Router {
		routeCreateRouterCalled++
		return dummyRouter
	}
	registerRoutesFuncExpected = 1
	registerRoutesFunc = func(router *mux.Router) {
		registerRoutesFuncCalled++
		assert.Equal(t, dummyRouter, router)
	}
	registerStaticsFuncExpected = 1
	registerStaticsFunc = func(router *mux.Router) {
		registerStaticsFuncCalled++
		assert.Equal(t, dummyRouter, router)
	}
	routeWalkRegisteredRoutesExpected = 1
	routeWalkRegisteredRoutes = func(router *mux.Router) error {
		routeWalkRegisteredRoutesCalled++
		assert.Equal(t, dummyRouter, router)
		return nil
	}

	// SUT + act
	result, err := Instantiate()

	// assert
	assert.Equal(t, dummyRouter, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}
