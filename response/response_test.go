package response

import (
	"errors"
	"go/types"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestGetAppError_NotAppError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetGeneralFailureError(dummyError)

	// mock
	createMock(t)

	// expect
	apperrorGetGeneralFailureErrorExpected = 1
	apperrorGetGeneralFailureError = func(innerError error) apperror.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, dummyError, innerError)
		return dummyAppError
	}

	// SUT + act
	var result = getAppError(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyAppError, result)

	// verify
	verifyAll(t)
}

func TestGetAppError_AppError(t *testing.T) {
	// arrange
	var dummyError = apperror.GetGeneralFailureError(errors.New("some error"))

	// mock
	createMock(t)

	// SUT + act
	var result = getAppError(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyError, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_GeneralFailure(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeGeneralFailure
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusInternalServerError, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_InvalidOperation(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeInvalidOperation
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusMethodNotAllowed, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_BadRequest(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeBadRequest
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusBadRequest, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_CircuitBreak(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeCircuitBreak
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusBadRequest, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_OperationLock(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeOperationLock
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusBadRequest, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_OtherCode(t *testing.T) {
	// arrange
	var dummyCode = apperror.Code(rand.Intn(math.MaxInt8) + 999)
	var dummyAppError = dummyAppError{
		t,
		&dummyCode,
		nil,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = getStatusCode(
		dummyAppError,
	)

	// assert
	assert.Equal(t, http.StatusInternalServerError, result)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_EmptyContent(t *testing.T) {
	// arrange
	var dummyResponseContent = ""

	// mock
	createMock(t)

	// expect
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyResponseContent, v)
		return ""
	}

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseContent,
	)

	// assert
	assert.Zero(t, result)
	assert.Equal(t, http.StatusNoContent, code)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_DirectNilContent(t *testing.T) {
	// arrange
	var dummyResponseContent types.Object

	// mock
	createMock(t)

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseContent,
	)

	// assert
	assert.Zero(t, result)
	assert.Equal(t, http.StatusNoContent, code)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_IndirectNilContent(t *testing.T) {
	// arrange
	var dummyNilObject types.Object
	var dummyResponseContent interface{} = dummyNilObject

	// mock
	createMock(t)

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseContent,
	)

	// assert
	assert.Zero(t, result)
	assert.Equal(t, http.StatusNoContent, code)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_ValidContent(t *testing.T) {
	// arrange
	var dummyResponseContent = "some response content"
	var dummyResponseMessage = "some response message"

	// mock
	createMock(t)

	// expect
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyResponseContent, v)
		return dummyResponseMessage
	}

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseContent,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, result)
	assert.Equal(t, http.StatusOK, code)

	// verify
	verifyAll(t)
}

func TestGenerateErrorResponse(t *testing.T) {
	// arrange
	var codeInteger = rand.Intn(math.MaxInt8)
	var expectedCode = apperror.Code(codeInteger)
	var expectedType = expectedCode.String()
	var expectedMessages = []string{"some", "message", "array"}
	var dummyAppError = dummyAppError{
		t,
		&expectedCode,
		&expectedMessages,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = generateErrorResponse(
		dummyAppError,
	)

	// assert
	assert.Equal(t, codeInteger, result.Code)
	assert.Equal(t, expectedType, result.Type)
	assert.Equal(t, expectedMessages, result.Messages)

	// verify
	verifyAll(t)
}

func TestCreateErrorResponse(t *testing.T) {
	// arrange
	var dummyAppError = dummyAppError{
		t,
		nil,
		nil,
	}
	var dummyErrorResponseModel = errorResponseModel{
		Code:     rand.Int(),
		Type:     "some type",
		Messages: []string{"some", "message", "array"},
	}
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	generateErrorResponseFuncExpected = 1
	generateErrorResponseFunc = func(appError apperror.AppError) errorResponseModel {
		generateErrorResponseFuncCalled++
		assert.Equal(t, dummyAppError, appError)
		return dummyErrorResponseModel
	}
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyErrorResponseModel, v)
		return dummyResponseMessage
	}
	getStatusCodeFuncExpected = 1
	getStatusCodeFunc = func(appError apperror.AppError) int {
		getStatusCodeFuncCalled++
		assert.Equal(t, dummyAppError, appError)
		return dummyStatusCode
	}

	// SUT + act
	var result, code = createErrorResponse(
		dummyAppError,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, result)
	assert.Equal(t, dummyStatusCode, code)

	// verify
	verifyAll(t)
}

func TestWriteResponse(t *testing.T) {
	// arrange
	var dummyHeader = make(http.Header)
	var dummyStatusCode = rand.Int()
	var dummyResponseMessage = "some response message"
	var dummyResponseBytes = []byte(dummyResponseMessage)
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		&dummyHeader,
		&dummyStatusCode,
		&dummyResponseBytes,
	}

	// mock
	createMock(t)

	// SUT + act
	writeResponse(
		dummyResponseWriter,
		dummyStatusCode,
		dummyResponseMessage,
	)

	// verify
	verifyAll(t)
}

func TestOk(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseContent = "some response content"
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		nil,
		nil,
		nil,
	}
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()
	var dummyStatusCodeString = strconv.Itoa(dummyStatusCode)

	// mock
	createMock(t)

	// expect
	strconvItoaExpected = 1
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	createOkResponseFuncExpected = 1
	createOkResponseFunc = func(responseContent interface{}) (string, int) {
		createOkResponseFuncCalled++
		assert.Equal(t, dummyResponseContent, responseContent)
		return dummyResponseMessage, dummyStatusCode
	}
	loggerAPIResponseExpected = 1
	loggerAPIResponse = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIResponseCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "response", category)
		assert.Equal(t, dummyStatusCodeString, subcategory)
		assert.Equal(t, dummyResponseMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	writeResponseFuncExpected = 1
	writeResponseFunc = func(responseWriter http.ResponseWriter, statusCode int, responseMessage string) {
		writeResponseFuncCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyStatusCode, statusCode)
		assert.Equal(t, dummyResponseMessage, responseMessage)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "response", category)
		assert.Equal(t, "Ok", subcategory)
		assert.Equal(t, "", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	Ok(
		dummySessionID,
		dummyResponseContent,
		dummyResponseWriter,
	)

	// verify
	verifyAll(t)
}

func TestError(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetGeneralFailureError(dummyError)
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		nil,
		nil,
		nil,
	}
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()
	var dummyStatusCodeString = strconv.Itoa(dummyStatusCode)

	// mock
	createMock(t)

	// expect
	strconvItoaExpected = 1
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	getAppErrorFuncExpected = 1
	getAppErrorFunc = func(err error) apperror.AppError {
		getAppErrorFuncCalled++
		assert.Equal(t, dummyError, err)
		return dummyAppError
	}
	createErrorResponseFuncExpected = 1
	createErrorResponseFunc = func(appError apperror.AppError) (string, int) {
		createErrorResponseFuncCalled++
		assert.Equal(t, dummyAppError, appError)
		return dummyResponseMessage, dummyStatusCode
	}
	loggerAPIResponseExpected = 1
	loggerAPIResponse = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIResponseCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "response", category)
		assert.Equal(t, dummyStatusCodeString, subcategory)
		assert.Equal(t, dummyResponseMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
	}
	writeResponseFuncExpected = 1
	writeResponseFunc = func(responseWriter http.ResponseWriter, statusCode int, responseMessage string) {
		writeResponseFuncCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyStatusCode, statusCode)
		assert.Equal(t, dummyResponseMessage, responseMessage)
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "response", category)
		assert.Equal(t, "Error", subcategory)
		assert.Equal(t, "", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	Error(
		dummySessionID,
		dummyError,
		dummyResponseWriter,
	)

	// verify
	verifyAll(t)
}
