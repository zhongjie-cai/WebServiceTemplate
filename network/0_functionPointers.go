package network

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
)

// func pointers for injection / testing: logCategory.go
var (
	stringsNewReader                = strings.NewReader
	httpNewRequest                  = http.NewRequest
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	loggerNetworkCall               = logger.NetworkCall
	loggerNetworkRequest            = logger.NetworkRequest
	loggerNetworkResponse           = logger.NetworkResponse
	loggerNetworkFinish             = logger.NetworkFinish
	loggerAppRoot                   = logger.AppRoot
	ioutilReadAll                   = ioutil.ReadAll
	httpStatusText                  = http.StatusText
	strconvItoa                     = strconv.Itoa
	ioutilNopCloser                 = ioutil.NopCloser
	bytesNewBuffer                  = bytes.NewBuffer
	timeSleep                       = time.Sleep
	headerutilLogHTTPHeader         = headerutil.LogHTTPHeader
	createHTTPRequestFunc           = createHTTPRequest
	clientDoFunc                    = clientDo
	delayForRetryFunc               = delayForRetry
	clientDoWithRetryFunc           = clientDoWithRetry
	logErrorResponseFunc            = logErrorResponse
	logHTTPResponseFunc             = logHTTPResponse
	doRequestProcessingFunc         = doRequestProcessing
	jsonutilTryUnmarshal            = jsonutil.TryUnmarshal
	parseResponseFunc               = parseResponse
	certificateGetClientCertificate = certificate.GetClientCertificate
	customizeRoundTripperFunc       = customizeRoundTripper
	getHTTPTransportFunc            = getHTTPTransport
	customizeHTTPRequestFunc        = customizeHTTPRequest
	getClientForRequestFunc         = getClientForRequest
)
