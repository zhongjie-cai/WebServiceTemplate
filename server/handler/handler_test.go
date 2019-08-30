package handler

import (
	"errors"
	"math/rand"
	"net/http"
	"testing"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

func TestHandleInSession_RouteError(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummyActionExpected = 0
	var dummyActionCalled = 0
	var dummyAction = func(sessionID uuid.UUID, requestBody string, parameters map[string]string, queries map[string][]string) (interface{}, apperror.AppError) {
		dummyActionCalled++
		return nil, nil
	}
	var dummyLoginID = uuid.New()
	var dummyCorrelationID = uuid.New()
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyRouteError = errors.New("some route error")
	var dummyResponseError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	routeGetRouteInfoExpected = 1
	routeGetRouteInfo = func(httpRequest *http.Request) (string, model.ActionFunc, error) {
		routeGetRouteInfoCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyEndpoint, dummyAction, dummyRouteError
	}
	requestGetLoginIDExpected = 1
	requestGetLoginID = func(httpRequest *http.Request) uuid.UUID {
		requestGetLoginIDCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyLoginID
	}
	requestGetCorrelationIDExpected = 1
	requestGetCorrelationID = func(httpRequest *http.Request) uuid.UUID {
		requestGetCorrelationIDCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyCorrelationID
	}
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, loginID uuid.UUID, correlationID uuid.UUID, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyLoginID, loginID)
		assert.Equal(t, dummyCorrelationID, correlationID)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionID
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyHTTPRequest.Method, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	apperrorGetInvalidOperationExpected = 1
	apperrorGetInvalidOperation = func(innerError error) apperror.AppError {
		apperrorGetInvalidOperationCalled++
		assert.Equal(t, dummyRouteError, innerError)
		return dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError apperror.AppError) {
		responseWriteCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyHTTPRequest.Method, messageFormat)
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
	Session(
		dummyResponseWriter,
		dummyHTTPRequest,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyActionExpected, dummyActionCalled, "Unexpected number of calls to dummyAction")
}

func TestHandleInSession_Success(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummyAction func(uuid.UUID, string, map[string]string, map[string][]string) (interface{}, apperror.AppError)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyLoginID = uuid.New()
	var dummyCorrelationID = uuid.New()
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyRequestBody = "some request body"
	var dummyParameters = map[string]string{"foo": "bar"}
	var dummyQueries = map[string][]string{"test": []string{"me", "you"}}
	var dummyResponseObject = "some response object"
	var dummyResponseError = apperror.GetGeneralFailureError(nil)

	// stub
	dummyHTTPRequest.Form = dummyQueries

	// mock
	createMock(t)

	// expect
	routeGetRouteInfoExpected = 1
	routeGetRouteInfo = func(httpRequest *http.Request) (string, model.ActionFunc, error) {
		routeGetRouteInfoCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyEndpoint, dummyAction, nil
	}
	requestGetLoginIDExpected = 1
	requestGetLoginID = func(httpRequest *http.Request) uuid.UUID {
		requestGetLoginIDCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyLoginID
	}
	requestGetCorrelationIDExpected = 1
	requestGetCorrelationID = func(httpRequest *http.Request) uuid.UUID {
		requestGetCorrelationIDCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyCorrelationID
	}
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, loginID uuid.UUID, correlationID uuid.UUID, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyLoginID, loginID)
		assert.Equal(t, dummyCorrelationID, correlationID)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionID
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyHTTPRequest.Method, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	requestGetRequestBodyExpected = 1
	requestGetRequestBody = func(sessionID uuid.UUID, httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestBody
	}
	muxVarsExpected = 1
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		assert.Equal(t, dummyHTTPRequest, r)
		return dummyParameters
	}
	dummyActionExpected = 1
	dummyAction = func(sessionID uuid.UUID, requestBody string, parameters map[string]string, queries map[string][]string) (interface{}, apperror.AppError) {
		dummyActionCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyRequestBody, requestBody)
		assert.Equal(t, dummyParameters, parameters)
		assert.Equal(t, dummyQueries, queries)
		return dummyResponseObject, dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError apperror.AppError) {
		responseWriteCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyResponseObject, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "handler", category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Equal(t, dummyHTTPRequest.Method, messageFormat)
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
	Session(
		dummyResponseWriter,
		dummyHTTPRequest,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyActionExpected, dummyActionCalled, "Unexpected number of calls to dummyAction")
}
