package main

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/stretchr/testify/assert"
)

func TestBootstrapApplication_ConfigError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to bootstrap application for configuration"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return dummyError
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
	var err = bootstrapApplication()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestBootstrapApplication_CertError(t *testing.T) {
	// arrange
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"
	var dummyError = errors.New("some error message")
	var dummyMessageFormat = "Failed to bootstrap application for certificates"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return nil
	}
	configSendClientCertExpected = 1
	configSendClientCert = func() bool {
		configSendClientCertCalled++
		return dummySendClientCert
	}
	configClientCertContentExpected = 1
	configClientCertContent = func() string {
		configClientCertContentCalled++
		return dummyClientCertContent
	}
	configClientKeyContentExpected = 1
	configClientKeyContent = func() string {
		configClientKeyContentCalled++
		return dummyClientKeyContent
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configServerCertContentExpected = 1
	configServerCertContent = func() string {
		configServerCertContentCalled++
		return dummyServerCertContent
	}
	configServerKeyContentExpected = 1
	configServerKeyContent = func() string {
		configServerKeyContentCalled++
		return dummyServerKeyContent
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	configCaCertContentExpected = 1
	configCaCertContent = func() string {
		configCaCertContentCalled++
		return dummyCaCertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(sendClientCert bool, clientCertContent string, clientKeyContent string, serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummySendClientCert, sendClientCert)
		assert.Equal(t, dummyClientCertContent, clientCertContent)
		assert.Equal(t, dummyClientKeyContent, clientKeyContent)
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyCaCertContent, caCertContent)
		return dummyError
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
	var err = bootstrapApplication()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestBootstrapApplication_Success(t *testing.T) {
	// arrange
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"

	// mock
	createMock(t)

	// expect
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return nil
	}
	configSendClientCertExpected = 1
	configSendClientCert = func() bool {
		configSendClientCertCalled++
		return dummySendClientCert
	}
	configClientCertContentExpected = 1
	configClientCertContent = func() string {
		configClientCertContentCalled++
		return dummyClientCertContent
	}
	configClientKeyContentExpected = 1
	configClientKeyContent = func() string {
		configClientKeyContentCalled++
		return dummyClientKeyContent
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configServerCertContentExpected = 1
	configServerCertContent = func() string {
		configServerCertContentCalled++
		return dummyServerCertContent
	}
	configServerKeyContentExpected = 1
	configServerKeyContent = func() string {
		configServerKeyContentCalled++
		return dummyServerKeyContent
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	configCaCertContentExpected = 1
	configCaCertContent = func() string {
		configCaCertContentCalled++
		return dummyCaCertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(sendClientCert bool, clientCertContent string, clientKeyContent string, serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummySendClientCert, sendClientCert)
		assert.Equal(t, dummyClientCertContent, clientCertContent)
		assert.Equal(t, dummyClientKeyContent, clientKeyContent)
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyCaCertContent, caCertContent)
		return nil
	}

	// SUT + act
	var err = bootstrapApplication()

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestConnectStorages(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var result = connectStorages()

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestDisconnectStorages(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var result = disconnectStorages()

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestMain_FailBootstrapServer(t *testing.T) {
	// arrange
	var dummyError = errors.New("some dummy error message")

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return dummyError
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		assert.Equal(t, "bootstrapApplicationFunc", subcategory)
		assert.Equal(t, "Failed to initialize server due to %v.", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}

func TestMain_FailConnectStorage(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyError = errors.New("dummy db error")

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 1
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return dummyError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Started server (v-%v) on port %v.", messageFormat)
			assert.Equal(t, "applicationStart", subcategory)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
			assert.Equal(t, dummyAppPort, parameters[1])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Failed to initialize server due to %v.", messageFormat)
			assert.Equal(t, "connectStorages", subcategory)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyError, parameters[0])
		}
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}

func TestMain_ErrorTerminateFailStorageDisconnect(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyError = errors.New("dummy final error")
	var dummyDBError = errors.New("dummy db error")

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 1
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return nil
	}
	disconnectStoragesFuncExpected = 1
	disconnectStoragesFunc = func() error {
		disconnectStoragesFuncCalled++
		return dummyDBError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	loggerAppRootExpected = 3
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Started server (v-%v) on port %v.", messageFormat)
			assert.Equal(t, "applicationStart", subcategory)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
			assert.Equal(t, dummyAppPort, parameters[1])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Stopped server due to %v.", messageFormat)
			assert.Equal(t, "applicationStop", subcategory)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyError, parameters[0])
		} else if loggerAppRootCalled == 3 {
			assert.Equal(t, "Failed to terminate server cleanly due to %v.", messageFormat)
			assert.Equal(t, "disconnectStorages", subcategory)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyDBError, parameters[0])
		}
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return dummyError
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}

func TestMain_ErrorTerminateFully(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyError = errors.New("dummy final error")

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 1
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return nil
	}
	disconnectStoragesFuncExpected = 1
	disconnectStoragesFunc = func() error {
		disconnectStoragesFuncCalled++
		return nil
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Started server (v-%v) on port %v.", messageFormat)
			assert.Equal(t, "applicationStart", subcategory)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
			assert.Equal(t, dummyAppPort, parameters[1])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Stopped server due to %v.", messageFormat)
			assert.Equal(t, "applicationStop", subcategory)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyError, parameters[0])
		}
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return dummyError
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}

func TestMain_PeaceTerminateFailStorageDisconnect(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyError = errors.New("dummy final error")

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 1
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return nil
	}
	disconnectStoragesFuncExpected = 1
	disconnectStoragesFunc = func() error {
		disconnectStoragesFuncCalled++
		return dummyError
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	loggerAppRootExpected = 3
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Started server (v-%v) on port %v.", messageFormat)
			assert.Equal(t, "applicationStart", subcategory)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
			assert.Equal(t, dummyAppPort, parameters[1])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Stopped server peacefully.", messageFormat)
			assert.Equal(t, "applicationStop", subcategory)
			assert.Equal(t, 0, len(parameters))
		} else if loggerAppRootCalled == 3 {
			assert.Equal(t, "Failed to terminate server cleanly due to %v.", messageFormat)
			assert.Equal(t, "disconnectStorages", subcategory)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyError, parameters[0])
		}
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return nil
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}

func TestMain_PeaceTerminateFull(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 1
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return nil
	}
	disconnectStoragesFuncExpected = 1
	disconnectStoragesFunc = func() error {
		disconnectStoragesFuncCalled++
		return nil
	}
	configAppPortExpected = 1
	configAppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Started server (v-%v) on port %v.", messageFormat)
			assert.Equal(t, "applicationStart", subcategory)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
			assert.Equal(t, dummyAppPort, parameters[1])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Stopped server peacefully.", messageFormat)
			assert.Equal(t, "applicationStop", subcategory)
			assert.Equal(t, 0, len(parameters))
		}
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return nil
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}
