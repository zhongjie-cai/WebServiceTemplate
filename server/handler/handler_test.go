package handler

import (
	"errors"
	"math/rand"
	"net/http"
	"testing"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
)

func TestVerifyAuthorization_NoAuthCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)

	// mock
	createMock(t)

	// expect
	customizationAuthorizationFuncExpected = 0
	customization.AuthorizationFunc = nil

	// SUT + act
	var err = verifyAuthorization(
		dummyHTTPRequest,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestVerifyAuthorization_WithAuthCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	customizationAuthorizationFuncExpected = 1
	customization.AuthorizationFunc = func(httpRequest *http.Request) error {
		customizationAuthorizationFuncCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyError
	}

	// SUT + act
	var err = verifyAuthorization(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

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
	var dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		return nil, nil
	}
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
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
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
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError error) {
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

func TestHandleInSession_AuthError(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAuthError = errors.New("some auth error")

	// mock
	createMock(t)

	// expect
	routeGetRouteInfoExpected = 1
	routeGetRouteInfo = func(httpRequest *http.Request) (string, model.ActionFunc, error) {
		routeGetRouteInfoCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyEndpoint, dummyAction, nil
	}
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
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
	verifyAuthorizationFuncExpected = 1
	verifyAuthorizationFunc = func(httpRequest *http.Request) error {
		verifyAuthorizationFuncCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAuthError
	}
	responseWriteExpected = 1
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyAuthError, responseError)
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
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyResponseObject = "some response object"
	var dummyResponseError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	routeGetRouteInfoExpected = 1
	routeGetRouteInfo = func(httpRequest *http.Request) (string, model.ActionFunc, error) {
		routeGetRouteInfoCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyEndpoint, dummyAction, nil
	}
	requestGetAllowedLogTypeExpected = 1
	requestGetAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		requestGetAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogType
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, httpRequest *http.Request, responseWriter http.ResponseWriter) uuid.UUID {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
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
	verifyAuthorizationFuncExpected = 1
	verifyAuthorizationFunc = func(httpRequest *http.Request) error {
		verifyAuthorizationFuncCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return nil
	}
	dummyActionExpected = 1
	dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseObject, dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError error) {
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
