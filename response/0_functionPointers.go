package response

import (
	"net/http"
	"strconv"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
)

// func pointers for injection / testing: swagger.go
var (
	strconvItoa                    = strconv.Itoa
	jsonutilMarshalIgnoreError     = jsonutil.MarshalIgnoreError
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	loggerAPIResponse              = logger.APIResponse
	httpStatusText                 = http.StatusText
	writeResponseFunc              = writeResponse
	getAppErrorFunc                = getAppError
	generateErrorResponseFunc      = generateErrorResponse
	createOkResponseFunc           = createOkResponse
	createErrorResponseFunc        = createErrorResponse
	constructResponseFunc          = constructResponse
)
