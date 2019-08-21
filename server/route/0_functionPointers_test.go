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
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

var (
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	apperrorConsolidateAllErrorsExpected    int
	apperrorConsolidateAllErrorsCalled      int
	apperrorGetNotImplementedErrorExpected  int
	apperrorGetNotImplementedErrorCalled    int
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
	responseWriteExpected                   int
	responseWriteCalled                     int
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
	apperrorGetNotImplementedErrorExpected = 0
	apperrorGetNotImplementedErrorCalled = 0
	apperrorGetNotImplementedError = func(innerError error) apperror.AppError {
		apperrorGetNotImplementedErrorCalled++
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
	responseWriteExpected = 0
	responseWriteCalled = 0
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError apperror.AppError) {
		responseWriteCalled++
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
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	assert.Equal(t, apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled, "Unexpected number of calls to apperrorConsolidateAllErrors")
	apperrorGetNotImplementedError = apperror.GetNotImplementedError
	assert.Equal(t, apperrorGetNotImplementedErrorExpected, apperrorGetNotImplementedErrorCalled, "Unexpected number of calls to apperrorGetNotImplementedError")
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
	responseWrite = response.Write
	assert.Equal(t, responseWriteExpected, responseWriteCalled, "Unexpected number of calls to responseWrite")
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
type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.WriteHeader")
}

type dummyHandler struct {
	t *testing.T
}

func (dh dummyHandler) ServeHTTP(responseWriter http.ResponseWriter, httphttpRequest *http.Request) {
	assert.Fail(dh.t, "Unexpected number of calls to ServeHTTP")
}
