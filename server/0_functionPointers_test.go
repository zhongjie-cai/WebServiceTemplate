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
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

var (
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
	routeRegisterEntriesExpected            int
	routeRegisterEntriesCalled              int
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
	faviconHostEntryExpected = 0
	faviconHostEntryCalled = 0
	faviconHostEntry = func(router *mux.Router) {
		faviconHostEntryCalled++
	}
	swaggerHostEntryExpected = 0
	swaggerHostEntryCalled = 0
	swaggerHostEntry = func(router *mux.Router) {
		swaggerHostEntryCalled++
	}
	healthHostEntryExpected = 0
	healthHostEntryCalled = 0
	healthHostEntry = func(router *mux.Router) {
		healthHostEntryCalled++
	}
	routeRegisterEntriesExpected = 0
	routeRegisterEntriesCalled = 0
	routeRegisterEntries = func(entryFuncs ...func(*mux.Router)) (*mux.Router, error) {
		routeRegisterEntriesCalled++
		return nil, nil
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
	assert.Equal(t, certificateGetServerCertificateExpected, certificateGetServerCertificateCalled, "Unexpected method call to certificateGetServerCertificate")
	certificateGetClientCertPool = certificate.GetClientCertPool
	assert.Equal(t, certificateGetClientCertPoolExpected, certificateGetClientCertPoolCalled, "Unexpected method call to certificateGetClientCertPool")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
	faviconHostEntry = favicon.HostEntry
	assert.Equal(t, faviconHostEntryExpected, faviconHostEntryCalled, "Unexpected method call to faviconHostEntry")
	swaggerHostEntry = swagger.HostEntry
	assert.Equal(t, swaggerHostEntryExpected, swaggerHostEntryCalled, "Unexpected method call to swaggerHostEntry")
	healthHostEntry = health.HostEntry
	assert.Equal(t, healthHostEntryExpected, healthHostEntryCalled, "Unexpected method call to healthHostEntry")
	routeRegisterEntries = route.RegisterEntries
	assert.Equal(t, routeRegisterEntriesExpected, routeRegisterEntriesCalled, "Unexpected method call to routeRegisterEntries")
	createServerFunc = createServer
	assert.Equal(t, createServerFuncExpected, createServerFuncCalled, "Unexpected method call to createServerFunc")
	listenAndServeFunc = listenAndServe
	assert.Equal(t, listenAndServeFuncExpected, listenAndServeFuncCalled, "Unexpected method call to listenAndServeFunc")
	runServerFunc = runServer
	assert.Equal(t, runServerFuncExpected, runServerFuncCalled, "Unexpected method call to runServerFunc")
}
