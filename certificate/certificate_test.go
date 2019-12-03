package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
)

func TestLoadTLSCertificate_ErrorTLSCert(t *testing.T) {
	// arrange
	var dummyCertBytes = []byte("some cert bytes")
	var dummyKeyBytes = []byte("some key bytes")
	var dummyTLSCert = tls.Certificate{}
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to load certificate content"
	var dummyAppError = apperror.GetCustomError(0, "")

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
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var cert, err = loadTLSCertificate(
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
	var cert, err = loadTLSCertificate(
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
	var dummyAppError = apperror.GetCustomError(0, "")

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
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
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

func TestInitializeTLSCertiticate_NoServeHTTPS(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = false
	var dummyCertContent = "some cert content"
	var dummyKeyContent = "some key content"

	// mock
	createMock(t)

	// SUT + act
	var result, err = initializeTLSCertiticate(
		dummyShouldLoadCert,
		dummyCertContent,
		dummyKeyContent,
	)

	// assert
	assert.Nil(t, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeTLSCertiticate_ServerCertError(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = true
	var dummyCertContent = "some cert content"
	var dummyKeyContent = "some key content"
	var dummyCert = &tls.Certificate{}
	var dummyCertError = errors.New("some cert error")
	var dummyMessageFormat = "Failed to initialize certificate by key-cert pair"
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 1
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		assert.Equal(t, []byte(dummyCertContent), certBytes)
		assert.Equal(t, []byte(dummyKeyContent), keyBytes)
		return dummyCert, dummyCertError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyCertError, innerErrors[0])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var result, err = initializeTLSCertiticate(
		dummyShouldLoadCert,
		dummyCertContent,
		dummyKeyContent,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitializeTLSCertiticate_Success(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = true
	var dummyCertContent = "some cert content"
	var dummyKeyContent = "some key content"
	var dummyCert = &tls.Certificate{}

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 1
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		assert.Equal(t, []byte(dummyCertContent), certBytes)
		assert.Equal(t, []byte(dummyKeyContent), keyBytes)
		return dummyCert, nil
	}

	// SUT + act
	var result, err = initializeTLSCertiticate(
		dummyShouldLoadCert,
		dummyCertContent,
		dummyKeyContent,
	)

	// assert
	assert.Equal(t, dummyCert, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeX509CertPool_NoValidateClientCert(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = false
	var dummyCertContent = "some cert content"

	// mock
	createMock(t)

	// SUT + act
	var result, err = initializeX509CertPool(
		dummyShouldLoadCert,
		dummyCertContent,
	)

	// assert
	assert.Nil(t, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitializeX509CertPool_CaCertPoolError(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = true
	var dummyCertContent = "some cert content"
	var dummyCertPool = &x509.CertPool{}
	var dummyCertPoolError = errors.New("some cert pool error")
	var dummyMessageFormat = "Failed to initialize cert pool by cert content"
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCertContent), certBytes)
		return dummyCertPool, dummyCertPoolError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyCertPoolError, innerErrors[0])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var result, err = initializeX509CertPool(
		dummyShouldLoadCert,
		dummyCertContent,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitializeX509CertPool_Success(t *testing.T) {
	// arrange
	var dummyShouldLoadCert = true
	var dummyCertContent = "some cert content"
	var dummyCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// expect
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCertContent), certBytes)
		return dummyCertPool, nil
	}

	// SUT + act
	var result, err = initializeX509CertPool(
		dummyShouldLoadCert,
		dummyCertContent,
	)

	// assert
	assert.Equal(t, dummyCertPool, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitialize_ErrorConsolidated(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyServerCert = &tls.Certificate{}
	var dummyServerCertError = errors.New("some server cert error")
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"
	var dummyCaCertPool = &x509.CertPool{}
	var dummyCaCertPoolError = errors.New("some ca cert pool error")
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyClientCert = &tls.Certificate{}
	var dummyClientCertError = errors.New("some client cert error")
	var dummyMessageFormat = "Failed to initialize certificates for application"
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	initializeTLSCertiticateFuncExpected = 2
	initializeTLSCertiticateFunc = func(shouldLoadCert bool, certContent string, keyContent string) (*tls.Certificate, error) {
		initializeTLSCertiticateFuncCalled++
		if initializeTLSCertiticateFuncCalled == 1 {
			assert.Equal(t, dummyServeHTTPS, shouldLoadCert)
			assert.Equal(t, dummyServerCertContent, certContent)
			assert.Equal(t, dummyServerKeyContent, keyContent)
			return dummyServerCert, dummyServerCertError
		} else if initializeTLSCertiticateFuncCalled == 2 {
			assert.Equal(t, dummySendClientCert, shouldLoadCert)
			assert.Equal(t, dummyClientCertContent, certContent)
			assert.Equal(t, dummyClientKeyContent, keyContent)
			return dummyClientCert, dummyClientCertError
		}
		return nil, nil
	}
	initializeX509CertPoolFuncExpected = 1
	initializeX509CertPoolFunc = func(shouldLoadCert bool, certContent string) (*x509.CertPool, error) {
		initializeX509CertPoolFuncCalled++
		assert.Equal(t, dummyValidateClientCert, shouldLoadCert)
		assert.Equal(t, dummyCaCertContent, certContent)
		return dummyCaCertPool, dummyCaCertPoolError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 3, len(innerErrors))
		assert.Equal(t, dummyServerCertError, innerErrors[0])
		assert.Equal(t, dummyCaCertPoolError, innerErrors[1])
		assert.Equal(t, dummyClientCertError, innerErrors[2])
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = Initialize(
		dummyServeHTTPS,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyValidateClientCert,
		dummyCaCertContent,
		dummySendClientCert,
		dummyClientCertContent,
		dummyClientKeyContent,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyServerCert, serverCertificate)
	assert.Equal(t, dummyCaCertPool, caCertPool)
	assert.Equal(t, dummyClientCert, clientCertificate)

	// verify
	verifyAll(t)
}

func TestGetServerCertificate(t *testing.T) {
	// arrange
	var dummyCert = &tls.Certificate{}

	// stub
	serverCertificate = dummyCert

	// SUT + act
	var cert = GetServerCertificate()

	// assert
	assert.Equal(t, dummyCert, cert)

	// verify
	verifyAll(t)
}

func TestGetCaCertPool(t *testing.T) {
	// arrange
	var dummyCertPool = &x509.CertPool{}

	// stub
	caCertPool = dummyCertPool

	// SUT + act
	var result = GetCaCertPool()

	// assert
	assert.Equal(t, dummyCertPool, result)

	// verify
	verifyAll(t)
}

func TestGetClientCertificate(t *testing.T) {
	// arrange
	var dummyCert = &tls.Certificate{}

	// stub
	clientCertificate = dummyCert

	// SUT + act
	var cert = GetClientCertificate()

	// assert
	assert.Equal(t, dummyCert, cert)

	// verify
	verifyAll(t)
}
