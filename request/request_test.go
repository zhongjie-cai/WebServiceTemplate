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

	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"

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
	var dummyLogType = logtype.LogType(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogTypeExpected = 1
	config.DefaultAllowedLogType = func() logtype.LogType {
		configDefaultAllowedLogTypeCalled++
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

func TestGetAllowedLogType_NoCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyLogType = logtype.LogType(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogTypeExpected = 1
	config.DefaultAllowedLogType = func() logtype.LogType {
		configDefaultAllowedLogTypeCalled++
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

func TestGetAllowedLogType_WithCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyLogType = logtype.LogType(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	customizationSessionAllowedLogTypeExpected = 1
	customization.SessionAllowedLogType = func(httpRequest *http.Request) logtype.LogType {
		customizationSessionAllowedLogTypeCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
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

func TestGetAllowedLogLevel_NilHTTPRequest(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request
	var dummyLogLevel = loglevel.LogLevel(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogLevelExpected = 1
	config.DefaultAllowedLogLevel = func() loglevel.LogLevel {
		configDefaultAllowedLogLevelCalled++
		return dummyLogLevel
	}

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyLogLevel, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_NoCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyLogLevel = loglevel.LogLevel(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	configDefaultAllowedLogLevelExpected = 1
	config.DefaultAllowedLogLevel = func() loglevel.LogLevel {
		configDefaultAllowedLogLevelCalled++
		return dummyLogLevel
	}

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyLogLevel, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetAllowedLogLevel_WithCustomization(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyLogLevel = loglevel.LogLevel(rand.Intn(255))

	// mock
	createMock(t)

	// expect
	customizationSessionAllowedLogLevelExpected = 1
	customization.SessionAllowedLogLevel = func(httpRequest *http.Request) loglevel.LogLevel {
		customizationSessionAllowedLogLevelCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyLogLevel
	}

	// SUT + act
	var allowedLogLevel = GetAllowedLogLevel(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyLogLevel, allowedLogLevel)

	// verify
	verifyAll(t)
}

func TestGetClientCertificates_RequestNil(t *testing.T) {
	// arrange
	var dummyHTTPRequest *http.Request
	var dummyMessageFormat = "Invalid request or insecure communication channel"
	var dummySyncError = apperror.GetCustomError(0, "")

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
	var dummySyncError = apperror.GetCustomError(0, "")

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
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}

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
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
		Body:       ioutil.NopCloser(strings.NewReader(bodyContent)),
	}
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
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
		Body:       ioutil.NopCloser(strings.NewReader(bodyContent)),
	}
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
