package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestCreateServer_NoHTTPS(t *testing.T) {
	// arrange
	var dummyServeHTTPS = false
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}

	// mock
	createMock(t)

	// SUT + act
	var server = createServer(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
		dummyRouter,
	)

	// assert
	assert.NotNil(t, server)
	assert.Equal(t, ":"+dummyAppPort, server.Addr)
	assert.NotNil(t, server.TLSConfig)
	assert.Empty(t, server.TLSConfig.Certificates)
	assert.Equal(t, tls.NoClientCert, server.TLSConfig.ClientAuth)
	assert.Nil(t, server.TLSConfig.ClientCAs)
	assert.Equal(t, 1, len(server.TLSConfig.CipherSuites))
	assert.Equal(t, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, server.TLSConfig.CipherSuites[0])
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Equal(t, time.Second*60, server.WriteTimeout)
	assert.Equal(t, time.Second*60, server.ReadTimeout)
	assert.Equal(t, time.Second*180, server.IdleTimeout)

	// verify
	verifyAll(t)
}

func TestCreateServer_HTTPS_NoValidateClientCert(t *testing.T) {
	// arrange
	var dummyServeHTTPS = true
	var dummyValidateClientCert = false
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}
	var dummyServerCert = &tls.Certificate{}

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() *tls.Certificate {
		certificateGetServerCertificateCalled++
		return dummyServerCert
	}

	// SUT + act
	var server = createServer(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
		dummyRouter,
	)

	// assert
	assert.NotNil(t, server)
	assert.Equal(t, ":"+dummyAppPort, server.Addr)
	assert.NotNil(t, server.TLSConfig)
	assert.Equal(t, 1, len(server.TLSConfig.Certificates))
	assert.Equal(t, *dummyServerCert, server.TLSConfig.Certificates[0])
	assert.Equal(t, tls.NoClientCert, server.TLSConfig.ClientAuth)
	assert.Nil(t, server.TLSConfig.ClientCAs)
	assert.Equal(t, 1, len(server.TLSConfig.CipherSuites))
	assert.Equal(t, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, server.TLSConfig.CipherSuites[0])
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Equal(t, time.Second*60, server.WriteTimeout)
	assert.Equal(t, time.Second*60, server.ReadTimeout)
	assert.Equal(t, time.Second*180, server.IdleTimeout)

	// verify
	verifyAll(t)
}

func TestCreateServer_HTTPS_ValidateClientCert(t *testing.T) {
	// arrange
	var dummyServeHTTPS = true
	var dummyValidateClientCert = true
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// expect
	certificateGetServerCertificateExpected = 1
	certificateGetServerCertificate = func() *tls.Certificate {
		certificateGetServerCertificateCalled++
		return dummyServerCert
	}
	certificateGetClientCertPoolExpected = 1
	certificateGetClientCertPool = func() *x509.CertPool {
		certificateGetClientCertPoolCalled++
		return dummyCertPool
	}

	// SUT + act
	var server = createServer(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
		dummyRouter,
	)

	// assert
	assert.NotNil(t, server)
	assert.Equal(t, ":"+dummyAppPort, server.Addr)
	assert.NotNil(t, server.TLSConfig)
	assert.Equal(t, 1, len(server.TLSConfig.Certificates))
	assert.Equal(t, *dummyServerCert, server.TLSConfig.Certificates[0])
	assert.Equal(t, tls.RequireAndVerifyClientCert, server.TLSConfig.ClientAuth)
	assert.Equal(t, dummyCertPool, server.TLSConfig.ClientCAs)
	assert.Equal(t, 1, len(server.TLSConfig.CipherSuites))
	assert.Equal(t, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, server.TLSConfig.CipherSuites[0])
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Equal(t, time.Second*60, server.WriteTimeout)
	assert.Equal(t, time.Second*60, server.ReadTimeout)
	assert.Equal(t, time.Second*180, server.IdleTimeout)

	// verify
	verifyAll(t)
}

func TestListenAndServe_HTTPS(t *testing.T) {
	// arrange
	var dummyServer = &http.Server{}
	var dummyServeHTTPS = true

	// mock
	createMock(t)

	// SUT + act
	var err = listenAndServe(
		dummyServer,
		dummyServeHTTPS,
	)

	// assert
	assert.NotNil(t, err)

	// verify
	verifyAll(t)
}

func TestListenAndServe_HTTP(t *testing.T) {
	// arrange
	var dummyServer = &http.Server{}
	var dummyServeHTTPS = false

	// mock
	createMock(t)

	// SUT + act
	var err = listenAndServe(
		dummyServer,
		dummyServeHTTPS,
	)

	// assert
	assert.NotNil(t, err)

	// verify
	verifyAll(t)
}

func TestShutDown(t *testing.T) {
	// arrange
	var dummyContext = context.TODO()
	var dummyServer = &http.Server{}

	// mock
	createMock(t)

	// SUT + act
	var err = shutDown(
		dummyContext,
		dummyServer,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestRunServer_HappyPath(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}
	var dummyServer = &http.Server{}
	var dummyHostError = errors.New("some host error message")
	var dummyBackgroundContext = context.Background()
	var dummyRuntimeContext = context.TODO()
	var dummyShutDownError = errors.New("some shut down error message")
	var dummyMessageFormat = "One or more errors have occurred during server hosting"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	createServerFuncExpected = 1
	createServerFunc = func(serveHTTPS bool, validateClientCert bool, appPort string, router *mux.Router) *http.Server {
		createServerFuncCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		assert.Equal(t, dummyRouter, router)
		return dummyServer
	}
	signalNotifyExpected = 1
	signalNotify = func(c chan<- os.Signal, sig ...os.Signal) {
		signalNotifyCalled++
		assert.Equal(t, 1, len(sig))
		assert.Equal(t, os.Interrupt, sig[0])
	}
	listenAndServeFuncExpected = 1
	listenAndServeFunc = func(server *http.Server, serveHTTPS bool) error {
		listenAndServeFuncCalled++
		assert.Equal(t, dummyServer, server)
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		return dummyHostError
	}
	contextBackgroundExpected = 1
	contextBackground = func() context.Context {
		contextBackgroundCalled++
		return dummyBackgroundContext
	}
	var cancelCallbackExpected = 1
	var cancelCallbackCalled = 0
	var cancelCallback = func() {
		cancelCallbackCalled++
	}
	contextWithTimeoutExpected = 1
	contextWithTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		contextWithTimeoutCalled++
		assert.Equal(t, dummyBackgroundContext, parent)
		assert.Equal(t, 15*time.Second, timeout)
		return dummyRuntimeContext, cancelCallback
	}
	shutDownFuncExpected = 1
	shutDownFunc = func(runtimeContext context.Context, server *http.Server) error {
		shutDownFuncCalled++
		assert.Equal(t, dummyRuntimeContext, runtimeContext)
		assert.Equal(t, dummyServer, server)
		return dummyShutDownError
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, dummyMessageFormat, baseErrorMessage)
		assert.Equal(t, 2, len(allErrors))
		assert.Equal(t, dummyHostError, allErrors[0])
		assert.Equal(t, dummyShutDownError, allErrors[1])
		return dummyAppError
	}

	// SUT + act
	var err = runServer(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
		dummyRouter,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, cancelCallbackExpected, cancelCallbackCalled, "Unexpected number of calls to cancelCallback")
}

func TestHost_ErrorRegisterRoutes(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to host entries on port %v"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	registerInstantiateExpected = 1
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return dummyRouter, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var err = Host(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHost_ErrorRunServer(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to run server on port %v"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	registerInstantiateExpected = 1
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return dummyRouter, nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "server", category)
		assert.Equal(t, "Host", subcategory)
		assert.Equal(t, "Targeting port [%v] HTTPS [%v] mTLS [%v]", messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		assert.Equal(t, dummyServeHTTPS, parameters[1])
		assert.Equal(t, dummyValidateClientCert, parameters[2])
	}
	runServerFuncExpected = 1
	runServerFunc = func(serveHTTPS bool, validateClientCert bool, appPort string, router *mux.Router) error {
		runServerFuncCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		assert.Equal(t, dummyRouter, router)
		return dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var err = Host(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestHost_Success(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyAppPort = "some app port"
	var dummyRouter = &mux.Router{}

	// mock
	createMock(t)

	// expect
	registerInstantiateExpected = 1
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return dummyRouter, nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "server", category)
		assert.Equal(t, "Host", subcategory)
		assert.Equal(t, "Targeting port [%v] HTTPS [%v] mTLS [%v]", messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyAppPort, parameters[0])
		assert.Equal(t, dummyServeHTTPS, parameters[1])
		assert.Equal(t, dummyValidateClientCert, parameters[2])
	}
	runServerFuncExpected = 1
	runServerFunc = func(serveHTTPS bool, validateClientCert bool, appPort string, router *mux.Router) error {
		runServerFuncCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		assert.Equal(t, dummyRouter, router)
		return nil
	}

	// SUT + act
	var err = Host(
		dummyServeHTTPS,
		dummyValidateClientCert,
		dummyAppPort,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}
