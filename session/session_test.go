package session

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestInitialize(t *testing.T) {
	// arrange
	assert.IsType(t, nil, sessionModel.NilSession)

	// mock
	createMock(t)

	// SUT + act
	Initialize()

	// assert
	assert.Nil(t, sessionModel.NilSession)
	assert.IsType(t, defaultSession, sessionModel.NilSession)

	// verify
	verifyAll(t)
}

func TestRegister(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "dummy name"
	var dummyAllowedLogType = logtype.LogType(rand.Intn(math.MaxInt8))
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Intn(math.MaxInt8))
	var dummyHTTPRequest = &http.Request{}
	var dummyResponseWriter = dummyResponseWriter{}

	// mock
	createMock(t)

	// expect
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return dummySessionID
	}
	getAllowedLogTypeFuncExpected = 1
	getAllowedLogTypeFunc = func(session *session) logtype.LogType {
		getAllowedLogTypeFuncCalled++
		assert.Equal(t, dummySessionID, session.ID)
		return dummyAllowedLogType
	}
	getAllowedLogLevelFuncExpected = 1
	getAllowedLogLevelFunc = func(session *session) loglevel.LogLevel {
		getAllowedLogLevelFuncCalled++
		assert.Equal(t, dummySessionID, session.ID)
		return dummyAllowedLogLevel
	}

	// SUT
	var result = Register(
		dummyName,
		dummyHTTPRequest,
		dummyResponseWriter,
	)

	// act
	var session, typeOK = result.(*session)

	// assert
	assert.True(t, typeOK)
	assert.Equal(t, dummySessionID, session.ID)
	assert.Equal(t, dummyName, session.Name)
	assert.Equal(t, dummyAllowedLogType, session.AllowedLogType)
	assert.Equal(t, dummyAllowedLogLevel, session.AllowedLogLevel)
	assert.Equal(t, dummyHTTPRequest, session.Request)
	assert.Equal(t, dummyResponseWriter, session.ResponseWriter)

	// verify
	verifyAll(t)
}

func TestGetID_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.GetID()

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetID_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// act
	var result = dummySessionObject.GetID()

	// assert
	assert.Equal(t, dummySessionID, result)

	// verify
	verifyAll(t)
}

func TestGetName_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.GetName()

	// assert
	assert.Equal(t, defaultName, result)

	// verify
	verifyAll(t)
}

func TestGetName_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Name: dummyName,
	}

	// act
	var result = dummySessionObject.GetName()

	// assert
	assert.Equal(t, dummyName, result)

	// verify
	verifyAll(t)
}

func TestGetRequest_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.GetRequest()

	// assert
	assert.Equal(t, defaultRequest, result)

	// verify
	verifyAll(t)
}

func TestGetRequest_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// act
	var result = dummySessionObject.GetRequest()

	// assert
	assert.Equal(t, dummyHTTPRequest, result)

	// verify
	verifyAll(t)
}

func TestGetResponseWriter_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.GetResponseWriter()

	// assert
	assert.Equal(t, defaultResponseWriter, result)

	// verify
	verifyAll(t)
}

func TestGetResponseWriter_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyResponseWriter = dummyResponseWriter{}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ResponseWriter: &dummyResponseWriter,
	}

	// act
	var result = dummySessionObject.GetResponseWriter()

	// assert
	assert.Equal(t, &dummyResponseWriter, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_EmptyBody(t *testing.T) {
	// arrange
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBody string
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// expect
	requestGetRequestBodyExpected = 1
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestBody
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "The request body is empty", format)
		assert.Equal(t, 0, len(a))
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestBody(
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_ValidBody(t *testing.T) {
	// arrange
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBody = "some request body"
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// expect
	requestGetRequestBodyExpected = 1
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestBody
	}
	loggerAPIRequestExpected = 1
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Body", category)
		assert.Zero(t, subcategory)
		assert.Equal(t, dummyRequestBody, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyRequestBody, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestBody(
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestParameter_ValueNotFound(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyParameters = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// expect
	muxVarsExpected = 1
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		assert.Equal(t, dummyHTTPRequest, r)
		return dummyParameters
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "The expected parameter [%v] is not found in request", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestParameter(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetRequestParameter_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = "some value"
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyParameters = map[string]string{
		"foo":     "bar",
		"test":    "123",
		dummyName: dummyValue,
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// expect
	muxVarsExpected = 1
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		assert.Equal(t, dummyHTTPRequest, r)
		return dummyParameters
	}
	loggerAPIRequestExpected = 1
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Parameter", category)
		assert.Equal(t, dummyName, subcategory)
		assert.Equal(t, dummyValue, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyValue, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestParameter(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAllQueries_NotFound(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyHTTPRequest = &http.Request{
		URL: &url.URL{
			RawQuery: "test=me&test=you",
		},
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getAllQueries(
		dummySessionObject,
		dummyName,
	)

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestGetAllQueries_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyHTTPRequest = &http.Request{
		URL: &url.URL{
			RawQuery: dummyName + "=me&" + dummyName + "=you",
		},
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getAllQueries(
		dummySessionObject,
		dummyName,
	)

	// assert
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "me", result[0])
	assert.Equal(t, "you", result[1])

	// verify
	verifyAll(t)
}

func TestGetRequestQuery_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries []string
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "The expected query string [%v] is not found in request", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestQuery(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestQuery_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries = []string{
		"some query string 1",
		"some query string 2",
		"some query string 3",
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	loggerAPIRequestExpected = 1
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Query", category)
		assert.Equal(t, dummyName, subcategory)
		assert.Equal(t, dummyQueries[0], messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyQueries[0], value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestQuery(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestQueries_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries []string
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	dummyFillCallbackExpected = 0
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 0, len(innerErrors))
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestQueries(
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestGetRequestQueries_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries = []string{
		"some query string 1",
		"some query string 2",
		"some query string 3",
	}
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var unmarshalErrors = []error{
		nil,
		errors.New("some error"),
		nil,
	}
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	loggerAPIRequestExpected = 3
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Query", category)
		assert.Equal(t, dummyName, subcategory)
		assert.Contains(t, dummyQueries, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 3
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyQueries[jsonutilTryUnmarshalCalled-1], value)
		*(dataTemplate.(*int)) = dummyResult
		return unmarshalErrors[jsonutilTryUnmarshalCalled-1]
	}
	dummyFillCallbackExpected = 2
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, unmarshalErrors[1], innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestQueries(
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestGetAllHeaders_NotFound(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyCanonicalName = "some conanical name"
	var dummyHTTPRequest = &http.Request{
		Header: http.Header{},
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// stub
	dummyHTTPRequest.Header.Add("test", "me")
	dummyHTTPRequest.Header.Add("test", "you")

	// mock
	createMock(t)

	// expect
	textprotoCanonicalMIMEHeaderKeyExpected = 1
	textprotoCanonicalMIMEHeaderKey = func(s string) string {
		textprotoCanonicalMIMEHeaderKeyCalled++
		assert.Equal(t, dummyName, s)
		return dummyCanonicalName
	}

	// SUT + act
	var result = getAllHeaders(
		dummySessionObject,
		dummyName,
	)

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestGetAllHeaders_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyCanonicalName = "some conanical name"
	var dummyHTTPRequest = &http.Request{
		Header: http.Header{},
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// stub
	dummyHTTPRequest.Header.Add(dummyCanonicalName, "me")
	dummyHTTPRequest.Header.Add(dummyCanonicalName, "you")

	// mock
	createMock(t)

	// expect
	textprotoCanonicalMIMEHeaderKeyExpected = 1
	textprotoCanonicalMIMEHeaderKey = func(s string) string {
		textprotoCanonicalMIMEHeaderKeyCalled++
		assert.Equal(t, dummyName, s)
		return dummyCanonicalName
	}

	// SUT + act
	var result = getAllHeaders(
		dummySessionObject,
		dummyName,
	)

	// assert
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "me", result[0])
	assert.Equal(t, "you", result[1])

	// verify
	verifyAll(t)
}

func TestGetRequestHeader_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders []string
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "The expected header string [%v] is not found in request", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestHeader(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestHeader_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders = []string{
		"some header string 1",
		"some header string 2",
		"some header string 3",
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	loggerAPIRequestExpected = 1
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Header", category)
		assert.Equal(t, dummyName, subcategory)
		assert.Contains(t, dummyHeaders, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyHeaders[0], value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestHeader(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestHeaders_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders []string
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	dummyFillCallbackExpected = 0
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 0, len(innerErrors))
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestHeaders(
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestGetRequestHeaders_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders = []string{
		"some header string 1",
		"some header string 2",
		"some header string 3",
	}
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var unmarshalErrors = []error{
		nil,
		errors.New("some error"),
		nil,
	}
	var dummyAppError = apperror.GetCustomError(0, "some app error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	loggerAPIRequestExpected = 3
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Header", category)
		assert.Equal(t, dummyName, subcategory)
		assert.Contains(t, dummyHeaders, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	jsonutilTryUnmarshalExpected = 3
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, dummyHeaders[jsonutilTryUnmarshalCalled-1], value)
		*(dataTemplate.(*int)) = dummyResult
		return unmarshalErrors[jsonutilTryUnmarshalCalled-1]
	}
	dummyFillCallbackExpected = 2
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, unmarshalErrors[1], innerErrors[0])
		return dummyAppError
	}

	// act
	var err = dummySessionObject.GetRequestHeaders(
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestAttach_NilSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestAttach_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		attachment: nil,
	}

	// act
	var result = dummySessionObject.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummySessionObject.attachment[dummyName])

	// verify
	verifyAll(t)
}

func TestAttach_WithAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// act
	var result = dummySessionObject.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummySessionObject.attachment[dummyName])

	// verify
	verifyAll(t)
}

func TestDetach_NilSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.Detach(
		dummyName,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestDetach_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		attachment: nil,
	}

	// act
	var result = dummySessionObject.Detach(
		dummyName,
	)

	// assert
	assert.True(t, result)
	var _, found = dummySessionObject.attachment[dummyName]
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestDetach_WithAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// act
	var result = dummySessionObject.Detach(
		dummyName,
	)

	// assert
	assert.True(t, result)
	var _, found = dummySessionObject.attachment[dummyName]
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestGetAttachment_NoSession(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAttachment_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{}

	// act
	var result = dummySessionObject.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAttachment_MarshalError(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return nil, errors.New("some marshal error")
	}

	// SUT
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySessionObject.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAttachment_UnmarshalError(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return json.Marshal(v)
	}
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return json.Unmarshal(data, v)
	}

	// SUT
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySessionObject.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAttachment_Success(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return json.Marshal(v)
	}
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return json.Unmarshal(data, v)
	}

	// SUT
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySessionObject.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_NoCustomization(t *testing.T) {
	// arrange
	var dummyLogType = logtype.LogType(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogTypeExpected = 1
	config.DefaultAllowedLogType = func() logtype.LogType {
		configDefaultAllowedLogTypeCalled++
		return dummyLogType
	}

	// SUT
	var dummySessionObject = &session{
		ID: uuid.New(),
	}

	// act
	var allowedLogType = getAllowedLogType(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyLogType, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_WithCustomization(t *testing.T) {
	// arrange
	var dummyLogType = logtype.LogType(rand.Intn(255))

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: uuid.New(),
	}

	// expect
	customizationSessionAllowedLogTypeExpected = 1
	customization.SessionAllowedLogType = func(session sessionModel.Session) logtype.LogType {
		customizationSessionAllowedLogTypeCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyLogType
	}

	// act
	var allowedLogType = getAllowedLogType(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyLogType, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_NoCustomization(t *testing.T) {
	// arrange
	var dummyLogLevel = loglevel.LogLevel(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogLevelExpected = 1
	config.DefaultAllowedLogLevel = func() loglevel.LogLevel {
		configDefaultAllowedLogLevelCalled++
		return dummyLogLevel
	}

	// SUT
	var dummySessionObject = &session{
		ID: uuid.New(),
	}

	// act
	var allowedLogLevel = getAllowedLogLevel(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyLogLevel, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_WithCustomization(t *testing.T) {
	// arrange
	var dummyLogLevel = loglevel.LogLevel(rand.Intn(255))

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: uuid.New(),
	}

	// expect
	customizationSessionAllowedLogLevelExpected = 1
	customization.SessionAllowedLogLevel = func(session sessionModel.Session) loglevel.LogLevel {
		customizationSessionAllowedLogLevelCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyLogLevel
	}

	// act
	var allowedLogLevel = getAllowedLogLevel(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyLogLevel, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestIsLoggingTypeMatch_NilSession(t *testing.T) {
	// arrange
	var dummySessionObject *session
	var dummyAllowedLogType = logtype.LogType(rand.Int())
	var dummyLogType = logtype.LogType(rand.Int())
	var expectedResult = dummyAllowedLogType.HasFlag(dummyLogType)

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogTypeExpected = 1
	config.DefaultAllowedLogType = func() logtype.LogType {
		configDefaultAllowedLogTypeCalled++
		return dummyAllowedLogType
	}

	// SUT + act
	var result = isLoggingTypeMatch(
		dummySessionObject,
		dummyLogType,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingTypeMatch_ValidSession(t *testing.T) {
	// arrange
	var dummyAllowedLogType = logtype.LogType(rand.Int())
	var dummySessionObject = &session{
		AllowedLogType: dummyAllowedLogType,
	}
	var dummyLogType = logtype.LogType(rand.Int())
	var expectedResult = dummyAllowedLogType.HasFlag(dummyLogType)

	// mock
	createMock(t)

	// SUT + act
	var result = isLoggingTypeMatch(
		dummySessionObject,
		dummyLogType,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingLevelMatch_NilSession(t *testing.T) {
	// arrange
	var dummySessionObject *session
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Int())
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var expectedResult = dummyAllowedLogLevel <= dummyLogLevel

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogLevelExpected = 1
	config.DefaultAllowedLogLevel = func() loglevel.LogLevel {
		configDefaultAllowedLogLevelCalled++
		return dummyAllowedLogLevel
	}

	// SUT + act
	var result = isLoggingLevelMatch(
		dummySessionObject,
		dummyLogLevel,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingLevelMatch_ValidSession(t *testing.T) {
	// arrange
	var dummyAllowedLogLevel = loglevel.LogLevel(rand.Int())
	var dummySessionObject = &session{
		AllowedLogLevel: dummyAllowedLogLevel,
	}
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var expectedResult = dummyAllowedLogLevel <= dummyLogLevel

	// mock
	createMock(t)

	// SUT + act
	var result = isLoggingLevelMatch(
		dummySessionObject,
		dummyLogLevel,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_IsLocalHost(t *testing.T) {
	// arrange
	var dummyLogType = logtype.LogType(rand.Int())
	var dummyLogLevel = loglevel.LogLevel(rand.Int())

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return true
	}

	// SUT
	var dummySessionObject = &session{}

	// act
	var result = dummySessionObject.IsLoggingAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_LoggingTypeNotMatch(t *testing.T) {
	// arrange
	var dummyLogType = logtype.LogType(rand.Int())
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var dummyLoggingTypeMatched = false
	var dummyLoggingLevelMatched = rand.Intn(100) < 50

	// SUT
	var dummySessionObject = &session{}

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	isLoggingTypeMatchFuncExpected = 1
	isLoggingTypeMatchFunc = func(session *session, logType logtype.LogType) bool {
		isLoggingTypeMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogType, logType)
		return dummyLoggingTypeMatched
	}
	isLoggingLevelMatchFuncExpected = 1
	isLoggingLevelMatchFunc = func(session *session, logLevel loglevel.LogLevel) bool {
		isLoggingLevelMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogLevel, logLevel)
		return dummyLoggingLevelMatched
	}

	// act
	var result = dummySessionObject.IsLoggingAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_LoggingLevelNotMatch_NotMethodLogic(t *testing.T) {
	// arrange
	var dummyLogType = logtype.LogType(rand.Int()) ^ logtype.MethodLogic
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var dummyLoggingTypeMatched = true
	var dummyLoggingLevelMatched = false

	// SUT
	var dummySessionObject = &session{}

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	isLoggingTypeMatchFuncExpected = 1
	isLoggingTypeMatchFunc = func(session *session, logType logtype.LogType) bool {
		isLoggingTypeMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogType, logType)
		return dummyLoggingTypeMatched
	}
	isLoggingLevelMatchFuncExpected = 1
	isLoggingLevelMatchFunc = func(session *session, logLevel loglevel.LogLevel) bool {
		isLoggingLevelMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogLevel, logLevel)
		return dummyLoggingLevelMatched
	}

	// act
	var result = dummySessionObject.IsLoggingAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_LoggingLevelNotMatch_IsMethodLogic(t *testing.T) {
	// arrange
	var dummyLogType = logtype.MethodLogic
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var dummyLoggingTypeMatched = true
	var dummyLoggingLevelMatched = false

	// SUT
	var dummySessionObject = &session{}

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	isLoggingTypeMatchFuncExpected = 1
	isLoggingTypeMatchFunc = func(session *session, logType logtype.LogType) bool {
		isLoggingTypeMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogType, logType)
		return dummyLoggingTypeMatched
	}
	isLoggingLevelMatchFuncExpected = 1
	isLoggingLevelMatchFunc = func(session *session, logLevel loglevel.LogLevel) bool {
		isLoggingLevelMatchFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogLevel, logLevel)
		return dummyLoggingLevelMatched
	}

	// act
	var result = dummySessionObject.IsLoggingAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestGetMethodName_UnknownCaller(t *testing.T) {
	// arrange
	var dummyPC = uintptr(rand.Int())
	var dummyFile = "some file"
	var dummyLine = rand.Int()
	var dummyOK = false

	// mock
	createMock(t)

	// expect
	runtimeCallerExpected = 1
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		assert.Equal(t, 3, skip)
		return dummyPC, dummyFile, dummyLine, dummyOK
	}

	// SUT + act
	var result = getMethodName()

	// assert
	assert.Equal(t, "?", result)

	// verify
	verifyAll(t)
}

func TestGetMethodName_HappyPath(t *testing.T) {
	// mock
	createMock(t)

	// expect
	runtimeCallerExpected = 1
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		assert.Equal(t, 3, skip)
		return runtime.Caller(2)
	}
	runtimeFuncForPCExpected = 1
	runtimeFuncForPC = func(pc uintptr) *runtime.Func {
		runtimeFuncForPCCalled++
		assert.NotZero(t, pc)
		return runtime.FuncForPC(pc)
	}

	// SUT + act
	var result = getMethodName()

	// assert
	assert.Contains(t, result, "TestGetMethodName_HappyPath")

	// verify
	verifyAll(t)
}

func TestLogMethodEnter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	loggerMethodEnterExpected = 1
	loggerMethodEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodEnterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Zero(t, subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// act
	dummySessionObject.LogMethodEnter()

	// verify
	verifyAll(t)
}

func TestLogMethodParameter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyParameter1 = "foo"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("test")
	var dummyParameters = []interface{}{
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	}
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	strconvItoaExpected = 3
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	loggerMethodParameterExpected = 3
	loggerMethodParameter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodParameterCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Equal(t, strconv.Itoa(loggerMethodParameterCalled-1), subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyParameters[loggerMethodParameterCalled-1], parameters[0])
	}

	// act
	dummySessionObject.LogMethodParameter(
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// verify
	verifyAll(t)
}

func TestLogMethodLogic(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogLevel = loglevel.LogLevel(rand.Int())
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("test")

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	loggerMethodLogicExpected = 1
	loggerMethodLogic = func(session sessionModel.Session, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodLogicCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogLevel, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubcategory, subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
	}

	// act
	dummySessionObject.LogMethodLogic(
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// verify
	verifyAll(t)
}

func TestLogMethodReturn(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyReturn1 = "foo"
	var dummyReturn2 = rand.Int()
	var dummyReturn3 = errors.New("test")
	var dummyReturns = []interface{}{
		dummyReturn1,
		dummyReturn2,
		dummyReturn3,
	}
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	strconvItoaExpected = 3
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	loggerMethodReturnExpected = 3
	loggerMethodReturn = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodReturnCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Equal(t, strconv.Itoa(loggerMethodReturnCalled-1), subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyReturns[loggerMethodReturnCalled-1], parameters[0])
	}

	// act
	dummySessionObject.LogMethodReturn(
		dummyReturn1,
		dummyReturn2,
		dummyReturn3,
	)

	// verify
	verifyAll(t)
}

func TestLogMethodExit(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	loggerMethodExitExpected = 1
	loggerMethodExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodExitCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Zero(t, subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// act
	dummySessionObject.LogMethodExit()

	// verify
	verifyAll(t)
}

func TestShouldSendClientCert_NoCustomization(t *testing.T) {
	// arrange
	var dummyURL = "some URL"
	var dummySendClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// expect
	certificateHasClientCertExpected = 1
	certificateHasClientCert = func() bool {
		certificateHasClientCertCalled++
		return dummySendClientCert
	}

	// SUT + act
	var result = shouldSendClientCert(
		dummyURL,
	)

	// assert
	assert.Equal(t, dummySendClientCert, result)

	// verify
	verifyAll(t)
}

func TestShouldSendClientCert_WithCustomization(t *testing.T) {
	// arrange
	var dummyURL = "some URL"
	var dummySendClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// expect
	customizationSendClientCertExpected = 1
	customization.SendClientCert = func(url string) bool {
		customizationSendClientCertCalled++
		assert.Equal(t, dummyURL, url)
		return dummySendClientCert
	}

	// SUT + act
	var result = shouldSendClientCert(
		dummyURL,
	)

	// assert
	assert.Equal(t, dummySendClientCert, result)

	// verify
	verifyAll(t)
}

func TestCreateNetworkRequest(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummyHeader = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyNetworkRequest = &dummyNetworkRequest{}

	// mock
	createMock(t)

	// SUT
	var dummySessionObject = &session{
		ID: dummySessionID,
	}

	// expect
	shouldSendClientCertFuncExpected = 1
	shouldSendClientCertFunc = func(url string) bool {
		shouldSendClientCertFuncCalled++
		assert.Equal(t, dummyURL, url)
		return dummySendClientCert
	}
	networkNewNetworkRequestExpected = 1
	networkNewNetworkRequest = func(session sessionModel.Session, method string, url string, payload string, header map[string]string, sendClientCert bool) networkModel.NetworkRequest {
		networkNewNetworkRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethod, method)
		assert.Equal(t, dummyURL, url)
		assert.Equal(t, dummyPayload, payload)
		assert.Equal(t, dummyHeader, header)
		assert.Equal(t, dummySendClientCert, sendClientCert)
		return dummyNetworkRequest
	}

	// act
	var result = dummySessionObject.CreateNetworkRequest(
		dummyMethod,
		dummyURL,
		dummyPayload,
		dummyHeader,
	)

	// assert
	assert.Equal(t, dummyNetworkRequest, result)

	// verify
	verifyAll(t)
}
