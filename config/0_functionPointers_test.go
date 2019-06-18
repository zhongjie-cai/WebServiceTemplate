package config

import (
	"fmt"
	"reflect"
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
	reflectValueOfExpected                   int
	reflectValueOfCalled                     int
	fmtSprintfExpected                       int
	fmtSprintfCalled                         int
	functionPointerEqualsFuncExpected        int
	functionPointerEqualsFuncCalled          int
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
	reflectValueOfExpected = 0
	reflectValueOfCalled = 0
	reflectValueOf = func(i interface{}) reflect.Value {
		reflectValueOfCalled++
		return reflect.Value{}
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	functionPointerEqualsFuncExpected = 0
	functionPointerEqualsFuncCalled = 0
	functionPointerEqualsFunc = func(left, right interface{}) bool {
		functionPointerEqualsFuncCalled++
		return false
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
	reflectValueOf = reflect.ValueOf
	assert.Equal(t, reflectValueOfExpected, reflectValueOfCalled, "Unexpected number of calls to reflectValueOf")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	functionPointerEqualsFunc = functionPointerEquals
	assert.Equal(t, functionPointerEqualsFuncExpected, functionPointerEqualsFuncCalled, "Unexpected number of calls to functionPointerEqualsFunc")
	isServerCertificateAvailableFunc = isServerCertificateAvailable
	assert.Equal(t, isServerCertificateAvailableFuncExpected, isServerCertificateAvailableFuncCalled, "Unexpected number of calls to isServerCertificateAvailableFunc")
	isCaCertificateAvailableFunc = isCaCertificateAvailable
	assert.Equal(t, isCaCertificateAvailableFuncExpected, isCaCertificateAvailableFuncCalled, "Unexpected number of calls to isCaCertificateAvailableFunc")
	validateStringFunctionFunc = validateStringFunction
	assert.Equal(t, validateStringFunctionFuncExpected, validateStringFunctionFuncCalled, "Unexpected number of calls to validateStringFunctionFunc")
	validateBooleanFunctionFunc = validateBooleanFunction
	assert.Equal(t, validateBooleanFunctionFuncExpected, validateBooleanFunctionFuncCalled, "Unexpected number of calls to validateBooleanFunctionFunc")

	AppVersion = defaultAppVersion
	AppPort = defaultAppPort
	AppName = defaultAppName
	AppPath = defaultAppPath
	IsLocalhost = defaultIsLocalhost
	ServeHTTPS = defaultServeHTTPS
	ServerCertContent = defaultServerCertContent
	ServerKeyContent = defaultServerKeyContent
	ValidateClientCert = defaultValidateClientCert
	CaCertContent = defaultCaCertContent
}
