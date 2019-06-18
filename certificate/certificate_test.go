package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestLoadTLSCertificate_ErrorTLSCert(t *testing.T) {
	// arrange
	var dummyCertBytes = []byte("some cert bytes")
	var dummyKeyBytes = []byte("some key bytes")
	var dummyTLSCert = tls.Certificate{}
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to load certificate content"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	tlsX509KeyPairExpected = 1
	tlsX509KeyPair = func(certPEMBlock, keyPEMBlock []byte) (tls.Certificate, error) {
		tlsX509KeyPairCalled++
		assert.Equal(t, dummyCertBytes, certPEMBlock)
		assert.Equal(t, dummyKeyBytes, keyPEMBlock)
		return dummyTLSCert, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	cert, err := loadTLSCertificate(
		dummyCertBytes,
		dummyKeyBytes,
	)

	// assert
	assert.Nil(t, cert)
	assert.NotNil(t, err)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestLoadTLSCertificate_Success(t *testing.T) {
	// arrange
	var dummyCertBytes = []byte("some cert bytes")
	var dummyKeyBytes = []byte("some key bytes")
	var dummyTLSCert = tls.Certificate{}

	// mock
	createMock(t)

	// expect
	tlsX509KeyPairExpected = 1
	tlsX509KeyPair = func(certPEMBlock, keyPEMBlock []byte) (tls.Certificate, error) {
		tlsX509KeyPairCalled++
		assert.Equal(t, dummyCertBytes, certPEMBlock)
		assert.Equal(t, dummyKeyBytes, keyPEMBlock)
		return dummyTLSCert, nil
	}

	// SUT + act
	cert, err := loadTLSCertificate(
		dummyCertBytes,
		dummyKeyBytes,
	)

	// assert
	assert.Equal(t, &dummyTLSCert, cert)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestLoadX509CertPool_ParseError(t *testing.T) {
	// arrange
	var dummyCertBytes []byte
	var dummyMessageFormat = "Failed to parse certificate bytes"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	x509NewCertPoolExpected = 1
	x509NewCertPool = func() *x509.CertPool {
		x509NewCertPoolCalled++
		return x509.NewCertPool()
	}
	appendCertsFromPEMFuncExpected = 1
	appendCertsFromPEMFunc = func(certPool *x509.CertPool, certBytes []byte) bool {
		appendCertsFromPEMFuncCalled++
		return appendCertsFromPEM(certPool, certBytes)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var result, err = loadX509CertPool(
		dummyCertBytes,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestLoadX509CertPool_Success(t *testing.T) {
	// arrange
	var dummyCertBytes = []byte("some cert bytes")
	var dummyCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// expect
	x509NewCertPoolExpected = 1
	x509NewCertPool = func() *x509.CertPool {
		x509NewCertPoolCalled++
		return dummyCertPool
	}
	appendCertsFromPEMFuncExpected = 1
	appendCertsFromPEMFunc = func(certPool *x509.CertPool, certBytes []byte) bool {
		appendCertsFromPEMFuncCalled++
		assert.Equal(t, dummyCertPool, certPool)
		assert.Equal(t, dummyCertBytes, certBytes)
		return true
	}

	// SUT + act
	var result, err = loadX509CertPool(
		dummyCertBytes,
	)

	// assert
	assert.Equal(t, dummyCertPool, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeServerCert_NoServeHTTPS(t *testing.T) {
	// arrange
	var dummyServeHTTPS = false
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"

	// stub
	serverCertificate = &tls.Certificate{}

	// mock
	createMock(t)

	// SUT + act
	err := initializeServerCert(
		dummyServeHTTPS,
		dummyServerCertContent,
		dummyServerKeyContent,
	)

	// assert
	assert.Nil(t, serverCertificate)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeServerCert_ServerCertError(t *testing.T) {
	// arrange
	var dummyServeHTTPS = true
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyServerCert = &tls.Certificate{}
	var dummyServerCertError = errors.New("some server cert error")
	var dummyMessageFormat = "Failed to initialize server certificate"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	serverCertificate = &tls.Certificate{}

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 1
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		assert.Equal(t, []byte(dummyServerCertContent), certBytes)
		assert.Equal(t, []byte(dummyServerKeyContent), keyBytes)
		return dummyServerCert, dummyServerCertError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyServerCertError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	err := initializeServerCert(
		dummyServeHTTPS,
		dummyServerCertContent,
		dummyServerKeyContent,
	)

	// assert
	assert.Nil(t, serverCertificate)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitializeServerCert_Success(t *testing.T) {
	// arrange
	var dummyServeHTTPS = true
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyServerCert = &tls.Certificate{}

	// stub
	serverCertificate = nil

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 1
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		assert.Equal(t, []byte(dummyServerCertContent), certBytes)
		assert.Equal(t, []byte(dummyServerKeyContent), keyBytes)
		return dummyServerCert, nil
	}

	// SUT + act
	err := initializeServerCert(
		dummyServeHTTPS,
		dummyServerCertContent,
		dummyServerKeyContent,
	)

	// assert
	assert.Equal(t, dummyServerCert, serverCertificate)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeCaCertPool_NoValidateClientCert(t *testing.T) {
	// arrange
	var dummyValidateClientCert = false
	var dummyCaCertContent = "some CA cert content"

	// stub
	caCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// SUT + act
	err := initializeCaCertPool(
		dummyValidateClientCert,
		dummyCaCertContent,
	)

	// assert
	assert.Nil(t, caCertPool)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeCaCertPool_CaCertPoolError(t *testing.T) {
	// arrange
	var dummyValidateClientCert = true
	var dummyCaCertContent = "some CA cert content"
	var dummyCaCertPool = &x509.CertPool{}
	var dummyCaCertPoolError = errors.New("some CA cert pool error")
	var dummyMessageFormat = "Failed to initialize CA cert pool"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	caCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// expect
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCaCertContent), certBytes)
		return dummyCaCertPool, dummyCaCertPoolError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyCaCertPoolError, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	err := initializeCaCertPool(
		dummyValidateClientCert,
		dummyCaCertContent,
	)

	// assert
	assert.Nil(t, caCertPool)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitializeCaCertPool_Success(t *testing.T) {
	// arrange
	var dummyValidateClientCert = true
	var dummyCaCertContent = "some CA cert content"
	var dummyCaCertPool = &x509.CertPool{}

	// stub
	caCertPool = nil

	// mock
	createMock(t)

	// expect
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCaCertContent), certBytes)
		return dummyCaCertPool, nil
	}

	// SUT + act
	err := initializeCaCertPool(
		dummyValidateClientCert,
		dummyCaCertContent,
	)

	// assert
	assert.Equal(t, dummyCaCertPool, caCertPool)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitialize_ErrorConsolidated(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyServerCertError = errors.New("some server cert error")
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"
	var dummyCaCertPoolError = errors.New("some ca cert pool error")
	var dummyBaseErrorMessage = "Failed to initialize certificates for application"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	initializeServerCertFuncExpected = 1
	initializeServerCertFunc = func(serveHTTPS bool, serverCertContent string, serverKeyContent string) error {
		initializeServerCertFuncCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		return dummyServerCertError
	}
	initializeCaCertPoolFuncExpected = 1
	initializeCaCertPoolFunc = func(validateClientCert bool, caCertContent string) error {
		initializeCaCertPoolFuncCalled++
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyCaCertContent, caCertContent)
		return dummyCaCertPoolError
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, dummyBaseErrorMessage, baseErrorMessage)
		assert.Equal(t, 2, len(allErrors))
		assert.Equal(t, dummyServerCertError, allErrors[0])
		assert.Equal(t, dummyCaCertPoolError, allErrors[1])
		return dummyAppError
	}

	// SUT + act
	var err = Initialize(
		dummyServeHTTPS,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyValidateClientCert,
		dummyCaCertContent,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetServerCertificate(t *testing.T) {
	// arrange
	var dummyCert = &tls.Certificate{}

	// stub
	serverCertificate = dummyCert

	// SUT + act
	cert := GetServerCertificate()

	// assert
	assert.Equal(t, dummyCert, cert)

	// verify
	verifyAll(t)
}

func TestGetClientCertPool(t *testing.T) {
	// arrange
	var dummyCertPool = &x509.CertPool{}

	// stub
	caCertPool = dummyCertPool

	// SUT + act
	var result = GetClientCertPool()

	// assert
	assert.Equal(t, dummyCertPool, result)

	// verify
	verifyAll(t)
}
