package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
)

var (
	tlsX509KeyPairExpected               int
	tlsX509KeyPairCalled                 int
	x509NewCertPoolExpected              int
	x509NewCertPoolCalled                int
	apperrorGetCustomErrorExpected       int
	apperrorGetCustomErrorCalled         int
	apperrorWrapSimpleErrorExpected      int
	apperrorWrapSimpleErrorCalled        int
	loadTLSCertificateFuncExpected       int
	loadTLSCertificateFuncCalled         int
	appendCertsFromPEMFuncExpected       int
	appendCertsFromPEMFuncCalled         int
	loadX509CertPoolFuncExpected         int
	loadX509CertPoolFuncCalled           int
	initializeTLSCertiticateFuncExpected int
	initializeTLSCertiticateFuncCalled   int
	initializeX509CertPoolFuncExpected   int
	initializeX509CertPoolFuncCalled     int
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
	apperrorGetCustomErrorExpected = 0
	apperrorGetCustomErrorCalled = 0
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
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
	initializeTLSCertiticateFuncExpected = 0
	initializeTLSCertiticateFuncCalled = 0
	initializeTLSCertiticateFunc = func(shouldLoadCert bool, certContent string, keyContent string) (*tls.Certificate, error) {
		initializeTLSCertiticateFuncCalled++
		return nil, nil
	}
	initializeX509CertPoolFuncExpected = 0
	initializeX509CertPoolFuncCalled = 0
	initializeX509CertPoolFunc = func(shouldLoadCert bool, certContent string) (*x509.CertPool, error) {
		initializeX509CertPoolFuncCalled++
		return nil, nil
	}
}

func verifyAll(t *testing.T) {
	tlsX509KeyPair = tls.X509KeyPair
	assert.Equal(t, tlsX509KeyPairExpected, tlsX509KeyPairCalled, "Unexpected number of calls to tlsX509KeyPair")
	x509NewCertPool = x509.NewCertPool
	assert.Equal(t, x509NewCertPoolExpected, x509NewCertPoolCalled, "Unexpected number of calls to x509NewCertPool")
	apperrorGetCustomError = apperror.GetCustomError
	assert.Equal(t, apperrorGetCustomErrorExpected, apperrorGetCustomErrorCalled, "Unexpected number of calls to apperrorGetCustomError")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	loadTLSCertificateFunc = loadTLSCertificate
	assert.Equal(t, loadTLSCertificateFuncExpected, loadTLSCertificateFuncCalled, "Unexpected number of calls to loadTLSCertificateFunc")
	appendCertsFromPEMFunc = appendCertsFromPEM
	assert.Equal(t, appendCertsFromPEMFuncExpected, appendCertsFromPEMFuncCalled, "Unexpected number of calls to appendCertsFromPEMFunc")
	loadX509CertPoolFunc = loadX509CertPool
	assert.Equal(t, loadX509CertPoolFuncExpected, loadX509CertPoolFuncCalled, "Unexpected number of calls to loadX509CertPoolFunc")
	initializeTLSCertiticateFunc = initializeTLSCertiticate
	assert.Equal(t, initializeTLSCertiticateFuncExpected, initializeTLSCertiticateFuncCalled, "Unexpected number of calls to initializeTLSCertiticateFunc")
	initializeX509CertPoolFunc = initializeX509CertPool
	assert.Equal(t, initializeX509CertPoolFuncExpected, initializeX509CertPoolFuncCalled, "Unexpected number of calls to initializeX509CertPoolFunc")

	serverCertificate = nil
	caCertPool = nil
	clientCertificate = nil
}
