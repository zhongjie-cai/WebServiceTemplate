package route

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

var (
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	apperrorGetNotImplementedErrorExpected  int
	apperrorGetNotImplementedErrorCalled    int
	apperrorGetCustomErrorExpected          int
	apperrorGetCustomErrorCalled            int
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
	getActionByNameFuncExpected             int
	getActionByNameFuncCalled               int
	printRegisteredRouteDetailsFuncExpected int
	printRegisteredRouteDetailsFuncCalled   int
)

func createMock(t *testing.T) {
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	apperrorGetNotImplementedErrorExpected = 0
	apperrorGetNotImplementedErrorCalled = 0
	apperrorGetNotImplementedError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetNotImplementedErrorCalled++
		return nil
	}
	apperrorGetCustomErrorExpected = 0
	apperrorGetCustomErrorCalled = 0
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
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
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
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
	muxCurrentRoute = func(httpRequest *http.Request) *mux.Route {
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
	getActionByNameFuncExpected = 0
	getActionByNameFuncCalled = 0
	getActionByNameFunc = func(name string) model.ActionFunc {
		getActionByNameFuncCalled++
		return nil
	}
	printRegisteredRouteDetailsFuncExpected = 0
	printRegisteredRouteDetailsFuncCalled = 0
	printRegisteredRouteDetailsFunc = func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		printRegisteredRouteDetailsFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	apperrorGetNotImplementedError = apperror.GetNotImplementedError
	assert.Equal(t, apperrorGetNotImplementedErrorExpected, apperrorGetNotImplementedErrorCalled, "Unexpected number of calls to apperrorGetNotImplementedError")
	apperrorGetCustomError = apperror.GetCustomError
	assert.Equal(t, apperrorGetCustomErrorExpected, apperrorGetCustomErrorCalled, "Unexpected number of calls to apperrorGetCustomError")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	muxNewRouter = mux.NewRouter
	assert.Equal(t, muxNewRouterExpected, muxNewRouterCalled, "Unexpected number of calls to muxNewRouter")
	muxCurrentRoute = mux.CurrentRoute
	assert.Equal(t, muxCurrentRouteExpected, muxCurrentRouteCalled, "Unexpected number of calls to muxCurrentRoute")
	getPathTemplateFunc = getPathTemplate
	assert.Equal(t, getPathTemplateFuncExpected, getPathTemplateFuncCalled, "Unexpected number of calls to getPathTemplateFunc")
	getNameFunc = getName
	assert.Equal(t, getNameFuncExpected, getNameFuncCalled, "Unexpected number of calls to getNameFunc")
	getPathRegexpFunc = getPathRegexp
	assert.Equal(t, getPathRegexpFuncExpected, getPathRegexpFuncCalled, "Unexpected number of calls to getPathRegexpFunc")
	getQueriesTemplatesFunc = getQueriesTemplates
	assert.Equal(t, getQueriesTemplatesFuncExpected, getQueriesTemplatesFuncCalled, "Unexpected number of calls to getQueriesTemplatesFunc")
	getQueriesRegexpFunc = getQueriesRegexp
	assert.Equal(t, getQueriesRegexpFuncExpected, getQueriesRegexpFuncCalled, "Unexpected number of calls to getQueriesRegexpFunc")
	getMethodsFunc = getMethods
	assert.Equal(t, getMethodsFuncExpected, getMethodsFuncCalled, "Unexpected number of calls to getMethodsFunc")
	getActionByNameFunc = getActionByName
	assert.Equal(t, getActionByNameFuncExpected, getActionByNameFuncCalled, "Unexpected number of calls to getActionByNameFunc")
	printRegisteredRouteDetailsFunc = printRegisteredRouteDetails
	assert.Equal(t, printRegisteredRouteDetailsFuncExpected, printRegisteredRouteDetailsFuncCalled, "Unexpected number of calls to printRegisteredRouteDetailsFunc")
}

// mock structs
type dummyHandler struct {
	t *testing.T
}

func (dh dummyHandler) ServeHTTP(responseWriter http.ResponseWriter, httphttpRequest *http.Request) {
	assert.Fail(dh.t, "Unexpected number of calls to ServeHTTP")
}
