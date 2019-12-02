package request

import (
	"bytes"
	"io/ioutil"

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
)
