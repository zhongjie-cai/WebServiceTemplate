package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

var (
	tlsX509KeyPairExpected               int
	tlsX509KeyPairCalled                 int
	x509NewCertPoolExpected              int
	x509NewCertPoolCalled                int
	apperrorWrapSimpleErrorExpected      int
	apperrorWrapSimpleErrorCalled        int
	apperrorConsolidateAllErrorsExpected int
	apperrorConsolidateAllErrorsCalled   int
	loadTLSCertificateFuncExpected       int
	loadTLSCertificateFuncCalled         int
	appendCertsFromPEMFuncExpected       int
	appendCertsFromPEMFuncCalled         int
	loadX509CertPoolFuncExpected         int
	loadX509CertPoolFuncCalled           int
	initializeServerCertFuncExpected     int
	initializeServerCertFuncCalled       int
	initializeCaCertPoolFuncExpected     int
	initializeCaCertPoolFuncCalled       int
)

func createMock(t *testing.T) {
	tlsX509KeyPairExpected = 0
	tlsX509KeyPairCalled = 0
	tlsX509KeyPair = func(certPEMBlock, keyPEMBlock []byte) (tls.Certificate, error) {
		tlsX509KeyPairCalled++
		return tls.Certificate{}, nil
	}
	x509NewCertPoolExpected = 0
	x509NewCertPoolCalled = 0
	x509NewCertPool = func() *x509.CertPool {
		x509NewCertPoolCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	apperrorConsolidateAllErrorsExpected = 0
	apperrorConsolidateAllErrorsCalled = 0
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		return nil
	}
	loadTLSCertificateFuncExpected = 0
	loadTLSCertificateFuncCalled = 0
	loadTLSCertificateFunc = func(certFile, keyFile []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		return nil, nil
	}
	appendCertsFromPEMFuncExpected = 0
	appendCertsFromPEMFuncCalled = 0
	appendCertsFromPEMFunc = func(certPool *x509.CertPool, certBytes []byte) bool {
		appendCertsFromPEMFuncCalled++
		return false
	}
	loadX509CertPoolFuncExpected = 0
	loadX509CertPoolFuncCalled = 0
	loadX509CertPoolFunc = func(certFile []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		return nil, nil
	}
	initializeServerCertFuncExpected = 0
	initializeServerCertFuncCalled = 0
	initializeServerCertFunc = func(serveHTTPS bool, serverCertContent string, serverKeyContent string) error {
		initializeServerCertFuncCalled++
		return nil
	}
	initializeCaCertPoolFuncExpected = 0
	initializeCaCertPoolFuncCalled = 0
	initializeCaCertPoolFunc = func(validateClientCert bool, caCertContent string) error {
		initializeCaCertPoolFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	tlsX509KeyPair = tls.X509KeyPair
	assert.Equal(t, tlsX509KeyPairExpected, tlsX509KeyPairCalled, "Unexpected number of calls to tlsX509KeyPair")
	x509NewCertPool = x509.NewCertPool
	assert.Equal(t, x509NewCertPoolExpected, x509NewCertPoolCalled, "Unexpected number of calls to x509NewCertPool")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	assert.Equal(t, apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled, "Unexpected number of calls to apperrorConsolidateAllErrors")
	loadTLSCertificateFunc = loadTLSCertificate
	assert.Equal(t, loadTLSCertificateFuncExpected, loadTLSCertificateFuncCalled, "Unexpected number of calls to loadTLSCertificateFunc")
	appendCertsFromPEMFunc = appendCertsFromPEM
	assert.Equal(t, appendCertsFromPEMFuncExpected, appendCertsFromPEMFuncCalled, "Unexpected number of calls to appendCertsFromPEMFunc")
	loadX509CertPoolFunc = loadX509CertPool
	assert.Equal(t, loadX509CertPoolFuncExpected, loadX509CertPoolFuncCalled, "Unexpected number of calls to loadX509CertPoolFunc")
	initializeServerCertFunc = initializeServerCert
	assert.Equal(t, initializeServerCertFuncExpected, initializeServerCertFuncCalled, "Unexpected number of calls to initializeServerCertFunc")
	initializeCaCertPoolFunc = initializeCaCertPool
	assert.Equal(t, initializeCaCertPoolFuncExpected, initializeCaCertPoolFuncCalled, "Unexpected number of calls to initializeCaCertPoolFunc")

	serverCertificate = nil
	caCertPool = nil
}