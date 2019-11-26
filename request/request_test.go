package request

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

func TestGetAllowedLogType_NilHTTPRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request

	// mock
	createMock(t)

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, logtype.GeneralTracing, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_NoMatchingHeaderFound(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("foo", "bar")

	// mock
	createMock(t)

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, logtype.GeneralTracing, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_MatchingHeaderEmpty(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("Foo", "bar")
	dummyHTTPRequest.Header["Log-Type"] = []string{}

	// mock
	createMock(t)

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, logtype.GeneralTracing, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_MatchingHeaderInvalid(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("Foo", "bar")
	dummyHTTPRequest.Header.Add("Log-Type", "abc")
	dummyHTTPRequest.Header.Add("Log-Type", "def")

	// mock
	createMock(t)

	// expect
	logtypeFromStringExpected = 2
	logtypeFromString = func(value string) logtype.LogType {
		logtypeFromStringCalled++
		if logtypeFromStringCalled == 1 {
			assert.Equal(t, "abc", value)
		} else if logtypeFromStringCalled == 2 {
			assert.Equal(t, "def", value)
		}
		return logtype.AppRoot
	}

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, logtype.GeneralTracing, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogType_MatchingHeaderValid(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("Foo", "bar")
	dummyHTTPRequest.Header.Add("Log-Type", logtype.MethodEnterName)
	dummyHTTPRequest.Header.Add("Log-Type", logtype.MethodExitName)

	// mock
	createMock(t)

	// expect
	logtypeFromStringExpected = 2
	logtypeFromString = func(value string) logtype.LogType {
		logtypeFromStringCalled++
		if logtypeFromStringCalled == 1 {
			assert.Equal(t, logtype.MethodEnterName, value)
			return logtype.MethodEnter
		} else if logtypeFromStringCalled == 2 {
			assert.Equal(t, logtype.MethodExitName, value)
			return logtype.MethodExit
		}
		return logtype.AppRoot
	}

	// SUT + act
	var allowedLogType = GetAllowedLogType(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, logtype.MethodEnter|logtype.MethodExit, allowedLogType)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_NilHTTPRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request

	// mock
	createMock(t)

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, loglevel.Warn, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_NoMatchingHeaderFound(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("foo", "bar")

	// mock
	createMock(t)

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, loglevel.Warn, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_MatchingHeaderEmpty(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("Foo", "bar")
	dummyHTTPRequest.Header["Log-Level"] = []string{}

	// mock
	createMock(t)

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, loglevel.Warn, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_MatchingHeader(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost/",
		nil,
	)

	// stub
	dummyHTTPRequest.Header.Add("Foo", "bar")
	dummyHTTPRequest.Header.Add("Log-Level", loglevel.FatalName)
	dummyHTTPRequest.Header.Add("Log-Level", loglevel.InfoName)

	// mock
	createMock(t)

	// expect
	loglevelFromStringExpected = 2
	loglevelFromString = func(value string) loglevel.LogLevel {
		loglevelFromStringCalled++
		if loglevelFromStringCalled == 1 {
			assert.Equal(t, loglevel.FatalName, value)
			return loglevel.Fatal
		} else if loglevelFromStringCalled == 2 {
			assert.Equal(t, loglevel.InfoName, value)
			return loglevel.Info
		}
		return loglevel.Warn
	}

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, loglevel.Info, allowedLogLevel)

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
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
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
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
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
