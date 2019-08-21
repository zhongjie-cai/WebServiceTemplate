package response

import (
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
	assert.Equal(t, http.StatusForbidden, result)

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
	assert.Equal(t, http.StatusLocked, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_AccessForbidden(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeAccessForbidden
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
	assert.Equal(t, http.StatusForbidden, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_DataCorruption(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeDataCorruption
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
	assert.Equal(t, http.StatusConflict, result)

	// verify
	verifyAll(t)
}

func TestGetStatusCode_NotImplemented(t *testing.T) {
	// arrange
	var dummyCode = apperror.CodeNotImplemented
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
	assert.Equal(t, http.StatusNotImplemented, result)

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
	var dummyResponseObject = ""

	// mock
	createMock(t)

	// expect
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyResponseObject, v)
		return ""
	}

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseObject,
	)

	// assert
	assert.Zero(t, result)
	assert.Equal(t, http.StatusNoContent, code)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_DirectNilContent(t *testing.T) {
	// arrange
	var dummyResponseObject types.Object

	// mock
	createMock(t)

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseObject,
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
	var dummyResponseObject interface{} = dummyNilObject

	// mock
	createMock(t)

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseObject,
	)

	// assert
	assert.Zero(t, result)
	assert.Equal(t, http.StatusNoContent, code)

	// verify
	verifyAll(t)
}

func TestCreateOkResponse_ValidContent(t *testing.T) {
	// arrange
	var dummyResponseObject = "some response content"
	var dummyResponseMessage = "some response message"

	// mock
	createMock(t)

	// expect
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyResponseObject, v)
		return dummyResponseMessage
	}

	// SUT + act
	var result, code = createOkResponse(
		dummyResponseObject,
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

func TestWrite_Ok(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseObject = "some response content"
	var dummyResponseError apperror.AppError
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
	sessionGetResponseWriterExpected = 1
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseWriter
	}
	createOkResponseFuncExpected = 1
	createOkResponseFunc = func(responseContent interface{}) (string, int) {
		createOkResponseFuncCalled++
		assert.Equal(t, dummyResponseObject, responseContent)
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
		assert.Equal(t, "Write", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyStatusCode, parameters[0])
	}

	// SUT + act
	Write(
		dummySessionID,
		dummyResponseObject,
		dummyResponseError,
	)

	// verify
	verifyAll(t)
}

func TestWrite_Error(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseObject = "some response content"
	var dummyResponseError = apperror.GetGeneralFailureError(nil)
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
	sessionGetResponseWriterExpected = 1
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseWriter
	}
	createErrorResponseFuncExpected = 1
	createErrorResponseFunc = func(appError apperror.AppError) (string, int) {
		createErrorResponseFuncCalled++
		assert.Equal(t, dummyResponseError, appError)
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
		assert.Equal(t, "Write", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyStatusCode, parameters[0])
	}

	// SUT + act
	Write(
		dummySessionID,
		dummyResponseObject,
		dummyResponseError,
	)

	// verify
	verifyAll(t)
}

func TestOverride(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyHTTPRequest, _ = http.NewRequest(http.MethodGet, "http://localhost", nil)
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		nil,
		nil,
		nil,
	}
	var dummyCallbackExpected int
	var dummyCallbackCalled int
	var dummyCallback func(*http.Request, http.ResponseWriter)

	// mock
	createMock(t)

	// expect
	sessionGetRequestExpected = 1
	sessionGetRequest = func(sessionID uuid.UUID) *http.Request {
		sessionGetRequestCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyHTTPRequest
	}
	sessionGetResponseWriterExpected = 1
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseWriter
	}
	dummyCallbackExpected = 1
	dummyCallback = func(httpRequest *http.Request, responseWriter http.ResponseWriter) {
		dummyCallbackCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyResponseWriter, responseWriter)
	}
	sessionClearResponseWriterExpected = 1
	sessionClearResponseWriter = func(sessionID uuid.UUID) {
		sessionClearResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
	}

	// SUT + act
	Override(
		dummySessionID,
		dummyCallback,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCallbackExpected, dummyCallbackCalled, "Unexpected number of calls to dummyCallback")
}
