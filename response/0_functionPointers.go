package response

import (
	"strconv"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// func pointers for injection / testing: swagger.go
var (
	strconvItoa                    = strconv.Itoa
	jsonutilMarshalIgnoreError     = jsonutil.MarshalIgnoreError
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	loggerAPIResponse              = logger.APIResponse
	loggerAPIExit                  = logger.APIExit
	sessionGetRequest              = session.GetRequest
	sessionGetResponseWriter       = session.GetResponseWriter
	sessionClearResponseWriter     = session.ClearResponseWriter
	getStatusCodeFunc              = getStatusCode
	writeResponseFunc              = writeResponse
	generateErrorResponseFunc      = generateErrorResponse
	createOkResponseFunc           = createOkResponse
	createErrorResponseFunc        = createErrorResponse
)
