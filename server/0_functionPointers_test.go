package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/register"
)

var (
	certificateGetServerCertificateExpected int
	certificateGetServerCertificateCalled   int
	certificateGetClientCertPoolExpected    int
	certificateGetClientCertPoolCalled      int
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	registerInstantiateExpected             int
	registerInstantiateCalled               int
	loggerAppRootExpected                   int
	loggerAppRootCalled                     int
	signalNotifyExpected                    int
	signalNotifyCalled                      int
	contextWithTimeoutExpected              int
	contextWithTimeoutCalled                int
	contextBackgroundExpected               int
	contextBackgroundCalled                 int
	createServerFuncExpected                int
	createServerFuncCalled                  int
	listenAndServeFuncExpected              int
	listenAndServeFuncCalled                int
	shutDownFuncExpected                    int
	shutDownFuncCalled                      int
	runServerFuncExpected                   int
	runServerFuncCalled                     int
)

func createMock(t *testing.T) {
	certificateGetServerCertificateExpected = 0
	certificateGetServerCertificateCalled = 0
	certificateGetServerCertificate = func() *tls.Certificate {
		certificateGetServerCertificateCalled++
		return nil
	}
	certificateGetClientCertPoolExpected = 0
	certificateGetClientCertPoolCalled = 0
	certificateGetClientCertPool = func() *x509.CertPool {
		certificateGetClientCertPoolCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	registerInstantiateExpected = 0
	registerInstantiateCalled = 0
	registerInstantiate = func() (*mux.Router, error) {
		registerInstantiateCalled++
		return nil, nil
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	signalNotifyExpected = 0
	signalNotifyCalled = 0
	signalNotify = func(c chan<- os.Signal, sig ...os.Signal) {
		signalNotifyCalled++
	}
	contextWithTimeoutExpected = 0
	contextWithTimeoutCalled = 0
	contextWithTimeout = func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
		contextWithTimeoutCalled++
		return nil, nil
	}
	contextBackgroundExpected = 0
	contextBackgroundCalled = 0
	contextBackground = func() context.Context {
		contextBackgroundCalled++
		return nil
	}
	createServerFuncExpected = 0
	createServerFuncCalled = 0
	createServerFunc = func(serveHTTPS bool, validateClientCert bool, appPort string, router *mux.Router) *http.Server {
		createServerFuncCalled++
		return nil
	}
	listenAndServeFuncExpected = 0
	listenAndServeFuncCalled = 0
	listenAndServeFunc = func(server *http.Server, serveHTTPS bool) error {
		listenAndServeFuncCalled++
		return nil
	}
	shutDownFuncExpected = 0
	shutDownFuncCalled = 0
	shutDownFunc = func(runtimeContext context.Context, server *http.Server) error {
		shutDownFuncCalled++
		return nil
	}
	runServerFuncExpected = 0
	runServerFuncCalled = 0
	runServerFunc = func(serveHTTPS bool, validateClientCert bool, appPort string, router *mux.Router) error {
		runServerFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	certificateGetServerCertificate = certificate.GetServerCertificate
	assert.Equal(t, certificateGetServerCertificateExpected, certificateGetServerCertificateCalled, "Unexpected number of calls to certificateGetServerCertificate")
	certificateGetClientCertPool = certificate.GetClientCertPool
	assert.Equal(t, certificateGetClientCertPoolExpected, certificateGetClientCertPoolCalled, "Unexpected number of calls to certificateGetClientCertPool")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	registerInstantiate = register.Instantiate
	assert.Equal(t, registerInstantiateExpected, registerInstantiateCalled, "Unexpected number of calls to registerInstantiate")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	signalNotify = signal.Notify
	assert.Equal(t, signalNotifyExpected, signalNotifyCalled, "Unexpected number of calls to signalNotify")
	contextWithTimeout = context.WithTimeout
	assert.Equal(t, contextWithTimeoutExpected, contextWithTimeoutCalled, "Unexpected number of calls to contextWithTimeout")
	contextBackground = context.Background
	assert.Equal(t, contextBackgroundExpected, contextBackgroundCalled, "Unexpected number of calls to contextBackground")
	createServerFunc = createServer
	assert.Equal(t, createServerFuncExpected, createServerFuncCalled, "Unexpected number of calls to createServerFunc")
	listenAndServeFunc = listenAndServe
	assert.Equal(t, listenAndServeFuncExpected, listenAndServeFuncCalled, "Unexpected number of calls to listenAndServeFunc")
	shutDownFunc = shutDown
	assert.Equal(t, shutDownFuncExpected, shutDownFuncCalled, "Unexpected number of calls to shutDownFunc")
	runServerFunc = runServer
	assert.Equal(t, runServerFuncExpected, runServerFuncCalled, "Unexpected number of calls to runServerFunc")
}
