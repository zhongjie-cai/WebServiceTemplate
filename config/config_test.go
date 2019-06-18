package config

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestDefaultAppVersion(t *testing.T) {
	// arrange
	var expectedResult = "0.0.0.0"

	// mock
	createMock(t)

	// SUT + act
	result := defaultAppVersion()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultAppPort(t *testing.T) {
	// arrange
	var expectedResult = "18605"

	// mock
	createMock(t)

	// SUT + act
	result := defaultAppPort()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultAppName(t *testing.T) {
	// arrange
	var expectedResult = "WebServiceTemplate"

	// mock
	createMock(t)

	// SUT + act
	result := defaultAppName()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultAppPath(t *testing.T) {
	// arrange
	var expectedResult = "."

	// mock
	createMock(t)

	// SUT + act
	result := defaultAppPath()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultIsLocalhost(t *testing.T) {
	// arrange
	var expectedResult = false

	// mock
	createMock(t)

	// SUT + act
	result := defaultIsLocalhost()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultServeHTTPS(t *testing.T) {
	// arrange
	var expectedResult = false

	// mock
	createMock(t)

	// SUT + act
	result := defaultServeHTTPS()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultServerCertContent(t *testing.T) {
	// arrange
	var expectedResult = ""

	// mock
	createMock(t)

	// SUT + act
	result := defaultServerCertContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultServerKeyContent(t *testing.T) {
	// arrange
	var expectedResult = ""

	// mock
	createMock(t)

	// SUT + act
	result := defaultServerKeyContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultValidateClientCert(t *testing.T) {
	// arrange
	var expectedResult = false

	// mock
	createMock(t)

	// SUT + act
	result := defaultValidateClientCert()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultCaCertContent(t *testing.T) {
	// arrange
	var expectedResult = ""

	// mock
	createMock(t)

	// SUT + act
	result := defaultCaCertContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestValidateStringFunction_ForcedToDefault(t *testing.T) {
	// arrange
	var dummyStringFuncExpected int
	var dummyStringFuncCalled int
	var dummyStringFuncReturn = "some string func return"
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = "some default func return"
	var dummyForceToDefault = true
	var dummyMessageFormat = "config.%v function is forced to default [%v] due to forceToDefault flag set"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	dummyStringFuncExpected = 0
	var dummyStringFunc = func() string {
		dummyStringFuncCalled++
		return dummyStringFuncReturn
	}
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() string {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	result, err := validateStringFunction(
		dummyStringFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyStringFuncExpected, dummyStringFuncCalled, "Unexpected number of calls to dummyStringFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateStringFunction_NilStringFunc(t *testing.T) {
	// arrange
	var dummyStringFuncExpected int
	var dummyStringFuncCalled int
	var dummyStringFunc func() string
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = "some default func return"
	var dummyForceToDefault = false
	var dummyMessageFormat = "config.%v function is not configured or is empty; fallback to default [%v]"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() string {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	result, err := validateStringFunction(
		dummyStringFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyStringFuncExpected, dummyStringFuncCalled, "Unexpected number of calls to dummyStringFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateStringFunction_EmptyStringFunc(t *testing.T) {
	// arrange
	var dummyStringFuncExpected int
	var dummyStringFuncCalled int
	var dummyStringFuncReturn string
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = "some default func return"
	var dummyForceToDefault = false
	var dummyMessageFormat = "config.%v function is not configured or is empty; fallback to default [%v]"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	dummyStringFuncExpected = 1
	var dummyStringFunc = func() string {
		dummyStringFuncCalled++
		return dummyStringFuncReturn
	}
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() string {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	result, err := validateStringFunction(
		dummyStringFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyStringFuncExpected, dummyStringFuncCalled, "Unexpected number of calls to dummyStringFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateStringFunction_ValidStringFunc(t *testing.T) {
	// arrange
	var dummyStringFuncExpected int
	var dummyStringFuncCalled int
	var dummyStringFuncReturn = "some string func return"
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = "some default func return"
	var dummyForceToDefault = false

	// mock
	createMock(t)

	// expect
	dummyStringFuncExpected = 1
	var dummyStringFunc = func() string {
		dummyStringFuncCalled++
		return dummyStringFuncReturn
	}
	dummyDefaultFuncExpected = 0
	var dummyDefaultFunc = func() string {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// SUT + act
	result, err := validateStringFunction(
		dummyStringFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyStringFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyStringFuncExpected, dummyStringFuncCalled, "Unexpected number of calls to dummyStringFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateBooleanFunction_ForcedToDefault(t *testing.T) {
	// arrange
	var dummyBooleanFuncExpected int
	var dummyBooleanFuncCalled int
	var dummyBooleanFuncReturn = rand.Intn(100) < 50
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = rand.Intn(100) < 50
	var dummyForceToDefault = true
	var dummyMessageFormat = "config.%v function is forced to default [%v] due to forceToDefault flag set"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	dummyBooleanFuncExpected = 0
	var dummyBooleanFunc = func() bool {
		dummyBooleanFuncCalled++
		return dummyBooleanFuncReturn
	}
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() bool {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	result, err := validateBooleanFunction(
		dummyBooleanFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyBooleanFuncExpected, dummyBooleanFuncCalled, "Unexpected number of calls to dummyBooleanFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateBooleanFunction_NilBooleanFunc(t *testing.T) {
	// arrange
	var dummyBooleanFuncExpected int
	var dummyBooleanFuncCalled int
	var dummyBooleanFunc func() bool
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = rand.Intn(100) < 50
	var dummyForceToDefault = false
	var dummyMessageFormat = "config.%v function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() bool {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	result, err := validateBooleanFunction(
		dummyBooleanFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyBooleanFuncExpected, dummyBooleanFuncCalled, "Unexpected number of calls to dummyBooleanFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateBooleanFunction_ValidBooleanFunc(t *testing.T) {
	// arrange
	var dummyBooleanFuncExpected int
	var dummyBooleanFuncCalled int
	var dummyBooleanFuncReturn = rand.Intn(100) < 50
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = rand.Intn(100) < 50
	var dummyForceToDefault = false

	// mock
	createMock(t)

	// expect
	dummyBooleanFuncExpected = 0
	var dummyBooleanFunc = func() bool {
		dummyBooleanFuncCalled++
		return dummyBooleanFuncReturn
	}
	dummyDefaultFuncExpected = 0
	var dummyDefaultFunc = func() bool {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// SUT + act
	result, err := validateBooleanFunction(
		dummyBooleanFunc,
		dummyName,
		dummyDefaultFunc,
		dummyForceToDefault,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyBooleanFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyBooleanFuncExpected, dummyBooleanFuncCalled, "Unexpected number of calls to dummyBooleanFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestIsServerCertificateAvailable_CertEmpty(t *testing.T) {
	// arrange
	var serverCertContentExpected int
	var serverCertContentCalled int
	var serverKeyContentExpected int
	var serverKeyContentCalled int

	// mock
	createMock(t)

	// expect
	serverCertContentExpected = 1
	ServerCertContent = func() string {
		serverCertContentCalled++
		return ""
	}

	// SUT + act
	result := isServerCertificateAvailable()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, serverCertContentExpected, serverCertContentCalled, "Unexpected number of calls to ServerCertContent")
	assert.Equal(t, serverKeyContentExpected, serverKeyContentCalled, "Unexpected number of calls to ServerKeyContent")
}

func TestIsServerCertificateAvailable_KeyEmpty(t *testing.T) {
	// arrange
	var serverCertContentExpected int
	var serverCertContentCalled int
	var serverKeyContentExpected int
	var serverKeyContentCalled int

	// mock
	createMock(t)

	// expect
	serverCertContentExpected = 1
	ServerCertContent = func() string {
		serverCertContentCalled++
		return "some cert content"
	}
	serverKeyContentExpected = 1
	ServerKeyContent = func() string {
		serverKeyContentCalled++
		return ""
	}

	// SUT + act
	result := isServerCertificateAvailable()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, serverCertContentExpected, serverCertContentCalled, "Unexpected number of calls to ServerCertContent")
	assert.Equal(t, serverKeyContentExpected, serverKeyContentCalled, "Unexpected number of calls to ServerKeyContent")
}

func TestIsServerCertificateAvailable_NotEmpty(t *testing.T) {
	// arrange
	var serverCertContentExpected int
	var serverCertContentCalled int
	var serverKeyContentExpected int
	var serverKeyContentCalled int

	// mock
	createMock(t)

	// expect
	serverCertContentExpected = 1
	ServerCertContent = func() string {
		serverCertContentCalled++
		return "some cert content"
	}
	serverKeyContentExpected = 1
	ServerKeyContent = func() string {
		serverKeyContentCalled++
		return "some key content"
	}

	// SUT + act
	result := isServerCertificateAvailable()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, serverCertContentExpected, serverCertContentCalled, "Unexpected number of calls to ServerCertContent")
	assert.Equal(t, serverKeyContentExpected, serverKeyContentCalled, "Unexpected number of calls to ServerKeyContent")
}

func TestIsCaCertificateAvailable_Empty(t *testing.T) {
	// arrange
	var caCertContentExpected int
	var caCertContentCalled int

	// mock
	createMock(t)

	// expect
	caCertContentExpected = 1
	CaCertContent = func() string {
		caCertContentCalled++
		return ""
	}

	// SUT + act
	result := isCaCertificateAvailable()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, caCertContentExpected, caCertContentCalled, "Unexpected number of calls to CaCertContent")
}

func TestIsCaCertificateAvailable_NotEmpty(t *testing.T) {
	// arrange
	var caCertContentExpected int
	var caCertContentCalled int

	// mock
	createMock(t)

	// expect
	caCertContentExpected = 1
	CaCertContent = func() string {
		caCertContentCalled++
		return "some ca cert content"
	}

	// SUT + act
	result := isCaCertificateAvailable()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, caCertContentExpected, caCertContentCalled, "Unexpected number of calls to CaCertContent")
}

func TestInitialize(t *testing.T) {
	// arrange
	var expectedValidateStringFunctionFuncParameter1 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(AppVersion)),
		fmt.Sprintf("%v", reflect.ValueOf(AppPort)),
		fmt.Sprintf("%v", reflect.ValueOf(AppName)),
		fmt.Sprintf("%v", reflect.ValueOf(AppPath)),
		fmt.Sprintf("%v", reflect.ValueOf(ServerCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(ServerKeyContent)),
		fmt.Sprintf("%v", reflect.ValueOf(CaCertContent)),
	}
	var expectedValidateStringFunctionFuncParameter2 = []string{
		"AppVersion",
		"AppPort",
		"AppName",
		"AppPath",
		"ServerCertContent",
		"ServerKeyContent",
		"CaCertContent",
	}
	var expectedValidateStringFunctionFuncParameter3 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppVersion)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppPort)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppName)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppPath)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServerCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServerKeyContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultCaCertContent)),
	}
	var expectedValidateStringFunctionFuncReturn1 = []func() string{
		defaultAppVersion,
		defaultAppPort,
		defaultAppName,
		defaultAppPath,
		defaultServerCertContent,
		defaultServerKeyContent,
		defaultCaCertContent,
	}
	var expectedValidateStringFunctionFuncReturn2 = []error{
		errors.New("some AppVersion error"),
		errors.New("some AppPort error"),
		errors.New("some AppName error"),
		errors.New("some AppPath error"),
		errors.New("some ServerCertContent error"),
		errors.New("some ServerKeyContent error"),
		errors.New("some CaCertContent error"),
	}
	var expectedValidateBooleanFunctionFuncParameter1 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(IsLocalhost)),
		fmt.Sprintf("%v", reflect.ValueOf(ServeHTTPS)),
		fmt.Sprintf("%v", reflect.ValueOf(ValidateClientCert)),
	}
	var expectedValidateBooleanFunctionFuncParameter2 = []string{
		"IsLocalhost",
		"ServeHTTPS",
		"ValidateClientCert",
	}
	var expectedValidateBooleanFunctionFuncParameter3 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(defaultIsLocalhost)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServeHTTPS)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultValidateClientCert)),
	}
	var dummyIsServerCertificateAvailable = rand.Intn(100) < 50
	var dummyIsCaCertificateAvailable = rand.Intn(100) < 50
	var expectedValidateBooleanFunctionFuncParameter4 = []bool{
		false,
		!dummyIsServerCertificateAvailable,
		!dummyIsCaCertificateAvailable,
	}
	var expectedValidateBooleanFunctionFuncReturn1 = []func() bool{
		defaultIsLocalhost,
		defaultServeHTTPS,
		defaultValidateClientCert,
	}
	var expectedValidateBooleanFunctionFuncReturn2 = []error{
		errors.New("some IsLocalhost error"),
		errors.New("some ServeHTTPS error"),
		errors.New("some ValidateClientCert error"),
	}
	var dummyBaseErrorMessage = "Unexpected errors occur during configuration initialization"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	validateStringFunctionFuncExpected = 7
	validateStringFunctionFunc = func(stringFunc func() string, name string, defaultFunc func() string, forceToDefault bool) (func() string, error) {
		validateStringFunctionFuncCalled++
		assert.Equal(t, expectedValidateStringFunctionFuncParameter1[validateStringFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(stringFunc)))
		assert.Equal(t, expectedValidateStringFunctionFuncParameter2[validateStringFunctionFuncCalled-1], name)
		assert.Equal(t, expectedValidateStringFunctionFuncParameter3[validateStringFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		assert.False(t, forceToDefault)
		return expectedValidateStringFunctionFuncReturn1[validateStringFunctionFuncCalled-1],
			expectedValidateStringFunctionFuncReturn2[validateStringFunctionFuncCalled-1]
	}
	validateBooleanFunctionFuncExpected = 3
	validateBooleanFunctionFunc = func(booleanFunc func() bool, name string, defaultFunc func() bool, forceToDefault bool) (func() bool, error) {
		validateBooleanFunctionFuncCalled++
		assert.Equal(t, expectedValidateBooleanFunctionFuncParameter1[validateBooleanFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(booleanFunc)))
		assert.Equal(t, expectedValidateBooleanFunctionFuncParameter2[validateBooleanFunctionFuncCalled-1], name)
		assert.Equal(t, expectedValidateBooleanFunctionFuncParameter3[validateBooleanFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		assert.Equal(t, expectedValidateBooleanFunctionFuncParameter4[validateBooleanFunctionFuncCalled-1], forceToDefault)
		return expectedValidateBooleanFunctionFuncReturn1[validateBooleanFunctionFuncCalled-1],
			expectedValidateBooleanFunctionFuncReturn2[validateBooleanFunctionFuncCalled-1]
	}
	isServerCertificateAvailableFuncExpected = 1
	isServerCertificateAvailableFunc = func() bool {
		isServerCertificateAvailableFuncCalled++
		return dummyIsServerCertificateAvailable
	}
	isCaCertificateAvailableFuncExpected = 1
	isCaCertificateAvailableFunc = func() bool {
		isCaCertificateAvailableFuncCalled++
		return dummyIsCaCertificateAvailable
	}
	apperrorConsolidateAllErrorsExpected = 1
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		assert.Equal(t, dummyBaseErrorMessage, baseErrorMessage)
		assert.Equal(t, 10, len(allErrors))
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[0], allErrors[0])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[1], allErrors[1])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[2], allErrors[2])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[3], allErrors[3])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[0], allErrors[4])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[4], allErrors[5])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[5], allErrors[6])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[1], allErrors[7])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[6], allErrors[8])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[2], allErrors[9])
		return dummyAppError
	}

	// SUT + act
	err := Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}
