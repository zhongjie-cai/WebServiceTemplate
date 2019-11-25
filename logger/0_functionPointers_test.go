package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
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
	configIsLocalhostExpected          int
	configIsLocalhostCalled            int
	isLoggingAllowedFuncExpected       int
	isLoggingAllowedFuncCalled         int
	sessionGetExpected                 int
	sessionGetCalled                   int
	apperrorWrapSimpleErrorExpected    int
	apperrorWrapSimpleErrorCalled      int
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
	configIsLocalhostExpected = 0
	configIsLocalhostCalled = 0
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	isLoggingAllowedFuncExpected = 0
	isLoggingAllowedFuncCalled = 0
	isLoggingAllowedFunc = func(session *session.Session, logType logtype.LogType, logLevel loglevel.LogLevel) bool {
		isLoggingAllowedFuncCalled++
		return false
	}
	sessionGetExpected = 0
	sessionGetCalled = 0
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	defaultLoggingFuncExpected = 0
	defaultLoggingFuncCalled = 0
	defaultLoggingFunc = func(session *session.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
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
	config.IsLocalhost = func() bool { return false }
	assert.Equal(t, configIsLocalhostExpected, configIsLocalhostCalled, "Unexpected number of calls to configIsLocalhost")
	isLoggingAllowedFunc = isLoggingAllowed
	assert.Equal(t, isLoggingAllowedFuncExpected, isLoggingAllowedFuncCalled, "Unexpected number of calls to isLoggingAllowedFunc")
	sessionGet = session.Get
	assert.Equal(t, sessionGetExpected, sessionGetCalled, "Unexpected number of calls to sessionGet")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	defaultLoggingFunc = defaultLogging
	assert.Equal(t, defaultLoggingFuncExpected, defaultLoggingFuncCalled, "Unexpected number of calls to defaultLoggingFunc")
	prepareLoggingFunc = prepareLogging
	assert.Equal(t, prepareLoggingFuncExpected, prepareLoggingFuncCalled, "Unexpected number of calls to prepareLoggingFunc")

	customization.LoggingFunc = nil
}
