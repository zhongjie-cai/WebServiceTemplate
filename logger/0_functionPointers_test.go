package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
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
	sessionGetExpected                 int
	sessionGetCalled                   int
	doLoggingFuncExpected              int
	doLoggingFuncCalled                int
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
	configAppName = func() string {
		configAppNameCalled++
		return ""
	}
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	configAppVersion = func() string {
		configAppVersionCalled++
		return ""
	}
	configIsLocalhostExpected = 0
	configIsLocalhostCalled = 0
	configIsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
	sessionGetExpected = 0
	sessionGetCalled = 0
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		return nil
	}
	doLoggingFuncExpected = 0
	doLoggingFuncCalled = 0
	doLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		doLoggingFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	fmtPrintln = fmt.Println
	assert.Equal(t, fmtPrintlnExpected, fmtPrintlnCalled, "Unexpected method call to fmtPrintln")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected method call to uuidNew")
	uuidParse = uuid.Parse
	assert.Equal(t, uuidParseExpected, uuidParseCalled, "Unexpected method call to uuidParse")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected method call to fmtSprintf")
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	assert.Equal(t, timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled, "Unexpected method call to timeutilGetTimeNowUTC")
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected method call to jsonutilMarshalIgnoreError")
	configAppName = config.AppName
	assert.Equal(t, configAppNameExpected, configAppNameCalled, "Unexpected method call to configAppName")
	configAppVersion = config.AppVersion
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected method call to configAppVersion")
	configIsLocalhost = config.IsLocalhost
	assert.Equal(t, configIsLocalhostExpected, configIsLocalhostCalled, "Unexpected method call to configIsLocalhost")
	sessionGet = session.Get
	assert.Equal(t, sessionGetExpected, sessionGetCalled, "Unexpected method call to sessionGet")
	doLoggingFunc = doLogging
	assert.Equal(t, doLoggingFuncExpected, doLoggingFuncCalled, "Unexpected method call to doLoggingFunc")
}
