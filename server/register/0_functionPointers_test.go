package register

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/handler"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

var (
	stringsReplaceExpected                 int
	stringsReplaceCalled                   int
	fmtSprintfExpected                     int
	fmtSprintfCalled                       int
	loggerAppRootExpected                  int
	loggerAppRootCalled                    int
	routeHandleFuncExpected                int
	routeHandleFuncCalled                  int
	routeHostStaticExpected                int
	routeHostStaticCalled                  int
	routeCreateRouterExpected              int
	routeCreateRouterCalled                int
	routeWalkRegisteredRoutesExpected      int
	routeWalkRegisteredRoutesCalled        int
	apperrorWrapSimpleErrorExpected        int
	apperrorWrapSimpleErrorCalled          int
	handlerSessionExpected                 int
	handlerSessionCalled                   int
	doParameterReplacementFuncExpected     int
	doParameterReplacementFuncCalled       int
	evaluatePathWithParametersFuncExpected int
	evaluatePathWithParametersFuncCalled   int
	registerRoutesFuncExpected             int
	registerRoutesFuncCalled               int
	registerStaticsFuncExpected            int
	registerStaticsFuncCalled              int
)

func createMock(t *testing.T) {
	stringsReplaceExpected = 0
	stringsReplaceCalled = 0
	stringsReplace = func(s, old, new string, n int) string {
		stringsReplaceCalled++
		return ""
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	routeHandleFuncExpected = 0
	routeHandleFuncCalled = 0
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handlerFunc func(http.ResponseWriter, *http.Request), actionFunc func(http.ResponseWriter, *http.Request, uuid.UUID)) *mux.Route {
		routeHandleFuncCalled++
		return nil
	}
	routeHostStaticExpected = 0
	routeHostStaticCalled = 0
	routeHostStatic = func(router *mux.Router, name string, path string, handler http.Handler) *mux.Route {
		routeHostStaticCalled++
		return nil
	}
	routeCreateRouterExpected = 0
	routeCreateRouterCalled = 0
	routeCreateRouter = func() *mux.Router {
		routeCreateRouterCalled++
		return nil
	}
	routeWalkRegisteredRoutesExpected = 0
	routeWalkRegisteredRoutesCalled = 0
	routeWalkRegisteredRoutes = func(router *mux.Router) error {
		routeWalkRegisteredRoutesCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	handlerSessionExpected = 0
	handlerSessionCalled = 0
	handlerSession = func(responseWriter http.ResponseWriter, httpRequst *http.Request) {
		handlerSessionCalled++
	}
	doParameterReplacementFuncExpected = 0
	doParameterReplacementFuncCalled = 0
	doParameterReplacementFunc = func(originalPath string, parameterName string, parameterType model.ParameterType, parameterReplacementsMap map[model.ParameterType]string) string {
		doParameterReplacementFuncCalled++
		return ""
	}
	evaluatePathWithParametersFuncExpected = 0
	evaluatePathWithParametersFuncCalled = 0
	evaluatePathWithParametersFunc = func(path string, parameters map[string]model.Parameter, replacementsMap map[model.ParameterType]string) string {
		evaluatePathWithParametersFuncCalled++
		return ""
	}
	registerRoutesFuncExpected = 0
	registerRoutesFuncCalled = 0
	registerRoutesFunc = func(router *mux.Router) {
		registerRoutesFuncCalled++
	}
	registerStaticsFuncExpected = 0
	registerStaticsFuncCalled = 0
	registerStaticsFunc = func(router *mux.Router) {
		registerStaticsFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	stringsReplace = strings.Replace
	assert.Equal(t, stringsReplaceExpected, stringsReplaceCalled, "Unexpected number of calls to stringsReplace")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	routeHandleFunc = route.HandleFunc
	assert.Equal(t, routeHandleFuncExpected, routeHandleFuncCalled, "Unexpected number of calls to routeHandleFunc")
	routeHostStatic = route.HostStatic
	assert.Equal(t, routeHostStaticExpected, routeHostStaticCalled, "Unexpected number of calls to routeHostStatic")
	routeCreateRouter = route.CreateRouter
	assert.Equal(t, routeCreateRouterExpected, routeCreateRouterCalled, "Unexpected number of calls to routeCreateRouter")
	routeWalkRegisteredRoutes = route.WalkRegisteredRoutes
	assert.Equal(t, routeWalkRegisteredRoutesExpected, routeWalkRegisteredRoutesCalled, "Unexpected number of calls to routeWalkRegisteredRoutes")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	handlerSession = handler.Session
	assert.Equal(t, handlerSessionExpected, handlerSessionCalled, "Unexpected number of calls to handlerSession")
	doParameterReplacementFunc = doParameterReplacement
	assert.Equal(t, doParameterReplacementFuncExpected, doParameterReplacementFuncCalled, "Unexpected number of calls to doParameterReplacementFunc")
	evaluatePathWithParametersFunc = evaluatePathWithParameters
	assert.Equal(t, evaluatePathWithParametersFuncExpected, evaluatePathWithParametersFuncCalled, "Unexpected number of calls to evaluatePathWithParametersFunc")
	registerRoutesFunc = registerRoutes
	assert.Equal(t, registerRoutesFuncExpected, registerRoutesFuncCalled, "Unexpected number of calls to registerRoutesFunc")
	registerStaticsFunc = registerStatics
	assert.Equal(t, registerStaticsFuncExpected, registerStaticsFuncCalled, "Unexpected number of calls to registerStaticsFunc")
}

// mock structs
type dummyHandler struct {
	t *testing.T
}

func (dh dummyHandler) ServeHTTP(responseWriter http.ResponseWriter, httphttpRequest *http.Request) {
	assert.Fail(dh.t, "Unexpected number of calls to ServeHTTP")
}
