package network

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestGetClientForRequest_SendClientCert(t *testing.T) {
	// arrange
	var dummyHTTPClient1 = &http.Client{Timeout: time.Duration(rand.Int())}
	var dummyHTTPClient2 = &http.Client{Timeout: time.Duration(rand.Int())}

	// stub
	httpClientWithCert = dummyHTTPClient1
	httpClientNoCert = dummyHTTPClient2

	// mock
	createMock(t)

	// SUT + act
	var result = getClientForRequest(true)

	// assert
	assert.Equal(t, dummyHTTPClient1, result)

	// verify
	verifyAll(t)
}

func TestGetClientForRequest_NoSendClientCert(t *testing.T) {
	// arrange
	var dummyHTTPClient1 = &http.Client{Timeout: time.Duration(rand.Int())}
	var dummyHTTPClient2 = &http.Client{Timeout: time.Duration(rand.Int())}

	// stub
	httpClientWithCert = dummyHTTPClient1
	httpClientNoCert = dummyHTTPClient2

	// mock
	createMock(t)

	// SUT + act
	var result = getClientForRequest(false)

	// assert
	assert.Equal(t, dummyHTTPClient2, result)

	// verify
	verifyAll(t)
}

func TestClientDo(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequest, _ = http.NewRequest(
		"",
		"",
		nil,
	)

	// assert
	assert.NotPanics(
		t,
		func() {
			// SUT + act
			clientDo(
				dummyClient,
				dummyRequest,
			)
		},
	)
}

func TestDelayForRetry_NoCustomization(t *testing.T) {
	// arrange
	var dummyDelay = 3 * time.Second

	// mock
	createMock(t)

	// expect
	timeSleepExpected = 1
	timeSleep = func(d time.Duration) {
		timeSleepCalled++
		assert.Equal(t, dummyDelay, d)
	}

	// SUT + act
	delayForRetry()

	// verify
	verifyAll(t)
}

func TestDelayForRetry_WithCustomization(t *testing.T) {
	// arrange
	var dummyDelay = time.Duration(rand.Int())

	// mock
	createMock(t)

	// expect
	customizationDefaultNetworkRetryDelayExpected = 1
	customization.DefaultNetworkRetryDelay = func() time.Duration {
		customizationDefaultNetworkRetryDelayCalled++
		return dummyDelay
	}
	timeSleepExpected = 1
	timeSleep = func(d time.Duration) {
		timeSleepCalled++
		assert.Equal(t, dummyDelay, d)
	}

	// SUT + act
	delayForRetry()

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_ConnError_NoRetry(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = 0
	var dummyHTTPRetry = map[int]int{}
	var dummyResponseObject = &http.Response{}
	var dummyResponseError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 1
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		return dummyResponseObject, dummyResponseError
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_ConnError_RetryOK(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = 2
	var dummyHTTPRetry = map[int]int{}
	var dummyResponseObject = &http.Response{}
	var dummyResponseError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 2
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		if clientDoFuncCalled == 1 {
			return dummyResponseObject, dummyResponseError
		} else if clientDoFuncCalled == 2 {
			return dummyResponseObject, nil
		}
		return nil, nil
	}
	delayForRetryFuncExpected = 1
	delayForRetryFunc = func() {
		delayForRetryFuncCalled++
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_ConnError_RetryFail(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = 2
	var dummyHTTPRetry = map[int]int{}
	var dummyResponseObject = &http.Response{}
	var dummyResponseError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 3
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		return dummyResponseObject, dummyResponseError
	}
	delayForRetryFuncExpected = 2
	delayForRetryFunc = func() {
		delayForRetryFuncCalled++
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_HTTPError_NilResponse(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{}
	var dummyResponseObject *http.Response

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 1
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		return dummyResponseObject, nil
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_HTTPError_NoRetry(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{}
	var dummyResponseObject = &http.Response{}

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 1
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		return dummyResponseObject, nil
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_HTTPError_RetryOK(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = rand.Int()
	var dummyStatusCode = rand.Int()
	var dummyHTTPRetry = map[int]int{
		dummyStatusCode: 2,
	}
	var dummyResponseObject1 = &http.Response{
		StatusCode: dummyStatusCode,
	}
	var dummyResponseObject2 = &http.Response{}

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 2
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		if clientDoFuncCalled == 1 {
			return dummyResponseObject1, nil
		} else if clientDoFuncCalled == 2 {
			return dummyResponseObject2, nil
		}
		return nil, nil
	}
	delayForRetryFuncExpected = 1
	delayForRetryFunc = func() {
		delayForRetryFuncCalled++
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject2, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestClientDoWithRetry_HTTPError_RetryFail(t *testing.T) {
	// arrange
	var dummyClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyConnRetry = rand.Int()
	var dummyStatusCode = rand.Int()
	var dummyHTTPRetry = map[int]int{
		dummyStatusCode: 2,
	}
	var dummyResponseObject = &http.Response{
		StatusCode: dummyStatusCode,
	}

	// mock
	createMock(t)

	// expect
	clientDoFuncExpected = 3
	clientDoFunc = func(client *http.Client, request *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		assert.Equal(t, dummyClient, client)
		assert.Equal(t, dummyRequestObject, request)
		return dummyResponseObject, nil
	}
	delayForRetryFuncExpected = 2
	delayForRetryFunc = func() {
		delayForRetryFuncCalled++
	}

	// SUT + act
	var result, err = clientDoWithRetry(
		dummyClient,
		dummyRequestObject,
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestCustomizeRoundTripper_NoCustomization(t *testing.T) {
	// arrange
	var dummyOriginal = http.DefaultTransport

	// mock
	createMock(t)

	// SUT + act
	var result = customizeRoundTripper(
		dummyOriginal,
	)

	// assert
	assert.Equal(t, dummyOriginal, result)

	// verify
	verifyAll(t)
}

func TestCustomizeRoundTripper_WithCustomization(t *testing.T) {
	// arrange
	var dummyOriginal = http.DefaultTransport
	var dummyCustomized = &http.Transport{}

	// mock
	createMock(t)

	// expect
	customizationHTTPRoundTripperExpected = 1
	customization.HTTPRoundTripper = func(originalTransport http.RoundTripper) http.RoundTripper {
		customizationHTTPRoundTripperCalled++
		assert.Equal(t, dummyOriginal, originalTransport)
		return dummyCustomized
	}

	// SUT + act
	var result = customizeRoundTripper(
		dummyOriginal,
	)

	// assert
	assert.Equal(t, dummyCustomized, result)

	// verify
	verifyAll(t)
}

func TestGetHTTPTransport_NoSendClientCert(t *testing.T) {
	// arrange
	var dummySendClientCert = false
	var dummyRoundTripper = &http.Transport{}

	// mock
	createMock(t)

	// expect
	customizeRoundTripperFuncExpected = 1
	customizeRoundTripperFunc = func(original http.RoundTripper) http.RoundTripper {
		customizeRoundTripperFuncCalled++
		assert.Equal(t, http.DefaultTransport, original)
		return dummyRoundTripper
	}

	// SUT + act
	var result = getHTTPTransport(
		dummySendClientCert,
	)

	// assert
	assert.Equal(t, dummyRoundTripper, result)

	// verify
	verifyAll(t)
}

func TestGetHTTPTransport_SendClientCert_NoCertFound(t *testing.T) {
	// arrange
	var dummySendClientCert = true
	var dummyClientCert *tls.Certificate
	var dummyRoundTripper = &http.Transport{}

	// mock
	createMock(t)

	// expect
	certificateGetClientCertificateExpected = 1
	certificateGetClientCertificate = func() *tls.Certificate {
		certificateGetClientCertificateCalled++
		return dummyClientCert
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "network", category)
		assert.Equal(t, "getHTTPTransport", subcategory)
		assert.Equal(t, "Failed to load client certificate for mTLS communications; fallback to default HTTP transport", messageFormat)
		assert.Empty(t, parameters)
	}
	customizeRoundTripperFuncExpected = 1
	customizeRoundTripperFunc = func(original http.RoundTripper) http.RoundTripper {
		customizeRoundTripperFuncCalled++
		assert.Equal(t, http.DefaultTransport, original)
		return dummyRoundTripper
	}

	// SUT + act
	var result = getHTTPTransport(
		dummySendClientCert,
	)

	// assert
	assert.Equal(t, dummyRoundTripper, result)

	// verify
	verifyAll(t)
}

func TestGetHTTPTransport_SendClientCert_CertFound(t *testing.T) {
	// arrange
	var dummySendClientCert = true
	var dummyClientCert = &tls.Certificate{}
	var dummyRoundTripper = &http.Transport{}

	// mock
	createMock(t)

	// expect
	certificateGetClientCertificateExpected = 1
	certificateGetClientCertificate = func() *tls.Certificate {
		certificateGetClientCertificateCalled++
		return dummyClientCert
	}
	customizeRoundTripperFuncExpected = 1
	customizeRoundTripperFunc = func(original http.RoundTripper) http.RoundTripper {
		customizeRoundTripperFuncCalled++
		assert.NotEqual(t, http.DefaultTransport, original)
		return dummyRoundTripper
	}

	// SUT + act
	var result = getHTTPTransport(
		dummySendClientCert,
	)

	// assert
	assert.Equal(t, dummyRoundTripper, result)

	// verify
	verifyAll(t)
}

func TestInitialize(t *testing.T) {
	// arrange
	var dummyNetworkTimeout = time.Duration(rand.Int())
	var dummyHTTPTransport1 = &http.Transport{MaxConnsPerHost: rand.Int()}
	var dummyHTTPTransport2 = &http.Transport{MaxConnsPerHost: rand.Int()}

	// mock
	createMock(t)

	// expect
	getHTTPTransportFuncExpected = 2
	getHTTPTransportFunc = func(sendClientCert bool) http.RoundTripper {
		getHTTPTransportFuncCalled++
		if getHTTPTransportFuncCalled == 1 {
			assert.True(t, sendClientCert)
			return dummyHTTPTransport1
		} else if getHTTPTransportFuncCalled == 2 {
			assert.False(t, sendClientCert)
			return dummyHTTPTransport2
		}
		return nil
	}

	// SUT + act
	Initialize(
		dummyNetworkTimeout,
	)

	// assert
	assert.NotNil(t, httpClientWithCert)
	assert.Equal(t, dummyHTTPTransport1, httpClientWithCert.Transport)
	assert.Equal(t, dummyNetworkTimeout, httpClientWithCert.Timeout)
	assert.NotNil(t, httpClientNoCert)
	assert.Equal(t, dummyHTTPTransport2, httpClientNoCert.Transport)
	assert.Equal(t, dummyNetworkTimeout, httpClientNoCert.Timeout)

	// verify
	verifyAll(t)
}

func TestNewNetworkRequest(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummyHeader = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummySendClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// SUT
	var result = NewNetworkRequest(
		dummySessionObject,
		dummyMethod,
		dummyURL,
		dummyPayload,
		dummyHeader,
		dummySendClientCert,
	)

	// act
	var typedResult, ok = result.(*networkRequest)

	// assert
	assert.True(t, ok)
	assert.NotNil(t, typedResult)
	assert.Equal(t, dummySessionObject, typedResult.session)
	assert.Equal(t, dummyMethod, typedResult.method)
	assert.Equal(t, dummyURL, typedResult.url)
	assert.Equal(t, dummyPayload, typedResult.payload)
	assert.Equal(t, dummyHeader, typedResult.header)
	assert.Equal(t, dummySendClientCert, typedResult.sendClientCert)

	// verify
	verifyAll(t)
}

func TestNetworkRequestEnableRetry(t *testing.T) {
	// arrange
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{
		rand.Int(): rand.Int(),
		rand.Int(): rand.Int(),
	}

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// act
	sut.EnableRetry(
		dummyConnRetry,
		dummyHTTPRetry,
	)

	// assert
	assert.Equal(t, dummyConnRetry, sut.connRetry)
	assert.Equal(t, dummyHTTPRetry, sut.httpRetry)

	// verify
	verifyAll(t)
}

func TestCustomizeHTTPRequest_NoCustomization(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHTTPRequest = &http.Request{
		RequestURI: "foo",
	}

	// mock
	createMock(t)

	// SUT + act
	var result = customizeHTTPRequest(
		dummySessionObject,
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyHTTPRequest, result)

	// verify
	verifyAll(t)
}

func TestCustomizeHTTPRequest_WithCustomization(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHTTPRequest = &http.Request{
		RequestURI: "foo",
	}
	var dummyCustomized = &http.Request{
		RequestURI: "bar",
	}

	// mock
	createMock(t)

	// expect
	customizationWrapHTTPRequestExpected = 1
	customization.WrapHTTPRequest = func(session sessionModel.Session, httpRequest *http.Request) *http.Request {
		customizationWrapHTTPRequestCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyCustomized
	}

	// SUT + act
	var result = customizeHTTPRequest(
		dummySessionObject,
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyCustomized, result)

	// verify
	verifyAll(t)
}

func TestCreateHTTPRequest_RequestError(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummyHeader = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{
		rand.Int(): rand.Int(),
		rand.Int(): rand.Int(),
	}
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyNetworkRequest = &networkRequest{
		dummySessionObject,
		dummyMethod,
		dummyURL,
		dummyPayload,
		dummyHeader,
		dummyConnRetry,
		dummyHTTPRetry,
		dummySendClientCert,
	}
	var dummyRequest *http.Request
	var dummyError = errors.New("some error message")
	var expectedErrorMessage = "Failed to generate request to [%v]"
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	stringsNewReaderExpected = 1
	stringsNewReader = func(s string) *strings.Reader {
		stringsNewReaderCalled++
		return strings.NewReader(s)
	}
	httpNewRequestExpected = 1
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		httpNewRequestCalled++
		assert.Equal(t, dummyMethod, method)
		assert.Equal(t, dummyURL, url)
		assert.NotNil(t, body)
		return dummyRequest, dummyError
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyURL, parameters[0])
		return dummyAppError
	}

	// SUT + act
	var result, err = createHTTPRequest(
		dummyNetworkRequest,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestCreateHTTPRequest_Success(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummyHeader = map[string]string{
		"foo":  "bar",
		"test": "123",
	}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{
		rand.Int(): rand.Int(),
		rand.Int(): rand.Int(),
	}
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyNetworkRequest = &networkRequest{
		dummySessionObject,
		dummyMethod,
		dummyURL,
		dummyPayload,
		dummyHeader,
		dummyConnRetry,
		dummyHTTPRetry,
		dummySendClientCert,
	}
	var dummyRequest = &http.Request{
		RequestURI: "abc",
	}
	var dummyCustomized = &http.Request{
		RequestURI: "xyz",
	}

	// mock
	createMock(t)

	// expect
	stringsNewReaderExpected = 1
	stringsNewReader = func(s string) *strings.Reader {
		stringsNewReaderCalled++
		return strings.NewReader(s)
	}
	httpNewRequestExpected = 1
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		httpNewRequestCalled++
		assert.Equal(t, dummyMethod, method)
		assert.Equal(t, dummyURL, url)
		assert.NotNil(t, body)
		return dummyRequest, nil
	}
	loggerNetworkCallExpected = 1
	loggerNetworkCall = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkCallCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyMethod, category)
		assert.Equal(t, dummyURL, messageFormat)
		assert.Zero(t, subcategory)
		assert.Empty(t, parameters)
	}
	loggerNetworkRequestExpected = 1
	loggerNetworkRequest = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkRequestCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Payload", category)
		assert.Zero(t, subcategory)
		assert.Equal(t, dummyPayload, messageFormat)
		assert.Empty(t, parameters)
	}
	headerutilLogHTTPHeaderExpected = 1
	headerutilLogHTTPHeader = func(session sessionModel.Session, header http.Header) {
		headerutilLogHTTPHeaderCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHeader["foo"], header["Foo"][0])
		assert.Equal(t, dummyHeader["test"], header["Test"][0])
	}
	customizeHTTPRequestFuncExpected = 1
	customizeHTTPRequestFunc = func(session sessionModel.Session, httpRequest *http.Request) *http.Request {
		customizeHTTPRequestFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyRequest, httpRequest)
		return dummyCustomized
	}

	// SUT + act
	var result, err = createHTTPRequest(
		dummyNetworkRequest,
	)

	// assert
	assert.Equal(t, dummyCustomized, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestLogErrorResponse(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	loggerNetworkResponseExpected = 1
	loggerNetworkResponse = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkResponseCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Message", category)
		assert.Zero(t, subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}
	loggerNetworkFinishExpected = 1
	loggerNetworkFinish = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkFinishCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Error", category)
		assert.Zero(t, subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	logErrorResponse(
		dummySessionObject,
		dummyError,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPResponse_NilResponse(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyResponse *http.Response

	// mock
	createMock(t)

	// SUT + act
	logHTTPResponse(
		dummySessionObject,
		dummyResponse,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPResponse_ValidResponse(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyStatus = "some status"
	var dummyStatusCode = rand.Intn(1000)
	var dummyBody = ioutil.NopCloser(bytes.NewBufferString("some body"))
	var dummyHeader = http.Header{
		"foo":  []string{"bar"},
		"test": []string{"123", "456", "789"},
	}
	var dummyResponse = &http.Response{
		StatusCode: dummyStatusCode,
		Body:       dummyBody,
		Header:     dummyHeader,
	}
	var dummyResponseBytes = []byte("some response bytes")
	var dummyResponseBody = string(dummyResponseBytes)
	var dummyError = errors.New("some error")
	var dummyBuffer = &bytes.Buffer{}
	var dummyNewBody = ioutil.NopCloser(bytes.NewBufferString("some new body"))

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummyBody, r)
		return dummyResponseBytes, dummyError
	}
	bytesNewBufferExpected = 1
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		assert.Equal(t, dummyResponseBytes, buf)
		return dummyBuffer
	}
	ioutilNopCloserExpected = 1
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		assert.Equal(t, dummyBuffer, r)
		return dummyNewBody
	}
	httpStatusTextExpected = 1
	httpStatusText = func(code int) string {
		httpStatusTextCalled++
		assert.Equal(t, dummyStatusCode, code)
		return dummyStatus
	}
	strconvItoaExpected = 1
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		assert.Equal(t, dummyStatusCode, i)
		return strconv.Itoa(i)
	}
	loggerNetworkResponseExpected = 1
	loggerNetworkResponse = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkResponseCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Body", category)
		assert.Zero(t, subcategory)
		assert.Equal(t, dummyResponseBody, messageFormat)
		assert.Empty(t, parameters)
	}
	headerutilLogHTTPHeaderExpected = 1
	headerutilLogHTTPHeader = func(session sessionModel.Session, header http.Header) {
		headerutilLogHTTPHeaderCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHeader, header)
	}
	loggerNetworkFinishExpected = 1
	loggerNetworkFinish = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerNetworkFinishCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyStatus, category)
		assert.Equal(t, strconv.Itoa(dummyStatusCode), subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	logHTTPResponse(
		dummySessionObject,
		dummyResponse,
	)

	// assert
	assert.Equal(t, dummyNewBody, dummyResponse.Body)

	// verify
	verifyAll(t)
}

func TestDoRequestProcessing_RequestError(t *testing.T) {
	// arrange
	var dummyNetworkRequest = &networkRequest{}
	var dummyRequestObject *http.Request
	var dummyRequestError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	createHTTPRequestFuncExpected = 1
	createHTTPRequestFunc = func(networkRequest *networkRequest) (*http.Request, error) {
		createHTTPRequestFuncCalled++
		assert.Equal(t, dummyNetworkRequest, networkRequest)
		return dummyRequestObject, dummyRequestError
	}

	// SUT + act
	var result, err = doRequestProcessing(
		dummyNetworkRequest,
	)

	// assert
	assert.Nil(t, result)
	assert.Equal(t, dummyRequestError, err)

	// verify
	verifyAll(t)
}

func TestDoRequestProcessing_ResponseError(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{
		rand.Int(): rand.Int(),
		rand.Int(): rand.Int(),
	}
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyNetworkRequest = &networkRequest{
		session:        dummySessionObject,
		connRetry:      dummyConnRetry,
		httpRetry:      dummyHTTPRetry,
		sendClientCert: dummySendClientCert,
	}
	var dummyHTTPClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyResponseObject *http.Response
	var dummyResponseError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	createHTTPRequestFuncExpected = 1
	createHTTPRequestFunc = func(networkRequest *networkRequest) (*http.Request, error) {
		createHTTPRequestFuncCalled++
		assert.Equal(t, dummyNetworkRequest, networkRequest)
		return dummyRequestObject, nil
	}
	getClientForRequestFuncExpected = 1
	getClientForRequestFunc = func(sendClientCert bool) *http.Client {
		getClientForRequestFuncCalled++
		assert.Equal(t, dummySendClientCert, sendClientCert)
		return dummyHTTPClient
	}
	clientDoWithRetryFuncExpected = 1
	clientDoWithRetryFunc = func(client *http.Client, request *http.Request, connRetry int, httpRetry map[int]int) (*http.Response, error) {
		clientDoWithRetryFuncCalled++
		assert.Equal(t, dummyHTTPClient, client)
		assert.Equal(t, dummyRequestObject, request)
		assert.Equal(t, dummyConnRetry, connRetry)
		assert.Equal(t, dummyHTTPRetry, httpRetry)
		return dummyResponseObject, dummyResponseError
	}
	logErrorResponseFuncExpected = 1
	logErrorResponseFunc = func(session sessionModel.Session, responseError error) {
		logErrorResponseFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyResponseError, responseError)
	}

	// SUT + act
	var result, err = doRequestProcessing(
		dummyNetworkRequest,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestDoRequestProcessing_ResponseSuccess(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyConnRetry = rand.Int()
	var dummyHTTPRetry = map[int]int{
		rand.Int(): rand.Int(),
		rand.Int(): rand.Int(),
	}
	var dummySendClientCert = rand.Intn(100) < 50
	var dummyNetworkRequest = &networkRequest{
		session:        dummySessionObject,
		connRetry:      dummyConnRetry,
		httpRetry:      dummyHTTPRetry,
		sendClientCert: dummySendClientCert,
	}
	var dummyHTTPClient = &http.Client{}
	var dummyRequestObject = &http.Request{}
	var dummyResponseObject = &http.Response{}

	// mock
	createMock(t)

	// expect
	createHTTPRequestFuncExpected = 1
	createHTTPRequestFunc = func(networkRequest *networkRequest) (*http.Request, error) {
		createHTTPRequestFuncCalled++
		assert.Equal(t, dummyNetworkRequest, networkRequest)
		return dummyRequestObject, nil
	}
	getClientForRequestFuncExpected = 1
	getClientForRequestFunc = func(sendClientCert bool) *http.Client {
		getClientForRequestFuncCalled++
		assert.Equal(t, dummySendClientCert, sendClientCert)
		return dummyHTTPClient
	}
	clientDoWithRetryFuncExpected = 1
	clientDoWithRetryFunc = func(client *http.Client, request *http.Request, connRetry int, httpRetry map[int]int) (*http.Response, error) {
		clientDoWithRetryFuncCalled++
		assert.Equal(t, dummyHTTPClient, client)
		assert.Equal(t, dummyRequestObject, request)
		assert.Equal(t, dummyConnRetry, connRetry)
		assert.Equal(t, dummyHTTPRetry, httpRetry)
		return dummyResponseObject, nil
	}
	logHTTPResponseFuncExpected = 1
	logHTTPResponseFunc = func(session sessionModel.Session, response *http.Response) {
		logHTTPResponseFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyResponseObject, response)
	}

	// SUT + act
	var result, err = doRequestProcessing(
		dummyNetworkRequest,
	)

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestNetworkRequestProcessRaw(t *testing.T) {
	// arrange
	var dummyResponseObject = &http.Response{}
	var dummyResponseError = errors.New("some error")

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// expect
	doRequestProcessingFuncExpected = 1
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		assert.Equal(t, sut, networkRequest)
		return dummyResponseObject, dummyResponseError
	}

	// act
	var result, err = sut.ProcessRaw()

	// assert
	assert.Equal(t, dummyResponseObject, result)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestParseResponse_ReadError(t *testing.T) {
	// arrange
	var dummyBody = ioutil.NopCloser(bytes.NewBufferString("some body"))
	var dummyBytes = []byte("some bytes")
	var dummyError = errors.New("some error")
	var dummyDataTemplate string

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummyBody, r)
		return dummyBytes, dummyError
	}

	// SUT + act
	var err = parseResponse(
		dummyBody,
		&dummyDataTemplate,
	)

	// assert
	assert.Zero(t, dummyDataTemplate)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestParseResponse_JSONError(t *testing.T) {
	// arrange
	var dummyBody = ioutil.NopCloser(bytes.NewBufferString("some body"))
	var dummyBytes = []byte("some bytes")
	var dummyError = errors.New("some error")
	var dummyDataTemplate string

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummyBody, r)
		return dummyBytes, nil
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, string(dummyBytes), value)
		return dummyError
	}

	// SUT + act
	var err = parseResponse(
		dummyBody,
		&dummyDataTemplate,
	)

	// assert
	assert.Zero(t, dummyDataTemplate)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestParseResponse_HappyPath(t *testing.T) {
	// arrange
	var dummyBody = ioutil.NopCloser(bytes.NewBufferString("some body"))
	var dummyData = "some data"
	var dummyBytes = []byte("\"" + dummyData + "\"")
	var dummyDataTemplate string

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummyBody, r)
		return dummyBytes, nil
	}
	jsonutilTryUnmarshalExpected = 1
	jsonutilTryUnmarshal = func(value string, dataTemplate interface{}) error {
		jsonutilTryUnmarshalCalled++
		assert.Equal(t, string(dummyBytes), value)
		(*(dataTemplate).(*string)) = dummyData
		return nil
	}

	// SUT + act
	var err = parseResponse(
		dummyBody,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyData, dummyDataTemplate)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestNetworkRequestProcess_Error_NilObject(t *testing.T) {
	// arrange
	var dummyResponseObject *http.Response
	var dummyResponseError = errors.New("some error")
	var dummyDataTemplate string

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// expect
	doRequestProcessingFuncExpected = 1
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		assert.Equal(t, sut, networkRequest)
		return dummyResponseObject, dummyResponseError
	}

	// act
	var result, header, err = sut.Process(
		&dummyDataTemplate,
	)

	// assert
	assert.Zero(t, dummyDataTemplate)
	assert.Equal(t, http.StatusInternalServerError, result)
	assert.Empty(t, header)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestNetworkRequestProcess_Error_ValidObject(t *testing.T) {
	// arrange
	var dummyStatusCode = rand.Int()
	var dummyHeader = map[string][]string{
		"foo":  []string{"bar"},
		"test": []string{"123", "456", "789"},
	}
	var dummyResponseObject = &http.Response{
		StatusCode: dummyStatusCode,
		Header:     dummyHeader,
	}
	var dummyResponseError = errors.New("some error")
	var dummyDataTemplate string

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// expect
	doRequestProcessingFuncExpected = 1
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		assert.Equal(t, sut, networkRequest)
		return dummyResponseObject, dummyResponseError
	}

	// act
	var result, header, err = sut.Process(
		&dummyDataTemplate,
	)

	// assert
	assert.Zero(t, dummyDataTemplate)
	assert.Equal(t, dummyStatusCode, result)
	assert.Equal(t, http.Header(dummyHeader), header)
	assert.Equal(t, dummyResponseError, err)

	// verify
	verifyAll(t)
}

func TestNetworkRequestProcess_Success_NilObject(t *testing.T) {
	// arrange
	var dummyResponseObject *http.Response
	var dummyResponseError error
	var dummyDataTemplate string

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// expect
	doRequestProcessingFuncExpected = 1
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		assert.Equal(t, sut, networkRequest)
		return dummyResponseObject, dummyResponseError
	}

	// act
	var result, header, err = sut.Process(
		&dummyDataTemplate,
	)

	// assert
	assert.Zero(t, dummyDataTemplate)
	assert.Zero(t, result)
	assert.Empty(t, header)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestNetworkRequestProcess_Success_ValidObject(t *testing.T) {
	// arrange
	var dummyStatusCode = rand.Int()
	var dummyHeader = map[string][]string{
		"foo":  []string{"bar"},
		"test": []string{"123", "456", "789"},
	}
	var dummyBody = ioutil.NopCloser(bytes.NewBufferString("some body"))
	var dummyResponseObject = &http.Response{
		StatusCode: dummyStatusCode,
		Header:     dummyHeader,
		Body:       dummyBody,
	}
	var dummyResponseError error
	var dummyParseError = errors.New("some parse error")
	var dummyDataTemplate string
	var dummyData = "some data"

	// SUT
	var sut = &networkRequest{}

	// mock
	createMock(t)

	// expect
	doRequestProcessingFuncExpected = 1
	doRequestProcessingFunc = func(networkRequest *networkRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		assert.Equal(t, sut, networkRequest)
		return dummyResponseObject, dummyResponseError
	}
	parseResponseFuncExpected = 1
	parseResponseFunc = func(body io.ReadCloser, dataTemplate interface{}) error {
		parseResponseFuncCalled++
		assert.Equal(t, dummyBody, body)
		(*(dataTemplate).(*string)) = dummyData
		return dummyParseError
	}

	// act
	var result, header, err = sut.Process(
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyData, dummyDataTemplate)
	assert.Equal(t, dummyStatusCode, result)
	assert.Equal(t, http.Header(dummyHeader), header)
	assert.Equal(t, dummyParseError, err)

	// verify
	verifyAll(t)
}
