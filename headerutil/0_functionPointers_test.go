package headerutil

import (
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	jsonutilMarshalIgnoreErrorExpected             int
	jsonutilMarshalIgnoreErrorCalled               int
	stringsJoinExpected                            int
	stringsJoinCalled                              int
	loggerAPIRequestExpected                       int
	loggerAPIRequestCalled                         int
	getHeaderLogStyleFuncExpected                  int
	getHeaderLogStyleFuncCalled                    int
	logCombinedHTTPHeaderFuncExpected              int
	logCombinedHTTPHeaderFuncCalled                int
	logPerNameHTTPHeaderFuncExpected               int
	logPerNameHTTPHeaderFuncCalled                 int
	logPerValueHTTPHeaderFuncExpected              int
	logPerValueHTTPHeaderFuncCalled                int
	logHTTPHeaderFuncExpected                      int
	logHTTPHeaderFuncCalled                        int
	customizationSessionHTTPHeaderLogStyleExpected int
	customizationSessionHTTPHeaderLogStyleCalled   int
	customizationDefaultHTTPHeaderLogStyleExpected int
	customizationDefaultHTTPHeaderLogStyleCalled   int
)

func createMock(t *testing.T) {
	jsonutilMarshalIgnoreErrorExpected = 0
	jsonutilMarshalIgnoreErrorCalled = 0
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		return ""
	}
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
	loggerAPIRequestExpected = 0
	loggerAPIRequestCalled = 0
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
	}
	getHeaderLogStyleFuncExpected = 0
	getHeaderLogStyleFuncCalled = 0
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		return 0
	}
	logCombinedHTTPHeaderFuncExpected = 0
	logCombinedHTTPHeaderFuncCalled = 0
	logCombinedHTTPHeaderFunc = func(session sessionModel.Session, header http.Header) {
		logCombinedHTTPHeaderFuncCalled++
	}
	logPerNameHTTPHeaderFuncExpected = 0
	logPerNameHTTPHeaderFuncCalled = 0
	logPerNameHTTPHeaderFunc = func(session sessionModel.Session, header http.Header) {
		logPerNameHTTPHeaderFuncCalled++
	}
	logPerValueHTTPHeaderFuncExpected = 0
	logPerValueHTTPHeaderFuncCalled = 0
	logPerValueHTTPHeaderFunc = func(session sessionModel.Session, header http.Header) {
		logPerValueHTTPHeaderFuncCalled++
	}
	logHTTPHeaderFuncExpected = 0
	logHTTPHeaderFuncCalled = 0
	logHTTPHeaderFunc = func(session sessionModel.Session, header http.Header) {
		logHTTPHeaderFuncCalled++
	}
	customizationSessionHTTPHeaderLogStyleExpected = 0
	customizationSessionHTTPHeaderLogStyleCalled = 0
	customization.SessionHTTPHeaderLogStyle = nil
	customizationDefaultHTTPHeaderLogStyleExpected = 0
	customizationDefaultHTTPHeaderLogStyleCalled = 0
	customization.DefaultHTTPHeaderLogStyle = nil
}

func verifyAll(t *testing.T) {
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected number of calls to jsonutilMarshalIgnoreError")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	loggerAPIRequest = logger.APIRequest
	assert.Equal(t, loggerAPIRequestExpected, loggerAPIRequestCalled, "Unexpected number of calls to loggerAPIRequest")
	getHeaderLogStyleFunc = getHeaderLogStyle
	assert.Equal(t, getHeaderLogStyleFuncExpected, getHeaderLogStyleFuncCalled, "Unexpected number of calls to getHeaderLogStyleFunc")
	logCombinedHTTPHeaderFunc = logCombinedHTTPHeader
	assert.Equal(t, logCombinedHTTPHeaderFuncExpected, logCombinedHTTPHeaderFuncCalled, "Unexpected number of calls to logCombinedHTTPHeaderFunc")
	logPerNameHTTPHeaderFunc = logPerNameHTTPHeader
	assert.Equal(t, logPerNameHTTPHeaderFuncExpected, logPerNameHTTPHeaderFuncCalled, "Unexpected number of calls to logPerNameHTTPHeaderFunc")
	logPerValueHTTPHeaderFunc = logPerValueHTTPHeader
	assert.Equal(t, logPerValueHTTPHeaderFuncExpected, logPerValueHTTPHeaderFuncCalled, "Unexpected number of calls to logPerValueHTTPHeaderFunc")
	logHTTPHeaderFunc = LogHTTPHeader
	assert.Equal(t, logHTTPHeaderFuncExpected, logHTTPHeaderFuncCalled, "Unexpected number of calls to logHTTPHeaderFunc")
	assert.Equal(t, customizationSessionHTTPHeaderLogStyleExpected, customizationSessionHTTPHeaderLogStyleCalled, "Unexpected number of calls to customization.SessionHTTPHeaderLogStyle")
	assert.Equal(t, customizationDefaultHTTPHeaderLogStyleExpected, customizationDefaultHTTPHeaderLogStyleCalled, "Unexpected number of calls to customization.DefaultHTTPHeaderLogStyle")
}

// mock structs
type dummySession struct {
	t *testing.T
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
	assert.Fail(session.t, "Unexpected call to GetRequest")
	return nil
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	assert.Fail(session.t, "Unexpected call to GetResponseWriter")
	return nil
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
