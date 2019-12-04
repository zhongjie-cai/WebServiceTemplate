package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
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
	strconvItoa                     = strconv.Itoa
	ioutilNopCloser                 = ioutil.NopCloser
	bytesNewBuffer                  = bytes.NewBuffer
	createHTTPRequestFunc           = createHTTPRequest
	clientDoFunc                    = clientDo
	clientDoWithRetryFunc           = clientDoWithRetry
	logErrorResponseFunc            = logErrorResponse
	logHTTPResponseFunc             = logHTTPResponse
	doRequestProcessingFunc         = doRequestProcessing
	jsonUnmarshal                   = json.Unmarshal
	parseResponseFunc               = parseResponse
	certificateGetClientCertificate = certificate.GetClientCertificate
	customizeRoundTripperFunc       = customizeRoundTripper
	getHTTPTransportFunc            = getHTTPTransport
	customizeHTTPRequestFunc        = customizeHTTPRequest
)
