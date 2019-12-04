package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	stringsNewReaderExpected                int
	stringsNewReaderCalled                  int
	httpNewRequestExpected                  int
	httpNewRequestCalled                    int
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	loggerNetworkCallExpected               int
	loggerNetworkCallCalled                 int
	loggerNetworkRequestExpected            int
	loggerNetworkRequestCalled              int
	loggerNetworkResponseExpected           int
	loggerNetworkResponseCalled             int
	loggerNetworkFinishExpected             int
	loggerNetworkFinishCalled               int
	loggerAppRootExpected                   int
	loggerAppRootCalled                     int
	ioutilReadAllExpected                   int
	ioutilReadAllCalled                     int
	strconvItoaExpected                     int
	strconvItoaCalled                       int
	ioutilNopCloserExpected                 int
	ioutilNopCloserCalled                   int
	bytesNewBufferExpected                  int
	bytesNewBufferCalled                    int
	createHTTPRequestFuncExpected           int
	createHTTPRequestFuncCalled             int
	clientDoFuncExpected                    int
	clientDoFuncCalled                      int
	logErrorResponseFuncExpected            int
	logErrorResponseFuncCalled              int
	logHTTPResponseFuncExpected             int
	logHTTPResponseFuncCalled               int
	doRequestProcessingFuncExpected         int
	doRequestProcessingFuncCalled           int
	jsonUnmarshalExpected                   int
	jsonUnmarshalCalled                     int
	parseResponseFuncExpected               int
	parseResponseFuncCalled                 int
	certificateGetClientCertificateExpected int
	certificateGetClientCertificateCalled   int
	customizeRoundTripperFuncExpected       int
	customizeRoundTripperFuncCalled         int
	getHTTPTransportFuncExpected            int
	getHTTPTransportFuncCalled              int
	customizeHTTPRequestFuncExpected        int
	customizeHTTPRequestFuncCalled          int
	customizationHTTPRoundTripperExpected   int
	customizationHTTPRoundTripperCalled     int
	customizationWrapHTTPRequestExpected    int
	customizationWrapHTTPRequestCalled      int
)

func createMock(t *testing.T) {
	stringsNewReaderExpected = 0
	stringsNewReaderCalled = 0
	stringsNewReader = func(s string) *strings.Reader {
		stringsNewReaderCalled++
		return nil
	}
	httpNewRequestExpected = 0
	httpNewRequestCalled = 0
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		httpNewRequestCalled++
		return nil, nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	loggerNetworkCallExpected = 0
	loggerNetworkCallCalled = 0
	loggerNetworkCall = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkCallCalled++
	}
	loggerNetworkRequestExpected = 0
	loggerNetworkRequestCalled = 0
	loggerNetworkRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkRequestCalled++
	}
	loggerNetworkResponseExpected = 0
	loggerNetworkResponseCalled = 0
	loggerNetworkResponse = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkResponseCalled++
	}
	loggerNetworkFinishExpected = 0
	loggerNetworkFinishCalled = 0
	loggerNetworkFinish = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkFinishCalled++
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	ioutilReadAllExpected = 0
	ioutilReadAllCalled = 0
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		return nil, nil
	}
	strconvItoaExpected = 0
	strconvItoaCalled = 0
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return ""
	}
	ioutilNopCloserExpected = 0
	ioutilNopCloserCalled = 0
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		return nil
	}
	bytesNewBufferExpected = 0
	bytesNewBufferCalled = 0
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		return nil
	}
	createHTTPRequestFuncExpected = 0
	createHTTPRequestFuncCalled = 0
	createHTTPRequestFunc = func(networkRequest *networkRequest) (*http.Request, error) {
		createHTTPRequestFuncCalled++
		return nil, nil
	}
	clientDoFuncExpected = 0
	clientDoFuncCalled = 0
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		return nil, nil
	}
	logErrorResponseFuncExpected = 0
	logErrorResponseFuncCalled = 0
	logErrorResponseFunc = func(session sessionModel.Session, responseError error) {
		logErrorResponseFuncCalled++
	}
	logHTTPResponseFuncExpected = 0
	logHTTPResponseFuncCalled = 0
	logHTTPResponseFunc = func(session sessionModel.Session, response *http.Response) {
		logHTTPResponseFuncCalled++
	}
	doRequestProcessingFuncExpected = 0
	doRequestProcessingFuncCalled = 0
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		return nil, nil
	}
	jsonUnmarshalExpected = 0
	jsonUnmarshalCalled = 0
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return nil
	}
	parseResponseFuncExpected = 0
	parseResponseFuncCalled = 0
	parseResponseFunc = func(body io.ReadCloser, dataTemplate interface{}) error {
		parseResponseFuncCalled++
		return nil
	}
	certificateGetClientCertificateExpected = 0
	certificateGetClientCertificateCalled = 0
	certificateGetClientCertificate = func() *tls.Certificate {
		certificateGetClientCertificateCalled++
		return nil
	}
	customizeRoundTripperFuncExpected = 0
	customizeRoundTripperFuncCalled = 0
	customizeRoundTripperFunc = func(original http.RoundTripper) http.RoundTripper {
		customizeRoundTripperFuncCalled++
		return nil
	}
	getHTTPTransportFuncExpected = 0
	getHTTPTransportFuncCalled = 0
	getHTTPTransportFunc = func(sendClientCert bool) http.RoundTripper {
		getHTTPTransportFuncCalled++
		return nil
	}
	customizeHTTPRequestFuncExpected = 0
	customizeHTTPRequestFuncCalled = 0
	customizeHTTPRequestFunc = func(httpRequest *http.Request) *http.Request {
		customizeHTTPRequestFuncCalled++
		return nil
	}
	customizationHTTPRoundTripperExpected = 0
	customizationHTTPRoundTripperCalled = 0
	customization.HTTPRoundTripper = nil
	customizationWrapHTTPRequestExpected = 0
	customizationWrapHTTPRequestCalled = 0
	customization.WrapHTTPRequest = nil
}

func verifyAll(t *testing.T) {
	stringsNewReader = strings.NewReader
	assert.Equal(t, stringsNewReaderExpected, stringsNewReaderCalled, "Unexpected number of calls to method stringsNewReader")
	httpNewRequest = http.NewRequest
	assert.Equal(t, httpNewRequestExpected, httpNewRequestCalled, "Unexpected number of calls to method httpNewRequest")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to method apperrorWrapSimpleError")
	loggerNetworkCall = logger.NetworkCall
	assert.Equal(t, loggerNetworkCallExpected, loggerNetworkCallCalled, "Unexpected number of calls to method loggerNetworkCall")
	loggerNetworkRequest = logger.NetworkRequest
	assert.Equal(t, loggerNetworkRequestExpected, loggerNetworkRequestCalled, "Unexpected number of calls to method loggerNetworkRequest")
	loggerNetworkResponse = logger.NetworkResponse
	assert.Equal(t, loggerNetworkResponseExpected, loggerNetworkResponseCalled, "Unexpected number of calls to method loggerNetworkResponse")
	loggerNetworkFinish = logger.NetworkFinish
	assert.Equal(t, loggerNetworkFinishExpected, loggerNetworkFinishCalled, "Unexpected number of calls to method loggerNetworkFinish")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to method loggerAppRoot")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected number of calls to method ioutilReadAll")
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected number of calls to method strconvItoa")
	ioutilNopCloser = ioutil.NopCloser
	assert.Equal(t, ioutilNopCloserExpected, ioutilNopCloserCalled, "Unexpected number of calls to method ioutilNopCloser")
	bytesNewBuffer = bytes.NewBuffer
	assert.Equal(t, bytesNewBufferExpected, bytesNewBufferCalled, "Unexpected number of calls to method bytesNewBuffer")
	createHTTPRequestFunc = createHTTPRequest
	assert.Equal(t, createHTTPRequestFuncExpected, createHTTPRequestFuncCalled, "Unexpected number of calls to method createHTTPRequestFunc")
	clientDoFunc = clientDo
	assert.Equal(t, clientDoFuncExpected, clientDoFuncCalled, "Unexpected number of calls to method clientDoFunc")
	logErrorResponseFunc = logErrorResponse
	assert.Equal(t, logErrorResponseFuncExpected, logErrorResponseFuncCalled, "Unexpected number of calls to method logErrorResponseFunc")
	logHTTPResponseFunc = logHTTPResponse
	assert.Equal(t, logHTTPResponseFuncExpected, logHTTPResponseFuncCalled, "Unexpected number of calls to method logHTTPResponseFunc")
	doRequestProcessingFunc = doRequestProcessing
	assert.Equal(t, doRequestProcessingFuncExpected, doRequestProcessingFuncCalled, "Unexpected number of calls to method doRequestProcessingFunc")
	jsonUnmarshal = json.Unmarshal
	assert.Equal(t, jsonUnmarshalExpected, jsonUnmarshalCalled, "Unexpected number of calls to method jsonUnmarshal")
	parseResponseFunc = parseResponse
	assert.Equal(t, parseResponseFuncExpected, parseResponseFuncCalled, "Unexpected number of calls to method parseResponseFunc")
	certificateGetClientCertificate = certificate.GetClientCertificate
	assert.Equal(t, certificateGetClientCertificateExpected, certificateGetClientCertificateCalled, "Unexpected number of calls to method certificateGetClientCertificate")
	customizeRoundTripperFunc = customizeRoundTripper
	assert.Equal(t, customizeRoundTripperFuncExpected, customizeRoundTripperFuncCalled, "Unexpected number of calls to method customizeRoundTripperFunc")
	getHTTPTransportFunc = getHTTPTransport
	assert.Equal(t, getHTTPTransportFuncExpected, getHTTPTransportFuncCalled, "Unexpected number of calls to method getHTTPTransportFunc")
	customizeHTTPRequestFunc = customizeHTTPRequest
	assert.Equal(t, customizeHTTPRequestFuncExpected, customizeHTTPRequestFuncCalled, "Unexpected number of calls to method customizeHTTPRequestFunc")

	httpClient = nil
	customization.HTTPRoundTripper = nil
	customization.WrapHTTPRequest = nil
}

// mock structs
type dummySession struct {
	t *testing.T
}

func (session *dummySession) GetID() uuid.UUID {
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySession) GetName() string {
	assert.Fail(session.t, "Unexpected call to GetName")
	return ""
}

func (session *dummySession) GetRequest() *http.Request {
	assert.Fail(session.t, "Unexpected call to GetRequest")
	return nil
}

func (session *dummySession) GetResponseWriter() http.ResponseWriter {
	assert.Fail(session.t, "Unexpected call to GetResponseWriter")
	return nil
}

func (session *dummySession) GetRequestBody(dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestBody")
	return nil
}

func (session *dummySession) GetRequestParameter(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestParameter")
	return nil
}

func (session *dummySession) GetRequestQuery(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQuery")
	return nil
}

func (session *dummySession) GetRequestQueries(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestQueries")
	return nil
}

func (session *dummySession) GetRequestHeader(name string, dataTemplate interface{}) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeader")
	return nil
}

func (session *dummySession) GetRequestHeaders(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	assert.Fail(session.t, "Unexpected call to GetRequestHeaders")
	return nil
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	assert.Fail(session.t, "Unexpected call to Attach")
	return false
}

func (session *dummySession) Detach(name string) bool {
	assert.Fail(session.t, "Unexpected call to Detach")
	return false
}

func (session *dummySession) GetAttachment(name string, dataTemplate interface{}) bool {
	assert.Fail(session.t, "Unexpected call to GetAttachment")
	return false
}

func (session *dummySession) IsLoggingAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool {
	assert.Fail(session.t, "Unexpected call to IsLoggingAllowed")
	return false
}