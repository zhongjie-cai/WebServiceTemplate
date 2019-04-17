package request

import (
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

// func pointers for injection / testing: logCategory.go
var (
	uuidParse               = uuid.Parse
	uuidNew                 = uuid.New
	logtypeFromString       = logtype.FromString
	apperrorWrapSimpleError = apperror.WrapSimpleError
	ioutilReadAll           = ioutil.ReadAll
	loggerAPIRequest        = logger.APIRequest
	getUUIDFromHeaderFunc   = getUUIDFromHeader
)
