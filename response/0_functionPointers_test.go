package response

import (
	"fmt"
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
	if strconvItoaExpected != strconvItoaCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to strconvItoa, expected %v, actual %v", strconvItoaExpected, strconvItoaCalled))
	}
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	if jsonutilMarshalIgnoreErrorExpected != jsonutilMarshalIgnoreErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to jsonutilMarshalIgnoreError, expected %v, actual %v", jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled))
	}
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	if apperrorGetGeneralFailureErrorExpected != apperrorGetGeneralFailureErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorGetGeneralFailureError, expected %v, actual %v", apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled))
	}
	loggerAPIResponse = logger.APIResponse
	if loggerAPIResponseExpected != loggerAPIResponseCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAPIResponse, expected %v, actual %v", loggerAPIResponseExpected, loggerAPIResponseCalled))
	}
	loggerAPIExit = logger.APIExit
	if loggerAPIExitExpected != loggerAPIExitCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAPIExit, expected %v, actual %v", loggerAPIExitExpected, loggerAPIExitCalled))
	}
	getStatusCodeFunc = getStatusCode
	if getStatusCodeFuncExpected != getStatusCodeFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getStatusCodeFunc, expected %v, actual %v", getStatusCodeFuncExpected, getStatusCodeFuncCalled))
	}
	getAppErrorFunc = getAppError
	if getAppErrorFuncExpected != getAppErrorFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getAppErrorFunc, expected %v, actual %v", getAppErrorFuncExpected, getAppErrorFuncCalled))
	}
	writeResponseFunc = writeResponse
	if writeResponseFuncExpected != writeResponseFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to writeResponseFunc, expected %v, actual %v", writeResponseFuncExpected, writeResponseFuncCalled))
	}
	generateErrorResponseFunc = generateErrorResponse
	if generateErrorResponseFuncExpected != generateErrorResponseFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to generateErrorResponseFunc, expected %v, actual %v", generateErrorResponseFuncExpected, generateErrorResponseFuncCalled))
	}
	createOkResponseFunc = createOkResponse
	if createOkResponseFuncExpected != createOkResponseFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to createOkResponseFunc, expected %v, actual %v", createOkResponseFuncExpected, createOkResponseFuncCalled))
	}
	createErrorResponseFunc = createErrorResponse
	if createErrorResponseFuncExpected != createErrorResponseFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to createErrorResponseFunc, expected %v, actual %v", createErrorResponseFuncExpected, createErrorResponseFuncCalled))
	}
}
