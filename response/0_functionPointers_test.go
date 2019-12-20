package response

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
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
	httpStatusTextExpected                       int
	httpStatusTextCalled                         int
	writeResponseFuncExpected                    int
	writeResponseFuncCalled                      int
	getAppErrorFuncExpected                      int
	getAppErrorFuncCalled                        int
	generateErrorResponseFuncExpected            int
	generateErrorResponseFuncCalled              int
	createOkResponseFuncExpected                 int
	createOkResponseFuncCalled                   int
	createErrorResponseFuncExpected              int
	createErrorResponseFuncCalled                int
	constructResponseFuncExpected                int
	constructResponseFuncCalled                  int
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
	apperrorGetGeneralFailureError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetGeneralFailureErrorCalled++
		return nil
	}
	loggerAPIResponseExpected = 0
	loggerAPIResponseCalled = 0
	loggerAPIResponse = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIResponseCalled++
	}
	httpStatusTextExpected = 0
	httpStatusTextCalled = 0
	httpStatusText = func(code int) string {
		httpStatusTextCalled++
		return ""
	}
	writeResponseFuncExpected = 0
	writeResponseFuncCalled = 0
	writeResponseFunc = func(session sessionModel.Session, statusCode int, responseMessage string) {
		writeResponseFuncCalled++
	}
	getAppErrorFuncExpected = 0
	getAppErrorFuncCalled = 0
	getAppErrorFunc = func(err error) apperrorModel.AppError {
		getAppErrorFuncCalled++
		return nil
	}
	generateErrorResponseFuncExpected = 0
	generateErrorResponseFuncCalled = 0
	generateErrorResponseFunc = func(appError apperrorModel.AppError) errorResponseModel {
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
	createErrorResponseFunc = func(err error) (string, int) {
		createErrorResponseFuncCalled++
		return "", 0
	}
	constructResponseFuncExpected = 0
	constructResponseFuncCalled = 0
	constructResponseFunc = func(responseObject interface{}, responseError error) (string, int) {
		constructResponseFuncCalled++
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
	httpStatusText = http.StatusText
	assert.Equal(t, httpStatusTextExpected, httpStatusTextCalled, "Unexpected number of calls to httpStatusText")
	writeResponseFunc = writeResponse
	assert.Equal(t, writeResponseFuncExpected, writeResponseFuncCalled, "Unexpected number of calls to writeResponseFunc")
	getAppErrorFunc = getAppError
	assert.Equal(t, getAppErrorFuncExpected, getAppErrorFuncCalled, "Unexpected number of calls to getAppErrorFunc")
	generateErrorResponseFunc = generateErrorResponse
	assert.Equal(t, generateErrorResponseFuncExpected, generateErrorResponseFuncCalled, "Unexpected number of calls to generateErrorResponseFunc")
	createOkResponseFunc = createOkResponse
	assert.Equal(t, createOkResponseFuncExpected, createOkResponseFuncCalled, "Unexpected number of calls to createOkResponseFunc")
	createErrorResponseFunc = createErrorResponse
	assert.Equal(t, createErrorResponseFuncExpected, createErrorResponseFuncCalled, "Unexpected number of calls to createErrorResponseFunc")
	constructResponseFunc = constructResponse
	assert.Equal(t, constructResponseFuncExpected, constructResponseFuncCalled, "Unexpected number of calls to constructResponseFunc")
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
	t                  *testing.T
	expectedCode       *string
	expectedHTTPStatus *int
	expectedMessages   *[]string
	expectedExtraData  *map[string]string
}

func (dae *dummyAppError) Code() string {
	if dae.expectedCode == nil {
		assert.Fail(dae.t, "Unexpected number of calls to Code")
		return "Unknown"
	}
	return *dae.expectedCode
}

func (dae *dummyAppError) HTTPStatusCode() int {
	if dae.expectedHTTPStatus == nil {
		assert.Fail(dae.t, "Unexpected number of calls to HTTPStatusCode")
		return -1
	}
	return *dae.expectedHTTPStatus
}

func (dae *dummyAppError) Error() string {
	assert.Fail(dae.t, "Unexpected number of calls to Error")
	return ""
}

func (dae *dummyAppError) InnerErrors() []error {
	assert.Fail(dae.t, "Unexpected number of calls to InnerErrors")
	return nil
}

func (dae *dummyAppError) Messages() []string {
	if dae.expectedMessages == nil {
		assert.Fail(dae.t, "Unexpected number of calls to Messages")
		return nil
	}
	return *dae.expectedMessages
}

func (dae *dummyAppError) ExtraData() map[string]string {
	if dae.expectedExtraData == nil {
		assert.Fail(dae.t, "Unexpected number of calls to ExtraData")
		return nil
	}
	return *dae.expectedExtraData
}

func (dae *dummyAppError) Append(innerErrors ...error) {
	assert.Fail(dae.t, "Unexpected number of calls to Append")
}

func (dae *dummyAppError) Attach(name string, value interface{}) {
	assert.Fail(dae.t, "Unexpected number of calls to Attach")
}

type dummySession struct {
	t              *testing.T
	httpRequest    *http.Request
	responseWriter *dummyResponseWriter
}

func (session *dummySession) GetID() uuid.UUID {
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySession) GetName() string {
	assert.Fail(session.t, "Unexpected call to GetName")
	return ""
}

func (session *dummySession) GetRequest() *http.Request {
	if session.httpRequest == nil {
		assert.Fail(session.t, "Unexpected call to GetRequest")
		return nil
	}
	return session.httpRequest
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	if session.responseWriter == nil {
		assert.Fail(session.t, "Unexpected call to GetResponseWriter")
		return nil
	}
	return session.responseWriter
}

func (session *dummySession) GetRequestBody(dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestBody")
	return nil
}

func (session *dummySession) GetRequestParameter(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestParameter")
	return nil
}

func (session *dummySession) GetRequestQuery(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQuery")
	return nil
}

func (session *dummySession) GetRequestQueries(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQueries")
	return nil
}

func (session *dummySession) GetRequestHeader(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeader")
	return nil
}

func (session *dummySession) GetRequestHeaders(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeaders")
	return nil
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	assert.Fail(session.t, "Unexpected call to Attach")
	return false
}

func (session *dummySession) Detach(name string) bool {
	assert.Fail(session.t, "Unexpected call to Detach")
	return false
}

func (session *dummySession) GetAttachment(name string, dataTemplate interface{}) bool {
	assert.Fail(session.t, "Unexpected call to GetAttachment")
	return false
}

func (session *dummySession) IsLoggingAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool {
	assert.Fail(session.t, "Unexpected call to IsLoggingAllowed")
	return false
}

// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
func (session *dummySession) LogMethodEnter() {
	assert.Fail(session.t, "Unexpected call to LogMethodEnter")
}

// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
func (session *dummySession) LogMethodParameter(parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodParameter")
}

// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
func (session *dummySession) LogMethodLogic(logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodLogic")
}

// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
func (session *dummySession) LogMethodReturn(returns ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodReturn")
}

// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
func (session *dummySession) LogMethodExit() {
	assert.Fail(session.t, "Unexpected call to LogMethodExit")
}

// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
func (session *dummySession) CreateNetworkRequest(method string, url string, payload string, header map[string]string) networkModel.NetworkRequest {
	assert.Fail(session.t, "Unexpected call to CreateNetworkRequest")
	return nil
}
