package network

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/network/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	httpClient *http.Client
)

func clientDo(
	httpClient *http.Client,
	httpRequest *http.Request,
) (*http.Response, error) {
	return httpClient.Do(
		httpRequest,
	)
}

func customizeRoundTripper(original http.RoundTripper) http.RoundTripper {
	if customization.HTTPRoundTripper == nil {
		return original
	}
	return customization.HTTPRoundTripper(
		original,
	)
}

func getHTTPTransport(sendClientCert bool) http.RoundTripper {
	var httpTransport = http.DefaultTransport
	if sendClientCert {
		var clientCert = certificateGetClientCertificate()
		if clientCert != nil {
			var tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{
					*clientCert,
				},
			}
			httpTransport = &http.Transport{
				TLSClientConfig: tlsConfig,
				Proxy:           http.ProxyFromEnvironment,
			}
		} else {
			loggerAppRoot(
				"network",
				"getHTTPTransport",
				"Failed to load client certificate for mTLS communications; fallback to default HTTP transport",
			)
		}
	}
	return customizeRoundTripperFunc(
		httpTransport,
	)
}

// Initialize creates a singleton instance for the network package to make HTTP request to external web services
func Initialize(
	sendClientCert bool,
	networkTimeout time.Duration,
) {
	var httpTransport = getHTTPTransportFunc(
		sendClientCert,
	)
	httpClient = &http.Client{
		Transport: httpTransport,
		Timeout:   networkTimeout,
	}
}

type networkRequest struct {
	session sessionModel.Session
	method  string
	url     string
	payload string
	header  map[string]string
}

// NewNetworkRequest creates a new network request for consumer to use
func NewNetworkRequest(
	session sessionModel.Session,
	method string,
	url string,
	payload string,
	header map[string]string,
) model.NetworkRequest {
	return &networkRequest{
		session,
		method,
		url,
		payload,
		header,
	}
}

func customizeHTTPRequest(httpRequest *http.Request) *http.Request {
	if customization.WrapHTTPRequest == nil {
		return httpRequest
	}
	return customization.WrapHTTPRequest(
		httpRequest,
	)
}

func createHTTPRequest(networkRequest *networkRequest) (*http.Request, error) {
	var requestBody = stringsNewReader(
		networkRequest.payload,
	)
	var requestObject, requestError = httpNewRequest(
		networkRequest.method,
		networkRequest.url,
		requestBody,
	)
	if requestError != nil {
		return nil,
			apperrorWrapSimpleError(
				[]error{requestError},
				"Failed to generate request to [%v]",
				networkRequest.url,
			)
	}
	loggerNetworkCall(
		networkRequest.session,
		networkRequest.method,
		networkRequest.url,
		"",
	)
	loggerNetworkRequest(
		networkRequest.session,
		"Payload",
		"",
		networkRequest.payload,
	)
	requestObject.Header = make(http.Header)
	for name, value := range networkRequest.header {
		loggerNetworkRequest(
			networkRequest.session,
			"Header",
			name,
			value,
		)
		requestObject.Header.Add(name, value)
	}
	return customizeHTTPRequestFunc(
		requestObject,
	), nil
}

func logErrorResponse(session sessionModel.Session, responseError error) {
	loggerNetworkFinish(
		session,
		"Error",
		"",
		"",
	)
	loggerNetworkResponse(
		session,
		"Message",
		"",
		"%v",
		responseError,
	)
}

func logHTTPResponse(session sessionModel.Session, response *http.Response) {
	if response == nil {
		return
	}
	var (
		responseStatus     = response.Status
		responseStatusCode = response.StatusCode
		responseBody, _    = ioutilReadAll(response.Body)
		responseHeaders    = response.Header
	)
	response.Body.Close()
	response.Body = ioutilNopCloser(
		bytesNewBuffer(
			responseBody,
		),
	)
	loggerNetworkFinish(
		session,
		responseStatus,
		strconvItoa(responseStatusCode),
		"",
	)
	loggerNetworkResponse(
		session,
		"Body",
		"",
		string(responseBody),
	)
	for name, values := range responseHeaders {
		for _, value := range values {
			loggerNetworkResponse(
				session,
				"Header",
				name,
				value,
			)
		}
	}
}

func doRequestProcessing(networkRequest *networkRequest) (*http.Response, error) {
	var requestObject, requestError = createHTTPRequestFunc(
		networkRequest,
	)
	if requestError != nil {
		return nil, requestError
	}
	var responseObject, responseError = clientDoFunc(
		httpClient,
		requestObject,
	)
	if responseError != nil {
		logErrorResponseFunc(
			networkRequest.session,
			responseError,
		)
	} else {
		logHTTPResponseFunc(
			networkRequest.session,
			responseObject,
		)
	}
	return responseObject, responseError
}

// ProcessRaw sends the network request over the wire, retrieves the response, and returns that response and error if applicable
func (networkRequest *networkRequest) ProcessRaw() (responseObject *http.Response, responseError error) {
	return doRequestProcessingFunc(
		networkRequest,
	)
}

func parseResponse(body io.ReadCloser, dataTemplate interface{}) error {
	var bodyBytes, bodyError = ioutilReadAll(
		body,
	)
	if bodyError != nil {
		return bodyError
	}
	return jsonUnmarshal(
		bodyBytes,
		dataTemplate,
	)
}

// Process sends the network request over the wire, retrieves and serialize the response to dataTemplate, and provides status code, header and error if applicable
func (networkRequest *networkRequest) Process(dataTemplate interface{}) (statusCode int, responseHeader http.Header, responseError error) {
	var responseObject *http.Response
	responseObject, responseError = doRequestProcessingFunc(
		networkRequest,
	)
	if responseError != nil {
		if responseObject == nil {
			return http.StatusInternalServerError, make(http.Header), responseError
		}
	} else {
		if responseObject == nil {
			return http.StatusNoContent, make(http.Header), nil
		}
		responseError = parseResponseFunc(
			responseObject.Body,
			dataTemplate,
		)
	}
	return responseObject.StatusCode, responseObject.Header, responseError
}