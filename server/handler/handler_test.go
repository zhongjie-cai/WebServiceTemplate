package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestExecuteCustomizedFunction_NoCustomization(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// expect
	var dummyCustomFuncExpected = 0
	var dummyCustomFuncCalled = 0
	var dummyCustomFunc func(sessionID uuid.UUID) error

	// SUT + act
	var err = executeCustomizedFunction(
		dummySessionID,
		dummyCustomFunc,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomFuncExpected, dummyCustomFuncCalled, "Unexpected number of calls to dummyCustomFunc")
}

func TestExecuteCustomizedFunction_WithCustomization(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	var dummyCustomFuncExpected = 1
	var dummyCustomFuncCalled = 0
	var dummyCustomFunc = func(sessionID uuid.UUID) error {
		dummyCustomFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyError
	}

	// SUT + act
	var err = executeCustomizedFunction(
		dummySessionID,
		dummyCustomFunc,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomFuncExpected, dummyCustomFuncCalled, "Unexpected number of calls to dummyCustomFunc")
}

func TestHandleInSession_RouteError(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummySessionObject = &dummySession{
		t,
		&dummySessionID,
	}
	var dummyActionExpected = 0
	var dummyActionCalled = 0
	var dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		return nil, nil
	}
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(256))
	var dummyRouteError = errors.New("some route error")
	var dummyResponseError = apperror.GetCustomError(0, "some app error")

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
	requestGetAllowedLogLevelExpected = 1
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogLevel
	}
	sessionRegisterExpected = 1
	sessionRegister = func(name string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, name)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyAllowedLogLevel, allowedLogLevel)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionObject
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	apperrorGetInvalidOperationExpected = 1
	apperrorGetInvalidOperation = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetInvalidOperationCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyRouteError, innerErrors[0])
		return dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	panicHandleExpected = 1
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, recover(), recoverResult)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionObject, session)
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

func TestHandleInSession_PreActionError(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummySessionObject = &dummySession{
		t,
		&dummySessionID,
	}
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(256))
	var dummyPreActionError = errors.New("some pre-action error")

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
	requestGetAllowedLogLevelExpected = 1
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogLevel
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyAllowedLogLevel, allowedLogLevel)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionObject
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	executeCustomizedFunctionFuncExpected = 1
	executeCustomizedFunctionFunc = func(sessionID uuid.UUID, customFunc func(uuid.UUID) error) error {
		executeCustomizedFunctionFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PreActionFunc))
		var pointerActual = fmt.Sprintf("%v", reflect.ValueOf(customFunc))
		assert.Equal(t, pointerExpect, pointerActual)
		return dummyPreActionError
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyPreActionError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	panicHandleExpected = 1
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, recover(), recoverResult)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionObject, session)
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

func TestHandleInSession_PostActionError_WithResponseError(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummySessionObject = &dummySession{
		t,
		&dummySessionID,
	}
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(256))
	var dummyResponseObject = "some response object"
	var dummyResponseError = apperror.GetCustomError(0, "some app error")
	var dummyPostActionError = errors.New("some post-action error")

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
	requestGetAllowedLogLevelExpected = 1
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogLevel
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyAllowedLogLevel, allowedLogLevel)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionObject
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	executeCustomizedFunctionFuncExpected = 2
	executeCustomizedFunctionFunc = func(sessionID uuid.UUID, customFunc func(uuid.UUID) error) error {
		executeCustomizedFunctionFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		var pointerActual = fmt.Sprintf("%v", reflect.ValueOf(customFunc))
		if executeCustomizedFunctionFuncCalled == 1 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PreActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
			return nil
		} else if executeCustomizedFunctionFuncCalled == 2 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PostActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
			return dummyPostActionError
		}
		return nil
	}
	dummyActionExpected = 1
	dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseObject, dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
	}
	loggerAPIExitExpected = 2
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		if loggerAPIExitCalled == 1 {
			assert.Equal(t, "Post-action error: %v", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyPostActionError, parameters[0])
		} else if loggerAPIExitCalled == 2 {
			assert.Zero(t, messageFormat)
			assert.Equal(t, 0, len(parameters))
		}
	}
	panicHandleExpected = 1
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, recover(), recoverResult)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionObject, session)
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

func TestHandleInSession_PostActionError_NoResponseError(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummySessionObject = &dummySession{
		t,
		&dummySessionID,
	}
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(256))
	var dummyResponseObject = "some response object"
	var dummyPostActionError = errors.New("some post-action error")

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
	requestGetAllowedLogLevelExpected = 1
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogLevel
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyAllowedLogLevel, allowedLogLevel)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionObject
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	executeCustomizedFunctionFuncExpected = 2
	executeCustomizedFunctionFunc = func(sessionID uuid.UUID, customFunc func(uuid.UUID) error) error {
		executeCustomizedFunctionFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		var pointerActual = fmt.Sprintf("%v", reflect.ValueOf(customFunc))
		if executeCustomizedFunctionFuncCalled == 1 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PreActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
			return nil
		} else if executeCustomizedFunctionFuncCalled == 2 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PostActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
			return dummyPostActionError
		}
		return nil
	}
	dummyActionExpected = 1
	dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseObject, nil
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyPostActionError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	panicHandleExpected = 1
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, recover(), recoverResult)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionObject, session)
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
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyEndpoint = "some endpoint"
	var dummySessionID = uuid.New()
	var dummySessionObject = &dummySession{
		t,
		&dummySessionID,
	}
	var dummyAction func(uuid.UUID) (interface{}, error)
	var dummyActionExpected int
	var dummyActionCalled int
	var dummyAllowedLogType = logtype.LogType(rand.Intn(256))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(256))
	var dummyResponseObject = "some response object"
	var dummyResponseError = apperror.GetCustomError(0, "some app error")

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
	requestGetAllowedLogLevelExpected = 1
	requestGetAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		requestGetAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyAllowedLogLevel
	}
	sessionRegisterExpected = 1
	sessionRegister = func(endpoint string, allowedLogType logtype.LogType, allowedLogLevel loglevel.LogLevel, httpRequest *http.Request, responseWriter http.ResponseWriter) sessionModel.Session {
		sessionRegisterCalled++
		assert.Equal(t, dummyEndpoint, endpoint)
		assert.Equal(t, dummyAllowedLogType, allowedLogType)
		assert.Equal(t, dummyAllowedLogLevel, allowedLogLevel)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		return dummySessionObject
	}
	loggerAPIEnterExpected = 1
	loggerAPIEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	executeCustomizedFunctionFuncExpected = 2
	executeCustomizedFunctionFunc = func(sessionID uuid.UUID, customFunc func(uuid.UUID) error) error {
		executeCustomizedFunctionFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		var pointerActual = fmt.Sprintf("%v", reflect.ValueOf(customFunc))
		if executeCustomizedFunctionFuncCalled == 1 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PreActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
		} else if executeCustomizedFunctionFuncCalled == 2 {
			var pointerExpect = fmt.Sprintf("%v", reflect.ValueOf(customization.PostActionFunc))
			assert.Equal(t, pointerExpect, pointerActual)
		}
		return nil
	}
	dummyActionExpected = 1
	dummyAction = func(sessionID uuid.UUID) (interface{}, error) {
		dummyActionCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseObject, dummyResponseError
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyResponseObject, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHTTPRequest.Method, category)
		assert.Equal(t, dummyEndpoint, subcategory)
		assert.Zero(t, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	panicHandleExpected = 1
	panicHandle = func(session sessionModel.Session, recoverResult interface{}) {
		panicHandleCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, recover(), recoverResult)
	}
	sessionUnregisterExpected = 1
	sessionUnregister = func(session sessionModel.Session) {
		sessionUnregisterCalled++
		assert.Equal(t, dummySessionObject, session)
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
