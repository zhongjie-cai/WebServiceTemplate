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
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
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
	assert.Empty(t, server.TLSConfig.CipherSuites)
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Zero(t, server.WriteTimeout)
	assert.Zero(t, server.ReadTimeout)
	assert.Zero(t, server.IdleTimeout)

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
	assert.Equal(t, tls.RequestClientCert, server.TLSConfig.ClientAuth)
	assert.Nil(t, server.TLSConfig.ClientCAs)
	assert.Empty(t, server.TLSConfig.CipherSuites)
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Zero(t, server.WriteTimeout)
	assert.Zero(t, server.ReadTimeout)
	assert.Zero(t, server.IdleTimeout)

	// verify
	verifyAll(t)
}

func TestCreateServer_HTTPS_NoCaCert(t *testing.T) {
	// arrange
	var dummyServeHTTPS = true
	var dummyValidateClientCert = true
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
	certificateGetCaCertPoolExpected = 1
	certificateGetCaCertPool = func() *x509.CertPool {
		certificateGetCaCertPoolCalled++
		return nil
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
	assert.Equal(t, tls.RequireAnyClientCert, server.TLSConfig.ClientAuth)
	assert.Nil(t, server.TLSConfig.ClientCAs)
	assert.Empty(t, server.TLSConfig.CipherSuites)
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Zero(t, server.WriteTimeout)
	assert.Zero(t, server.ReadTimeout)
	assert.Zero(t, server.IdleTimeout)

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
	certificateGetCaCertPoolExpected = 1
	certificateGetCaCertPool = func() *x509.CertPool {
		certificateGetCaCertPoolCalled++
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
	assert.Empty(t, server.TLSConfig.CipherSuites)
	assert.Equal(t, true, server.TLSConfig.PreferServerCipherSuites)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion)
	assert.Zero(t, server.WriteTimeout)
	assert.Zero(t, server.ReadTimeout)
	assert.Zero(t, server.IdleTimeout)

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
	assert.NotPanics(
		t,
		func() {
			go listenAndServe(
				dummyServer,
				dummyServeHTTPS,
			)
			var err = dummyServer.Close()

			// assert
			assert.NoError(t, err)
		},
	)

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
	assert.NotPanics(
		t,
		func() {
			go listenAndServe(
				dummyServer,
				dummyServeHTTPS,
			)
			var err = dummyServer.Close()

			// assert
			assert.NoError(t, err)
		},
	)

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

func TestConsolidateError_SkipServerClosedErrors(t *testing.T) {
	// arrange
	var dummyHostError = http.ErrServerClosed
	var dummyShutDownError = http.ErrServerClosed
	var dummyMessageFormat = "One or more errors have occurred during server hosting"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 2, len(innerErrors))
		assert.Nil(t, innerErrors[0])
		assert.Nil(t, innerErrors[1])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = consolidateError(
		dummyHostError,
		dummyShutDownError,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestConsolidateError_NoSkipOtherErrors(t *testing.T) {
	// arrange
	var dummyHostError = errors.New("some host error")
	var dummyShutDownError = errors.New("some shutdown error")
	var dummyMessageFormat = "One or more errors have occurred during server hosting"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 2, len(innerErrors))
		assert.Equal(t, dummyHostError, innerErrors[0])
		assert.Equal(t, dummyShutDownError, innerErrors[1])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = consolidateError(
		dummyHostError,
		dummyShutDownError,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

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
	var dummyGraceShutdownWaitTime = time.Duration(rand.Intn(100)) * time.Second
	var dummyShutDownError = errors.New("some shut down error message")
	var dummyAppError = errors.New("some app error")

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
		assert.Equal(t, 2, len(sig))
		assert.Equal(t, os.Interrupt, sig[0])
		assert.Equal(t, os.Kill, sig[1])
	}
	listenAndServeFuncExpected = 1
	listenAndServeFunc = func(server *http.Server, serveHTTPS bool) error {
		listenAndServeFuncCalled++
		assert.Equal(t, dummyServer, server)
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		return dummyHostError
	}
	haltFuncExpected = 1
	haltFunc = func() {
		haltFuncCalled++
		shutdownSignal <- os.Interrupt
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "server", category)
		assert.Equal(t, "Host", subcategory)
		assert.Equal(t, "Interrupt signal received: Terminating server", messageFormat)
		assert.Empty(t, parameters)
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
	configGraceShutdownWaitTimeExpected = 1
	configGraceShutdownWaitTime = func() time.Duration {
		configGraceShutdownWaitTimeCalled++
		return dummyGraceShutdownWaitTime
	}
	contextWithTimeoutExpected = 1
	contextWithTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		contextWithTimeoutCalled++
		assert.Equal(t, dummyBackgroundContext, parent)
		assert.Equal(t, dummyGraceShutdownWaitTime, timeout)
		return dummyRuntimeContext, cancelCallback
	}
	shutDownFuncExpected = 1
	shutDownFunc = func(runtimeContext context.Context, server *http.Server) error {
		shutDownFuncCalled++
		assert.Equal(t, dummyRuntimeContext, runtimeContext)
		assert.Equal(t, dummyServer, server)
		return dummyShutDownError
	}
	consolidateErrorFuncExpected = 1
	consolidateErrorFunc = func(hostError error, shutdownError error) error {
		consolidateErrorFuncCalled++
		assert.Equal(t, dummyHostError, hostError)
		assert.Equal(t, dummyShutDownError, shutdownError)
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
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	registerInstantiateExpected = 1
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return dummyRouter, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
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
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	registerInstantiateExpected = 1
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return dummyRouter, nil
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "server", category)
		assert.Equal(t, "Host", subcategory)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Targeting port [%v] HTTPS [%v] mTLS [%v]", messageFormat)
			assert.Equal(t, 3, len(parameters))
			assert.Equal(t, dummyAppPort, parameters[0])
			assert.Equal(t, dummyServeHTTPS, parameters[1])
			assert.Equal(t, dummyValidateClientCert, parameters[2])
		} else {
			assert.Equal(t, "Server terminated", messageFormat)
			assert.Empty(t, parameters)
		}
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
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
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
	loggerAppRootExpected = 2
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "server", category)
		assert.Equal(t, "Host", subcategory)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Targeting port [%v] HTTPS [%v] mTLS [%v]", messageFormat)
			assert.Equal(t, 3, len(parameters))
			assert.Equal(t, dummyAppPort, parameters[0])
			assert.Equal(t, dummyServeHTTPS, parameters[1])
			assert.Equal(t, dummyValidateClientCert, parameters[2])
		} else {
			assert.Equal(t, "Server terminated", messageFormat)
			assert.Empty(t, parameters)
		}
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

func TestHalt(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	Halt()

	// act
	var result, ok = <-shutdownSignal

	// assert
	assert.True(t, ok)
	assert.Equal(t, os.Interrupt, result)

	// verify
	verifyAll(t)
}
