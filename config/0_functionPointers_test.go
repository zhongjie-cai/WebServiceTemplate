package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

var (
	timeutilGetTimeNowUTCExpected            int
	timeutilGetTimeNowUTCCalled              int
	timeutilFormatDateTimeExpected           int
	timeutilFormatDateTimeCalled             int
	apperrorWrapSimpleErrorExpected          int
	apperrorWrapSimpleErrorCalled            int
	apperrorConsolidateAllErrorsExpected     int
	apperrorConsolidateAllErrorsCalled       int
	isServerCertificateAvailableFuncExpected int
	isServerCertificateAvailableFuncCalled   int
	isCaCertificateAvailableFuncExpected     int
	isCaCertificateAvailableFuncCalled       int
	validateStringFunctionFuncExpected       int
	validateStringFunctionFuncCalled         int
	validateBooleanFunctionFuncExpected      int
	validateBooleanFunctionFuncCalled        int
)

func createMock(t *testing.T) {
	timeutilGetTimeNowUTCExpected = 0
	timeutilGetTimeNowUTCCalled = 0
	timeutilGetTimeNowUTC = func() time.Time {
		timeutilGetTimeNowUTCCalled++
		return time.Time{}
	}
	timeutilFormatDateTimeExpected = 0
	timeutilFormatDateTimeCalled = 0
	timeutilFormatDateTime = func(value time.Time) string {
		timeutilFormatDateTimeCalled++
		return ""
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	apperrorConsolidateAllErrorsExpected = 0
	apperrorConsolidateAllErrorsCalled = 0
	apperrorConsolidateAllErrors = func(baseErrorMessage string, allErrors ...error) apperror.AppError {
		apperrorConsolidateAllErrorsCalled++
		return nil
	}
	isServerCertificateAvailableFuncExpected = 0
	isServerCertificateAvailableFuncCalled = 0
	isServerCertificateAvailableFunc = func() bool {
		isServerCertificateAvailableFuncCalled++
		return false
	}
	isCaCertificateAvailableFuncExpected = 0
	isCaCertificateAvailableFuncCalled = 0
	isCaCertificateAvailableFunc = func() bool {
		isCaCertificateAvailableFuncCalled++
		return false
	}
	validateStringFunctionFuncExpected = 0
	validateStringFunctionFuncCalled = 0
	validateStringFunctionFunc = func(stringFunc func() string, name string, defaultFunc func() string, forceToDefault bool) (func() string, error) {
		validateStringFunctionFuncCalled++
		return nil, nil
	}
	validateBooleanFunctionFuncExpected = 0
	validateBooleanFunctionFuncCalled = 0
	validateBooleanFunctionFunc = func(booleanFunc func() bool, name string, defaultFunc func() bool, forceToDefault bool) (func() bool, error) {
		validateBooleanFunctionFuncCalled++
		return nil, nil
	}
}

func verifyAll(t *testing.T) {
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	assert.Equal(t, timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled, "Unexpected number of calls to timeutilGetTimeNowUTC")
	timeutilFormatDateTime = timeutil.FormatDateTime
	assert.Equal(t, timeutilFormatDateTimeExpected, timeutilFormatDateTimeCalled, "Unexpected number of calls to timeutilFormatDateTime")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	assert.Equal(t, apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled, "Unexpected number of calls to apperrorConsolidateAllErrors")
	isServerCertificateAvailableFunc = isServerCertificateAvailable
	assert.Equal(t, isServerCertificateAvailableFuncExpected, isServerCertificateAvailableFuncCalled, "Unexpected number of calls to isServerCertificateAvailableFunc")
	isCaCertificateAvailableFunc = isCaCertificateAvailable
	assert.Equal(t, isCaCertificateAvailableFuncExpected, isCaCertificateAvailableFuncCalled, "Unexpected number of calls to isCaCertificateAvailableFunc")
	validateStringFunctionFunc = validateStringFunction
	assert.Equal(t, validateStringFunctionFuncExpected, validateStringFunctionFuncCalled, "Unexpected number of calls to validateStringFunctionFunc")
	validateBooleanFunctionFunc = validateBooleanFunction
	assert.Equal(t, validateBooleanFunctionFuncExpected, validateBooleanFunctionFuncCalled, "Unexpected number of calls to validateBooleanFunctionFunc")

	AppVersion = nil
	AppPort = nil
	AppName = nil
	AppPath = nil
	IsLocalhost = nil
	ServeHTTPS = nil
	ServerCertContent = nil
	ServerKeyContent = nil
	ValidateClientCert = nil
	CaCertContent = nil
}
