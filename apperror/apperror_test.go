package apperror

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeEnumString_UnknownNegative(t *testing.T) {
	// arrange
	var testCode Code

	// mock
	createMock(t)

	// SUT
	testCode = -1

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unknown", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GeneralFailure(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeGeneralFailure

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "GeneralFailure", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_Unauthorized(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeUnauthorized

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unauthorized", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_InvalidOperation(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeInvalidOperation

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "InvalidOperation", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_BadRequest(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeBadRequest

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "BadRequest", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_NotFound(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeNotFound

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "NotFound", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_CircuitBreak(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeCircuitBreak

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "CircuitBreak", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_OperationLock(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeOperationLock

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "OperationLock", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_AccessForbidden(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeAccessForbidden

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "AccessForbidden", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GetDataCorruption(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeDataCorruption

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "DataCorruption", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GetNotImplemented(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeNotImplemented

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "NotImplemented", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_UnknownTooBig(t *testing.T) {
	// arrange
	var testCode Code

	// mock
	createMock(t)

	// SUT
	testCode = 999

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unknown", convertedString)

	// verify
	verifyAll(t)
}

func TestAppErrorGetCode(t *testing.T) {
	// arrange
	var expectedError = errors.New("dummy error")
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
		expectedError,
		expectedCode,
		expectedInnerErrors,
	}

	// act
	var code = appError.Code()

	// assert
	assert.Equal(t, expectedCode, code)

	// verify
	verifyAll(t)
}

func TestAppErrorGetError_NoInner(t *testing.T) {
	// arrange
	var dummyMessage = "dummy error"
	var expectedError = errors.New(dummyMessage)
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure

	// mock
	createMock(t)

	// SUT
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure
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
	var appError = appError{
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
	var expectedCode = CodeGeneralFailure
	var expectedMessage = "(GeneralFailure)dummy error"
	var dummyInnerErrorMessage = "dummy inner error"
	var dummyInnerMostErrorMessage = "dummy inner most error"
	var expectedInnerError1 = errors.New("dummy inner error 1")
	var expectedInnerError2 = appError{
		errors.New(dummyInnerErrorMessage),
		CodeGeneralFailure,
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
	var appError = appError{
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

func TestGetGeneralFailureError(t *testing.T) {
	// arrange
	var expectedInnerError = errors.New("dummy inner error")
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeGeneralFailure, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeUnauthorized, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeInvalidOperation, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeBadRequest, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeNotFound, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeCircuitBreak, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeOperationLock, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeAccessForbidden, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeDataCorruption, errorCode)
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
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, expectedInnerError, innerError)
		assert.Equal(t, CodeNotImplemented, errorCode)
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

func TestConsolidateAllErrors_NilList(t *testing.T) {
	// arrange
	var baseErrorMessage = "some base error message"

	// mock
	createMock(t)

	// SUT + act
	var err = ConsolidateAllErrors(baseErrorMessage, nil)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestConsolidateAllErrors_EmptyList(t *testing.T) {
	// arrange
	var baseErrorMessage = "some base error message"
	var allErrors = []error{}

	// mock
	createMock(t)

	// SUT + act
	var err = ConsolidateAllErrors(baseErrorMessage, allErrors...)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestConsolidateAllErrors_ListOfNil(t *testing.T) {
	// arrange
	var baseErrorMessage = "some base error message"
	var allErrors = []error{
		nil,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var err = ConsolidateAllErrors(baseErrorMessage, allErrors...)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestConsolidateAllErrors_ListOfEmptyErrors(t *testing.T) {
	// arrange
	var baseErrorMessage = "some base error message"
	var allErrors = []error{
		errors.New(""),
		nil,
		errors.New(""),
	}
	var dummyMessageFormat = "Unknown Error | Unknown Error"
	var dummyAppError = GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}
	wrapSimpleErrorFuncExpected = 1
	wrapSimpleErrorFunc = func(innerError error, messageFormat string, parameters ...interface{}) AppError {
		wrapSimpleErrorFuncCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, dummyMessageFormat, innerError.Error())
		assert.Equal(t, baseErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = ConsolidateAllErrors(baseErrorMessage, allErrors...)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestConsolidateAllErrors_ListOfValidErrors(t *testing.T) {
	// arrange
	var baseErrorMessage = "some base error message"
	var errorMessage1 = "dummy error 1"
	var errorMessage3 = "dummy error 3"
	var allErrors = []error{
		errors.New(errorMessage1),
		nil,
		errors.New(errorMessage3),
	}
	var dummyMessageFormat = errorMessage1 + " | " + errorMessage3
	var dummyAppError = GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}
	wrapSimpleErrorFuncExpected = 1
	wrapSimpleErrorFunc = func(innerError error, messageFormat string, parameters ...interface{}) AppError {
		wrapSimpleErrorFuncCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, dummyMessageFormat, innerError.Error())
		assert.Equal(t, baseErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = ConsolidateAllErrors(baseErrorMessage, allErrors...)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestWrapError(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("some random error")
	var dummyErrorCode = Code(rand.Int())
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
	var appError, ok = WrapError(
		dummyInnerError,
		dummyErrorCode,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).(appError)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummyErrorMessage, appError.error.Error())
	assert.Equal(t, dummyErrorCode, appError.code)
	assert.Equal(t, 1, len(appError.innerErrors))
	assert.Equal(t, dummyInnerError.Error(), appError.innerErrors[0].Error())

	// verify
	verifyAll(t)
}

func TestWrapSimpleError(t *testing.T) {
	// arrange
	var dummyInnerError = errors.New("some random error")
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = 123
	var dummyParameter3 = errors.New("dummy")
	var expectedResult = appError{}

	// mock
	createMock(t)

	// expect
	wrapErrorFuncExpected = 1
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		assert.Equal(t, dummyInnerError, innerError)
		assert.Equal(t, CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
		return expectedResult
	}

	// SUT + act
	var appError = WrapSimpleError(
		dummyInnerError,
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
	var dummyError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
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
	var dummyError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
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
	var dummyThirdLayerError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		nil,
	}
	var dummySecondLayerError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyThirdLayerError},
	}
	var dummyError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
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
	var dummyThirdLayerError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyInnerError},
	}
	var dummySecondLayerError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
		[]error{dummyThirdLayerError},
	}
	var dummyError = appError{
		errors.New("dummy WebServiceTemplate error"),
		CodeGeneralFailure,
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
