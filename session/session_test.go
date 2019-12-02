package session

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestInit_AllValuesSet(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	Init()
	var result, found = sessionCache.Get(uuid.Nil.String())

	// assert
	assert.True(t, found)
	assert.Equal(t, defaultSession, result)
	assert.Zero(t, defaultSession.ID)
	assert.Equal(t, defaultName, defaultSession.Name)
	assert.Equal(t, logtype.BasicLogging, defaultSession.AllowedLogType)
	assert.Equal(t, loglevel.Debug, defaultSession.AllowedLogLevel)

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

	// stub
	sessionCache.Delete(dummySessionID.String())

	// mock
	createMock(t)

	// expect
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return dummySessionID
	}

	// SUT
	var result = Register(
		dummyName,
		dummyAllowedLogType,
		dummyAllowedLogLevel,
		dummyHTTPRequest,
		dummyResponseWriter,
	)

	// act
	var cacheItem, cacheOK = sessionCache.Get(dummySessionID.String())
	var session, typeOK = cacheItem.(session)

	// assert
	assert.Equal(t, dummySessionID, result)
	assert.True(t, cacheOK)
	assert.True(t, typeOK)
	assert.Equal(t, dummySessionID, session.ID)
	assert.Equal(t, dummyName, session.Name)
	assert.Equal(t, dummyAllowedLogType, session.AllowedLogType)
	assert.Equal(t, dummyAllowedLogLevel, session.AllowedLogLevel)
	assert.Equal(t, dummyHTTPRequest, session.Request)
	assert.Equal(t, dummyResponseWriter, session.ResponseWriter)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
}

func TestUnregister(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// stub
	sessionCache.Set(dummySessionID.String(), 123, cache.NoExpiration)

	// mock
	createMock(t)

	// SUT
	Unregister(dummySessionID)

	// act
	var _, cacheOK = sessionCache.Get(dummySessionID.String())

	// assert
	assert.False(t, cacheOK)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
}

func TestGet_CacheNotLoaded(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// SUT + act
	var session = Get(dummySessionID)

	// assert
	assert.Equal(t, &defaultSession, session)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
}

func TestGet_CacheItemInvalid(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// stub
	sessionCache.SetDefault(dummySessionID.String(), 123)

	// SUT + act
	var session = Get(dummySessionID)

	// assert
	assert.Equal(t, &defaultSession, session)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
}

func TestGet_CacheItemValid(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var expectedSession = &session{
		Name:           "dummy name",
		AllowedLogType: logtype.BasicLogging,
	}

	// stub
	sessionCache.SetDefault(dummySessionID.String(), *expectedSession)

	// mock
	createMock(t)

	// SUT + act
	var session = Get(dummySessionID)

	// assert
	assert.Equal(t, expectedSession, session)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
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
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetName(
		dummySessionID,
	)

	// assert
	assert.Equal(t, defaultName, result)

	// verify
	verifyAll(t)
}

func TestGetName_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummySessionObject = &session{
		Name: dummyName,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetName(
		dummySessionID,
	)

	// assert
	assert.Equal(t, dummyName, result)

	// verify
	verifyAll(t)
}

func TestGetRequest_NilSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetRequest(
		dummySessionID,
	)

	// assert
	assert.Equal(t, defaultRequest, result)

	// verify
	verifyAll(t)
}

func TestGetRequest_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetRequest(
		dummySessionID,
	)

	// assert
	assert.Equal(t, dummyHTTPRequest, result)

	// verify
	verifyAll(t)
}

func TestGetResponseWriter_NilSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetResponseWriter(
		dummySessionID,
	)

	// assert
	assert.Equal(t, defaultResponseWriter, result)

	// verify
	verifyAll(t)
}

func TestGetResponseWriter_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseWriter = dummyResponseWriter{}
	var dummySessionObject = &session{
		ResponseWriter: &dummyResponseWriter,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetResponseWriter(
		dummySessionID,
	)

	// assert
	assert.Equal(t, &dummyResponseWriter, result)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_NoQuoteJSONEmpty(t *testing.T) {
	// arrange
	var dummyValue string
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValue), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Primitive(t *testing.T) {
	// arrange
	var dummyValue = rand.Int()
	var dummyValueString = strconv.Itoa(dummyValue)
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValueString), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Struct(t *testing.T) {
	// arrange
	var dummyValueString = "{\"foo\":\"bar\",\"test\":123}"
	var dummyDataTemplate struct {
		Foo  string `json:"foo"`
		Test int    `json:"test"`
	}

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValueString), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "bar", dummyDataTemplate.Foo)
	assert.Equal(t, 123, dummyDataTemplate.Test)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_WithQuoteJSONEmpty(t *testing.T) {
	// arrange
	var dummyValueString string
	var dummyDataTemplate struct {
		Foo  string `json:"foo"`
		Test int    `json:"test"`
	}

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValueString), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Zero(t, dummyDataTemplate)
	assert.Zero(t, dummyDataTemplate.Foo)
	assert.Zero(t, dummyDataTemplate.Test)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_WithQuoteJSONSuccess(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate string

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 2
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		if jsonUnmarshalCalled == 1 {
			assert.Equal(t, []byte(dummyValue), data)
		} else if jsonUnmarshalCalled == 2 {
			assert.Equal(t, []byte("\""+dummyValue+"\""), data)
		}
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_Failure(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate uuid.UUID
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 2
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		if jsonUnmarshalCalled == 1 {
			assert.Equal(t, []byte(dummyValue), data)
		} else if jsonUnmarshalCalled == 2 {
			assert.Equal(t, []byte("\""+dummyValue+"\""), data)
		}
		return json.Unmarshal(data, v)
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "Unable to unmarshal value [%v] into data template", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyValue, a[0])
		return dummyError
	}
	apperrorGetBadRequestErrorExpected = 1
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_EmptyBody(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBody string
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var err = GetRequestBody(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBody = "some request body"
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	requestGetRequestBodyExpected = 1
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestBody
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyRequestBody, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestBody(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyParameters = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var err = GetRequestParameter(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = "some value"
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyParameters = map[string]string{
		"foo":     "bar",
		"test":    "123",
		dummyName: dummyValue,
	}
	var dummySessionObject = &session{
		Request: dummyHTTPRequest,
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	muxVarsExpected = 1
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		assert.Equal(t, dummyHTTPRequest, r)
		return dummyParameters
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyValue, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestParameter(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries []string
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var err = GetRequestQuery(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries = []string{
		"some query string 1",
		"some query string 2",
		"some query string 3",
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyQueries[0], value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestQuery(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueries []string
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var dummyError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 0, len(innerErrors))
		assert.Equal(t, "Failed to get request query strings", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyError
	}

	// SUT + act
	var err = GetRequestQueries(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestGetRequestQueries_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
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
	var unmarshalErrors = []apperrorModel.AppError{
		nil,
		apperror.GetCustomError(0, ""),
		nil,
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	getAllQueriesFuncExpected = 1
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyQueries
	}
	tryUnmarshalFuncExpected = 3
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyQueries[tryUnmarshalFuncCalled-1], value)
		*(dataTemplate.(*int)) = dummyResult
		return unmarshalErrors[tryUnmarshalFuncCalled-1]
	}
	dummyFillCallbackExpected = 2
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, unmarshalErrors[1], innerErrors[0])
		assert.Equal(t, "Failed to get request query strings", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestQueries(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders []string
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var err = GetRequestHeader(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders = []string{
		"some header string 1",
		"some header string 2",
		"some header string 3",
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyHeaders[0], value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestHeader(
		dummySessionID,
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
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyHeaders []string
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var dummyError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 0, len(innerErrors))
		assert.Equal(t, "Failed to get request header strings", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyError
	}

	// SUT + act
	var err = GetRequestHeaders(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}

func TestGetRequestHeaders_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject = &session{
		ID: dummySessionID,
	}
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
	var unmarshalErrors = []apperrorModel.AppError{
		nil,
		apperror.GetCustomError(0, ""),
		nil,
	}
	var dummyAppError = apperror.GetCustomError(0, "")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	getAllHeadersFuncExpected = 1
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyName, name)
		return dummyHeaders
	}
	tryUnmarshalFuncExpected = 3
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyHeaders[tryUnmarshalFuncCalled-1], value)
		*(dataTemplate.(*int)) = dummyResult
		return unmarshalErrors[tryUnmarshalFuncCalled-1]
	}
	dummyFillCallbackExpected = 2
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, unmarshalErrors[1], innerErrors[0])
		assert.Equal(t, "Failed to get request header strings", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = GetRequestHeaders(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Attach(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}
	var dummySessionObject = &session{
		attachment: nil,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Attach(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Attach(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Detach(
		dummySessionID,
		dummyName,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestDetach_NoAttachment(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummySessionObject = &session{}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Detach(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = Detach(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment
	var dummySessionObject *session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetAttachment(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummySessionObject = &session{}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	var result = GetAttachment(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return nil, errors.New("some marshal error")
	}

	// SUT + act
	var result = GetAttachment(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var result = GetAttachment(
		dummySessionID,
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
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummySessionObject = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
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

	// SUT + act
	var result = GetAttachment(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_NilSession(t *testing.T) {
	// arrange
	var dummyLogType = logtype.APIEnter
	var dummyLogLevel = loglevel.Warn

	// mock
	createMock(t)

	// SUT
	var dummySessionObject *session

	// act
	var result = dummySessionObject.IsLogAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_IsLocalHost(t *testing.T) {
	// arrange
	var dummyLogType = logtype.APIEnter
	var dummyLogLevel = loglevel.Warn

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
	var result = dummySessionObject.IsLogAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_FlagNotMatch(t *testing.T) {
	// arrange
	var dummyAllowedLogType = logtype.GeneralTracing
	var dummyAllowedLogLevel = loglevel.Warn
	var dummyLogType = logtype.MethodEnter
	var dummyLogLevel = loglevel.Warn

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}

	// SUT
	var dummySessionObject = &session{
		AllowedLogType:  dummyAllowedLogType,
		AllowedLogLevel: dummyAllowedLogLevel,
	}

	// act
	var result = dummySessionObject.IsLogAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_FlagMatch_LevelNotMatch(t *testing.T) {
	// arrange
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyAllowedLogLevel = loglevel.Warn
	var dummyLogType = logtype.MethodLogic
	var dummyLogLevel = loglevel.Info

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}

	// SUT
	var dummySessionObject = &session{
		AllowedLogType:  dummyAllowedLogType,
		AllowedLogLevel: dummyAllowedLogLevel,
	}

	// act
	var result = dummySessionObject.IsLogAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsLoggingAllowed_FlagMatch_LevelMatch(t *testing.T) {
	// arrange
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyAllowedLogLevel = loglevel.Warn
	var dummyLogType = logtype.MethodLogic
	var dummyLogLevel = loglevel.Warn

	// mock
	createMock(t)

	// expect
	configIsLocalhostExpected = 1
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}

	// SUT
	var dummySessionObject = &session{
		AllowedLogType:  dummyAllowedLogType,
		AllowedLogLevel: dummyAllowedLogLevel,
	}

	// act
	var result = dummySessionObject.IsLogAllowed(
		dummyLogType,
		dummyLogLevel,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}
