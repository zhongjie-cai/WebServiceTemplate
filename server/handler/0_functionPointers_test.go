package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

var (
	routeGetRouteInfoExpected             int
	routeGetRouteInfoCalled               int
	sessionRegisterExpected               int
	sessionRegisterCalled                 int
	panicHandleExpected                   int
	panicHandleCalled                     int
	responseWriteExpected                 int
	responseWriteCalled                   int
	loggerAPIEnterExpected                int
	loggerAPIEnterCalled                  int
	loggerAPIExitExpected                 int
	loggerAPIExitCalled                   int
	apperrorGetInvalidOperationExpected   int
	apperrorGetInvalidOperationCalled     int
	timeutilGetTimeNowUTCExpected         int
	timeutilGetTimeNowUTCCalled           int
	timeSinceExpected                     int
	timeSinceCalled                       int
	executeCustomizedFunctionFuncExpected int
	executeCustomizedFunctionFuncCalled   int
	customizationPreActionFuncExpected    int
	customizationPreActionFuncCalled      int
	customizationPostActionFuncExpected   int
	customizationPostActionFuncCalled     int
	requestFullDumpExpected               int
	requestFullDumpCalled                 int
	loggerAppRootExpected                 int
	loggerAppRootCalled                   int
	httpErrorExpected                     int
	httpErrorCalled                       int
)

func createMock(t *testing.T) {
	routeGetRouteInfoExpected = 0
	routeGetRouteInfoCalled = 0
	routeGetRouteInfo = func(httpRequest *http.Request) (string, model.ActionFunc, error) {
		routeGetRouteInfoCalled++
		return "", nil, nil
	}
	sessionRegisterExpected = 0
	sessionRegisterCalled = 0
	sessionRegister = func(name string, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		return nil
	}
	panicHandleExpected = 0
	panicHandleCalled = 0
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
	}
	responseWriteExpected = 0
	responseWriteCalled = 0
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
	}
	loggerAPIEnterExpected = 0
	loggerAPIEnterCalled = 0
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
	}
	loggerAPIExitExpected = 0
	loggerAPIExitCalled = 0
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
	}
	apperrorGetInvalidOperationExpected = 0
	apperrorGetInvalidOperationCalled = 0
	apperrorGetInvalidOperation = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetInvalidOperationCalled++
		return nil
	}
	timeutilGetTimeNowUTCExpected = 0
	timeutilGetTimeNowUTCCalled = 0
	timeutilGetTimeNowUTC = func() time.Time {
		timeutilGetTimeNowUTCCalled++
		return time.Time{}
	}
	timeSinceExpected = 0
	timeSinceCalled = 0
	timeSince = func(ts time.Time) time.Duration {
		timeSinceCalled++
		return 0
	}
	executeCustomizedFunctionFuncExpected = 0
	executeCustomizedFunctionFuncCalled = 0
	executeCustomizedFunctionFunc = func(session sessionModel.Session, customFunc func(sessionModel.Session) error) error {
		executeCustomizedFunctionFuncCalled++
		return nil
	}
	customizationPreActionFuncExpected = 0
	customizationPreActionFuncCalled = 0
	customization.PreActionFunc = func(session sessionModel.Session) error {
		customizationPreActionFuncCalled++
		return nil
	}
	customizationPostActionFuncExpected = 0
	customizationPostActionFuncCalled = 0
	customization.PostActionFunc = func(session sessionModel.Session) error {
		customizationPostActionFuncCalled++
		return nil
	}
	requestFullDumpExpected = 0
	requestFullDumpCalled = 0
	requestFullDump = func(httpRequest *http.Request) string {
		requestFullDumpCalled++
		return ""
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	httpErrorExpected = 0
	httpErrorCalled = 0
	httpError = func(w http.ResponseWriter, error string, code int) {
		httpErrorCalled++
	}
}

func verifyAll(t *testing.T) {
	routeGetRouteInfo = route.GetRouteInfo
	assert.Equal(t, routeGetRouteInfoExpected, routeGetRouteInfoCalled, "Unexpected number of calls to routeGetRouteInfo")
	sessionRegister = session.Register
	assert.Equal(t, sessionRegisterExpected, sessionRegisterCalled, "Unexpected number of calls to sessionRegister")
	panicHandle = panic.Handle
	assert.Equal(t, panicHandleExpected, panicHandleCalled, "Unexpected number of calls to panicHandle")
	responseWrite = response.Write
	assert.Equal(t, responseWriteExpected, responseWriteCalled, "Unexpected number of calls to responseWrite")
	loggerAPIEnter = logger.APIEnter
	assert.Equal(t, loggerAPIEnterExpected, loggerAPIEnterCalled, "Unexpected number of calls to loggerAPIEnter")
	loggerAPIExit = logger.APIExit
	assert.Equal(t, loggerAPIExitExpected, loggerAPIExitCalled, "Unexpected number of calls to loggerAPIExit")
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
	assert.Equal(t, apperrorGetInvalidOperationExpected, apperrorGetInvalidOperationCalled, "Unexpected number of calls to apperrorGetInvalidOperation")
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	assert.Equal(t, timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled, "Unexpected number of calls to timeutilGetTimeNowUTC")
	timeSince = time.Since
	assert.Equal(t, timeSinceExpected, timeSinceCalled, "Unexpected number of calls to timeSince")
	executeCustomizedFunctionFunc = executeCustomizedFunction
	assert.Equal(t, executeCustomizedFunctionFuncExpected, executeCustomizedFunctionFuncCalled, "Unexpected number of calls to executeCustomizedFunctionFunc")
	customization.PreActionFunc = nil
	assert.Equal(t, customizationPreActionFuncExpected, customizationPreActionFuncCalled, "Unexpected number of calls to customization.PreActionFunc")
	customization.PostActionFunc = nil
	assert.Equal(t, customizationPostActionFuncExpected, customizationPostActionFuncCalled, "Unexpected number of calls to customization.PostActionFunc")
	requestFullDump = request.FullDump
	assert.Equal(t, requestFullDumpExpected, requestFullDumpCalled, "Unexpected number of calls to requestFullDump")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	httpError = http.Error
	assert.Equal(t, httpErrorExpected, httpErrorCalled, "Unexpected number of calls to httpError")
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
