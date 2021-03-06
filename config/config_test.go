package config

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

func TestDefaultAppVersion(t *testing.T) {
	// arrange
	var expectedResult = "0.0.0.0"

	// mock
	createMock(t)

	// SUT + act
	var result = defaultAppVersion()

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
	var result = defaultAppPort()

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
	var result = defaultAppName()

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
	var result = defaultAppPath()

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
	var result = defaultIsLocalhost()

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
	var result = defaultServeHTTPS()

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
	var result = defaultServerCertContent()

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
	var result = defaultServerKeyContent()

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
	var result = defaultValidateClientCert()

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
	var result = defaultCaCertContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultClientCertContent(t *testing.T) {
	// arrange
	var expectedResult = ""

	// mock
	createMock(t)

	// SUT + act
	var result = defaultClientCertContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultClientKeyContent(t *testing.T) {
	// arrange
	var expectedResult = ""

	// mock
	createMock(t)

	// SUT + act
	var result = defaultClientKeyContent()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultAllowedLogType(t *testing.T) {
	// arrange
	var expectedResult = logtype.BasicLogging

	// mock
	createMock(t)

	// SUT + act
	var result = defaultAllowedLogType()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultAllowedLogLevel(t *testing.T) {
	// arrange
	var expectedResult = loglevel.Warn

	// mock
	createMock(t)

	// SUT + act
	var result = defaultAllowedLogLevel()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestDefaultNetworkTimeout(t *testing.T) {
	// arrange
	var expectedResult = 3 * time.Minute

	// mock
	createMock(t)

	// SUT + act
	var result = defaultNetworkTimeout()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestGraceShutdownWaitTime(t *testing.T) {
	// arrange
	var expectedResult = 15 * time.Second

	// mock
	createMock(t)

	// SUT + act
	var result = graceShutdownWaitTime()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestSkipServerCertVerification(t *testing.T) {
	// arrange
	var expectedResult = false

	// mock
	createMock(t)

	// SUT + act
	var result = defaultSkipServerCertVerification()

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestFunctionPointerEquals_AllDifferent(t *testing.T) {
	// arrange
	var dummyLeft = func(foo int) string { return "bar" }
	var dummyRight = func(test string) int { return 123 }

	// mock
	createMock(t)

	// expect
	reflectValueOfExpected = 2
	reflectValueOf = func(i interface{}) reflect.Value {
		reflectValueOfCalled++
		return reflect.ValueOf(i)
	}
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT + act
	var result = functionPointerEquals(
		dummyLeft,
		dummyRight,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestFunctionPointerEquals_PointerDifferent(t *testing.T) {
	// arrange
	var dummyLeft = func(foo int) string { return "bar" }
	var dummyRight = func(foo int) string { return "bar" }

	// mock
	createMock(t)

	// expect
	reflectValueOfExpected = 2
	reflectValueOf = func(i interface{}) reflect.Value {
		reflectValueOfCalled++
		return reflect.ValueOf(i)
	}
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT + act
	var result = functionPointerEquals(
		dummyLeft,
		dummyRight,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestFunctionPointerEquals_NothingDifferent(t *testing.T) {
	// arrange
	var dummyLeft = func(foo int) string { return "bar" }
	var dummyRight = dummyLeft

	// mock
	createMock(t)

	// expect
	reflectValueOfExpected = 2
	reflectValueOf = func(i interface{}) reflect.Value {
		reflectValueOfCalled++
		return reflect.ValueOf(i)
	}
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT + act
	var result = functionPointerEquals(
		dummyLeft,
		dummyRight,
	)

	// assert
	assert.True(t, result)

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
	var dummyMessageFormat = "customization.%v function is forced to default [%v] due to forceToDefault flag set"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateStringFunction(
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
	var dummyMessageFormat = "customization.%v function is not configured or is empty; fallback to default [%v]"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() string {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateStringFunction(
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

func TestValidateStringFunction_DefaultStringFunc(t *testing.T) {
	// arrange
	var dummyStringFuncExpected int
	var dummyStringFuncCalled int
	var dummyStringFuncReturn string
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = "some default func return"
	var dummyForceToDefault = false
	var dummyMessageFormat = "customization.%v function is not configured or is empty; fallback to default [%v]"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	functionPointerEqualsFuncExpected = 1
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyStringFunc)), fmt.Sprintf("%v", reflect.ValueOf(left)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(right)))
		return true
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateStringFunction(
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
	var dummyMessageFormat = "customization.%v function is not configured or is empty; fallback to default [%v]"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	functionPointerEqualsFuncExpected = 1
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyStringFunc)), fmt.Sprintf("%v", reflect.ValueOf(left)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(right)))
		return false
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateStringFunction(
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
	functionPointerEqualsFuncExpected = 1
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyStringFunc)), fmt.Sprintf("%v", reflect.ValueOf(left)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(right)))
		return false
	}

	// SUT + act
	var result, err = validateStringFunction(
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
	var dummyMessageFormat = "customization.%v function is forced to default [%v] due to forceToDefault flag set"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateBooleanFunction(
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
	var dummyMessageFormat = "customization.%v function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() bool {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateBooleanFunction(
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

func TestValidateBooleanFunction_DefaultBooleanFunc(t *testing.T) {
	// arrange
	var dummyBooleanFuncExpected int
	var dummyBooleanFuncCalled int
	var dummyBooleanFuncReturn = rand.Intn(100) < 50
	var dummyName = "some name"
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = rand.Intn(100) < 50
	var dummyForceToDefault = false
	var dummyMessageFormat = "customization.%v function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	functionPointerEqualsFuncExpected = 1
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyBooleanFunc)), fmt.Sprintf("%v", reflect.ValueOf(left)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(right)))
		return true
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyName, parameters[0])
		assert.Equal(t, dummyDefaultFuncReturn, parameters[1])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateBooleanFunction(
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
	functionPointerEqualsFuncExpected = 1
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyBooleanFunc)), fmt.Sprintf("%v", reflect.ValueOf(left)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(right)))
		return false
	}

	// SUT + act
	var result, err = validateBooleanFunction(
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

func TestValidateDefaultAllowedLogType_NilFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFunc func() logtype.LogType
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = logtype.LogType(rand.Intn(255))
	var dummyMessageFormat = "customization.DefaultAllowedLogType function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() logtype.LogType {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyDefaultFuncReturn, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateDefaultAllowedLogType(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateDefaultAllowedLogType_ValidFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFuncReturn = logtype.LogType(rand.Intn(255))
	var dummyCustomizedFunc = func() logtype.LogType {
		dummyCustomizedFuncCalled++
		return dummyCustomizedFuncReturn
	}
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = logtype.LogType(rand.Intn(255))
	var dummyDefaultFunc = func() logtype.LogType {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// mock
	createMock(t)

	// SUT + act
	var result, err = validateDefaultAllowedLogType(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyCustomizedFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateDefaultAllowedLogLevel_NilFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFunc func() loglevel.LogLevel
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = loglevel.LogLevel(rand.Intn(255))
	var dummyMessageFormat = "customization.DefaultAllowedLogLevel function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() loglevel.LogLevel {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyDefaultFuncReturn, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateDefaultAllowedLogLevel(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateDefaultAllowedLogLevel_ValidFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFuncReturn = loglevel.LogLevel(rand.Intn(255))
	var dummyCustomizedFunc = func() loglevel.LogLevel {
		dummyCustomizedFuncCalled++
		return dummyCustomizedFuncReturn
	}
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = loglevel.LogLevel(rand.Intn(255))
	var dummyDefaultFunc = func() loglevel.LogLevel {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// mock
	createMock(t)

	// SUT + act
	var result, err = validateDefaultAllowedLogLevel(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyCustomizedFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateDefaultNetworkTimeout_NilFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFunc func() time.Duration
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = time.Duration(rand.Intn(255))
	var dummyMessageFormat = "customization.DefaultNetworkTimeout function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() time.Duration {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyDefaultFuncReturn, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateDefaultNetworkTimeout(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateDefaultNetworkTimeout_ValidFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFuncReturn = time.Duration(rand.Intn(255))
	var dummyCustomizedFunc = func() time.Duration {
		dummyCustomizedFuncCalled++
		return dummyCustomizedFuncReturn
	}
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = time.Duration(rand.Intn(255))
	var dummyDefaultFunc = func() time.Duration {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// mock
	createMock(t)

	// SUT + act
	var result, err = validateDefaultNetworkTimeout(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyCustomizedFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateGraceShutdownWaitTime_NilFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFunc func() time.Duration
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = time.Duration(rand.Intn(255))
	var dummyMessageFormat = "customization.GraceShutdownWaitTime function is not configured; fallback to default [%v]."
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	dummyDefaultFuncExpected = 1
	var dummyDefaultFunc = func() time.Duration {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyDefaultFuncReturn, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var result, err = validateGraceShutdownWaitTime(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyDefaultFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
	assert.Equal(t, dummyDefaultFuncExpected, dummyDefaultFuncCalled, "Unexpected number of calls to dummyDefaultFunc")
}

func TestValidateGraceShutdownWaitTime_ValidFunc(t *testing.T) {
	// arrange
	var dummyCustomizedFuncExpected int
	var dummyCustomizedFuncCalled int
	var dummyCustomizedFuncReturn = time.Duration(rand.Intn(255))
	var dummyCustomizedFunc = func() time.Duration {
		dummyCustomizedFuncCalled++
		return dummyCustomizedFuncReturn
	}
	var dummyDefaultFuncExpected int
	var dummyDefaultFuncCalled int
	var dummyDefaultFuncReturn = time.Duration(rand.Intn(255))
	var dummyDefaultFunc = func() time.Duration {
		dummyDefaultFuncCalled++
		return dummyDefaultFuncReturn
	}

	// mock
	createMock(t)

	// SUT + act
	var result, err = validateGraceShutdownWaitTime(
		dummyCustomizedFunc,
		dummyDefaultFunc,
	)

	// assert
	assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(dummyCustomizedFunc)), fmt.Sprintf("%v", reflect.ValueOf(result)))
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCustomizedFuncExpected, dummyCustomizedFuncCalled, "Unexpected number of calls to dummyCustomizedFunc")
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
	var result = isServerCertificateAvailable()

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
	var result = isServerCertificateAvailable()

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
	var result = isServerCertificateAvailable()

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
	var result = isCaCertificateAvailable()

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
	var result = isCaCertificateAvailable()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, caCertContentExpected, caCertContentCalled, "Unexpected number of calls to CaCertContent")
}

func TestInitialize(t *testing.T) {
	// arrange
	var expectedValidateStringFunctionFuncParameter1 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(customization.AppVersion)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.AppPort)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.AppName)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.AppPath)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ServerCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ServerKeyContent)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.CaCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ClientCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ClientKeyContent)),
	}
	var expectedValidateStringFunctionFuncParameter2 = []string{
		"AppVersion",
		"AppPort",
		"AppName",
		"AppPath",
		"ServerCertContent",
		"ServerKeyContent",
		"CaCertContent",
		"ClientCertContent",
		"ClientKeyContent",
	}
	var expectedValidateStringFunctionFuncParameter3 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppVersion)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppPort)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppName)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultAppPath)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServerCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServerKeyContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultCaCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultClientCertContent)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultClientKeyContent)),
	}
	var expectedValidateStringFunctionFuncReturn1 = []func() string{
		defaultAppVersion,
		defaultAppPort,
		defaultAppName,
		defaultAppPath,
		defaultServerCertContent,
		defaultServerKeyContent,
		defaultCaCertContent,
		defaultClientCertContent,
		defaultClientKeyContent,
	}
	var expectedValidateStringFunctionFuncReturn2 = []error{
		errors.New("some AppVersion error"),
		errors.New("some AppPort error"),
		errors.New("some AppName error"),
		errors.New("some AppPath error"),
		errors.New("some ServerCertContent error"),
		errors.New("some ServerKeyContent error"),
		errors.New("some CaCertContent error"),
		errors.New("some ClientCertContent error"),
		errors.New("some ClientKeyContent error"),
	}
	var expectedValidateBooleanFunctionFuncParameter1 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(customization.IsLocalhost)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ServeHTTPS)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.ValidateClientCert)),
		fmt.Sprintf("%v", reflect.ValueOf(customization.SkipServerCertVerification)),
	}
	var expectedValidateBooleanFunctionFuncParameter2 = []string{
		"IsLocalhost",
		"ServeHTTPS",
		"ValidateClientCert",
		"SkipServerCertVerification",
	}
	var expectedValidateBooleanFunctionFuncParameter3 = []string{
		fmt.Sprintf("%v", reflect.ValueOf(defaultIsLocalhost)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultServeHTTPS)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultValidateClientCert)),
		fmt.Sprintf("%v", reflect.ValueOf(defaultSkipServerCertVerification)),
	}
	var dummyIsServerCertificateAvailable = rand.Intn(100) < 50
	var dummyIsCaCertificateAvailable = rand.Intn(100) < 50
	var expectedValidateBooleanFunctionFuncParameter4 = []bool{
		false,
		!dummyIsServerCertificateAvailable,
		!dummyIsCaCertificateAvailable,
		false,
	}
	var expectedValidateBooleanFunctionFuncReturn1 = []func() bool{
		defaultIsLocalhost,
		defaultServeHTTPS,
		defaultValidateClientCert,
		defaultSkipServerCertVerification,
	}
	var expectedValidateBooleanFunctionFuncReturn2 = []error{
		errors.New("some IsLocalhost error"),
		errors.New("some ServeHTTPS error"),
		errors.New("some ValidateClientCert error"),
		errors.New("some SkipServerCertVerification error"),
	}
	var expectedDefaultAllowedLogTypeError = errors.New("some default allowed log type error")
	var expectedDefaultAllowedLogLevelError = errors.New("some default allowed log level error")
	var expectedDefaultNetworkTimeoutError = errors.New("some default network timeout error")
	var expectedGraceShutdownWaitTimeError = errors.New("some grace shutdown wait time error")
	var dummyMessageFormat = "Unexpected errors occur during configuration initialization"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	validateStringFunctionFuncExpected = 9
	validateStringFunctionFunc = func(stringFunc func() string, name string, defaultFunc func() string, forceToDefault bool) (func() string, error) {
		validateStringFunctionFuncCalled++
		assert.Equal(t, expectedValidateStringFunctionFuncParameter1[validateStringFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(stringFunc)))
		assert.Equal(t, expectedValidateStringFunctionFuncParameter2[validateStringFunctionFuncCalled-1], name)
		assert.Equal(t, expectedValidateStringFunctionFuncParameter3[validateStringFunctionFuncCalled-1], fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		assert.False(t, forceToDefault)
		return expectedValidateStringFunctionFuncReturn1[validateStringFunctionFuncCalled-1],
			expectedValidateStringFunctionFuncReturn2[validateStringFunctionFuncCalled-1]
	}
	validateBooleanFunctionFuncExpected = 4
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
	validateDefaultAllowedLogTypeFuncExpected = 1
	validateDefaultAllowedLogTypeFunc = func(customizedFunc func() logtype.LogType, defaultFunc func() logtype.LogType) (func() logtype.LogType, error) {
		validateDefaultAllowedLogTypeFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(customization.DefaultAllowedLogType)), fmt.Sprintf("%v", reflect.ValueOf(customizedFunc)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(defaultAllowedLogType)), fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		return defaultAllowedLogType, expectedDefaultAllowedLogTypeError
	}
	validateDefaultAllowedLogLevelFuncExpected = 1
	validateDefaultAllowedLogLevelFunc = func(customizedFunc func() loglevel.LogLevel, defaultFunc func() loglevel.LogLevel) (func() loglevel.LogLevel, error) {
		validateDefaultAllowedLogLevelFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(customization.DefaultAllowedLogLevel)), fmt.Sprintf("%v", reflect.ValueOf(customizedFunc)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(defaultAllowedLogLevel)), fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		return defaultAllowedLogLevel, expectedDefaultAllowedLogLevelError
	}
	validateDefaultNetworkTimeoutFuncExpected = 1
	validateDefaultNetworkTimeoutFunc = func(customizedFunc func() time.Duration, defaultFunc func() time.Duration) (func() time.Duration, error) {
		validateDefaultNetworkTimeoutFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(customization.DefaultNetworkTimeout)), fmt.Sprintf("%v", reflect.ValueOf(customizedFunc)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(defaultNetworkTimeout)), fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		return defaultNetworkTimeout, expectedDefaultNetworkTimeoutError
	}
	validateGraceShutdownWaitTimeFuncExpected = 1
	validateGraceShutdownWaitTimeFunc = func(customizedFunc func() time.Duration, defaultFunc func() time.Duration) (func() time.Duration, error) {
		validateGraceShutdownWaitTimeFuncCalled++
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(customization.GraceShutdownWaitTime)), fmt.Sprintf("%v", reflect.ValueOf(customizedFunc)))
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(graceShutdownWaitTime)), fmt.Sprintf("%v", reflect.ValueOf(defaultFunc)))
		return graceShutdownWaitTime, expectedGraceShutdownWaitTimeError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 17, len(innerErrors))
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[0], innerErrors[0])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[1], innerErrors[1])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[2], innerErrors[2])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[3], innerErrors[3])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[0], innerErrors[4])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[4], innerErrors[5])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[5], innerErrors[6])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[1], innerErrors[7])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[6], innerErrors[8])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[2], innerErrors[9])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[7], innerErrors[10])
		assert.Equal(t, expectedValidateStringFunctionFuncReturn2[8], innerErrors[11])
		assert.Equal(t, expectedDefaultAllowedLogTypeError, innerErrors[12])
		assert.Equal(t, expectedDefaultAllowedLogLevelError, innerErrors[13])
		assert.Equal(t, expectedDefaultNetworkTimeoutError, innerErrors[14])
		assert.Equal(t, expectedValidateBooleanFunctionFuncReturn2[3], innerErrors[15])
		assert.Equal(t, expectedGraceShutdownWaitTimeError, innerErrors[16])
		assert.Equal(t, dummyMessageFormat, messageFormat)
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
