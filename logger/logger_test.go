package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/session"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInitialize_NotSet(t *testing.T) {
	// arrange
	var dummyMessageFormat = "logger.LoggingFunc is not configured; fallback to default logging function."
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NoError(t, innerError)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	err := Initialize()

	// assert
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestInitialize_Set(t *testing.T) {
	// stub
	LoggingFunc = defaultLogging

	// mock
	createMock(t)

	// SUT + act
	err := Initialize()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLoginID = uuid.New()
	var dummyEndpoint = "some endpoint"
	var dummyLogSession = &session.Session{
		ID:             dummySessionID,
		AllowedLogType: dummyAllowedLogType,
		LoginID:        dummyLoginID,
		Endpoint:       dummyEndpoint,
	}
	var dummyLogType = logtype.MethodLogic
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
		Login:       dummyLoginID,
		Endpoint:    dummyEndpoint,
		Level:       dummyLogType,
		Category:    dummyCategory,
		Subcategory: dummySubCategory,
		Description: dummyDescription,
	}
	var dummyLogEntryString = "some log entry string"

	// mock
	createMock(t)

	// expect
	configAppNameExpected = 1
	configAppName = func() string {
		configAppNameCalled++
		return dummyAppName
	}
	configAppVersionExpected = 1
	configAppVersion = func() string {
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
		dummyLogSession,
		dummyLogType,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_NotLocalhost_FlagNotMatch(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLogSession = &session.Session{
		AllowedLogType: dummyAllowedLogType,
	}
	var dummyLogType = logtype.APIEnter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyLogSession
	}
	configIsLocalhostExpected = 1
	configIsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_IsLocalhost_DefaultLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLoginID = uuid.New()
	var dummyEndpoint = "some endpoint"
	var dummyLogSession = &session.Session{
		ID:             dummySessionID,
		AllowedLogType: dummyAllowedLogType,
		LoginID:        dummyLoginID,
		Endpoint:       dummyEndpoint,
	}
	var dummyLogType = logtype.MethodEnter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// stub
	LoggingFunc = nil

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyLogSession
	}
	configIsLocalhostExpected = 1
	configIsLocalhost = func() bool {
		configIsLocalhostCalled++
		return true
	}
	defaultLoggingFuncExpected = 1
	defaultLoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		defaultLoggingFuncCalled++
		assert.Equal(t, dummyLogSession, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_FlagMatch_DefaultLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLoginID = uuid.New()
	var dummyEndpoint = "some endpoint"
	var dummyLogSession = &session.Session{
		ID:             dummySessionID,
		AllowedLogType: dummyAllowedLogType,
		LoginID:        dummyLoginID,
		Endpoint:       dummyEndpoint,
	}
	var dummyLogType = logtype.MethodLogic
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"

	// stub
	LoggingFunc = nil

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyLogSession
	}
	defaultLoggingFuncExpected = 1
	defaultLoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		defaultLoggingFuncCalled++
		assert.Equal(t, dummyLogSession, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestPrepareLogging_IsLocalhost_CustomLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLoginID = uuid.New()
	var dummyEndpoint = "some endpoint"
	var dummyLogSession = &session.Session{
		ID:             dummySessionID,
		AllowedLogType: dummyAllowedLogType,
		LoginID:        dummyLoginID,
		Endpoint:       dummyEndpoint,
	}
	var dummyLogType = logtype.MethodEnter
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"
	var loggingFuncExpected int
	var loggingFuncCalled int

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyLogSession
	}
	configIsLocalhostExpected = 1
	configIsLocalhost = func() bool {
		configIsLocalhostCalled++
		return true
	}
	loggingFuncExpected = 1
	LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		loggingFuncCalled++
		assert.Equal(t, dummyLogSession, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
		dummyCategory,
		dummySubCategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, loggingFuncExpected, loggingFuncCalled, "Unexpected number of calls to LoggingFunc")
}

func TestPrepareLogging_FlagMatch_CustomLogging(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyAllowedLogType = logtype.BasicLogging
	var dummyLoginID = uuid.New()
	var dummyEndpoint = "some endpoint"
	var dummyLogSession = &session.Session{
		ID:             dummySessionID,
		AllowedLogType: dummyAllowedLogType,
		LoginID:        dummyLoginID,
		Endpoint:       dummyEndpoint,
	}
	var dummyLogType = logtype.MethodLogic
	var dummyCategory = "some category"
	var dummySubCategory = "some sub category"
	var dummyDescription = "some description"
	var loggingFuncExpected int
	var loggingFuncCalled int

	// mock
	createMock(t)

	// expect
	sessionGetExpected = 1
	sessionGet = func(sessionID uuid.UUID) *session.Session {
		sessionGetCalled++
		assert.Equal(t, dummySessionID, sessionID)
		return dummyLogSession
	}
	loggingFuncExpected = 1
	LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		loggingFuncCalled++
		assert.Equal(t, dummyLogSession, session)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	prepareLogging(
		dummySessionID,
		dummyLogType,
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, uuid.Nil, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubCategory, subcategory)
		assert.Equal(t, dummyDescription, description)
	}

	// SUT + act
	MethodLogic(
		dummySessionID,
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
	prepareLoggingFunc = func(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
		prepareLoggingFuncCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyLogType, logType)
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
