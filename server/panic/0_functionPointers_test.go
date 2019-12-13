package panic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	fmtErrorfExpected                      int
	fmtErrorfCalled                        int
	getRecoverErrorFuncExpected            int
	getRecoverErrorFuncCalled              int
	loggerAppRootExpected                  int
	loggerAppRootCalled                    int
	responseWriteExpected                  int
	responseWriteCalled                    int
	apperrorGetGeneralFailureErrorExpected int
	apperrorGetGeneralFailureErrorCalled   int
	getDebugStackFuncExpected              int
	getDebugStackFuncCalled                int
)

func createMock(t *testing.T) {
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	getRecoverErrorFuncExpected = 0
	getRecoverErrorFuncCalled = 0
	getRecoverErrorFunc = func(recoverResult interface{}) apperrorModel.AppError {
		getRecoverErrorFuncCalled++
		return nil
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	responseWriteExpected = 0
	responseWriteCalled = 0
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
	}
	apperrorGetGeneralFailureErrorExpected = 0
	apperrorGetGeneralFailureErrorCalled = 0
	apperrorGetGeneralFailureError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetGeneralFailureErrorCalled++
		return nil
	}
	getDebugStackFuncExpected = 0
	getDebugStackFuncCalled = 0
	getDebugStackFunc = func() string {
		getDebugStackFuncCalled++
		return ""
	}
}

func verifyAll(t *testing.T) {
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	getRecoverErrorFunc = getRecoverError
	assert.Equal(t, getRecoverErrorFuncExpected, getRecoverErrorFuncCalled, "Unexpected number of calls to getRecoverErrorFunc")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	responseWrite = response.Write
	assert.Equal(t, responseWriteExpected, responseWriteCalled, "Unexpected number of calls to responseWrite")
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	assert.Equal(t, apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled, "Unexpected number of calls to apperrorGetGeneralFailureError")
	getDebugStackFunc = getDebugStack
	assert.Equal(t, getDebugStackFuncExpected, getDebugStackFuncCalled, "Unexpected number of calls to getDebugStackFunc")
}

// mock structs
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
