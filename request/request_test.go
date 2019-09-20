package request

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

func TestGetUUIDFromHeader_EmptyHeader(t *testing.T) {
	// arrange
	var dummyHeader = make(http.Header)
	var dummyName = "some name"
	var expectedError = errors.New("some error")
	var expectedUUID = uuid.New()

	// mock
	createMock(t)

	// expect
	uuidParseExpected = 1
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		assert.Equal(t, "", s)
		return uuid.Nil, expectedError
	}
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return expectedUUID
	}

	// SUT + act
	var parsedUUID = getUUIDFromHeader(
		dummyHeader,
		dummyName,
	)

	// assert
	assert.NotZero(t, parsedUUID)

	// verify
	verifyAll(t)
}

func TestGetUUIDFromHeader_HeaderNoUUID(t *testing.T) {
	// arrange
	var dummyHeader = make(http.Header)
	var dummyName = "some name"
	var expectedError = errors.New("some error")
	var expectedUUID = uuid.New()

	// stub
	dummyHeader.Add("foo", "bar")

	// mock
	createMock(t)

	// expect
	uuidParseExpected = 1
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		assert.Equal(t, "", s)
		return uuid.Nil, expectedError
	}
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return expectedUUID
	}

	// SUT + act
	var parsedUUID = getUUIDFromHeader(
		dummyHeader,
		dummyName,
	)

	// assert
	assert.Equal(t, expectedUUID, parsedUUID)

	// verify
	verifyAll(t)
}

func TestGetUUIDFromHeader_HeaderInvalidUUID(t *testing.T) {
	// arrange
	var dummyHeader = make(http.Header)
	var dummyName = "some name"
	var invalidUUID = "some invalid UUID"
	var expectedError = errors.New("some error")
	var expectedUUID = uuid.New()

	// stub
	dummyHeader.Add("foo", "bar")
	dummyHeader.Add(dummyName, invalidUUID)

	// mock
	createMock(t)

	// expect
	uuidParseExpected = 1
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		assert.Equal(t, invalidUUID, s)
		return uuid.Nil, expectedError
	}
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return expectedUUID
	}

	// SUT + act
	var parsedUUID = getUUIDFromHeader(
		dummyHeader,
		dummyName,
	)

	// assert
	assert.Equal(t, expectedUUID, parsedUUID)

	// verify
	verifyAll(t)
}

func TestGetUUIDFromHeader_HeaderValidUUID(t *testing.T) {
	// arrange
	var dummyHeader = make(http.Header)
	var dummyName = "some name"
	var expectedUUID, _ = uuid.NewUUID()

	// stub
	dummyHeader.Add("foo", "bar")
	dummyHeader.Add(dummyName, expectedUUID.String())

	// mock
	createMock(t)

	// expect
	uuidParseExpected = 1
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		return uuid.Parse(s)
	}

	// SUT + act
	var parsedUUID = getUUIDFromHeader(
		dummyHeader,
		dummyName,
	)

	// assert
	assert.Equal(t, expectedUUID.String(), parsedUUID.String())

	// verify
	verifyAll(t)
}

func TestGetLoginID_NilRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request
	var expectedCorrelationID = uuid.New()

	// mock
	createMock(t)

	// expect
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return expectedCorrelationID
	}

	// SUT + act
	var result = GetLoginID(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, expectedCorrelationID, result)

	// tear down
	verifyAll(t)
}

func TestGetLoginID_ValidRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)
	var expectedCorrelationID = uuid.New()

	// mock
	createMock(t)

	// expect
	getUUIDFromHeaderFuncExpected = 1
	getUUIDFromHeaderFunc = func(header http.Header, name string) uuid.UUID {
		getUUIDFromHeaderFuncCalled++
		assert.Equal(t, dummyHTTPRequest.Header, header)
		assert.Equal(t, "login-id", name)
		return expectedCorrelationID
	}

	// SUT + act
	var result = GetLoginID(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, expectedCorrelationID, result)

	// tear down
	verifyAll(t)
}

func TestGetCorrelationID_NilRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request
	var expectedCorrelationID = uuid.New()

	// mock
	createMock(t)

	// expect
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return expectedCorrelationID
	}

	// SUT + act
	var result = GetCorrelationID(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, expectedCorrelationID, result)

	// tear down
	verifyAll(t)
}

func TestGetCorrelationID_ValidRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)
	var expectedCorrelationID = uuid.New()

	// mock
	createMock(t)

	// expect
	getUUIDFromHeaderFuncExpected = 1
	getUUIDFromHeaderFunc = func(header http.Header, name string) uuid.UUID {
		getUUIDFromHeaderFuncCalled++
		assert.Equal(t, dummyHTTPRequest.Header, header)
		assert.Equal(t, "correlation-id", name)
		return expectedCorrelationID
	}

	// SUT + act
	var result = GetCorrelationID(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, expectedCorrelationID, result)

	// tear down
	verifyAll(t)
}

func TestGetAllowedLogType(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)
	var dummyHeaderValue = "some header value"
	var dummyLogType = logtype.LogType(rand.Intn(256))

	// stub
	dummyHTTPRequest.Header.Add("foo", "bar")
	dummyHTTPRequest.Header.Add("log-type", dummyHeaderValue)

	// mock
	createMock(t)

	// expect
	logtypeFromStringExpected = 1
	logtypeFromString = func(value string) logtype.LogType {
		logtypeFromStringCalled++
		assert.Equal(t, dummyHeaderValue, value)
		return dummyLogType
	}

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyLogType, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetClientCertificates_RequestNil(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request
	var dummyMessageFormat = "Invalid request or insecure communication channel"
	var dummySyncError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummySyncError
	}

	// SUT + act
	var result, err = GetClientCertificates(
		dummyHTTPRequest,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummySyncError, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertificates_TLSNil(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyMessageFormat = "Invalid request or insecure communication channel"
	var dummySyncError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummySyncError
	}

	// SUT + act
	var result, err = GetClientCertificates(
		dummyHTTPRequest,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummySyncError, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertificates_Success(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		TLS: &tls.ConnectionState{
			PeerCertificates: []*x509.Certificate{
				&x509.Certificate{},
			},
		},
	}

	// mock
	createMock(t)

	// SUT + act
	var result, err = GetClientCertificates(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyHTTPRequest.TLS.PeerCertificates, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_NilBody(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1",
		nil,
	)

	// mock
	createMock(t)

	// SUT + act
	var result = GetRequestBody(
		dummyHTTPRequest,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_ErrorBody(t *testing.T) {
	// arrange
	var bodyContent = "some body content"
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/featuretoggle",
		strings.NewReader(bodyContent),
	)
	var dummyError = errors.New("some error message")

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummyHTTPRequest.Body, r)
		return nil, dummyError
	}

	// SUT + act
	var result = GetRequestBody(
		dummyHTTPRequest,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_Success(t *testing.T) {
	// arrange
	var bodyContent = "some body content"
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/featuretoggle",
		strings.NewReader(bodyContent),
	)
	var dummyBuffer = &bytes.Buffer{}
	var dummyReadCloser = ioutil.NopCloser(nil)

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		return ioutil.ReadAll(r)
	}
	bytesNewBufferExpected = 1
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		assert.Equal(t, []byte(bodyContent), buf)
		return dummyBuffer
	}
	ioutilNopCloserExpected = 1
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		assert.Equal(t, dummyBuffer, r)
		return dummyReadCloser
	}

	// SUT + act
	var result = GetRequestBody(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, bodyContent, result)
	assert.Equal(t, dummyReadCloser, dummyHTTPRequest.Body)

	// verify
	verifyAll(t)
}
