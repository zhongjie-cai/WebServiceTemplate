package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

var (
	tlsX509KeyPairExpected          int
	tlsX509KeyPairCalled            int
	x509NewCertPoolExpected         int
	x509NewCertPoolCalled           int
	apperrorWrapSimpleErrorExpected int
	apperrorWrapSimpleErrorCalled   int
	loadTLSCertificateFuncExpected  int
	loadTLSCertificateFuncCalled    int
	appendCertsFromPEMFuncExpected  int
	appendCertsFromPEMFuncCalled    int
	loadX509CertPoolFuncExpected    int
	loadX509CertPoolFuncCalled      int
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
}

func verifyAll(t *testing.T) {
	tlsX509KeyPair = tls.X509KeyPair
	if tlsX509KeyPairExpected != tlsX509KeyPairCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to tlsX509KeyPair, expected %v, actual %v", tlsX509KeyPairExpected, tlsX509KeyPairCalled))
	}
	x509NewCertPool = x509.NewCertPool
	if x509NewCertPoolExpected != x509NewCertPoolCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to x509NewCertPool, expected %v, actual %v", x509NewCertPoolExpected, x509NewCertPoolCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
	loadTLSCertificateFunc = loadTLSCertificate
	if loadTLSCertificateFuncExpected != loadTLSCertificateFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loadTLSCertificateFunc, expected %v, actual %v", loadTLSCertificateFuncExpected, loadTLSCertificateFuncCalled))
	}
	appendCertsFromPEMFunc = appendCertsFromPEM
	if appendCertsFromPEMFuncExpected != appendCertsFromPEMFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to appendCertsFromPEMFunc, expected %v, actual %v", appendCertsFromPEMFuncExpected, appendCertsFromPEMFuncCalled))
	}
	loadX509CertPoolFunc = loadX509CertPool
	if loadX509CertPoolFuncExpected != loadX509CertPoolFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loadX509CertPoolFunc, expected %v, actual %v", loadX509CertPoolFuncExpected, loadX509CertPoolFuncCalled))
	}

	clientCertificate = nil
	serverCertificate = nil
	caCertPool = nil
}
