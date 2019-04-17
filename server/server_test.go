package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestCreateServer(t *testing.T) {
	// arrange
	var dummyServerCert = &tls.Certificate{}
	var dummyClientCertPool = &x509.CertPool{}
	var dummyAppPort = "some app port"

	// mock
	createMock(t)

	// expect
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}

	// SUT + act
	var server = createServer(
		dummyServerCert,
		dummyClientCertPool,
	)

	// assert
	assert.NotNil(t, server)
	assert.Equal(t, ":"+dummyAppPort, server.Addr)
	assert.NotNil(t, server.TLSConfig)
	assert.Equal(t, 1, len(server.TLSConfig.Certificates))
	assert.Equal(t, *dummyServerCert, server.TLSConfig.Certificates[0])
	assert.Equal(t, tls.RequireAndVerifyClientCert, server.TLSConfig.ClientAuth)
	assert.Equal(t, dummyClientCertPool, server.TLSConfig.ClientCAs)
	assert.Equal(t, 1, len(server.TLSConfig.CipherSuites))
	assert.Equal(t, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, server.TLSConfig.CipherSuites[0])
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)

	// verify
	verifyAll(t)
}

func TestListenAndServeTLS(t *testing.T) {
	// arrange
	var dummyServer = &http.Server{}

	// mock
	createMock(t)

	// SUT + act
	var err = listenAndServeTLS(
		dummyServer,
	)

	// assert
	assert.NotNil(t, err)

	// verify
	verifyAll(t)
}

func TestRunServer_ServerCertError(t *testing.T) {
	// arrange
	var dummyServerCert = &tls.Certificate{}
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to run server due to server cert error"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() (*tls.Certificate, error) {
		certificateGetServerCertificateCalled++
		return dummyServerCert, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = runServer()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestRunServer_CertPoolError(t *testing.T) {
	// arrange
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to run server due to client cert pool error"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() (*tls.Certificate, error) {
		certificateGetServerCertificateCalled++
		return dummyServerCert, nil
	}
	certificateGetClientCertPoolExpected = 1
	certificateGetClientCertPool = func() (*x509.CertPool, error) {
		certificateGetClientCertPoolCalled++
		return dummyCertPool, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = runServer()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestRunServer_ServeError(t *testing.T) {
	// arrange
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}
	var dummyServer = &http.Server{}
	var dummyAppPort = "some app port"
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to host service on port %v"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() (*tls.Certificate, error) {
		certificateGetServerCertificateCalled++
		return dummyServerCert, nil
	}
	certificateGetClientCertPoolExpected = 1
	certificateGetClientCertPool = func() (*x509.CertPool, error) {
		certificateGetClientCertPoolCalled++
		return dummyCertPool, nil
	}
	createServerFuncExpected = 1
	createServerFunc = func(serverCert *tls.Certificate, clientCertPool *x509.CertPool) *http.Server {
		createServerFuncCalled++
		assert.Equal(t, dummyServerCert, serverCert)
		assert.Equal(t, dummyCertPool, clientCertPool)
		return dummyServer
	}
	listenAndServeTLSFuncExpected = 1
	listenAndServeTLSFunc = func(server *http.Server) error {
		listenAndServeTLSFuncCalled++
		assert.Equal(t, dummyServer, server)
		return dummyError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var err = runServer()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestRunServer_Success(t *testing.T) {
	// arrange
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}
	var dummyServer = &http.Server{}

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() (*tls.Certificate, error) {
		certificateGetServerCertificateCalled++
		return dummyServerCert, nil
	}
	certificateGetClientCertPoolExpected = 1
	certificateGetClientCertPool = func() (*x509.CertPool, error) {
		certificateGetClientCertPoolCalled++
		return dummyCertPool, nil
	}
	createServerFuncExpected = 1
	createServerFunc = func(serverCert *tls.Certificate, clientCertPool *x509.CertPool) *http.Server {
		createServerFuncCalled++
		assert.Equal(t, dummyServerCert, serverCert)
		assert.Equal(t, dummyCertPool, clientCertPool)
		return dummyServer
	}
	listenAndServeTLSFuncExpected = 1
	listenAndServeTLSFunc = func(server *http.Server) error {
		listenAndServeTLSFuncCalled++
		assert.Equal(t, dummyServer, server)
		return nil
	}

	// SUT + act
	var err = runServer()

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestHostEntries_NilEntries(t *testing.T) {
	// arrange
	var expectedErrorMessage = "No host entries found"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = hostEntries()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHostEntries_EmptyEntries(t *testing.T) {
	// arrange
	var expectedErrorMessage = "No host entries found"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	var dummyEntryFuncs = []func(){}

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = hostEntries(
		dummyEntryFuncs...,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHostEntries_ValidEntries(t *testing.T) {
	// arrange
	var entryFuncsCalled = 0

	// stub
	var dummyEntryFuncs = []func(){
		func() { entryFuncsCalled++ },
		func() { entryFuncsCalled++ },
		func() { entryFuncsCalled++ },
	}

	// mock
	createMock(t)

	// SUT + act
	var err = hostEntries(
		dummyEntryFuncs...,
	)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, len(dummyEntryFuncs), entryFuncsCalled)

	// verify
	verifyAll(t)
}

func TestHost_ErrorHostEntries(t *testing.T) {
	// arrange
	var dummyAppPort = "some app port"
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to host entries on port %v"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	hostEntriesFuncExpected = 1
	hostEntriesFunc = func(entryFuncs ...func()) error {
		hostEntriesFuncCalled++
		var pointers = []string{}
		for _, entryFunc := range entryFuncs {
			var pointer = fmt.Sprintf("%v", reflect.ValueOf(entryFunc))
			pointers = append(pointers, pointer)
		}
		assert.Equal(t, 3, len(pointers))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(healthHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(faviconHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(swaggerHostEntry)))
		return dummyError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var err = Host()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHost_ErrorRunServer(t *testing.T) {
	// arrange
	var dummyAppPort = "some app port"
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to run server on port %v"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	hostEntriesFuncExpected = 1
	hostEntriesFunc = func(entryFuncs ...func()) error {
		hostEntriesFuncCalled++
		var pointers = []string{}
		for _, entryFunc := range entryFuncs {
			var pointer = fmt.Sprintf("%v", reflect.ValueOf(entryFunc))
			pointers = append(pointers, pointer)
		}
		assert.Equal(t, 3, len(pointers))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(healthHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(faviconHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(swaggerHostEntry)))
		return nil
	}
	runServerFuncExpected = 1
	runServerFunc = func() error {
		runServerFuncCalled++
		return dummyError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var err = Host()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHost_Success(t *testing.T) {
	// mock
	createMock(t)

	// expect
	hostEntriesFuncExpected = 1
	hostEntriesFunc = func(entryFuncs ...func()) error {
		hostEntriesFuncCalled++
		var pointers = []string{}
		for _, entryFunc := range entryFuncs {
			var pointer = fmt.Sprintf("%v", reflect.ValueOf(entryFunc))
			pointers = append(pointers, pointer)
		}
		assert.Equal(t, 3, len(pointers))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(healthHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(faviconHostEntry)))
		assert.Contains(t, pointers, fmt.Sprintf("%v", reflect.ValueOf(swaggerHostEntry)))
		return nil
	}
	runServerFuncExpected = 1
	runServerFunc = func() error {
		runServerFuncCalled++
		return nil
	}

	// SUT + act
	var err = Host()

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}
