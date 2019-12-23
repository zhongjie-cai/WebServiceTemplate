package session

import (
	"encoding/json"
	"net/http"
	"net/textproto"
	"runtime"
	"strconv"
	"testing"

	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/network"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	uuidNewExpected                             int
	uuidNewCalled                               int
	jsonMarshalExpected                         int
	jsonMarshalCalled                           int
	jsonUnmarshalExpected                       int
	jsonUnmarshalCalled                         int
	fmtErrorfExpected                           int
	fmtErrorfCalled                             int
	muxVarsExpected                             int
	muxVarsCalled                               int
	loggerAPIRequestExpected                    int
	loggerAPIRequestCalled                      int
	requestGetRequestBodyExpected               int
	requestGetRequestBodyCalled                 int
	apperrorGetBadRequestErrorExpected          int
	apperrorGetBadRequestErrorCalled            int
	textprotoCanonicalMIMEHeaderKeyExpected     int
	textprotoCanonicalMIMEHeaderKeyCalled       int
	jsonutilTryUnmarshalExpected                int
	jsonutilTryUnmarshalCalled                  int
	headerutilLogHTTPHeaderForNameExpected      int
	headerutilLogHTTPHeaderForNameCalled        int
	getAllQueriesFuncExpected                   int
	getAllQueriesFuncCalled                     int
	getAllHeadersFuncExpected                   int
	getAllHeadersFuncCalled                     int
	isLoggingTypeMatchFuncExpected              int
	isLoggingTypeMatchFuncCalled                int
	isLoggingLevelMatchFuncExpected             int
	isLoggingLevelMatchFuncCalled               int
	configIsLocalhostExpected                   int
	configIsLocalhostCalled                     int
	configDefaultAllowedLogTypeExpected         int
	configDefaultAllowedLogTypeCalled           int
	configDefaultAllowedLogLevelExpected        int
	configDefaultAllowedLogLevelCalled          int
	customizationSessionAllowedLogTypeExpected  int
	customizationSessionAllowedLogTypeCalled    int
	customizationSessionAllowedLogLevelExpected int
	customizationSessionAllowedLogLevelCalled   int
	runtimeCallerExpected                       int
	runtimeCallerCalled                         int
	runtimeFuncForPCExpected                    int
	runtimeFuncForPCCalled                      int
	getMethodNameFuncExpected                   int
	getMethodNameFuncCalled                     int
	strconvItoaExpected                         int
	strconvItoaCalled                           int
	loggerMethodEnterExpected                   int
	loggerMethodEnterCalled                     int
	loggerMethodParameterExpected               int
	loggerMethodParameterCalled                 int
	loggerMethodLogicExpected                   int
	loggerMethodLogicCalled                     int
	loggerMethodReturnExpected                  int
	loggerMethodReturnCalled                    int
	loggerMethodExitExpected                    int
	loggerMethodExitCalled                      int
	networkNewNetworkRequestExpected            int
	networkNewNetworkRequestCalled              int
	getAllowedLogTypeFuncExpected               int
	getAllowedLogTypeFuncCalled                 int
	getAllowedLogLevelFuncExpected              int
	getAllowedLogLevelFuncCalled                int
	certificateHasClientCertExpected            int
	certificateHasClientCertCalled              int
	customizationSendClientCertExpected         int
	customizationSendClientCertCalled           int
	shouldSendClientCertFuncExpected            int
	shouldSendClientCertFuncCalled              int
)

func createMock(t *testing.T) {
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	jsonMarshalExpected = 0
	jsonMarshalCalled = 0
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		return nil, nil
	}
	jsonUnmarshalExpected = 0
	jsonUnmarshalCalled = 0
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return nil
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	muxVarsExpected = 0
	muxVarsCalled = 0
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		return nil
	}
	loggerAPIRequestExpected = 0
	loggerAPIRequestCalled = 0
	loggerAPIRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
	}
	requestGetRequestBodyExpected = 0
	requestGetRequestBodyCalled = 0
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		return ""
	}
	apperrorGetBadRequestErrorExpected = 0
	apperrorGetBadRequestErrorCalled = 0
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		return nil
	}
	textprotoCanonicalMIMEHeaderKeyExpected = 0
	textprotoCanonicalMIMEHeaderKeyCalled = 0
	textprotoCanonicalMIMEHeaderKey = func(s string) string {
		textprotoCanonicalMIMEHeaderKeyCalled++
		return ""
	}
	jsonutilTryUnmarshalExpected = 0
	jsonutilTryUnmarshalCalled = 0
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		return nil
	}
	headerutilLogHTTPHeaderForNameExpected = 0
	headerutilLogHTTPHeaderForNameCalled = 0
	headerutilLogHTTPHeaderForName = func(session sessionModel.Session, name string, values []string) {
		headerutilLogHTTPHeaderForNameCalled++
	}
	getAllQueriesFuncExpected = 0
	getAllQueriesFuncCalled = 0
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		return nil
	}
	getAllHeadersFuncExpected = 0
	getAllHeadersFuncCalled = 0
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		return nil
	}
	isLoggingTypeMatchFuncExpected = 0
	isLoggingTypeMatchFuncCalled = 0
	isLoggingTypeMatchFunc = func(session *session, logType logtype.LogType) bool {
		isLoggingTypeMatchFuncCalled++
		return false
	}
	isLoggingLevelMatchFuncExpected = 0
	isLoggingLevelMatchFuncCalled = 0
	isLoggingLevelMatchFunc = func(session *session, logLevel loglevel.LogLevel) bool {
		isLoggingLevelMatchFuncCalled++
		return false
	}
	configIsLocalhostExpected = 0
	configIsLocalhostCalled = 0
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	configDefaultAllowedLogTypeExpected = 0
	configDefaultAllowedLogTypeCalled = 0
	config.DefaultAllowedLogType = nil
	configDefaultAllowedLogLevelExpected = 0
	configDefaultAllowedLogLevelCalled = 0
	config.DefaultAllowedLogLevel = nil
	customizationSessionAllowedLogTypeExpected = 0
	customizationSessionAllowedLogTypeCalled = 0
	customization.SessionAllowedLogType = nil
	customizationSessionAllowedLogLevelExpected = 0
	customizationSessionAllowedLogLevelCalled = 0
	customization.SessionAllowedLogLevel = nil
	runtimeCallerExpected = 0
	runtimeCallerCalled = 0
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		return 0, "", 0, false
	}
	runtimeFuncForPCExpected = 0
	runtimeFuncForPCCalled = 0
	runtimeFuncForPC = func(pc uintptr) *runtime.Func {
		runtimeFuncForPCCalled++
		return nil
	}
	getMethodNameFuncExpected = 0
	getMethodNameFuncCalled = 0
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return ""
	}
	strconvItoaExpected = 0
	strconvItoaCalled = 0
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return ""
	}
	loggerMethodEnterExpected = 0
	loggerMethodEnterCalled = 0
	loggerMethodEnter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodEnterCalled++
	}
	loggerMethodParameterExpected = 0
	loggerMethodParameterCalled = 0
	loggerMethodParameter = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodParameterCalled++
	}
	loggerMethodLogicExpected = 0
	loggerMethodLogicCalled = 0
	loggerMethodLogic = func(session sessionModel.Session, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodLogicCalled++
	}
	loggerMethodReturnExpected = 0
	loggerMethodReturnCalled = 0
	loggerMethodReturn = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodReturnCalled++
	}
	loggerMethodExitExpected = 0
	loggerMethodExitCalled = 0
	loggerMethodExit = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerMethodExitCalled++
	}
	networkNewNetworkRequestExpected = 0
	networkNewNetworkRequestCalled = 0
	networkNewNetworkRequest = func(session sessionModel.Session, method string, url string, payload string, header map[string]string, sendClientCert bool) networkModel.NetworkRequest {
		networkNewNetworkRequestCalled++
		return nil
	}
	getAllowedLogTypeFuncExpected = 0
	getAllowedLogTypeFuncCalled = 0
	getAllowedLogTypeFunc = func(session *session) logtype.LogType {
		getAllowedLogTypeFuncCalled++
		return 0
	}
	getAllowedLogLevelFuncExpected = 0
	getAllowedLogLevelFuncCalled = 0
	getAllowedLogLevelFunc = func(session *session) loglevel.LogLevel {
		getAllowedLogLevelFuncCalled++
		return 0
	}
	certificateHasClientCertExpected = 0
	certificateHasClientCertCalled = 0
	certificateHasClientCert = func() bool {
		certificateHasClientCertCalled++
		return false
	}
	customizationSendClientCertExpected = 0
	customizationSendClientCertCalled = 0
	customization.SendClientCert = nil
	shouldSendClientCertFuncExpected = 0
	shouldSendClientCertFuncCalled = 0
	shouldSendClientCertFunc = func(url string) bool {
		shouldSendClientCertFuncCalled++
		return false
	}
}

func verifyAll(t *testing.T) {
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	jsonMarshal = json.Marshal
	assert.Equal(t, jsonMarshalExpected, jsonMarshalCalled, "Unexpected number of calls to jsonMarshal")
	jsonUnmarshal = json.Unmarshal
	assert.Equal(t, jsonUnmarshalExpected, jsonUnmarshalCalled, "Unexpected number of calls to jsonUnmarshal")
	muxVars = mux.Vars
	assert.Equal(t, muxVarsExpected, muxVarsCalled, "Unexpected number of calls to muxVars")
	loggerAPIRequest = logger.APIRequest
	assert.Equal(t, loggerAPIRequestExpected, loggerAPIRequestCalled, "Unexpected number of calls to loggerAPIRequest")
	requestGetRequestBody = request.GetRequestBody
	assert.Equal(t, requestGetRequestBodyExpected, requestGetRequestBodyCalled, "Unexpected number of calls to requestGetRequestBody")
	apperrorGetBadRequestError = apperror.GetBadRequestError
	assert.Equal(t, apperrorGetBadRequestErrorExpected, apperrorGetBadRequestErrorCalled, "Unexpected number of calls to apperrorGetBadRequestError")
	textprotoCanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey
	assert.Equal(t, textprotoCanonicalMIMEHeaderKeyExpected, textprotoCanonicalMIMEHeaderKeyCalled, "Unexpected number of calls to textprotoCanonicalMIMEHeaderKey")
	jsonutilTryUnmarshal = jsonutil.TryUnmarshal
	assert.Equal(t, jsonutilTryUnmarshalExpected, jsonutilTryUnmarshalCalled, "Unexpected number of calls to jsonutilTryUnmarshal")
	headerutilLogHTTPHeaderForName = headerutil.LogHTTPHeaderForName
	assert.Equal(t, headerutilLogHTTPHeaderForNameExpected, headerutilLogHTTPHeaderForNameCalled, "Unexpected number of calls to headerutilLogHTTPHeaderForName")
	getAllQueriesFunc = getAllQueries
	assert.Equal(t, getAllQueriesFuncExpected, getAllQueriesFuncCalled, "Unexpected number of calls to getAllQueriesFunc")
	getAllHeadersFunc = getAllHeaders
	assert.Equal(t, getAllHeadersFuncExpected, getAllHeadersFuncCalled, "Unexpected number of calls to getAllHeadersFunc")
	isLoggingTypeMatchFunc = isLoggingTypeMatch
	assert.Equal(t, isLoggingTypeMatchFuncExpected, isLoggingTypeMatchFuncCalled, "Unexpected number of calls to isLoggingTypeMatchFunc")
	isLoggingLevelMatchFunc = isLoggingLevelMatch
	assert.Equal(t, isLoggingLevelMatchFuncExpected, isLoggingLevelMatchFuncCalled, "Unexpected number of calls to isLoggingLevelMatchFunc")
	config.IsLocalhost = nil
	assert.Equal(t, configIsLocalhostExpected, configIsLocalhostCalled, "Unexpected number of calls to configIsLocalhost")
	config.DefaultAllowedLogType = nil
	assert.Equal(t, configDefaultAllowedLogTypeExpected, configDefaultAllowedLogTypeCalled, "Unexpected number of calls to configDefaultAllowedLogType")
	config.DefaultAllowedLogLevel = nil
	assert.Equal(t, configDefaultAllowedLogLevelExpected, configDefaultAllowedLogLevelCalled, "Unexpected number of calls to configDefaultAllowedLogLevel")
	customization.SessionAllowedLogType = nil
	assert.Equal(t, customizationSessionAllowedLogTypeExpected, customizationSessionAllowedLogTypeCalled, "Unexpected number of calls to customizationSessionAllowedLogType")
	customization.SessionAllowedLogLevel = nil
	assert.Equal(t, customizationSessionAllowedLogLevelExpected, customizationSessionAllowedLogLevelCalled, "Unexpected number of calls to customizationSessionAllowedLogLevel")
	runtimeCaller = runtime.Caller
	assert.Equal(t, runtimeCallerExpected, runtimeCallerCalled, "Unexpected number of calls to runtimeCaller")
	runtimeFuncForPC = runtime.FuncForPC
	assert.Equal(t, runtimeFuncForPCExpected, runtimeFuncForPCCalled, "Unexpected number of calls to runtimeFuncForPC")
	getMethodNameFunc = getMethodName
	assert.Equal(t, getMethodNameFuncExpected, getMethodNameFuncCalled, "Unexpected number of calls to getMethodNameFunc")
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected number of calls to strconvItoa")
	loggerMethodEnter = logger.MethodEnter
	assert.Equal(t, loggerMethodEnterExpected, loggerMethodEnterCalled, "Unexpected number of calls to loggerMethodEnter")
	loggerMethodParameter = logger.MethodParameter
	assert.Equal(t, loggerMethodParameterExpected, loggerMethodParameterCalled, "Unexpected number of calls to loggerMethodParameter")
	loggerMethodLogic = logger.MethodLogic
	assert.Equal(t, loggerMethodLogicExpected, loggerMethodLogicCalled, "Unexpected number of calls to loggerMethodLogic")
	loggerMethodReturn = logger.MethodReturn
	assert.Equal(t, loggerMethodReturnExpected, loggerMethodReturnCalled, "Unexpected number of calls to loggerMethodReturn")
	loggerMethodExit = logger.MethodExit
	assert.Equal(t, loggerMethodExitExpected, loggerMethodExitCalled, "Unexpected number of calls to loggerMethodExit")
	networkNewNetworkRequest = network.NewNetworkRequest
	assert.Equal(t, networkNewNetworkRequestExpected, networkNewNetworkRequestCalled, "Unexpected number of calls to networkNewNetworkRequest")
	getAllowedLogTypeFunc = getAllowedLogType
	assert.Equal(t, getAllowedLogTypeFuncExpected, getAllowedLogTypeFuncCalled, "Unexpected number of calls to getAllowedLogTypeFunc")
	getAllowedLogLevelFunc = getAllowedLogLevel
	assert.Equal(t, getAllowedLogLevelFuncExpected, getAllowedLogLevelFuncCalled, "Unexpected number of calls to getAllowedLogLevelFunc")
	certificateHasClientCert = certificate.HasClientCert
	assert.Equal(t, certificateHasClientCertExpected, certificateHasClientCertCalled, "Unexpected number of calls to certificateHasClientCert")
	customization.SendClientCert = nil
	assert.Equal(t, customizationSendClientCertExpected, customizationSendClientCertCalled, "Unexpected number of calls to customization.SendClientCert")
	shouldSendClientCertFunc = shouldSendClientCert
	assert.Equal(t, shouldSendClientCertFuncExpected, shouldSendClientCertFuncCalled, "Unexpected number of calls to shouldSendClientCertFunc")

	defaultSession = nil
}

// mock structs
type dummyResponseWriter struct {
	t *testing.T
}

func (drw dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to Header")
	return nil
}

func (drw dummyResponseWriter) Write(bytes []byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to Write")
	return 0, nil
}

func (drw dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to WriteHeader")
}

type dummyAttachment struct {
	ID   uuid.UUID
	Foo  string
	Test int
}

type dummyNetworkRequest struct {
	t *testing.T
}

func (dnr *dummyNetworkRequest) EnableRetry(connectivityRetryCount int, httpStatusRetryCount map[int]int) {
	assert.Fail(dnr.t, "Unexpected number of calls to EnableRetry")
}

func (dnr *dummyNetworkRequest) Process(dataTemplate interface{}) (statusCode int, responseHeader http.Header, responseError error) {
	assert.Fail(dnr.t, "Unexpected number of calls to Process")
	return 0, nil, nil
}

func (dnr *dummyNetworkRequest) ProcessRaw() (responseObject *http.Response, responseError error) {
	assert.Fail(dnr.t, "Unexpected number of calls to Process")
	return nil, nil
}
