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
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
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
	stringsSplitExpected                    int
	stringsSplitCalled                      int
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
	getEndpointByNameFuncExpected           int
	getEndpointByNameFuncCalled             int
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
	stringsSplitExpected = 0
	stringsSplitCalled = 0
	stringsSplit = func(s string, sep string) []string {
		stringsSplitCalled++
		return nil
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
	getEndpointByNameFuncExpected = 0
	getEndpointByNameFuncCalled = 0
	getEndpointByNameFunc = func(name string) string {
		getEndpointByNameFuncCalled++
		return ""
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
	stringsSplit = strings.Split
	assert.Equal(t, stringsSplitExpected, stringsSplitCalled, "Unexpected number of calls to stringsSplit")
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
	getEndpointByNameFunc = getEndpointByName
	assert.Equal(t, getEndpointByNameFuncExpected, getEndpointByNameFuncCalled, "Unexpected number of calls to getEndpointByNameFunc")
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

type dummySession struct {
	t *testing.T
}

func (session *dummySession) GetID() uuid.UUID {
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySession) GetName() string {
	assert.Fail(session.t, "Unexpected call to GetName")
	return ""
}

func (session *dummySession) GetRequest() *http.Request {
	assert.Fail(session.t, "Unexpected call to GetRequest")
	return nil
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	assert.Fail(session.t, "Unexpected call to GetResponseWriter")
	return nil
}

func (session *dummySession) GetRequestBody(dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestBody")
	return nil
}

func (session *dummySession) GetRequestParameter(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestParameter")
	return nil
}

func (session *dummySession) GetRequestQuery(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQuery")
	return nil
}

func (session *dummySession) GetRequestQueries(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQueries")
	return nil
}

func (session *dummySession) GetRequestHeader(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeader")
	return nil
}

func (session *dummySession) GetRequestHeaders(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeaders")
	return nil
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	assert.Fail(session.t, "Unexpected call to Attach")
	return false
}

func (session *dummySession) Detach(name string) bool {
	assert.Fail(session.t, "Unexpected call to Detach")
	return false
}

func (session *dummySession) GetRawAttachment(name string) (interface{}, bool) {
	assert.Fail(session.t, "Unexpected call to GetRawAttachment")
	return nil, false
}

func (session *dummySession) GetAttachment(name string, dataTemplate interface{}) bool {
	assert.Fail(session.t, "Unexpected call to GetAttachment")
	return false
}

func (session *dummySession) IsLoggingAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool {
	assert.Fail(session.t, "Unexpected call to IsLoggingAllowed")
	return false
}

// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
func (session *dummySession) LogMethodEnter() {
	assert.Fail(session.t, "Unexpected call to LogMethodEnter")
}

// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
func (session *dummySession) LogMethodParameter(parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodParameter")
}

// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
func (session *dummySession) LogMethodLogic(logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodLogic")
}

// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
func (session *dummySession) LogMethodReturn(returns ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodReturn")
}

// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
func (session *dummySession) LogMethodExit() {
	assert.Fail(session.t, "Unexpected call to LogMethodExit")
}

// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
func (session *dummySession) CreateNetworkRequest(method string, url string, payload string, header map[string]string) networkModel.NetworkRequest {
	assert.Fail(session.t, "Unexpected call to CreateNetworkRequest")
	return nil
}
