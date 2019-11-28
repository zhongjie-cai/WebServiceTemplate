package apperror

import (
	"fmt"
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
)

// func pointers for injection / testing: apperror.go
var (
	fmtSprintf                 = fmt.Sprintf
	fmtErrorf                  = fmt.Errorf
	stringsJoin                = strings.Join
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	cleanupInnerErrorsFunc     = cleanupInnerErrors
	wrapErrorFunc              = WrapError
	wrapSimpleErrorFunc        = WrapSimpleError
)
