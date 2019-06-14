package common

import (
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
	routeGetEndpointNameExpected     int
	routeGetEndpointNameCalled       int
	sessionRegisterExpected          int
	sessionRegisterCalled            int
	sessionUnregisterExpected        int
	sessionUnregisterCalled          int
	panicHandleExpected              int
	panicHandleCalled                int
	requestGetLoginIDExpected        int
	requestGetLoginIDCalled          int
	requestGetCorrelationIDExpected  int
	requestGetCorrelationIDCalled    int
	requestGetAllowedLogTypeExpected int
	requestGetAllowedLogTypeCalled   int
	responseErrorExpected            int
	responseErrorCalled              int
	loggerAPIEnterExpected           int
	loggerAPIEnterCalled             int
	loggerAPIExitExpected            int
	loggerAPIExitCalled              int
)

func createMock(t *testing.T) {
	routeGetEndpointNameExpected = 0
	routeGetEndpointNameCalled = 0
	routeGetEndpointName = func(request *http.Request) string {
		routeGetEndpointNameCalled++
		return ""
	}
	sessionRegisterExpected = 0
	sessionRegisterCalled = 0
	sessionRegister = func(endpoint string, loginID uuid.UUID, correlationID uuid.UUID, allowedLogType logtype.LogType, request *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
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
	requestGetLoginIDExpected = 0
	requestGetLoginIDCalled = 0
	requestGetLoginID = func(request *http.Request) uuid.UUID {
		requestGetLoginIDCalled++
		return uuid.Nil
	}
	requestGetCorrelationIDExpected = 0
	requestGetCorrelationIDCalled = 0
	requestGetCorrelationID = func(request *http.Request) uuid.UUID {
		requestGetCorrelationIDCalled++
		return uuid.Nil
	}
	requestGetAllowedLogTypeExpected = 0
	requestGetAllowedLogTypeCalled = 0
	requestGetAllowedLogType = func(request *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		return 0
	}
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error, responseWriter http.ResponseWriter) {
		responseErrorCalled++
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
}

func verifyAll(t *testing.T) {
	routeGetEndpointName = route.GetEndpointName
	assert.Equal(t, routeGetEndpointNameExpected, routeGetEndpointNameCalled, "Unexpected method call to routeGetEndpointName")
	sessionRegister = session.Register
	assert.Equal(t, sessionRegisterExpected, sessionRegisterCalled, "Unexpected method call to sessionRegister")
	sessionUnregister = session.Unregister
	assert.Equal(t, sessionUnregisterExpected, sessionUnregisterCalled, "Unexpected method call to sessionUnregister")
	panicHandle = panic.Handle
	assert.Equal(t, panicHandleExpected, panicHandleCalled, "Unexpected method call to panicHandle")
	requestGetLoginID = request.GetLoginID
	assert.Equal(t, requestGetLoginIDExpected, requestGetLoginIDCalled, "Unexpected method call to requestGetLoginID")
	requestGetCorrelationID = request.GetCorrelationID
	assert.Equal(t, requestGetCorrelationIDExpected, requestGetCorrelationIDCalled, "Unexpected method call to requestGetCorrelationID")
	requestGetAllowedLogType = request.GetAllowedLogType
	assert.Equal(t, requestGetAllowedLogTypeExpected, requestGetAllowedLogTypeCalled, "Unexpected method call to requestGetAllowedLogType")
	responseError = response.Error
	assert.Equal(t, responseErrorExpected, responseErrorCalled, "Unexpected method call to responseError")
	loggerAPIEnter = logger.APIEnter
	assert.Equal(t, loggerAPIEnterExpected, loggerAPIEnterCalled, "Unexpected method call to loggerAPIEnter")
	loggerAPIExit = logger.APIExit
	assert.Equal(t, loggerAPIExitExpected, loggerAPIExitCalled, "Unexpected method call to loggerAPIExit")
}
