package response

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
	strconvItoaExpected                          int
	strconvItoaCalled                            int
	jsonutilMarshalIgnoreErrorExpected           int
	jsonutilMarshalIgnoreErrorCalled             int
	apperrorGetGeneralFailureErrorExpected       int
	apperrorGetGeneralFailureErrorCalled         int
	loggerAPIResponseExpected                    int
	loggerAPIResponseCalled                      int
	loggerAPIExitExpected                        int
	loggerAPIExitCalled                          int
	sessionGetRequestExpected                    int
	sessionGetRequestCalled                      int
	sessionGetResponseWriterExpected             int
	sessionGetResponseWriterCalled               int
	sessionClearResponseWriterExpected           int
	sessionClearResponseWriterCalled             int
	getStatusCodeFuncExpected                    int
	getStatusCodeFuncCalled                      int
	writeResponseFuncExpected                    int
	writeResponseFuncCalled                      int
	generateErrorResponseFuncExpected            int
	generateErrorResponseFuncCalled              int
	createOkResponseFuncExpected                 int
	createOkResponseFuncCalled                   int
	createErrorResponseFuncExpected              int
	createErrorResponseFuncCalled                int
	customizationCreateErrorResponseFuncExpected int
	customizationCreateErrorResponseFuncCalled   int
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
	sessionGetRequestExpected = 0
	sessionGetRequestCalled = 0
	sessionGetRequest = func(sessionID uuid.UUID) *http.Request {
		sessionGetRequestCalled++
		return nil
	}
	sessionGetResponseWriterExpected = 0
	sessionGetResponseWriterCalled = 0
	sessionGetResponseWriter = func(sessionID uuid.UUID) http.ResponseWriter {
		sessionGetResponseWriterCalled++
		return nil
	}
	sessionClearResponseWriterExpected = 0
	sessionClearResponseWriterCalled = 0
	sessionClearResponseWriter = func(sessionID uuid.UUID) {
		sessionClearResponseWriterCalled++
	}
	getStatusCodeFuncExpected = 0
	getStatusCodeFuncCalled = 0
	getStatusCodeFunc = func(appError apperror.AppError) int {
		getStatusCodeFuncCalled++
		return 0
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
	customizationCreateErrorResponseFuncExpected = 0
	customizationCreateErrorResponseFuncCalled = 0
	customization.CreateErrorResponseFunc = nil
}

func verifyAll(t *testing.T) {
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected number of calls to strconvItoa")
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected number of calls to jsonutilMarshalIgnoreError")
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	assert.Equal(t, apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled, "Unexpected number of calls to apperrorGetGeneralFailureError")
	loggerAPIResponse = logger.APIResponse
	assert.Equal(t, loggerAPIResponseExpected, loggerAPIResponseCalled, "Unexpected number of calls to loggerAPIResponse")
	loggerAPIExit = logger.APIExit
	assert.Equal(t, loggerAPIExitExpected, loggerAPIExitCalled, "Unexpected number of calls to loggerAPIExit")
	sessionGetRequest = session.GetRequest
	assert.Equal(t, sessionGetRequestExpected, sessionGetRequestCalled, "Unexpected number of calls to sessionGetRequest")
	sessionGetResponseWriter = session.GetResponseWriter
	assert.Equal(t, sessionGetResponseWriterExpected, sessionGetResponseWriterCalled, "Unexpected number of calls to sessionGetResponseWriter")
	sessionClearResponseWriter = session.ClearResponseWriter
	assert.Equal(t, sessionClearResponseWriterExpected, sessionClearResponseWriterCalled, "Unexpected number of calls to sessionClearResponseWriter")
	getStatusCodeFunc = getStatusCode
	assert.Equal(t, getStatusCodeFuncExpected, getStatusCodeFuncCalled, "Unexpected number of calls to getStatusCodeFunc")
	writeResponseFunc = writeResponse
	assert.Equal(t, writeResponseFuncExpected, writeResponseFuncCalled, "Unexpected number of calls to writeResponseFunc")
	generateErrorResponseFunc = generateErrorResponse
	assert.Equal(t, generateErrorResponseFuncExpected, generateErrorResponseFuncCalled, "Unexpected number of calls to generateErrorResponseFunc")
	createOkResponseFunc = createOkResponse
	assert.Equal(t, createOkResponseFuncExpected, createOkResponseFuncCalled, "Unexpected number of calls to createOkResponseFunc")
	createErrorResponseFunc = createErrorResponse
	assert.Equal(t, createErrorResponseFuncExpected, createErrorResponseFuncCalled, "Unexpected number of calls to createErrorResponseFunc")
	customization.CreateErrorResponseFunc = nil
	assert.Equal(t, customizationCreateErrorResponseFuncExpected, customizationCreateErrorResponseFuncCalled, "Unexpected number of calls to customization.CreateErrorResponseFunc")
}

// mock structs
type dummyResponseWriter struct {
	t               *testing.T
	expectedHeader  *http.Header
	expectedCode    *int
	expectedContent *[]byte
}

func (drw *dummyResponseWriter) Header() http.Header {
	if drw.expectedHeader == nil {
		assert.Fail(drw.t, "Unexpected number of calls to Header")
		return nil
	}
	return *drw.expectedHeader
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	if drw.expectedCode == nil {
		assert.Fail(drw.t, "Unexpected number of calls to WriteHeader")
	} else {
		assert.Equal(drw.t, *drw.expectedCode, statusCode)
	}
}

func (drw *dummyResponseWriter) Write(bytes []byte) (int, error) {
	if drw.expectedContent == nil {
		assert.Fail(drw.t, "Unexpected number of calls to Write")
	} else {
		assert.Equal(drw.t, *drw.expectedContent, bytes)
	}
	return 0, nil
}

type dummyAppError struct {
	t                *testing.T
	expectedCode     *apperror.Code
	expectedMessages *[]string
}

func (dae dummyAppError) Code() apperror.Code {
	if dae.expectedCode == nil {
		assert.Fail(dae.t, "Unexpected number of calls to Code")
		return apperror.Code(-1)
	}
	return *dae.expectedCode
}

func (dae dummyAppError) Error() string {
	assert.Fail(dae.t, "Unexpected number of calls to Error")
	return ""
}

func (dae dummyAppError) InnerErrors() []error {
	assert.Fail(dae.t, "Unexpected number of calls to InnerErrors")
	return nil
}

func (dae dummyAppError) Messages() []string {
	if dae.expectedMessages == nil {
		assert.Fail(dae.t, "Unexpected number of calls to Messages")
		return nil
	}
	return *dae.expectedMessages
}
