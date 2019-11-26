package request

import (
	"bytes"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

// func pointers for injection / testing: logCategory.go
var (
	uuidParse              = uuid.Parse
	uuidNew                = uuid.New
	logtypeFromString      = logtype.FromString
	loglevelFromString     = loglevel.FromString
	apperrorGetCustomError = apperror.GetCustomError
	ioutilReadAll          = ioutil.ReadAll
	ioutilNopCloser        = ioutil.NopCloser
	bytesNewBuffer         = bytes.NewBuffer
)
