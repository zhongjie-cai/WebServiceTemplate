package server

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
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
	createServerFuncExpected                int
	createServerFuncCalled                  int
	listenAndServeFuncExpected              int
	listenAndServeFuncCalled                int
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
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
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
	createServerFunc = createServer
	assert.Equal(t, createServerFuncExpected, createServerFuncCalled, "Unexpected number of calls to createServerFunc")
	listenAndServeFunc = listenAndServe
	assert.Equal(t, listenAndServeFuncExpected, listenAndServeFuncCalled, "Unexpected number of calls to listenAndServeFunc")
	runServerFunc = runServer
	assert.Equal(t, runServerFuncExpected, runServerFuncCalled, "Unexpected number of calls to runServerFunc")
}
