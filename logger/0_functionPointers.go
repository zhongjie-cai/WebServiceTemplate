package logger

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/timeutil"
)

// func pointers for injection / testing: logger.go
var (
	fmtPrintf                  = fmt.Printf
	fmtPrintln                 = fmt.Println
	uuidNew                    = uuid.New
	uuidParse                  = uuid.Parse
	fmtSprintf                 = fmt.Sprintf
	timeutilGetTimeNowUTC      = timeutil.GetTimeNowUTC
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	apperrorGetCustomError     = apperror.GetCustomError
	defaultLoggingFunc         = defaultLogging
	prepareLoggingFunc         = prepareLogging
)
