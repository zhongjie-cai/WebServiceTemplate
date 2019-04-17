package certificate

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
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
	var expectedErrorMessage = "Failed to load certificate content"
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
		assert.Equal(t, expectedErrorMessage, messageFormat)
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
	assert.Nil(t, clientCertificate)

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
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestLoadX509CertPool_ParseError(t *testing.T) {
	// arrange
	var dummyCertBytes []byte
	var expectedErrorMessage = "Failed to parse certificate bytes"
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
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
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
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertificate_Failure(t *testing.T) {
	// arrange
	var expectedErrorMessage = "Client certificate not initialized"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	clientCertificate = nil

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	cert, err := GetClientCertificate()

	// assert
	assert.Nil(t, cert)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertificate_Success(t *testing.T) {
	// arrange
	var dummyCert = &tls.Certificate{}

	// stub
	clientCertificate = dummyCert

	// SUT + act
	cert, err := GetClientCertificate()

	// assert
	assert.Equal(t, dummyCert, cert)
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestGetServerCertificate_Failure(t *testing.T) {
	// arrange
	var expectedErrorMessage = "Server certificate not initialized"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	serverCertificate = nil

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	cert, err := GetServerCertificate()

	// assert
	assert.Nil(t, cert)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetServerCertificate_Success(t *testing.T) {
	// arrange
	var dummyCert = &tls.Certificate{}

	// stub
	serverCertificate = dummyCert

	// SUT + act
	cert, err := GetServerCertificate()

	// assert
	assert.Equal(t, dummyCert, cert)
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertPool_Failure(t *testing.T) {
	// arrange
	var expectedErrorMessage = "CA cert pool not initialized"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	caCertPool = nil

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var result, err = GetClientCertPool()

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGetClientCertPool_Success(t *testing.T) {
	// arrange
	var dummyCertPool = &x509.CertPool{}

	// stub
	caCertPool = dummyCertPool

	// SUT + act
	var result, err = GetClientCertPool()

	// assert
	assert.Equal(t, dummyCertPool, result)
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestInitialize_ClientCertError(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyClientCert = &tls.Certificate{}
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to initialize client certificate"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 1
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		assert.Equal(t, []byte(dummyClientCertContent), certBytes)
		assert.Equal(t, []byte(dummyClientKeyContent), keyBytes)
		return dummyClientCert, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = Initialize(
		dummyClientCertContent,
		dummyClientKeyContent,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyCACertContent,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_ServerCertError(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyClientCert = &tls.Certificate{}
	var dummyServerCert = &tls.Certificate{}
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to initialize server certificate"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 2
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		if loadTLSCertificateFuncCalled == 1 {
			assert.Equal(t, []byte(dummyClientCertContent), certBytes)
			assert.Equal(t, []byte(dummyClientKeyContent), keyBytes)
			return dummyClientCert, nil
		} else if loadTLSCertificateFuncCalled == 2 {
			assert.Equal(t, []byte(dummyServerCertContent), certBytes)
			assert.Equal(t, []byte(dummyServerKeyContent), keyBytes)
			return dummyServerCert, dummyError
		}
		return nil, nil
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = Initialize(
		dummyClientCertContent,
		dummyClientKeyContent,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyCACertContent,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyClientCert, clientCertificate)

	// verify
	verifyAll(t)
}

func TestInitialize_CertPoolError(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyClientCert = &tls.Certificate{}
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to initialize CA cert pool"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 2
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		if loadTLSCertificateFuncCalled == 1 {
			assert.Equal(t, []byte(dummyClientCertContent), certBytes)
			assert.Equal(t, []byte(dummyClientKeyContent), keyBytes)
			return dummyClientCert, nil
		} else if loadTLSCertificateFuncCalled == 2 {
			assert.Equal(t, []byte(dummyServerCertContent), certBytes)
			assert.Equal(t, []byte(dummyServerKeyContent), keyBytes)
			return dummyServerCert, nil
		}
		return nil, nil
	}
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCACertContent), certBytes)
		return dummyCertPool, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = Initialize(
		dummyClientCertContent,
		dummyClientKeyContent,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyCACertContent,
	)

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Equal(t, dummyClientCert, clientCertificate)
	assert.Equal(t, dummyServerCert, serverCertificate)

	// verify
	verifyAll(t)
}

func TestInitialize_Success(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyClientCert = &tls.Certificate{}
	var dummyServerCert = &tls.Certificate{}
	var dummyCertPool = &x509.CertPool{}

	// mock
	createMock(t)

	// expect
	loadTLSCertificateFuncExpected = 2
	loadTLSCertificateFunc = func(certBytes, keyBytes []byte) (*tls.Certificate, error) {
		loadTLSCertificateFuncCalled++
		if loadTLSCertificateFuncCalled == 1 {
			assert.Equal(t, []byte(dummyClientCertContent), certBytes)
			assert.Equal(t, []byte(dummyClientKeyContent), keyBytes)
			return dummyClientCert, nil
		} else if loadTLSCertificateFuncCalled == 2 {
			assert.Equal(t, []byte(dummyServerCertContent), certBytes)
			assert.Equal(t, []byte(dummyServerKeyContent), keyBytes)
			return dummyServerCert, nil
		}
		return nil, nil
	}
	loadX509CertPoolFuncExpected = 1
	loadX509CertPoolFunc = func(certBytes []byte) (*x509.CertPool, error) {
		loadX509CertPoolFuncCalled++
		assert.Equal(t, []byte(dummyCACertContent), certBytes)
		return dummyCertPool, nil
	}

	// SUT + act
	var err = Initialize(
		dummyClientCertContent,
		dummyClientKeyContent,
		dummyServerCertContent,
		dummyServerKeyContent,
		dummyCACertContent,
	)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, dummyClientCert, clientCertificate)
	assert.Equal(t, dummyServerCert, serverCertificate)
	assert.Equal(t, dummyCertPool, caCertPool)

	// verify
	verifyAll(t)
}
