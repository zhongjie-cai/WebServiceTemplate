package session

import (
	"encoding/json"
	"fmt"
	"net/textproto"
	"runtime"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/network"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
)

// func pointers for injection / testing: logger.go
var (
	uuidNew                         = uuid.New
	jsonMarshal                     = json.Marshal
	jsonUnmarshal                   = json.Unmarshal
	fmtErrorf                       = fmt.Errorf
	muxVars                         = mux.Vars
	loggerAPIRequest                = logger.APIRequest
	requestGetRequestBody           = request.GetRequestBody
	apperrorGetBadRequestError      = apperror.GetBadRequestError
	textprotoCanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey
	getFunc                         = Get
	jsonutilTryUnmarshal            = jsonutil.TryUnmarshal
	getRequestFunc                  = GetRequest
	getAllQueriesFunc               = getAllQueries
	getAllHeadersFunc               = getAllHeaders
	isLoggingTypeMatchFunc          = isLoggingTypeMatch
	isLoggingLevelMatchFunc         = isLoggingLevelMatch
	runtimeCaller                   = runtime.Caller
	runtimeFuncForPC                = runtime.FuncForPC
	getMethodNameFunc               = getMethodName
	strconvItoa                     = strconv.Itoa
	loggerMethodEnter               = logger.MethodEnter
	loggerMethodParameter           = logger.MethodParameter
	loggerMethodLogic               = logger.MethodLogic
	loggerMethodReturn              = logger.MethodReturn
	loggerMethodExit                = logger.MethodExit
	networkNewNetworkRequest        = network.NewNetworkRequest
)
