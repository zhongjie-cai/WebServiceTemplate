package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestInitialize_NotSet(t *testing.T) {
	// arrange
	var dummyMessageFormat = "customization.LoggingFunc is not configured; fallback to default logging function."
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// stub
	customization.LoggingFunc = nil

	// mock
	createMock(t)

	// expect
	apperrorGetCustomErrorExpected = 1
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
		assert.Equal(t, apperrorEnum.CodeGeneralFailure, errorCode)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var err = Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_Set(t *testing.T) {
	// stub
	customization.LoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
	}

	// mock
	createMock(t)

	// SUT + act
	var err = Initialize()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyName = "some Name"
	var dummySessionObject = &dummySession{
		t:    t,
		id:   &dummySessionID,
		name: &dummyName,
	}
	var dummyLogType = logtype.MethodLogic
	var dummyLogLevel = loglevel.Warn
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"
	var dummyAppName = "some app name"
	var dummyAppVersion = "some app version"
	var dummyTimestamp = time.Now().UTC()
	var dummyLogEntry = logEntry{
		Application: dummyAppName,
		Version:     dummyAppVersion,
		Timestamp:   dummyTimestamp,
		Session:     dummySessionID,
		Name:        dummyName,
		Type:        dummyLogType,
		Level:       dummyLogLevel,
		Category:    dummyCategory,
		Subcategory: dummySubCategory,
		Description: dummyDescription,
	}
	var dummyLogEntryString = "some log entry string"

	// mock
	createMock(t)

	// expect
	configAppNameExpected = 1
	config.AppName = func() string {
		configAppNameCalled++
		return dummyAppName
	}
	configAppVersionExpected = 1
	config.AppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	timeutilGetTimeNowUTCExpected = 1
	timeutilGetTimeNowUTC = func() time.Time {
		timeutilGetTimeNowUTCCalled++
		return dummyTimestamp
	}
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyLogEntry, v)
		return dummyLogEntryString
	}
	fmtPrintlnExpected = 1
	fmtPrintln = func(a ...interface{}) (n int, err error) {
		fmtPrintlnCalled++
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyLogEntryString, a[0])
		return 0, nil
	}

	// SUT + act
	defaultLogging(
		dummySessionObject,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_LoggingNotAllowed(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyIsLogAllowed = false
	var dummySessionObject = &dummySession{
		t:            t,
		isLogAllowed: &dummyIsLogAllowed,
	}
	var dummyLogType = logtype.APIEnter
	var dummyLogLevel = loglevel.Warn
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) sessionModel.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_LogAllowed_DefaultLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyIsLogAllowed = true
	var dummySessionObject = &dummySession{
		t:            t,
		isLogAllowed: &dummyIsLogAllowed,
	}
	var dummyLogType = logtype.MethodEnter
	var dummyLogLevel = loglevel.Error
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// stub
	customization.LoggingFunc = nil

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) sessionModel.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	defaultLoggingFuncExpected = 1
	defaultLoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		defaultLoggingFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyLogLevel, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_LogAllowed_CustomLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyIsLogAllowed = true
	var dummySessionObject = &dummySession{
		t:            t,
		isLogAllowed: &dummyIsLogAllowed,
	}
	var dummyLogType = logtype.MethodEnter
	var dummyLogLevel = loglevel.Error
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"
	var loggingFuncExpected int
	var loggingFuncCalled int

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) sessionModel.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummySessionObject
	}
	loggingFuncExpected = 1
	customization.LoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		loggingFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, loggingFuncExpected, loggingFuncCalled, "Unexpected number of calls to LoggingFunc")
}

func TestAppRoot(t *testing.T) {
	// arrange
	var dummyLogType = logtype.AppRoot
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	AppRoot(
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestAPIEnter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.APIEnter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	APIEnter(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestAPIRequest(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.APIRequest
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	APIRequest(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestMethodEnter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.MethodEnter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodEnter(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestMethodParameter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.MethodParameter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodParameter(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestMethodLogic(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.MethodLogic
	var dummyLogLevel = loglevel.Error
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyLogLevel, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodLogic(
		dummySessionID,
		dummyLogLevel,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestDependencyCall(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.DependencyCall
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	DependencyCall(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestDependencyRequest(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.DependencyRequest
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	DependencyRequest(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestDependencyResponse(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.DependencyResponse
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	DependencyResponse(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestDependencyFinish(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.DependencyFinish
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	DependencyFinish(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestMethodReturn(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.MethodReturn
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodReturn(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestMethodExit(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.MethodExit
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodExit(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestAPIResponse(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.APIResponse
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	APIResponse(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestAPIExit(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogType = logtype.APIExit
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return fmt.Sprintf(format, a...)
	}
	prepareLoggingFuncExpected = 1
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, loglevel.Info, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	APIExit(
		dummySessionID,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}
