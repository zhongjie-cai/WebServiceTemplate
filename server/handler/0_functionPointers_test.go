package handler

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	routeGetRouteInfoExpected             int
	routeGetRouteInfoCalled               int
	sessionRegisterExpected               int
	sessionRegisterCalled                 int
	sessionUnregisterExpected             int
	sessionUnregisterCalled               int
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
	executeCustomizedFunctionFuncExpected int
	executeCustomizedFunctionFuncCalled   int
	customizationPreActionFuncExpected    int
	customizationPreActionFuncCalled      int
	customizationPostActionFuncExpected   int
	customizationPostActionFuncCalled     int
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
	sessionUnregisterExpected = 0
	sessionUnregisterCalled = 0
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
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
	executeCustomizedFunctionFuncExpected = 0
	executeCustomizedFunctionFuncCalled = 0
	executeCustomizedFunctionFunc = func(sessionID uuid.UUID, customFunc func(uuid.UUID) error) error {
		executeCustomizedFunctionFuncCalled++
		return nil
	}
	customizationPreActionFuncExpected = 0
	customizationPreActionFuncCalled = 0
	customization.PreActionFunc = func(sessionID uuid.UUID) error {
		customizationPreActionFuncCalled++
		return nil
	}
	customizationPostActionFuncExpected = 0
	customizationPostActionFuncCalled = 0
	customization.PostActionFunc = func(sessionID uuid.UUID) error {
		customizationPostActionFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	routeGetRouteInfo = route.GetRouteInfo
	assert.Equal(t, routeGetRouteInfoExpected, routeGetRouteInfoCalled, "Unexpected number of calls to routeGetRouteInfo")
	sessionRegister = session.Register
	assert.Equal(t, sessionRegisterExpected, sessionRegisterCalled, "Unexpected number of calls to sessionRegister")
	sessionUnregister = session.Unregister
	assert.Equal(t, sessionUnregisterExpected, sessionUnregisterCalled, "Unexpected number of calls to sessionUnregister")
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
	executeCustomizedFunctionFunc = executeCustomizedFunction
	assert.Equal(t, executeCustomizedFunctionFuncExpected, executeCustomizedFunctionFuncCalled, "Unexpected number of calls to executeCustomizedFunctionFunc")
	customization.PreActionFunc = nil
	assert.Equal(t, customizationPreActionFuncExpected, customizationPreActionFuncCalled, "Unexpected number of calls to customization.PreActionFunc")
	customization.PostActionFunc = nil
	assert.Equal(t, customizationPostActionFuncExpected, customizationPostActionFuncCalled, "Unexpected number of calls to customization.PostActionFunc")
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
	t  *testing.T
	id *uuid.UUID
}

func (session *dummySession) GetID() uuid.UUID {
	if session.id == nil {
		assert.Fail(session.t, "Unexpected call to GetID")
		return uuid.Nil
	}
	return *session.id
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
