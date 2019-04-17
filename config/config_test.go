package config

import (
	"errors"
	"testing"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/stretchr/testify/assert"
)

func TestStateVariables(t *testing.T) {
	// mock
	createMock(t)

	// assert
	assert.Equal(t, "0.0.0.0", appVersion)
	assert.Equal(t, "443", appPort)
	assert.Equal(t, "WebServiceTemplate", appName)
	assert.Equal(t, ".", appPath)
	assert.Equal(t, "", cryptoKey)
	assert.Equal(t, "UEvaxQGW6YC9aeCs", CryptoKeyPartial)

	// verify
	verifyAll(t)
}

func TestInitializeBootTime(t *testing.T) {
	// arrange
	var dummyTimeNowUTC = time.Now().UTC()
	var dummyFormattedTime = "some formatted time"

	// mock
	createMock(t)

	// expect
	timeutilGetTimeNowUTCExpected = 1
	timeutilGetTimeNowUTC = func() time.Time {
		timeutilGetTimeNowUTCCalled++
		return dummyTimeNowUTC
	}
	timeutilFormatDateTimeExpected = 1
	timeutilFormatDateTime = func(value time.Time) string {
		timeutilFormatDateTimeCalled++
		assert.Equal(t, dummyTimeNowUTC, value)
		return dummyFormattedTime
	}

	// SUT + act
	initializeBootTime()

	// assert
	assert.Equal(t, dummyFormattedTime, bootTime)

	// verify
	verifyAll(t)
}

func TestInitializeCryptoKey_InvalidKeyLength(t *testing.T) {
	// arrange
	var dummyEnvCryptoKey = "some env crypto key"
	var expectedErrorMessage = "Invalid crypto key length: make sure environment variable is set properly"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	getEnvironmentVariableExpected = 1
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		assert.Equal(t, "CryptoKey", key)
		return dummyEnvCryptoKey
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
	var err = initializeCryptoKey()

	// assert
	assert.Equal(t, dummyAppError, err)
	assert.Zero(t, cryptoKey)

	// verify
	verifyAll(t)
}

func TestInitializeCryptoKey_Success(t *testing.T) {
	// arrange
	var dummyEnvCryptoKey = "valid crypto key"

	// mock
	createMock(t)

	// expect
	getEnvironmentVariableExpected = 1
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		assert.Equal(t, "CryptoKey", key)
		return dummyEnvCryptoKey
	}

	// SUT + act
	var err = initializeCryptoKey()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, CryptoKeyPartial+dummyEnvCryptoKey, cryptoKey)

	// verify
	verifyAll(t)
}

func TestInitializeEnvironmentVariables(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	err := initializeEnvironmentVariables()

	// assert
	assert.Nil(t, err)

	// tear down
	verifyAll(t)
}

func TestDecryptFromEnvironmentVariable_EmptyEnvVar(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = ""

	// mock
	createMock(t)

	// expect
	getEnvironmentVariableExpected = 1
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		assert.Equal(t, dummyName, key)
		return dummyValue
	}

	// SUT + act
	result, err := decryptFromEnvironmentVariable(
		dummyName,
	)

	// assert
	assert.Zero(t, result)
	assert.Nil(t, err)

	// tear down
	verifyAll(t)
}

func TestDecryptFromEnvironmentVariable_DecryptError(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = "some value"
	var dummyCryptoKey = "some crypto key"
	var dummyResult = "some result"
	var dummyError = errors.New("some error")
	var expectedErrorMessage = "Failed to decrypt environment variable [%v]"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	cryptoKey = dummyCryptoKey

	// mock
	createMock(t)

	// expect
	getEnvironmentVariableExpected = 1
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		assert.Equal(t, dummyName, key)
		return dummyValue
	}
	cryptoDecryptExpected = 1
	cryptoDecrypt = func(cipherText string, key string) (string, error) {
		cryptoDecryptCalled++
		assert.Equal(t, dummyValue, cipherText)
		assert.Equal(t, dummyCryptoKey, key)
		return dummyResult, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, dummyError, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		return dummyAppError
	}

	// SUT + act
	result, err := decryptFromEnvironmentVariable(
		dummyName,
	)

	// assert
	assert.Equal(t, dummyValue, result)
	assert.Equal(t, dummyAppError, err)

	// tear down
	verifyAll(t)
}

func TestDecryptFromEnvironmentVariable_Success(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = "some value"
	var dummyCryptoKey = "some crypto key"
	var dummyResult = "some result"

	// stub
	cryptoKey = dummyCryptoKey

	// mock
	createMock(t)

	// expect
	getEnvironmentVariableExpected = 1
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		assert.Equal(t, dummyName, key)
		return dummyValue
	}
	cryptoDecryptExpected = 1
	cryptoDecrypt = func(cipherText string, key string) (string, error) {
		cryptoDecryptCalled++
		assert.Equal(t, dummyValue, cipherText)
		assert.Equal(t, dummyCryptoKey, key)
		return dummyResult, nil
	}

	// SUT + act
	result, err := decryptFromEnvironmentVariable(
		dummyName,
	)

	// assert
	assert.Equal(t, dummyResult, result)
	assert.Nil(t, err)

	// tear down
	verifyAll(t)
}

func TestInitializeEncryptedEnvironmentVariables_WithErrors(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"
	var dummyClientCertError = errors.New("some client cert error")
	var dummyClientKeyError = errors.New("some client key error")
	var dummyServerCertError = errors.New("some server cert error")
	var dummyServerKeyError = errors.New("some server key error")
	var dummyCACertError = errors.New("some CA cert error")
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	decryptFromEnvironmentVariableFuncExpected = 5
	decryptFromEnvironmentVariableFunc = func(name string) (string, error) {
		decryptFromEnvironmentVariableFuncCalled++
		if decryptFromEnvironmentVariableFuncCalled == 1 {
			assert.Equal(t, "ClientCertContent", name)
			return dummyClientCertContent, dummyClientCertError
		} else if decryptFromEnvironmentVariableFuncCalled == 2 {
			assert.Equal(t, "ClientKeyContent", name)
			return dummyClientKeyContent, dummyClientKeyError
		} else if decryptFromEnvironmentVariableFuncCalled == 3 {
			assert.Equal(t, "ServerCertContent", name)
			return dummyServerCertContent, dummyServerCertError
		} else if decryptFromEnvironmentVariableFuncCalled == 4 {
			assert.Equal(t, "ServerKeyContent", name)
			return dummyServerKeyContent, dummyServerKeyError
		} else if decryptFromEnvironmentVariableFuncCalled == 5 {
			assert.Equal(t, "CACertContent", name)
			return dummyCACertContent, dummyCACertError
		}
		return "", nil
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, "Failed to decrypt environment variables", baseErrorMessage)
		assert.Equal(t, 5, len(allErrors))
		assert.Equal(t, dummyClientCertError, allErrors[0])
		assert.Equal(t, dummyClientKeyError, allErrors[1])
		assert.Equal(t, dummyServerCertError, allErrors[2])
		assert.Equal(t, dummyServerKeyError, allErrors[3])
		assert.Equal(t, dummyCACertError, allErrors[4])
		return dummyAppError
	}

	// SUT + act
	err := initializeEncryptedEnvironmentVariables()

	// assert
	assert.Equal(t, dummyAppError, err)

	// tear down
	verifyAll(t)
}

func TestInitializeEncryptedEnvironmentVariables_NoError(t *testing.T) {
	// arrange
	var dummyClientCertContent = "some client cert content"
	var dummyClientKeyContent = "some client key content"
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyCACertContent = "some CA cert content"

	// mock
	createMock(t)

	// expect
	decryptFromEnvironmentVariableFuncExpected = 5
	decryptFromEnvironmentVariableFunc = func(name string) (string, error) {
		decryptFromEnvironmentVariableFuncCalled++
		if decryptFromEnvironmentVariableFuncCalled == 1 {
			assert.Equal(t, "ClientCertContent", name)
			return dummyClientCertContent, nil
		} else if decryptFromEnvironmentVariableFuncCalled == 2 {
			assert.Equal(t, "ClientKeyContent", name)
			return dummyClientKeyContent, nil
		} else if decryptFromEnvironmentVariableFuncCalled == 3 {
			assert.Equal(t, "ServerCertContent", name)
			return dummyServerCertContent, nil
		} else if decryptFromEnvironmentVariableFuncCalled == 4 {
			assert.Equal(t, "ServerKeyContent", name)
			return dummyServerKeyContent, nil
		} else if decryptFromEnvironmentVariableFuncCalled == 5 {
			assert.Equal(t, "CACertContent", name)
			return dummyCACertContent, nil
		}
		return "", nil
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, "Failed to decrypt environment variables", baseErrorMessage)
		assert.Equal(t, 5, len(allErrors))
		assert.Nil(t, allErrors[0])
		assert.Nil(t, allErrors[1])
		assert.Nil(t, allErrors[2])
		assert.Nil(t, allErrors[3])
		assert.Nil(t, allErrors[4])
		return nil
	}

	// SUT + act
	err := initializeEncryptedEnvironmentVariables()

	// assert
	assert.Nil(t, err)

	// tear down
	verifyAll(t)
}

func TestInitialize_CryptoKeyError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var expectedErrorMessage = "Failed to initialize crypto key"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	initializeBootTimeFuncExpected = 1
	initializeBootTimeFunc = func() {
		initializeBootTimeFuncCalled++
	}
	initializeCryptoKeyFuncExpected = 1
	initializeCryptoKeyFunc = func() error {
		initializeCryptoKeyFuncCalled++
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
	var err = Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_EnvVarError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var expectedErrorMessage = "Failed to load environment variables"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	initializeBootTimeFuncExpected = 1
	initializeBootTimeFunc = func() {
		initializeBootTimeFuncCalled++
	}
	initializeCryptoKeyFuncExpected = 1
	initializeCryptoKeyFunc = func() error {
		initializeCryptoKeyFuncCalled++
		return nil
	}
	initializeEnvironmentVariablesFuncExpected = 1
	initializeEnvironmentVariablesFunc = func() error {
		initializeEnvironmentVariablesFuncCalled++
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
	var err = Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_EncryptedEnvVarError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var expectedErrorMessage = "Failed to load encrypted environment variables"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	initializeBootTimeFuncExpected = 1
	initializeBootTimeFunc = func() {
		initializeBootTimeFuncCalled++
	}
	initializeCryptoKeyFuncExpected = 1
	initializeCryptoKeyFunc = func() error {
		initializeCryptoKeyFuncCalled++
		return nil
	}
	initializeEnvironmentVariablesFuncExpected = 1
	initializeEnvironmentVariablesFunc = func() error {
		initializeEnvironmentVariablesFuncCalled++
		return nil
	}
	initializeEncryptedEnvironmentVariablesFuncExpected = 1
	initializeEncryptedEnvironmentVariablesFunc = func() error {
		initializeEncryptedEnvironmentVariablesFuncCalled++
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
	var err = Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_Success(t *testing.T) {
	// mock
	createMock(t)

	// expect
	initializeBootTimeFuncExpected = 1
	initializeBootTimeFunc = func() {
		initializeBootTimeFuncCalled++
	}
	initializeCryptoKeyFuncExpected = 1
	initializeCryptoKeyFunc = func() error {
		initializeCryptoKeyFuncCalled++
		return nil
	}
	initializeEnvironmentVariablesFuncExpected = 1
	initializeEnvironmentVariablesFunc = func() error {
		initializeEnvironmentVariablesFuncCalled++
		return nil
	}
	initializeEncryptedEnvironmentVariablesFuncExpected = 1
	initializeEncryptedEnvironmentVariablesFunc = func() error {
		initializeEncryptedEnvironmentVariablesFuncCalled++
		return nil
	}

	// SUT + act
	var err = Initialize()

	// assert
	assert.Nil(t, err)

	// verify
	verifyAll(t)
}

func TestAppVersion(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	appVersion = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = AppVersion()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestAppPort(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	appPort = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = AppPort()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestAppName(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	appName = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = AppName()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestAppPath(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	appPath = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = AppPath()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestCryptoKey(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	cryptoKey = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = CryptoKey()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestClientCertContent(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	clientCertContent = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = ClientCertContent()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestClientKeyContent(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	clientKeyContent = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = ClientKeyContent()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestServerCertContent(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	serverCertContent = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = ServerCertContent()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestServerKeyContent(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	serverKeyContent = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = ServerKeyContent()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}

func TestCACertContent(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// stub
	caCertContent = dummyValue

	// mock
	createMock(t)

	// SUT + act
	var result = CACertContent()

	// assert
	assert.Equal(t, dummyValue, result)

	// verify
	verifyAll(t)
}
