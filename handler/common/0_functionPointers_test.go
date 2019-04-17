package common

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
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
	sessionRegister = session.Register
	if sessionRegisterExpected != sessionRegisterCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to sessionRegister, expected %v, actual %v", sessionRegisterExpected, sessionRegisterCalled))
	}
	sessionUnregister = session.Unregister
	if sessionUnregisterExpected != sessionUnregisterCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to sessionUnregister, expected %v, actual %v", sessionUnregisterExpected, sessionUnregisterCalled))
	}
	panicHandle = panic.Handle
	if panicHandleExpected != panicHandleCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to panicHandle, expected %v, actual %v", panicHandleExpected, panicHandleCalled))
	}
	requestGetLoginID = request.GetLoginID
	if requestGetLoginIDExpected != requestGetLoginIDCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to requestGetLoginID, expected %v, actual %v", requestGetLoginIDExpected, requestGetLoginIDCalled))
	}
	requestGetCorrelationID = request.GetCorrelationID
	if requestGetCorrelationIDExpected != requestGetCorrelationIDCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to requestGetCorrelationID, expected %v, actual %v", requestGetCorrelationIDExpected, requestGetCorrelationIDCalled))
	}
	requestGetAllowedLogType = request.GetAllowedLogType
	if requestGetAllowedLogTypeExpected != requestGetAllowedLogTypeCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to requestGetAllowedLogType, expected %v, actual %v", requestGetAllowedLogTypeExpected, requestGetAllowedLogTypeCalled))
	}
	responseError = response.Error
	if responseErrorExpected != responseErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to responseError, expected %v, actual %v", responseErrorExpected, responseErrorCalled))
	}
	loggerAPIEnter = logger.APIEnter
	if loggerAPIEnterExpected != loggerAPIEnterCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAPIEnter, expected %v, actual %v", loggerAPIEnterExpected, loggerAPIEnterCalled))
	}
	loggerAPIExit = logger.APIExit
	if loggerAPIExitExpected != loggerAPIExitCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAPIExit, expected %v, actual %v", loggerAPIExitExpected, loggerAPIExitCalled))
	}
}
