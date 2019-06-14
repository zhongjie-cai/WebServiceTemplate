package route

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
)

var (
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	apperrorConsolidateAllErrorsExpected    int
	apperrorConsolidateAllErrorsCalled      int
	stringsJoinExpected                     int
	stringsJoinCalled                       int
	fmtSprintfExpected                      int
	fmtSprintfCalled                        int
	loggerAppRootExpected                   int
	loggerAppRootCalled                     int
	muxNewRouterExpected                    int
	muxNewRouterCalled                      int
	muxCurrentRouteExpected                 int
	muxCurrentRouteCalled                   int
	getNameFuncExpected                     int
	getNameFuncCalled                       int
	getPathTemplateFuncExpected             int
	getPathTemplateFuncCalled               int
	getPathRegexpFuncExpected               int
	getPathRegexpFuncCalled                 int
	getQueriesTemplatesFuncExpected         int
	getQueriesTemplatesFuncCalled           int
	getQueriesRegexpFuncExpected            int
	getQueriesRegexpFuncCalled              int
	getMethodsFuncExpected                  int
	getMethodsFuncCalled                    int
	printRegisteredRouteDetailsFuncExpected int
	printRegisteredRouteDetailsFuncCalled   int
	walkRegisteredRoutesFuncExpected        int
	walkRegisteredRoutesFuncCalled          int
)

func createMock(t *testing.T) {
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	apperrorConsolidateAllErrorsExpected = 0
	apperrorConsolidateAllErrorsCalled = 0
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		return nil
	}
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
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
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	muxNewRouterExpected = 0
	muxNewRouterCalled = 0
	muxNewRouter = func() *mux.Router {
		muxNewRouterCalled++
		return nil
	}
	muxCurrentRouteExpected = 0
	muxCurrentRouteCalled = 0
	muxCurrentRoute = func(request *http.Request) *mux.Route {
		muxCurrentRouteCalled++
		return nil
	}
	getNameFuncExpected = 0
	getNameFuncCalled = 0
	getNameFunc = func(route *mux.Route) string {
		getNameFuncCalled++
		return ""
	}
	getPathTemplateFuncExpected = 0
	getPathTemplateFuncCalled = 0
	getPathTemplateFunc = func(route *mux.Route) (string, error) {
		getPathTemplateFuncCalled++
		return "", nil
	}
	getPathRegexpFuncExpected = 0
	getPathRegexpFuncCalled = 0
	getPathRegexpFunc = func(route *mux.Route) (string, error) {
		getPathRegexpFuncCalled++
		return "", nil
	}
	getQueriesTemplatesFuncExpected = 0
	getQueriesTemplatesFuncCalled = 0
	getQueriesTemplatesFunc = func(route *mux.Route) string {
		getQueriesTemplatesFuncCalled++
		return ""
	}
	getQueriesRegexpFuncExpected = 0
	getQueriesRegexpFuncCalled = 0
	getQueriesRegexpFunc = func(route *mux.Route) string {
		getQueriesRegexpFuncCalled++
		return ""
	}
	getMethodsFuncExpected = 0
	getMethodsFuncCalled = 0
	getMethodsFunc = func(route *mux.Route) string {
		getMethodsFuncCalled++
		return ""
	}
	printRegisteredRouteDetailsFuncExpected = 0
	printRegisteredRouteDetailsFuncCalled = 0
	printRegisteredRouteDetailsFunc = func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		printRegisteredRouteDetailsFuncCalled++
		return nil
	}
	walkRegisteredRoutesFuncExpected = 0
	walkRegisteredRoutesFuncCalled = 0
	walkRegisteredRoutesFunc = func(router *mux.Router) error {
		walkRegisteredRoutesFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	assert.Equal(t, apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled, "Unexpected method call to apperrorConsolidateAllErrors")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected method call to stringsJoin")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected method call to fmtSprintf")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected method call to loggerAppRoot")
	muxNewRouter = mux.NewRouter
	assert.Equal(t, muxNewRouterExpected, muxNewRouterCalled, "Unexpected method call to muxNewRouter")
	muxCurrentRoute = mux.CurrentRoute
	assert.Equal(t, muxCurrentRouteExpected, muxCurrentRouteCalled, "Unexpected method call to muxCurrentRoute")
	getPathTemplateFunc = getPathTemplate
	assert.Equal(t, getPathTemplateFuncExpected, getPathTemplateFuncCalled, "Unexpected method call to getPathTemplateFunc")
	getNameFunc = getName
	assert.Equal(t, getNameFuncExpected, getNameFuncCalled, "Unexpected method call to getNameFunc")
	getPathRegexpFunc = getPathRegexp
	assert.Equal(t, getPathRegexpFuncExpected, getPathRegexpFuncCalled, "Unexpected method call to getPathRegexpFunc")
	getQueriesTemplatesFunc = getQueriesTemplates
	assert.Equal(t, getQueriesTemplatesFuncExpected, getQueriesTemplatesFuncCalled, "Unexpected method call to getQueriesTemplatesFunc")
	getQueriesRegexpFunc = getQueriesRegexp
	assert.Equal(t, getQueriesRegexpFuncExpected, getQueriesRegexpFuncCalled, "Unexpected method call to getQueriesRegexpFunc")
	getMethodsFunc = getMethods
	assert.Equal(t, getMethodsFuncExpected, getMethodsFuncCalled, "Unexpected method call to getMethodsFunc")
	printRegisteredRouteDetailsFunc = printRegisteredRouteDetails
	assert.Equal(t, printRegisteredRouteDetailsFuncExpected, printRegisteredRouteDetailsFuncCalled, "Unexpected method call to printRegisteredRouteDetailsFunc")
	walkRegisteredRoutesFunc = walkRegisteredRoutes
	assert.Equal(t, walkRegisteredRoutesFuncExpected, walkRegisteredRoutesFuncCalled, "Unexpected method call to walkRegisteredRoutesFunc")
}
