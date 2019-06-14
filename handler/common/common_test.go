package common

import (
	"math/rand"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.WriteHeader")
}

func TestHandleInSession_DoNotRequireClientCert(t *testing.T) {
	// arrange
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummyAction func(w http.ResponseWriter, r *http.Request, sessionID uuid.UUID)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyLoginID = uuid.New()
	var dummyCorrelationID = uuid.New()
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))

	// mock
	createMock(t)

	// expect
	routeGetEndpointNameExpected = 1
	routeGetEndpointName = func(request *http.Request) string {
		routeGetEndpointNameCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyEndpoint
	}
	requestGetLoginIDExpected = 1
	requestGetLoginID = func(request *http.Request) uuid.UUID {
		requestGetLoginIDCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyLoginID
	}
	requestGetCorrelationIDExpected = 1
	requestGetCorrelationID = func(request *http.Request) uuid.UUID {
		requestGetCorrelationIDCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyCorrelationID
	}
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(request *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyRequest, request)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, loginID uuid.UUID, correlationID uuid.UUID, allowedLogType logtype.LogType, request *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyLoginID, loginID)
		assert.Equal(t, dummyCorrelationID, correlationID)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyRequest, request)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionID
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyRequest.Method, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	dummyActionExpected = 1
	dummyAction = func(w http.ResponseWriter, r *http.Request, sessionID uuid.UUID) {
		dummyActionCalled++
		assert.Equal(t, dummyResponseWriter, w)
		assert.Equal(t, dummyRequest, r)
		assert.Equal(t, dummySessionID, sessionID)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyRequest.Method, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	panicHandleExpected = 1
	panicHandle = func(endpointName string, sessionID uuid.UUID, recoverResult interface{}, w http.ResponseWriter) {
		panicHandleCalled++
		assert.Equal(t, dummyEndpoint, endpointName)
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, recover(), recoverResult)
		assert.Equal(t, dummyResponseWriter, w)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(sessionID uuid.UUID) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionID, sessionID)
	}

	// SUT + act
	HandleInSession(
		dummyResponseWriter,
		dummyRequest,
		dummyAction,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyActionExpected, dummyActionCalled, "Unexpected method call to dummyAction")
}
