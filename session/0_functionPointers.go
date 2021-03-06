package session

import (
	"encoding/json"
	"fmt"
	"net/textproto"
	"reflect"
	"runtime"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/network"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
)

// func pointers for injection / testing: session.go
var (
	reflectValueOf                  = reflect.ValueOf
	isInterfaceValueNilFunc         = isInterfaceValueNil
	uuidNew                         = uuid.New
	jsonMarshal                     = json.Marshal
	jsonUnmarshal                   = json.Unmarshal
	fmtErrorf                       = fmt.Errorf
	muxVars                         = mux.Vars
	loggerAPIRequest                = logger.APIRequest
	requestGetRequestBody           = request.GetRequestBody
	apperrorGetBadRequestError      = apperror.GetBadRequestError
	textprotoCanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey
	jsonutilTryUnmarshal            = jsonutil.TryUnmarshal
	headerutilLogHTTPHeaderForName  = headerutil.LogHTTPHeaderForName
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
	getAllowedLogTypeFunc           = getAllowedLogType
	getAllowedLogLevelFunc          = getAllowedLogLevel
	certificateHasClientCert        = certificate.HasClientCert
	shouldSendClientCertFunc        = shouldSendClientCert
)
