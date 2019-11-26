package config

import (
	"fmt"
	"reflect"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

// func pointers for injection / testing: config.go
var (
	timeutilGetTimeNowUTC            = timeutil.GetTimeNowUTC
	timeutilFormatDateTime           = timeutil.FormatDateTime
	apperrorGetCustomError           = apperror.GetCustomError
	apperrorWrapSimpleError          = apperror.WrapSimpleError
	reflectValueOf                   = reflect.ValueOf
	fmtSprintf                       = fmt.Sprintf
	functionPointerEqualsFunc        = functionPointerEquals
	isServerCertificateAvailableFunc = isServerCertificateAvailable
	isCaCertificateAvailableFunc     = isCaCertificateAvailable
	validateStringFunctionFunc       = validateStringFunction
	validateBooleanFunctionFunc      = validateBooleanFunction
)
