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
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
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
	requestGetAllowedLogTypeExpected      int
	requestGetAllowedLogTypeCalled        int
	requestGetAllowedLogLevelExpected     int
	requestGetAllowedLogLevelCalled       int
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
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		return uuid.Nil
	}
	sessionUnregisterExpected = 0
	sessionUnregisterCalled = 0
	sessionUnregister = func(sessionID uuid.UUID) {
		sessionUnregisterCalled++
	}
	panicHandleExpected = 0
	panicHandleCalled = 0
	panicHandle = func(endpointName string, sessionID uuid.UUID, recoverResult interface{}, responseWriter http.ResponseWriter) {
		panicHandleCalled++
	}
	requestGetAllowedLogTypeExpected = 0
	requestGetAllowedLogTypeCalled = 0
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		return 0
	}
	requestGetAllowedLogLevelExpected = 0
	requestGetAllowedLogLevelCalled = 0
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		return 0
	}
	responseWriteExpected = 0
	responseWriteCalled = 0
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError error) {
		responseWriteCalled++
	}
	loggerAPIEnterExpected = 0
	loggerAPIEnterCalled = 0
	loggerAPIEnter = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
	}
	loggerAPIExitExpected = 0
	loggerAPIExitCalled = 0
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
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
	requestGetAllowedLogType = request.GetAllowedLogType
	assert.Equal(t, requestGetAllowedLogTypeExpected, requestGetAllowedLogTypeCalled, "Unexpected number of calls to requestGetAllowedLogType")
	requestGetAllowedLogLevel = request.GetAllowedLogLevel
	assert.Equal(t, requestGetAllowedLogLevelExpected, requestGetAllowedLogLevelCalled, "Unexpected number of calls to requestGetAllowedLogLevel")
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
