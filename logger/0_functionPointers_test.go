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
	if fmtPrintlnExpected != fmtPrintlnCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to fmtPrintln, expected %v, actual %v", fmtPrintlnExpected, fmtPrintlnCalled))
	}
	uuidNew = uuid.New
	if uuidNewExpected != uuidNewCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to uuidNew, expected %v, actual %v", uuidNewExpected, uuidNewCalled))
	}
	uuidParse = uuid.Parse
	if uuidParseExpected != uuidParseCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to uuidParse, expected %v, actual %v", uuidParseExpected, uuidParseCalled))
	}
	fmtSprintf = fmt.Sprintf
	if fmtSprintfExpected != fmtSprintfCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to fmtSprintf, expected %v, actual %v", fmtSprintfExpected, fmtSprintfCalled))
	}
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	if timeutilGetTimeNowUTCExpected != timeutilGetTimeNowUTCCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to timeutilGetTimeNowUTC, expected %v, actual %v", timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled))
	}
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	if jsonutilMarshalIgnoreErrorExpected != jsonutilMarshalIgnoreErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to jsonutilMarshalIgnoreError, expected %v, actual %v", jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled))
	}
	configAppName = config.AppName
	if configAppNameExpected != configAppNameCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppName, expected %v, actual %v", configAppNameExpected, configAppNameCalled))
	}
	configAppVersion = config.AppVersion
	if configAppVersionExpected != configAppVersionCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppVersion, expected %v, actual %v", configAppVersionExpected, configAppVersionCalled))
	}
	sessionGet = session.Get
	if sessionGetExpected != sessionGetCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to sessionGet, expected %v, actual %v", sessionGetExpected, sessionGetCalled))
	}
	doLoggingFunc = doLogging
	if doLoggingFuncExpected != doLoggingFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to doLoggingFunc, expected %v, actual %v", doLoggingFuncExpected, doLoggingFuncCalled))
	}
}
