package apperror

import (
	"fmt"
	"strings"
)

// func pointers for injection / testing: apperror.go
var (
	fmtSprintf             = fmt.Sprintf
	fmtErrorf              = fmt.Errorf
	stringsJoin            = strings.Join
	cleanupInnerErrorsFunc = cleanupInnerErrors
	wrapErrorFunc          = WrapError
	wrapSimpleErrorFunc    = WrapSimpleError
)
