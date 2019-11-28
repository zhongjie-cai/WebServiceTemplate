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
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

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

func TestGetAppError_IsAppError(t *testing.T) {
	// arrange
	var dummyError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// SUT + act
	var err = getAppError(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestGetAppError_IsNotAppError(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "")

	// mock
	createMock(t)

	// expect
	apperrorGetGeneralFailureErrorExpected = 1
	apperrorGetGeneralFailureError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		return dummyAppError
	}

	// SUT + act
	var err = getAppError(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestGenerateErrorResponse(t *testing.T) {
	// arrange
	var codeInteger = rand.Intn(math.MaxInt8)
	var expectedCode = apperrorEnum.Code(codeInteger).String()
	var expectedMessages = []string{"some", "message", "array"}
	var expectedExtraData = map[string]string{"foo": "bar", "test": "me"}
	var dummyAppError = &dummyAppError{
		t,
		&expectedCode,
		nil,
		&expectedMessages,
		&expectedExtraData,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = generateErrorResponse(
		dummyAppError,
	)

	// assert
	assert.Equal(t, expectedCode, result.Code)
	assert.Equal(t, expectedMessages, result.Messages)
	assert.Equal(t, expectedExtraData, result.ExtraData)

	// verify
	verifyAll(t)
}

func TestCreateErrorResponse(t *testing.T) {
	// arrange
	var dummyError = errors.New("some error")
	var dummyHTTPStatusCode = rand.Intn(1000)
	var dummyAppError = &dummyAppError{
		t,
		nil,
		&dummyHTTPStatusCode,
		nil,
		nil,
	}
	var dummyErrorResponseModel = errorResponseModel{
		Code:      "some type",
		Messages:  []string{"some", "message", "array"},
		ExtraData: map[string]string{"foo": "bar", "test": "me"},
	}
	var dummyResponseMessage = "some response message"

	// mock
	createMock(t)

	// expect
	getAppErrorFuncExpected = 1
	getAppErrorFunc = func(err error) apperrorModel.AppError {
		getAppErrorFuncCalled++
		assert.Equal(t, dummyError, err)
		return dummyAppError
	}
	generateErrorResponseFuncExpected = 1
	generateErrorResponseFunc = func(appError apperrorModel.AppError) errorResponseModel {
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

	// SUT + act
	var result, code = createErrorResponse(
		dummyError,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, result)
	assert.Equal(t, dummyHTTPStatusCode, code)

	// verify
	verifyAll(t)
}

func TestWriteResponse(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyHeader = make(http.Header)
	var dummyStatusCode = rand.Int()
	var dummyStatusCodeString = strconv.Itoa(dummyStatusCode)
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

	// expect
	strconvItoaExpected = 1
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
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
	writeResponse(
		dummySessionID,
		dummyResponseWriter,
		dummyStatusCode,
		dummyResponseMessage,
	)

	// verify
	verifyAll(t)
}

func TestConstructResponse_Error_WithCustomization(t *testing.T) {
	// arrange
	var dummyResponseObject = "some response content"
	var dummyResponseError = errors.New("some response error")
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	customizationCreateErrorResponseFuncExpected = 1
	customization.CreateErrorResponseFunc = func(err error) (string, int) {
		customizationCreateErrorResponseFuncCalled++
		assert.Equal(t, dummyResponseError, err)
		return dummyResponseMessage, dummyStatusCode
	}

	// SUT + act
	var message, code = constructResponse(
		dummyResponseObject,
		dummyResponseError,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, message)
	assert.Equal(t, dummyStatusCode, code)

	// verify
	verifyAll(t)
}

func TestConstructResponse_Error_NoCustomization(t *testing.T) {
	// arrange
	var dummyResponseObject = "some response content"
	var dummyResponseError = errors.New("some response error")
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	createErrorResponseFuncExpected = 1
	createErrorResponseFunc = func(err error) (string, int) {
		createErrorResponseFuncCalled++
		assert.Equal(t, dummyResponseError, err)
		return dummyResponseMessage, dummyStatusCode
	}

	// SUT + act
	var message, code = constructResponse(
		dummyResponseObject,
		dummyResponseError,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, message)
	assert.Equal(t, dummyStatusCode, code)

	// verify
	verifyAll(t)
}

func TestConstructResponse_NoError(t *testing.T) {
	// arrange
	var dummyResponseObject = "some response content"
	var dummyResponseError apperrorModel.AppError
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	createOkResponseFuncExpected = 1
	createOkResponseFunc = func(responseContent interface{}) (string, int) {
		createOkResponseFuncCalled++
		assert.Equal(t, dummyResponseObject, responseContent)
		return dummyResponseMessage, dummyStatusCode
	}

	// SUT + act
	var message, code = constructResponse(
		dummyResponseObject,
		dummyResponseError,
	)

	// assert
	assert.Equal(t, dummyResponseMessage, message)
	assert.Equal(t, dummyStatusCode, code)

	// verify
	verifyAll(t)
}

func TestWrite_NotOverrided(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseObject = "some response content"
	var dummyResponseError = errors.New("some response error")
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		nil,
		nil,
		nil,
	}
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	sessionGetResponseWriterExpected = 1
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseWriter
	}
	constructResponseFuncExpected = 1
	constructResponseFunc = func(responseObject interface{}, responseError error) (string, int) {
		constructResponseFuncCalled++
		assert.Equal(t, dummyResponseObject, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
		return dummyResponseMessage, dummyStatusCode
	}
	writeResponseFuncExpected = 1
	writeResponseFunc = func(sessionID uuid.UUID, responseWriter http.ResponseWriter, statusCode int, responseMessage string) {
		writeResponseFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyStatusCode, statusCode)
		assert.Equal(t, dummyResponseMessage, responseMessage)
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

func TestWrite_Overrided(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyResponseObject = overrideResponse{}
	var dummyResponseError = errors.New("some response error")
	var dummyResponseWriter = &dummyResponseWriter{
		t,
		nil,
		nil,
		nil,
	}
	var dummyResponseMessage = "some response message"
	var dummyStatusCode = rand.Int()

	// mock
	createMock(t)

	// expect
	sessionGetResponseWriterExpected = 1
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyResponseWriter
	}
	constructResponseFuncExpected = 1
	constructResponseFunc = func(responseObject interface{}, responseError error) (string, int) {
		constructResponseFuncCalled++
		assert.Equal(t, dummyResponseObject, responseObject)
		assert.Equal(t, dummyResponseError, responseError)
		return dummyResponseMessage, dummyStatusCode
	}
	loggerAPIExitExpected = 1
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, "response", category)
		assert.Equal(t, "Write", subcategory)
		assert.Equal(t, "Overrided", messageFormat)
		assert.Equal(t, 0, len(parameters))
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
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
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

	// SUT + act
	var result, err = Override(
		dummySessionID,
		dummyCallback,
	)

	// assert
	assert.IsType(t, overrideResponse{}, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, dummyCallbackExpected, dummyCallbackCalled, "Unexpected number of calls to dummyCallback")
}
