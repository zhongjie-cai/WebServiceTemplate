package apperror

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func TestAppErrorGetCode_NoCustomization(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}

	// mock
	createMock(t)

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.Code()

	// assert
	assert.Equal(t, expectedCode.String(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetCode_WithCustomization_NoFoundMatch(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyCustomizedNameMap = map[enum.Code]string{}

	// mock
	createMock(t)

	// expect
	customizationAppErrorsExpected = 1
	customization.AppErrors = func() (map[enum.Code]string, map[enum.Code]int) {
		customizationAppErrorsCalled++
		return dummyCustomizedNameMap, nil
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.Code()

	// assert
	assert.Equal(t, expectedCode.String(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetCode_WithCustomization_FoundMatch(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyCodeName = "some code name"
	var dummyCustomizedNameMap = map[enum.Code]string{
		expectedCode: dummyCodeName,
	}

	// mock
	createMock(t)

	// expect
	customizationAppErrorsExpected = 1
	customization.AppErrors = func() (map[enum.Code]string, map[enum.Code]int) {
		customizationAppErrorsCalled++
		return dummyCustomizedNameMap, nil
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.Code()

	// assert
	assert.Equal(t, dummyCodeName, code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetHTTPStatusCode_NoCustomization(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}

	// mock
	createMock(t)

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.HTTPStatusCode()

	// assert
	assert.Equal(t, expectedCode.HTTPStatusCode(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetHTTPStatusCode_WithCustomization_NoFoundMatch(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyCustomizedStatusMap = map[enum.Code]int{}

	// mock
	createMock(t)

	// expect
	customizationAppErrorsExpected = 1
	customization.AppErrors = func() (map[enum.Code]string, map[enum.Code]int) {
		customizationAppErrorsCalled++
		return nil, dummyCustomizedStatusMap
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.HTTPStatusCode()

	// assert
	assert.Equal(t, expectedCode.HTTPStatusCode(), code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetHTTPStatusCode_WithCustomization_FoundMatch(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyHTTPStatusCode = rand.Int()
	var dummyCustomizedStatusMap = map[enum.Code]int{
		expectedCode: dummyHTTPStatusCode,
	}

	// mock
	createMock(t)

	// expect
	customizationAppErrorsExpected = 1
	customization.AppErrors = func() (map[enum.Code]string, map[enum.Code]int) {
		customizationAppErrorsCalled++
		return nil, dummyCustomizedStatusMap
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.HTTPStatusCode()

	// assert
	assert.Equal(t, dummyHTTPStatusCode, code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetError_NoInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var expectedMessage = "(GeneralFailure)dummy error"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		nil,
	}

	// act
	var message = appError.Error()

	// assert
	assert.Equal(t, expectedMessage, message)

	// verify
	verifyAll(t)
}

func TestAppErrorGetError_WithInner(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedMessage = "(GeneralFailure)dummy error -> [ [ dummy inner error 1 ] | [ dummy inner error 2 ] | [ dummy inner error 3 ] ]"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var message = appError.Error()

	// assert
	assert.Equal(t, expectedMessage, message)

	// verify
	verifyAll(t)
}

func TestAppErrorGetInnerErrors_NoInner(t *testing.T) {
	// arrange
	var expectedMessage = "dummy error"
	var expectedError = errors.New(expectedMessage)
	var expectedCode = enum.CodeGeneralFailure

	// mock
	createMock(t)

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		nil,
	}

	// act
	var innerErrors = appError.InnerErrors()

	// assert
	assert.Equal(t, 0, len(innerErrors))

	// verify
	verifyAll(t)
}

func TestAppErrorGetInnerErrors_WithInner(t *testing.T) {
	// arrange
	var expectedMessage = "dummy error"
	var expectedError = errors.New(expectedMessage)
	var expectedCode = enum.CodeGeneralFailure
	var expectedInnerError1 = errors.New("some inner error 1")
	var expectedInnerError2 = errors.New("some inner error 2")
	var expectedInnerError3 = errors.New("some inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}

	// mock
	createMock(t)

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var innerErrors = appError.InnerErrors()

	// assert
	assert.Equal(t, expectedInnerErrors, innerErrors)

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_NoInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var expectedMessage = "(GeneralFailure)dummy error"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		nil,
	}

	// act
	var messages = appError.Messages()

	// assert
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, expectedMessage, messages[0])

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_WithNormalInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var expectedMessage = "(GeneralFailure)dummy error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = errors.New("dummy inner error 2")
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedInnerMessage1 = "  " + expectedInnerError1.Error()
	var expectedInnerMessage2 = "  " + expectedInnerError2.Error()
	var expectedInnerMessage3 = "  " + expectedInnerError3.Error()

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var messages = appError.Messages()

	// assert
	assert.Equal(t, 4, len(messages))
	assert.Equal(t, expectedMessage, messages[0])
	assert.Equal(t, expectedInnerMessage1, messages[1])
	assert.Equal(t, expectedInnerMessage2, messages[2])
	assert.Equal(t, expectedInnerMessage3, messages[3])

	// verify
	verifyAll(t)
}

func TestAppErrorGetMessages_WithAppErrorInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var expectedMessage = "(GeneralFailure)dummy error"
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &appError{
		errors.New(dummyInnerErrorMessage),
		enum.CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var expectedInnerMessage1 = "  " + expectedInnerError1.Error()
	var expectedInnerMessage2 = "  (GeneralFailure)" + dummyInnerErrorMessage
	var expectedInnerMostMessage = "    " + dummyInnerMostErrorMessage
	var expectedInnerMessage3 = "  " + expectedInnerError3.Error()

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 2
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var messages = appError.Messages()

	// assert
	assert.Equal(t, 5, len(messages))
	assert.Equal(t, expectedMessage, messages[0])
	assert.Equal(t, expectedInnerMessage1, messages[1])
	assert.Equal(t, expectedInnerMessage2, messages[2])
	assert.Equal(t, expectedInnerMostMessage, messages[3])
	assert.Equal(t, expectedInnerMessage3, messages[4])

	// verify
	verifyAll(t)
}

func TestAppErrorAppend_NoInnerError(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &appError{
		errors.New(dummyInnerErrorMessage),
		enum.CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}
	var cleanedInnerErrors = []error{}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	appError.Append(
		dummyInnerErrors...,
	)

	// assert
	assert.Equal(t, expectedInnerErrors, appError.innerErrors)

	// verify
	verifyAll(t)
}

func TestAppErrorAppend_HasInnerError(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = enum.CodeGeneralFailure
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = &appError{
		errors.New(dummyInnerErrorMessage),
		enum.CodeGeneralFailure,
		[]error{errors.New(dummyInnerMostErrorMessage)},
	}
	var expectedInnerError3 = errors.New("dummy inner error 3")
	var expectedInnerErrors = []error{
		expectedInnerError1,
		expectedInnerError2,
		expectedInnerError3,
	}
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}
	var cleanedInnerErrors = []error{
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT
	var appError = &appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	appError.Append(
		dummyInnerErrors...,
	)

	// assert
	assert.Equal(t, 6, len(appError.innerErrors))
	assert.Equal(t, expectedInnerErrors[0], appError.innerErrors[0])
	assert.Equal(t, expectedInnerErrors[1], appError.innerErrors[1])
	assert.Equal(t, expectedInnerErrors[2], appError.innerErrors[2])
	assert.Equal(t, dummyInnerError1, appError.innerErrors[3])
	assert.Equal(t, dummyInnerError2, appError.innerErrors[4])
	assert.Equal(t, dummyInnerError3, appError.innerErrors[5])

	// verify
	verifyAll(t)
}

func TestInitialize_NoCustomization(t *testing.T) {
	// mock
	createMock(t)

	// expect
	customization.AppErrors = nil

	// SUT + act
	var err = Initialize()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitialize_WithCustomization(t *testing.T) {
	// arrange
	var dummyCodeNames = map[enum.Code]string{
		0:                          "some code 0",
		enum.CodeReservedCount - 1: "some code before reserved count",
		enum.CodeReservedCount:     "some code at reserved count",
		enum.CodeReservedCount + 1: "some code after reserved count",
		math.MaxInt32:              "some code max int",
	}
	var dummyHTTPStatusCodes = map[enum.Code]int{
		0:                          0,
		enum.CodeReservedCount - 1: 100,
		enum.CodeReservedCount:     200,
		enum.CodeReservedCount + 1: 300,
		math.MaxInt32:              400,
	}
	var dummyError1 = errors.New("some error 1")
	var dummyError2 = errors.New("some error 2")
	var dummyError3 = errors.New("some error 3")
	var dummyError4 = errors.New("some error 4")
	var dummyFinalError = &appError{}

	// mock
	createMock(t)

	// expect
	customizationAppErrorsExpected = 1
	customization.AppErrors = func() (map[enum.Code]string, map[enum.Code]int) {
		customizationAppErrorsCalled++
		return dummyCodeNames, dummyHTTPStatusCodes
	}
	fmtErrorfExpected = 4
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		if dummyCodeNames[0] == a[1] {
			assert.Equal(t, "AppError code [%v] configured for code name [%v] is conflicting with reserved error codes", format)
			assert.Equal(t, enum.Code(0), a[0])
			return dummyError1
		} else if dummyCodeNames[enum.CodeReservedCount-1] == a[1] {
			assert.Equal(t, "AppError code [%v] configured for code name [%v] is conflicting with reserved error codes", format)
			assert.Equal(t, enum.CodeReservedCount-1, a[0])
			return dummyError2
		} else if dummyHTTPStatusCodes[0] == a[1] {
			assert.Equal(t, "AppError code [%v] configured for HTTP status code [%v] is conflicting with reserved error codes", format)
			assert.Equal(t, enum.Code(0), a[0])
			return dummyError3
		} else if dummyHTTPStatusCodes[enum.CodeReservedCount-1] == a[1] {
			assert.Equal(t, "AppError code [%v] configured for HTTP status code [%v] is conflicting with reserved error codes", format)
			assert.Equal(t, enum.CodeReservedCount-1, a[0])
			return dummyError4
		}
		return nil
	}
	wrapSimpleErrorFuncExpected = 1
	wrapSimpleErrorFunc = func(innerErrors []error, messageFormat string, parameters ...interface{}) model.AppError {
		wrapSimpleErrorFuncCalled++
		assert.Equal(t, 4, len(innerErrors))
		assert.Contains(t, innerErrors, dummyError1)
		assert.Contains(t, innerErrors, dummyError2)
		assert.Contains(t, innerErrors, dummyError3)
		assert.Contains(t, innerErrors, dummyError4)
		assert.Equal(t, "Failed to initialize AppError customization", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyFinalError
	}

	// SUT + act
	var err = Initialize()

	// assert
	assert.Equal(t, dummyFinalError, err)

	// verify
	verifyAll(t)
}

func TestGetGeneralFailureError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeGeneralFailure, errorCode)
		assert.Equal(t, "An error occurred during execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetGeneralFailureError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetUnauthorized(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeUnauthorized, errorCode)
		assert.Equal(t, "Access denied due to authorization error", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetUnauthorized(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetInvalidOperation(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeInvalidOperation, errorCode)
		assert.Equal(t, "Operation (method) not allowed", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetInvalidOperation(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetBadRequestError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeBadRequest, errorCode)
		assert.Equal(t, "Request URI or body is invalid", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetBadRequestError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetNotFoundError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeNotFound, errorCode)
		assert.Equal(t, "Requested resource is not found in the storage", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetNotFoundError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetCircuitBreakError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeCircuitBreak, errorCode)
		assert.Equal(t, "Operation refused due to internal circuit break on correlation ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetCircuitBreakError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetOperationLockError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeOperationLock, errorCode)
		assert.Equal(t, "Operation refused due to mutex lock on correlation ID or trip ID", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetOperationLockError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetAccessForbiddenError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeAccessForbidden, errorCode)
		assert.Equal(t, "Operation failed due to access forbidden", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetAccessForbiddenError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetDataCorruptionError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeDataCorruption, errorCode)
		assert.Equal(t, "Operation failed due to internal storage data corruption", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetDataCorruptionError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetNotImplementedError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, expectedInnerError, innerErrors[0])
		assert.Equal(t, enum.CodeNotImplemented, errorCode)
		assert.Equal(t, "Operation failed due to internal business logic not implemented", messageFormat)
		assert.Equal(t, 0, len(parameters))
		return expectedResult
	}

	// SUT + act
	var appError = GetNotImplementedError(expectedInnerError)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetCustomError(t *testing.T) {
	// arrange
	var dummyErrorCode = enum.Code(rand.Intn(255))
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var dummyErrorMessage = "some error message"

	// mock
	createMock(t)

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, parameters ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return errors.New(dummyErrorMessage)
	}

	// SUT + act
	var appError, ok = GetCustomError(
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).(*appError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyErrorMessage, appError.error.Error())
	assert.Equal(t, dummyErrorCode, appError.code)
	assert.Equal(t, 0, len(appError.innerErrors))

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_NilInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors []error

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_EmptyInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_NoValidInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Empty(t, result)

	// verify
	verifyAll(t)
}

func TestCleanupInnerErrors_HasValidInnerErrors(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = cleanupInnerErrors(
		dummyInnerErrors,
	)

	// assert
	assert.Equal(t, 3, len(result))
	assert.Equal(t, dummyInnerError1, result[0])
	assert.Equal(t, dummyInnerError2, result[1])
	assert.Equal(t, dummyInnerError3, result[2])

	// verify
	verifyAll(t)
}

func TestWrapError_Empty(t *testing.T) {
	// arrange
	var dummyInnerErrors = []error{
		nil,
		nil,
		nil,
	}
	var dummyErrorCode = enum.Code(rand.Int())
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var cleanedInnerErrors = []error{}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// SUT + act
	var result = WrapError(
		dummyInnerErrors,
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestWrapError_NotEmpty(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyInnerErrors = []error{
		dummyInnerError1,
		nil,
		dummyInnerError2,
		nil,
		dummyInnerError3,
	}
	var dummyErrorCode = enum.Code(rand.Int())
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var dummyErrorMessage = "some error message"
	var cleanedInnerErrors = []error{
		dummyInnerError1,
		dummyInnerError2,
		dummyInnerError3,
	}

	// mock
	createMock(t)

	// expect
	cleanupInnerErrorsFuncExpected = 1
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		assert.Equal(t, dummyInnerErrors, innerErrors)
		return cleanedInnerErrors
	}

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, parameters ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyMessageFormat, format)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return errors.New(dummyErrorMessage)
	}

	// SUT + act
	var appError, ok = WrapError(
		dummyInnerErrors,
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).(*appError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyErrorMessage, appError.error.Error())
	assert.Equal(t, dummyErrorCode, appError.code)
	assert.Equal(t, 3, len(appError.innerErrors))
	assert.Equal(t, dummyInnerError1, appError.innerErrors[0])
	assert.Equal(t, dummyInnerError2, appError.innerErrors[1])
	assert.Equal(t, dummyInnerError3, appError.innerErrors[2])

	// verify
	verifyAll(t)
}

func TestWrapSimpleError(t *testing.T) {
	// arrange
	var dummyInnerError1 = errors.New("some random error 1")
	var dummyInnerError2 = errors.New("some random error 2")
	var dummyInnerError3 = errors.New("some random error 3")
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var expectedResult = &appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, 3, len(innerErrors))
		assert.Equal(t, dummyInnerError1, innerErrors[0])
		assert.Equal(t, dummyInnerError2, innerErrors[1])
		assert.Equal(t, dummyInnerError3, innerErrors[2])
		assert.Equal(t, enum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return expectedResult
	}

	// SUT + act
	var appError = WrapSimpleError(
		[]error{
			dummyInnerError1,
			dummyInnerError2,
			dummyInnerError3,
		},
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// assert
	assert.Equal(t, expectedResult, appError)

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_NonAppError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some dummy error")

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyError, errs[0])

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_NoInner(t *testing.T) {
	// arrange
	var dummyError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 0, len(errs))

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_WithInner(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("dummy inner error")
	var dummyError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummyInnerError},
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyInnerError, errs[0])

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_MultiLayer_NoInner(t *testing.T) {
	// arrange
	var dummyThirdLayerError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		nil,
	}
	var dummySecondLayerError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummyThirdLayerError},
	}
	var dummyError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummySecondLayerError},
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 0, len(errs))

	// verify
	verifyAll(t)
}

func TestGetInnermostErrors_AppError_MultiLayer_WithInner(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("dummy inner error")
	var dummyThirdLayerError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummyInnerError},
	}
	var dummySecondLayerError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummyThirdLayerError},
	}
	var dummyError = &appError{
		errors.New("dummy WebServiceTemplate error"),
		enum.CodeGeneralFailure,
		[]error{dummySecondLayerError},
	}

	// mock
	createMock(t)

	// SUT + act
	var errs = GetInnermostErrors(
		dummyError,
	)

	// assert
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, dummyInnerError, errs[0])

	// verify
	verifyAll(t)
}
