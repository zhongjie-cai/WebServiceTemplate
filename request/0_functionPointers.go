package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httputil"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// func pointers for injection / testing: logCategory.go
var (
	uuidParse              = uuid.Parse
	uuidNew                = uuid.New
	apperrorGetCustomError = apperror.GetCustomError
	ioutilReadAll          = ioutil.ReadAll
	ioutilNopCloser        = ioutil.NopCloser
	bytesNewBuffer         = bytes.NewBuffer
	httputilDumpRequest    = httputil.DumpRequest
	fmtSprintf             = fmt.Sprintf
)
