package logger

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

var (
	fmtPrintlnExpected                 int
	fmtPrintlnCalled                   int
	uuidNewExpected                    int
	uuidNewCalled                      int
	uuidParseExpected                  int
	uuidParseCalled                    int
	fmtSprintfExpected                 int
	fmtSprintfCalled                   int
	timeutilGetTimeNowUTCExpected      int
	timeutilGetTimeNowUTCCalled        int
	jsonutilMarshalIgnoreErrorExpected int
	jsonutilMarshalIgnoreErrorCalled   int
	configAppNameExpected              int
	configAppNameCalled                int
	configAppVersionExpected           int
	configAppVersionCalled             int
	sessionGetExpected                 int
	sessionGetCalled                   int
	apperrorGetCustomErrorExpected     int
	apperrorGetCustomErrorCalled       int
	defaultLoggingFuncExpected         int
	defaultLoggingFuncCalled           int
	prepareLoggingFuncExpected         int
	prepareLoggingFuncCalled           int
)

func createMock(t *testing.T) {
	fmtPrintlnExpected = 0
	fmtPrintlnCalled = 0
	fmtPrintln = func(a ...interface{}) (n int, err error) {
		fmtPrintlnCalled++
		return 0, nil
	}
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	uuidParseExpected = 0
	uuidParseCalled = 0
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		return uuid.Nil, nil
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	timeutilGetTimeNowUTCExpected = 0
	timeutilGetTimeNowUTCCalled = 0
	timeutilGetTimeNowUTC = func() time.Time {
		timeutilGetTimeNowUTCCalled++
		return time.Time{}
	}
	jsonutilMarshalIgnoreErrorExpected = 0
	jsonutilMarshalIgnoreErrorCalled = 0
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		return ""
	}
	configAppNameExpected = 0
	configAppNameCalled = 0
	config.AppName = func() string {
		configAppNameCalled++
		return ""
	}
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	config.AppVersion = func() string {
		configAppVersionCalled++
		return ""
	}
	sessionGetExpected = 0
	sessionGetCalled = 0
	sessionGet = func(sessionID uuid.UUID) sessionModel.Session {
		sessionGetCalled++
		return nil
	}
	apperrorGetCustomErrorExpected = 0
	apperrorGetCustomErrorCalled = 0
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		return nil
	}
	defaultLoggingFuncExpected = 0
	defaultLoggingFuncCalled = 0
	defaultLoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		defaultLoggingFuncCalled++
	}
	prepareLoggingFuncExpected = 0
	prepareLoggingFuncCalled = 0
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	fmtPrintln = fmt.Println
	assert.Equal(t, fmtPrintlnExpected, fmtPrintlnCalled, "Unexpected number of calls to fmtPrintln")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	uuidParse = uuid.Parse
	assert.Equal(t, uuidParseExpected, uuidParseCalled, "Unexpected number of calls to uuidParse")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	assert.Equal(t, timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled, "Unexpected number of calls to timeutilGetTimeNowUTC")
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected number of calls to jsonutilMarshalIgnoreError")
	config.AppName = func() string { return "" }
	assert.Equal(t, configAppNameExpected, configAppNameCalled, "Unexpected number of calls to configAppName")
	config.AppVersion = func() string { return "" }
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected number of calls to configAppVersion")
	sessionGet = session.Get
	assert.Equal(t, sessionGetExpected, sessionGetCalled, "Unexpected number of calls to sessionGet")
	apperrorGetCustomError = apperror.GetCustomError
	assert.Equal(t, apperrorGetCustomErrorExpected, apperrorGetCustomErrorCalled, "Unexpected number of calls to apperrorGetCustomError")
	defaultLoggingFunc = defaultLogging
	assert.Equal(t, defaultLoggingFuncExpected, defaultLoggingFuncCalled, "Unexpected number of calls to defaultLoggingFunc")
	prepareLoggingFunc = prepareLogging
	assert.Equal(t, prepareLoggingFuncExpected, prepareLoggingFuncCalled, "Unexpected number of calls to prepareLoggingFunc")
	customization.LoggingFunc = nil
}

// mock structs
type dummySession struct {
	t            *testing.T
	id           *uuid.UUID
	name         *string
	isLogAllowed *bool
}

func (session *dummySession) GetID() uuid.UUID {
	if session.id == nil {
		assert.Fail(session.t, "Unexpected call to GetID")
		return uuid.Nil
	}
	return *session.id
}

func (session *dummySession) GetName() string {
	if session.name == nil {
		assert.Fail(session.t, "Unexpected call to GetName")
		return ""
	}
	return *session.name
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

func (session *dummySession) IsLogAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool {
	if session.isLogAllowed == nil {
		assert.Fail(session.t, "Unexpected call to IsLogAllowed")
	}
	return *session.isLogAllowed
}
