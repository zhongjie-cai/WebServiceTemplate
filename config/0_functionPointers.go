package config

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

// func pointers for injection / testing: config.go
var (
	timeutilGetTimeNowUTC            = timeutil.GetTimeNowUTC
	timeutilFormatDateTime           = timeutil.FormatDateTime
	apperrorWrapSimpleError          = apperror.WrapSimpleError
	apperrorConsolidateAllErrors     = apperror.ConsolidateAllErrors
	isServerCertificateAvailableFunc = isServerCertificateAvailable
	isCaCertificateAvailableFunc     = isCaCertificateAvailable
	validateStringFunctionFunc       = validateStringFunction
	validateBooleanFunctionFunc      = validateBooleanFunction
)
