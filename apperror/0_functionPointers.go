package apperror

import (
	"fmt"
	"strings"
)

// func pointers for injection / testing: error.go
var (
	fmtSprintf          = fmt.Sprintf
	fmtErrorf           = fmt.Errorf
	stringsJoin         = strings.Join
	wrapErrorFunc       = WrapError
	wrapSimpleErrorFunc = WrapSimpleError
)
