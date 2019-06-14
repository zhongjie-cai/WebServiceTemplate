package config

import (
	"os"
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/crypto"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

// func pointers for injection / testing: config.go
var (
	timeutilGetTimeNowUTC                       = timeutil.GetTimeNowUTC
	timeutilFormatDateTime                      = timeutil.FormatDateTime
	apperrorWrapSimpleError                     = apperror.WrapSimpleError
	apperrorConsolidateAllErrors                = apperror.ConsolidateAllErrors
	getEnvironmentVariable                      = os.Getenv
	cryptoDecrypt                               = crypto.Decrypt
	stringsEqualFold                            = strings.EqualFold
	initializeBootTimeFunc                      = initializeBootTime
	initializeCryptoKeyFunc                     = initializeCryptoKey
	decryptFromEnvironmentVariableFunc          = decryptFromEnvironmentVariable
	initializeGeneralEnvironmentVariablesFunc   = initializeGeneralEnvironmentVariables
	initializeEncryptedEnvironmentVariablesFunc = initializeEncryptedEnvironmentVariables
)
