package config

import (
	"os"
	"strings"
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
	stringsEqualFoldExpected                            int
	stringsEqualFoldCalled                              int
	initializeBootTimeFuncExpected                      int
	initializeBootTimeFuncCalled                        int
	initializeCryptoKeyFuncExpected                     int
	initializeCryptoKeyFuncCalled                       int
	decryptFromEnvironmentVariableFuncExpected          int
	decryptFromEnvironmentVariableFuncCalled            int
	initializeGeneralEnvironmentVariablesFuncExpected   int
	initializeGeneralEnvironmentVariablesFuncCalled     int
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
	stringsEqualFoldExpected = 0
	stringsEqualFoldCalled = 0
	stringsEqualFold = func(s, b string) bool {
		stringsEqualFoldCalled++
		return false
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
	initializeGeneralEnvironmentVariablesFuncExpected = 0
	initializeGeneralEnvironmentVariablesFuncCalled = 0
	initializeGeneralEnvironmentVariablesFunc = func() error {
		initializeGeneralEnvironmentVariablesFuncCalled++
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
	assert.Equal(t, timeutilGetTimeNowUTCExpected, timeutilGetTimeNowUTCCalled, "Unexpected method call to timeutilGetTimeNowUTC")
	timeutilFormatDateTime = timeutil.FormatDateTime
	assert.Equal(t, timeutilFormatDateTimeExpected, timeutilFormatDateTimeCalled, "Unexpected method call to timeutilFormatDateTime")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
	apperrorConsolidateAllErrors = apperror.ConsolidateAllErrors
	assert.Equal(t, apperrorConsolidateAllErrorsExpected, apperrorConsolidateAllErrorsCalled, "Unexpected method call to apperrorConsolidateAllErrors")
	getEnvironmentVariable = os.Getenv
	assert.Equal(t, getEnvironmentVariableExpected, getEnvironmentVariableCalled, "Unexpected method call to getEnvironmentVariable")
	cryptoDecrypt = crypto.Decrypt
	assert.Equal(t, cryptoDecryptExpected, cryptoDecryptCalled, "Unexpected method call to cryptoDecrypt")
	stringsEqualFold = strings.EqualFold
	assert.Equal(t, stringsEqualFoldExpected, stringsEqualFoldCalled, "Unexpected method call to stringsEqualFold")
	initializeBootTimeFunc = initializeBootTime
	assert.Equal(t, initializeBootTimeFuncExpected, initializeBootTimeFuncCalled, "Unexpected method call to initializeBootTimeFunc")
	initializeCryptoKeyFunc = initializeCryptoKey
	assert.Equal(t, initializeCryptoKeyFuncExpected, initializeCryptoKeyFuncCalled, "Unexpected method call to initializeCryptoKeyFunc")
	decryptFromEnvironmentVariableFunc = decryptFromEnvironmentVariable
	assert.Equal(t, decryptFromEnvironmentVariableFuncExpected, decryptFromEnvironmentVariableFuncCalled, "Unexpected method call to decryptFromEnvironmentVariableFunc")
	initializeGeneralEnvironmentVariablesFunc = initializeGeneralEnvironmentVariables
	assert.Equal(t, initializeGeneralEnvironmentVariablesFuncExpected, initializeGeneralEnvironmentVariablesFuncCalled, "Unexpected method call to initializeGeneralEnvironmentVariablesFunc")
	initializeEncryptedEnvironmentVariablesFunc = initializeEncryptedEnvironmentVariables
	assert.Equal(t, initializeEncryptedEnvironmentVariablesFuncExpected, initializeEncryptedEnvironmentVariablesFuncCalled, "Unexpected method call to initializeEncryptedEnvironmentVariablesFunc")

	// Reset state variables
	appVersion = "0.0.0.0"
	appPort = "18605"
	appName = "WebServiceTemplate"
	appPath = "."
	cryptoKey = ""
	bootTime = ""
	isLocalhost = false
	sendClientCert = false
	clientCertContent = ""
	clientKeyContent = ""
	serveHTTPS = false
	serverCertContent = ""
	serverKeyContent = ""
	validateClientCert = false
	caCertContent = ""
}
