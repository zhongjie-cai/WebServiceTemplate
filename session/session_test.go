package session

import (
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/patrickmn/go-cache"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNilResponseWriter(t *testing.T) {
	// arrange
	var dummyBody = []byte("some body")
	var dummyStatus = rand.Int()

	// mock
	createMock(t)

	// SUT
	var nilResponseWriter = &nilResponseWriter{}

	// act
	var header = nilResponseWriter.Header()
	var result, err = nilResponseWriter.Write(dummyBody)
	nilResponseWriter.WriteHeader(dummyStatus)

	// assert
	assert.Empty(t, header)
	assert.Zero(t, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInit_AllValuesSet(t *testing.T) {
	// arrange
	var appName = "dummyAppName"
	var roleType = "dummyRoleType"
	var hostName = "dummyHostName"
	var version = "dummyVersion"
	var buildTime = "dummyBuildTime"

	// mock
	createMock(t)

	// SUT + act
	Init(
		appName,
		roleType,
		hostName,
		version,
		buildTime,
	)
	var result, found = sessionCache.Get(uuid.Nil.String())

	// assert
	assert.True(t, found)
	assert.Equal(t, defaultSession, result)
	assert.Zero(t, defaultSession.ID)
	assert.Zero(t, defaultSession.Endpoint)
	assert.Zero(t, defaultSession.LoginID)
	assert.Equal(t, logtype.BasicLogging, defaultSession.AllowedLogType)

	// verify
	verifyAll(t)
}

func TestRegister_NilLoginID(t *testing.T) {
	// arrange
	var dummyEndpoint = "dummy endpoint"
	var dummyLoginID = uuid.Nil
	var dummyCorrelationID = uuid.New()
	var dummyAllowedLogType = logtype.LogType(rand.Intn(math.MaxInt8))
	var dummyHTTPRequest = &http.Request{}
	var dummyResponseWriter = dummyResponseWriter{}

	// stub
	sessionCache.Delete(uuid.Nil.String())

	// mock
	createMock(t)

	// SUT
	var result = Register(
		dummyEndpoint,
		dummyLoginID,
		dummyCorrelationID,
		dummyAllowedLogType,
		dummyHTTPRequest,
		dummyResponseWriter,
	)

	// act
	var _, cacheOK = sessionCache.Get(result.String())

	// assert
	assert.Equal(t, uuid.Nil, result)
	assert.False(t, cacheOK)

	// verify
	verifyAll(t)
	sessionCache.Delete(uuid.Nil.String())
}

func TestRegister_ValidLoginID(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyEndpoint = "dummy endpoint"
	var dummyLoginID = uuid.New()
	var dummyCorrelationID = uuid.New()
	var dummyAllowedLogType = logtype.LogType(rand.Intn(math.MaxInt8))
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
		dummyEndpoint,
		dummyLoginID,
		dummyCorrelationID,
		dummyAllowedLogType,
		dummyHTTPRequest,
		dummyResponseWriter,
	)

	// act
	var cacheItem, cacheOK = sessionCache.Get(dummySessionID.String())
	var session, typeOK = cacheItem.(*Session)

	// assert
	assert.Equal(t, dummySessionID, result)
	assert.True(t, cacheOK)
	assert.True(t, typeOK)
	assert.Equal(t, dummySessionID, session.ID)
	assert.Equal(t, dummyEndpoint, session.Endpoint)
	assert.Equal(t, dummyLoginID, session.LoginID)
	assert.Equal(t, dummyCorrelationID, session.CorrelationID)
	assert.Equal(t, dummyAllowedLogType, session.AllowedLogType)
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
	assert.Equal(t, defaultSession, session)

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
	assert.Equal(t, defaultSession, session)

	// verify
	verifyAll(t)
	sessionCache.Delete(dummySessionID.String())
}

func TestGet_CacheItemValid(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLoginID = uuid.New()
	var expectedSession = &Session{
		Endpoint:       "dummy endpoint",
		LoginID:        dummyLoginID,
		AllowedLogType: logtype.BasicLogging,
	}

	// stub
	sessionCache.SetDefault(dummySessionID.String(), expectedSession)

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

func TestGetRequest_NilSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject *Session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
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
	var dummyHTTPRequest, _ = http.NewRequest("FOO", "bar", nil)
	var dummySessionObject = &Session{
		Request: dummyHTTPRequest,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
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
	var dummySessionObject *Session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
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
	var dummySessionObject = &Session{
		ResponseWriter: &dummyResponseWriter,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
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

func TestClearResponseWriter_NilSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject *Session

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT
	ClearResponseWriter(
		dummySessionID,
	)

	// act
	var _, cacheOK = sessionCache.Get(dummySessionID.String())

	// assert
	assert.False(t, cacheOK)

	// verify
	verifyAll(t)
}

func TestClearResponseWriter_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummySessionObject = &Session{
		ResponseWriter: nil,
	}

	// mock
	createMock(t)

	// expect
	getFuncExpected = 1
	getFunc = func(sessionID uuid.UUID) *Session {
		getFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT
	ClearResponseWriter(
		dummySessionID,
	)

	// act
	var cacheItem, cacheOK = sessionCache.Get(dummySessionID.String())
	var session, typeOK = cacheItem.(*Session)

	// assert
	assert.True(t, cacheOK)
	assert.True(t, typeOK)
	assert.Equal(t, defaultResponseWriter, session.ResponseWriter)

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
	var dummyDataTemplate int
	var dummyError = errors.New("some error")

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

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestBody(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyDataTemplate int
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBody = "some request body"
	var dummyError = errors.New("some error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getRequestFuncExpected = 1
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
	}
	requestGetRequestBodyExpected = 1
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestBody
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) error {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyRequestBody, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}

	// SUT + act
	var err = GetRequestBody(
		dummySessionID,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
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
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	getRequestFuncExpected = 1
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
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

	// SUT + act
	var err = GetRequestParameter(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)

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
	var dummyError = errors.New("some error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getRequestFuncExpected = 1
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
	}
	muxVarsExpected = 1
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		assert.Equal(t, dummyHTTPRequest, r)
		return dummyParameters
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) error {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyValue, value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}

	// SUT + act
	var err = GetRequestParameter(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetAllQueryStrings_NotFound(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyQueries = map[string][]string{
		"test": []string{"me", "you"},
	}
	var dummyHTTPRequest = &http.Request{
		Form: dummyQueries,
	}

	// mock
	createMock(t)

	// expect
	getRequestFuncExpected = 1
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
	}

	// SUT + act
	var result = getAllQueryStrings(
		dummySessionID,
		dummyName,
	)

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestGetAllQueryStrings_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyQueries = map[string][]string{
		dummyName: []string{"me", "you"},
	}
	var dummyHTTPRequest = &http.Request{
		Form: dummyQueries,
	}

	// mock
	createMock(t)

	// expect
	getRequestFuncExpected = 1
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
	}

	// SUT + act
	var result = getAllQueryStrings(
		dummySessionID,
		dummyName,
	)

	// assert
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "me", result[0])
	assert.Equal(t, "you", result[1])

	// verify
	verifyAll(t)
}

func TestGetRequestQueryString_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueryStrings []string
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	getAllQueryStringsFuncExpected = 1
	getAllQueryStringsFunc = func(sessionID uuid.UUID, name string) []string {
		getAllQueryStringsFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyName, name)
		return dummyQueryStrings
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "The expected query string [%v] is not found in request", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyName, a[0])
		return dummyError
	}

	// SUT + act
	var err = GetRequestQueryString(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestQueryString_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueryStrings = []string{
		"some query string 1",
		"some query string 2",
		"some query string 3",
	}
	var dummyError = errors.New("some error")
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getAllQueryStringsFuncExpected = 1
	getAllQueryStringsFunc = func(sessionID uuid.UUID, name string) []string {
		getAllQueryStringsFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyName, name)
		return dummyQueryStrings
	}
	tryUnmarshalFuncExpected = 1
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) error {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyQueryStrings[0], value)
		*(dataTemplate.(*int)) = dummyResult
		return dummyError
	}

	// SUT + act
	var err = GetRequestQueryString(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestGetRequestQueryStrings_EmptyList(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueryStrings []string
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var dummyError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	getAllQueryStringsFuncExpected = 1
	getAllQueryStringsFunc = func(sessionID uuid.UUID, name string) []string {
		getAllQueryStringsFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyName, name)
		return dummyQueryStrings
	}
	dummyFillCallbackExpected = 0
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, "Failed to get request query strings", baseErrorMessage)
		assert.Equal(t, 0, len(allErrors))
		return dummyError
	}

	// SUT + act
	var err = GetRequestQueryStrings(
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

func TestGetRequestQueryStrings_HappyPath(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some name"
	var dummyDataTemplate int
	var dummyQueryStrings = []string{
		"some query string 1",
		"some query string 2",
		"some query string 3",
	}
	var dummyFillCallbackExpected int
	var dummyFillCallbackCalled int
	var dummyFillCallback func()
	var unmarshalErrors = []error{
		nil,
		errors.New("some unmarshal error"),
		nil,
	}
	var dummyError = apperror.GetGeneralFailureError(nil)
	var dummyResult = rand.Int()

	// mock
	createMock(t)

	// expect
	getAllQueryStringsFuncExpected = 1
	getAllQueryStringsFunc = func(sessionID uuid.UUID, name string) []string {
		getAllQueryStringsFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyName, name)
		return dummyQueryStrings
	}
	tryUnmarshalFuncExpected = 3
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) error {
		tryUnmarshalFuncCalled++
		assert.Equal(t, dummyQueryStrings[tryUnmarshalFuncCalled-1], value)
		*(dataTemplate.(*int)) = dummyResult
		return unmarshalErrors[tryUnmarshalFuncCalled-1]
	}
	dummyFillCallbackExpected = 2
	dummyFillCallback = func() {
		dummyFillCallbackCalled++
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, "Failed to get request query strings", baseErrorMessage)
		assert.Equal(t, 1, len(allErrors))
		assert.Equal(t, unmarshalErrors[1], allErrors[0])
		return dummyError
	}

	// SUT + act
	var err = GetRequestQueryStrings(
		dummySessionID,
		dummyName,
		&dummyDataTemplate,
		dummyFillCallback,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Equal(t, dummyResult, dummyDataTemplate)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyFillCallbackExpected, dummyFillCallbackCalled, "Unexpected number of calls to dummyFillCallback")
}
