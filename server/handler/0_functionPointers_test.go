package handler

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/panic"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
	muxVarsExpected                  int
	muxVarsCalled                    int
	routeGetRouteInfoExpected        int
	routeGetRouteInfoCalled          int
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
	requestGetRequestBodyExpected    int
	requestGetRequestBodyCalled      int
	responseErrorExpected            int
	responseErrorCalled              int
	loggerAPIEnterExpected           int
	loggerAPIEnterCalled             int
	loggerAPIExitExpected            int
	loggerAPIExitCalled              int
)

func createMock(t *testing.T) {
	muxVarsExpected = 0
	muxVarsCalled = 0
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		return nil
	}
	routeGetRouteInfoExpected = 0
	routeGetRouteInfoCalled = 0
	routeGetRouteInfo = func(httpRequest *http.Request) (string, func(uuid.UUID, string, map[string]string), error) {
		routeGetRouteInfoCalled++
		return "", nil, nil
	}
	sessionRegisterExpected = 0
	sessionRegisterCalled = 0
	sessionRegister = func(endpoint string, loginID uuid.UUID, correlationID uuid.UUID, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
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
	requestGetLoginID = func(httpRequest *http.Request) uuid.UUID {
		requestGetLoginIDCalled++
		return uuid.Nil
	}
	requestGetCorrelationIDExpected = 0
	requestGetCorrelationIDCalled = 0
	requestGetCorrelationID = func(httpRequest *http.Request) uuid.UUID {
		requestGetCorrelationIDCalled++
		return uuid.Nil
	}
	requestGetAllowedLogTypeExpected = 0
	requestGetAllowedLogTypeCalled = 0
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		return 0
	}
	requestGetRequestBodyExpected = 0
	requestGetRequestBodyCalled = 0
	requestGetRequestBody = func(sessionID uuid.UUID, httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		return ""
	}
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error) {
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
	muxVars = mux.Vars
	assert.Equal(t, muxVarsExpected, muxVarsCalled, "Unexpected number of calls to muxVars")
	routeGetRouteInfo = route.GetRouteInfo
	assert.Equal(t, routeGetRouteInfoExpected, routeGetRouteInfoCalled, "Unexpected number of calls to routeGetRouteInfo")
	sessionRegister = session.Register
	assert.Equal(t, sessionRegisterExpected, sessionRegisterCalled, "Unexpected number of calls to sessionRegister")
	sessionUnregister = session.Unregister
	assert.Equal(t, sessionUnregisterExpected, sessionUnregisterCalled, "Unexpected number of calls to sessionUnregister")
	panicHandle = panic.Handle
	assert.Equal(t, panicHandleExpected, panicHandleCalled, "Unexpected number of calls to panicHandle")
	requestGetLoginID = request.GetLoginID
	assert.Equal(t, requestGetLoginIDExpected, requestGetLoginIDCalled, "Unexpected number of calls to requestGetLoginID")
	requestGetCorrelationID = request.GetCorrelationID
	assert.Equal(t, requestGetCorrelationIDExpected, requestGetCorrelationIDCalled, "Unexpected number of calls to requestGetCorrelationID")
	requestGetAllowedLogType = request.GetAllowedLogType
	assert.Equal(t, requestGetAllowedLogTypeExpected, requestGetAllowedLogTypeCalled, "Unexpected number of calls to requestGetAllowedLogType")
	requestGetRequestBody = request.GetRequestBody
	assert.Equal(t, requestGetRequestBodyExpected, requestGetRequestBodyCalled, "Unexpected number of calls to requestGetRequestBody")
	responseError = response.Error
	assert.Equal(t, responseErrorExpected, responseErrorCalled, "Unexpected number of calls to responseError")
	loggerAPIEnter = logger.APIEnter
	assert.Equal(t, loggerAPIEnterExpected, loggerAPIEnterCalled, "Unexpected number of calls to loggerAPIEnter")
	loggerAPIExit = logger.APIExit
	assert.Equal(t, loggerAPIExitExpected, loggerAPIExitCalled, "Unexpected number of calls to loggerAPIExit")
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
