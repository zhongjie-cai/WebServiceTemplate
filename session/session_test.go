package session

import (
	"math"
	"math/rand"
	"net/http"
	"testing"

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
