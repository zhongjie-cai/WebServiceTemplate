package response

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
)

var (
	strconvItoaExpected                    int
	strconvItoaCalled                      int
	jsonutilMarshalIgnoreErrorExpected     int
	jsonutilMarshalIgnoreErrorCalled       int
	apperrorGetGeneralFailureErrorExpected int
	apperrorGetGeneralFailureErrorCalled   int
	loggerAPIResponseExpected              int
	loggerAPIResponseCalled                int
	loggerAPIExitExpected                  int
	loggerAPIExitCalled                    int
	getStatusCodeFuncExpected              int
	getStatusCodeFuncCalled                int
	getAppErrorFuncExpected                int
	getAppErrorFuncCalled                  int
	writeResponseFuncExpected              int
	writeResponseFuncCalled                int
	generateErrorResponseFuncExpected      int
	generateErrorResponseFuncCalled        int
	createOkResponseFuncExpected           int
	createOkResponseFuncCalled             int
	createErrorResponseFuncExpected        int
	createErrorResponseFuncCalled          int
)

func createMock(t *testing.T) {
	strconvItoaExpected = 0
	strconvItoaCalled = 0
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return ""
	}
	jsonutilMarshalIgnoreErrorExpected = 0
	jsonutilMarshalIgnoreErrorCalled = 0
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		return ""
	}
	apperrorGetGeneralFailureErrorExpected = 0
	apperrorGetGeneralFailureErrorCalled = 0
	apperrorGetGeneralFailureError = func(innerError error) apperror.AppError {
		apperrorGetGeneralFailureErrorCalled++
		return nil
	}
	loggerAPIResponseExpected = 0
	loggerAPIResponseCalled = 0
	loggerAPIResponse = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIResponseCalled++
	}
	loggerAPIExitExpected = 0
	loggerAPIExitCalled = 0
	loggerAPIExit = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIExitCalled++
	}
	getStatusCodeFuncExpected = 0
	getStatusCodeFuncCalled = 0
	getStatusCodeFunc = func(appError apperror.AppError) int {
		getStatusCodeFuncCalled++
		return 0
	}
	getAppErrorFuncExpected = 0
	getAppErrorFuncCalled = 0
	getAppErrorFunc = func(err error) apperror.AppError {
		getAppErrorFuncCalled++
		return nil
	}
	writeResponseFuncExpected = 0
	writeResponseFuncCalled = 0
	writeResponseFunc = func(responseWriter http.ResponseWriter, statusCode int, responseMessage string) {
		writeResponseFuncCalled++
	}
	generateErrorResponseFuncExpected = 0
	generateErrorResponseFuncCalled = 0
	generateErrorResponseFunc = func(appError apperror.AppError) errorResponseModel {
		generateErrorResponseFuncCalled++
		return errorResponseModel{}
	}
	createOkResponseFuncExpected = 0
	createOkResponseFuncCalled = 0
	createOkResponseFunc = func(responseContent interface{}) (string, int) {
		createOkResponseFuncCalled++
		return "", 0
	}
	createErrorResponseFuncExpected = 0
	createErrorResponseFuncCalled = 0
	createErrorResponseFunc = func(appError apperror.AppError) (string, int) {
		createErrorResponseFuncCalled++
		return "", 0
	}
}

func verifyAll(t *testing.T) {
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected method call to strconvItoa")
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected method call to jsonutilMarshalIgnoreError")
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	assert.Equal(t, apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled, "Unexpected method call to apperrorGetGeneralFailureError")
	loggerAPIResponse = logger.APIResponse
	assert.Equal(t, loggerAPIResponseExpected, loggerAPIResponseCalled, "Unexpected method call to loggerAPIResponse")
	loggerAPIExit = logger.APIExit
	assert.Equal(t, loggerAPIExitExpected, loggerAPIExitCalled, "Unexpected method call to loggerAPIExit")
	getStatusCodeFunc = getStatusCode
	assert.Equal(t, getStatusCodeFuncExpected, getStatusCodeFuncCalled, "Unexpected method call to getStatusCodeFunc")
	getAppErrorFunc = getAppError
	assert.Equal(t, getAppErrorFuncExpected, getAppErrorFuncCalled, "Unexpected method call to getAppErrorFunc")
	writeResponseFunc = writeResponse
	assert.Equal(t, writeResponseFuncExpected, writeResponseFuncCalled, "Unexpected method call to writeResponseFunc")
	generateErrorResponseFunc = generateErrorResponse
	assert.Equal(t, generateErrorResponseFuncExpected, generateErrorResponseFuncCalled, "Unexpected method call to generateErrorResponseFunc")
	createOkResponseFunc = createOkResponse
	assert.Equal(t, createOkResponseFuncExpected, createOkResponseFuncCalled, "Unexpected method call to createOkResponseFunc")
	createErrorResponseFunc = createErrorResponse
	assert.Equal(t, createErrorResponseFuncExpected, createErrorResponseFuncCalled, "Unexpected method call to createErrorResponseFunc")
}
