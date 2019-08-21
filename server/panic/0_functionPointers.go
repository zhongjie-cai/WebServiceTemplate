package panic

import (
	"fmt"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// func pointers for injection / testing: panic.go
var (
	fmtErrorf                      = fmt.Errorf
	getRecoverErrorFunc            = getRecoverError
	loggerAppRoot                  = logger.AppRoot
	responseWrite                  = response.Write
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	getDebugStackFunc              = getDebugStack
)
