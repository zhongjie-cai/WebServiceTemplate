package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/crypto"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

var (
	timeutilGetTimeNowUTCExpected                       int
	timeutilGetTimeNowUTCCalled                         int
	timeutilFormatDateTimeExpected                      int
	timeutilFormatDateTimeCalled                        int
	apperrorWrapSimpleErrorExpected                     int
	apperrorWrapSimpleErrorCalled                       int
	apperrorConsolidateAllErrorsExpected                int
	apperrorConsolidateAllErrorsCalled                  int
	getEnvironmentVariableExpected                      int
	getEnvironmentVariableCalled                        int
	cryptoDecryptExpected                               int
	cryptoDecryptCalled                                 int
	initializeBootTimeFuncExpected                      int
	initializeBootTimeFuncCalled                        int
	initializeCryptoKeyFuncExpected                     int
	initializeCryptoKeyFuncCalled                       int
	decryptFromEnvironmentVariableFuncExpected          int
	decryptFromEnvironmentVariableFuncCalled            int
	initializeEnvironmentVariablesFuncExpected          int
	initializeEnvironmentVariablesFuncCalled            int
	initializeEncryptedEnvironmentVariablesFuncExpected int
	initializeEncryptedEnvironmentVariablesFuncCalled   int
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
	getEnvironmentVariableExpected = 0
	getEnvironmentVariableCalled = 0
	getEnvironmentVariable = func(key string) string {
		getEnvironmentVariableCalled++
		return ""
	}
	cryptoDecryptExpected = 0
	cryptoDecryptCalled = 0
	cryptoDecrypt = func(cipherText string, key string) (string, error) {
		cryptoDecryptCalled++
		return "", nil
	}
	initializeBootTimeFuncExpected = 0
	initializeBootTimeFuncCalled = 0
	initializeBootTimeFunc = func() {
		initializeBootTimeFuncCalled++
	}
	initializeCryptoKeyFuncExpected = 0
	initializeCryptoKeyFuncCalled = 0
	initializeCryptoKeyFunc = func() error {
		initializeCryptoKeyFuncCalled++
		return nil
	}
	decryptFromEnvironmentVariableFuncExpected = 0
	decryptFromEnvironmentVariableFuncCalled = 0
	decryptFromEnvironmentVariableFunc = func(name string) (string, error) {
		decryptFromEnvironmentVariableFuncCalled++
		return "", nil
	}
	initializeEnvironmentVariablesFuncExpected = 0
	initializeEnvironmentVariablesFuncCalled = 0
	initializeEnvironmentVariablesFunc = func() error {
		initializeEnvironmentVariablesFuncCalled++
		return nil
	}
	initializeEncryptedEnvironmentVariablesFuncExpected = 0
	initializeEncryptedEnvironmentVariablesFuncCalled = 0
	initializeEncryptedEnvironmentVariablesFunc = func() error {
		initializeEncryptedEnvironmentVariablesFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	timeutilGetTimeNowUTC = timeutil.GetTimeNowUTC
	if timeutilGetTimeNowUTCExpected != timeutilGetTimeNowUTCCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to timeutilGetTimeNowUTC, expected %v, actual %v", timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled))
	}
	timeutilFormatDateTime = timeutil.FormatDateTime
	if timeutilFormatDateTimeExpected != timeutilFormatDateTimeCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to timeutilFormatDateTime, expected %v, actual %v", timeutilFormatDateTimeExpected, timeutilFormatDateTimeCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	if apperrorConsolidateAllErrorsExpected != apperrorConsolidateAllErrorsCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorConsolidateAllErrors, expected %v, actual %v", apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled))
	}
	getEnvironmentVariable = os.Getenv
	if getEnvironmentVariableExpected != getEnvironmentVariableCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getEnvironmentVariable, expected %v, actual %v", getEnvironmentVariableExpected, getEnvironmentVariableCalled))
	}
	cryptoDecrypt = crypto.Decrypt
	if cryptoDecryptExpected != cryptoDecryptCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to cryptoDecrypt, expected %v, actual %v", cryptoDecryptExpected, cryptoDecryptCalled))
	}
	initializeBootTimeFunc = initializeBootTime
	if initializeBootTimeFuncExpected != initializeBootTimeFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to initializeBootTimeFunc, expected %v, actual %v", initializeBootTimeFuncExpected, initializeBootTimeFuncCalled))
	}
	initializeCryptoKeyFunc = initializeCryptoKey
	if initializeCryptoKeyFuncExpected != initializeCryptoKeyFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to initializeCryptoKeyFunc, expected %v, actual %v", initializeCryptoKeyFuncExpected, initializeCryptoKeyFuncCalled))
	}
	decryptFromEnvironmentVariableFunc = decryptFromEnvironmentVariable
	if decryptFromEnvironmentVariableFuncExpected != decryptFromEnvironmentVariableFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to decryptFromEnvironmentVariableFunc, expected %v, actual %v", decryptFromEnvironmentVariableFuncExpected, decryptFromEnvironmentVariableFuncCalled))
	}
	initializeEnvironmentVariablesFunc = initializeEnvironmentVariables
	if initializeEnvironmentVariablesFuncExpected != initializeEnvironmentVariablesFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to initializeEnvironmentVariablesFunc, expected %v, actual %v", initializeEnvironmentVariablesFuncExpected, initializeEnvironmentVariablesFuncCalled))
	}
	initializeEncryptedEnvironmentVariablesFunc = initializeEncryptedEnvironmentVariables
	if initializeEncryptedEnvironmentVariablesFuncExpected != initializeEncryptedEnvironmentVariablesFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to initializeEncryptedEnvironmentVariablesFunc, expected %v, actual %v", initializeEncryptedEnvironmentVariablesFuncExpected, initializeEncryptedEnvironmentVariablesFuncCalled))
	}

	// Reset state variables
	appVersion = "0.0.0.0"
	appPort = "443"
	appName = "WebServiceTemplate"
	appPath = "."
	cryptoKey = ""
	bootTime = ""
}
