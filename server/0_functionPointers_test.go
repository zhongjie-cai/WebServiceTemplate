package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
)

var (
	configAppPortExpected                   int
	configAppPortCalled                     int
	certificateGetServerCertificateExpected int
	certificateGetServerCertificateCalled   int
	certificateGetClientCertPoolExpected    int
	certificateGetClientCertPoolCalled      int
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	faviconHostEntryExpected                int
	faviconHostEntryCalled                  int
	swaggerHostEntryExpected                int
	swaggerHostEntryCalled                  int
	healthHostEntryExpected                 int
	healthHostEntryCalled                   int
	createServerFuncExpected                int
	createServerFuncCalled                  int
	listenAndServeTLSFuncExpected           int
	listenAndServeTLSFuncCalled             int
	hostEntriesFuncExpected                 int
	hostEntriesFuncCalled                   int
	runServerFuncExpected                   int
	runServerFuncCalled                     int
)

func createMock(t *testing.T) {
	configAppPortExpected = 0
	configAppPortCalled = 0
	configAppPort = func() string {
		configAppPortCalled++
		return ""
	}
	certificateGetServerCertificateExpected = 0
	certificateGetServerCertificateCalled = 0
	certificateGetServerCertificate = func() (*tls.Certificate, error) {
		certificateGetServerCertificateCalled++
		return nil, nil
	}
	certificateGetClientCertPoolExpected = 0
	certificateGetClientCertPoolCalled = 0
	certificateGetClientCertPool = func() (*x509.CertPool, error) {
		certificateGetClientCertPoolCalled++
		return nil, nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	faviconHostEntryExpected = 0
	faviconHostEntryCalled = 0
	faviconHostEntry = func() {
		faviconHostEntryCalled++
	}
	swaggerHostEntryExpected = 0
	swaggerHostEntryCalled = 0
	swaggerHostEntry = func() {
		swaggerHostEntryCalled++
	}
	healthHostEntryExpected = 0
	healthHostEntryCalled = 0
	healthHostEntry = func() {
		healthHostEntryCalled++
	}
	createServerFuncExpected = 0
	createServerFuncCalled = 0
	createServerFunc = func(serverCert *tls.Certificate, clientCertPool *x509.CertPool) *http.Server {
		createServerFuncCalled++
		return nil
	}
	listenAndServeTLSFuncExpected = 0
	listenAndServeTLSFuncCalled = 0
	listenAndServeTLSFunc = func(server *http.Server) error {
		listenAndServeTLSFuncCalled++
		return nil
	}
	hostEntriesFuncExpected = 0
	hostEntriesFuncCalled = 0
	hostEntriesFunc = func(entryFuncs ...func()) error {
		hostEntriesFuncCalled++
		return nil
	}
	runServerFuncExpected = 0
	runServerFuncCalled = 0
	runServerFunc = func() error {
		runServerFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	configAppPort = config.AppPort
	if configAppPortExpected != configAppPortCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppPort, expected %v, actual %v", configAppPortExpected, configAppPortCalled))
	}
	certificateGetServerCertificate = certificate.GetServerCertificate
	if certificateGetServerCertificateExpected != certificateGetServerCertificateCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to certificateGetServerCertificate, expected %v, actual %v", certificateGetServerCertificateExpected, certificateGetServerCertificateCalled))
	}
	certificateGetClientCertPool = certificate.GetClientCertPool
	if certificateGetClientCertPoolExpected != certificateGetClientCertPoolCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to certificateGetClientCertPool, expected %v, actual %v", certificateGetClientCertPoolExpected, certificateGetClientCertPoolCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
	faviconHostEntry = favicon.HostEntry
	if faviconHostEntryExpected != faviconHostEntryCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to faviconHostEntry, expected %v, actual %v", faviconHostEntryExpected, faviconHostEntryCalled))
	}
	swaggerHostEntry = swagger.HostEntry
	if swaggerHostEntryExpected != swaggerHostEntryCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to swaggerHostEntry, expected %v, actual %v", swaggerHostEntryExpected, swaggerHostEntryCalled))
	}
	healthHostEntry = health.HostEntry
	if healthHostEntryExpected != healthHostEntryCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to healthHostEntry, expected %v, actual %v", healthHostEntryExpected, healthHostEntryCalled))
	}
	createServerFunc = createServer
	if createServerFuncExpected != createServerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to createServerFunc, expected %v, actual %v", createServerFuncExpected, createServerFuncCalled))
	}
	listenAndServeTLSFunc = listenAndServeTLS
	if listenAndServeTLSFuncExpected != listenAndServeTLSFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to listenAndServeTLSFunc, expected %v, actual %v", listenAndServeTLSFuncExpected, listenAndServeTLSFuncCalled))
	}
	hostEntriesFunc = hostEntries
	if hostEntriesFuncExpected != hostEntriesFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to hostEntriesFunc, expected %v, actual %v", hostEntriesFuncExpected, hostEntriesFuncCalled))
	}
	runServerFunc = runServer
	if runServerFuncExpected != runServerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to runServerFunc, expected %v, actual %v", runServerFuncExpected, runServerFuncCalled))
	}
}
