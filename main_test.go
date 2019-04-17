package main

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/stretchr/testify/assert"
)

func TestBootstrapApplication_ConfigError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to bootstrap application for configuration"
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
		assert.Equal(t, expectedErrorMessage, messageFormat)
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
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to bootstrap application for certificates"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return nil
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
	configCACertContentExpected = 1
	configCACertContent = func() string {
		configCACertContentCalled++
		return dummyCACertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(clientCertContent string, clientKeyContent string, serverCertContent string, serverKeyContent string, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummyClientCertContent, clientCertContent)
		assert.Equal(t, dummyClientKeyContent, clientKeyContent)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyCACertContent, caCertContent)
		return dummyError
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
	var err = bootstrapApplication()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestBootstrapApplication_Success(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"

	// mock
	createMock(t)

	// expect
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return nil
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
	configCACertContentExpected = 1
	configCACertContent = func() string {
		configCACertContentCalled++
		return dummyCACertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(clientCertContent string, clientKeyContent string, serverCertContent string, serverKeyContent string, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummyClientCertContent, clientCertContent)
		assert.Equal(t, dummyClientKeyContent, clientKeyContent)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyCACertContent, caCertContent)
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
	var expectedErrorMessage = "some dummy error message"

	// mock
	createMock(t)

	// expect
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return errors.New(expectedErrorMessage)
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, "main", category)
		assert.Equal(t, "bootstrapApplicationFunc", subcategory)
		assert.Equal(t, "Failed to initialize server due to %v.", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, expectedErrorMessage, parameters[0].(error).Error())
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
			assert.Equal(t, "dummy db error", parameters[0].(error).Error())
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
	serverHost = func() error {
		serverHostCalled++
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
			assert.Equal(t, "dummy final error", parameters[0].(error).Error())
		}
	}
	serverHostExpected = 1
	serverHost = func() error {
		serverHostCalled++
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
	serverHost = func() error {
		serverHostCalled++
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
	serverHost = func() error {
		serverHostCalled++
		return nil
	}

	// SUT + act
	main()

	// verify
	verifyAll(t)
}
