package request

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody_NilRequest(t *testing.T) {
	// arrange
	var dummySessionID *http.Request

	// mock
	createMock(t)

	// SUT + act
	var result = GetRequestBody(
		dummySessionID,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_NilBody(t *testing.T) {
	// arrange
	var dummySessionID = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}

	// mock
	createMock(t)

	// SUT + act
	var result = GetRequestBody(
		dummySessionID,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_ErrorBody(t *testing.T) {
	// arrange
	var bodyContent = "some body content"
	var dummySessionID = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
		Body:       ioutil.NopCloser(strings.NewReader(bodyContent)),
	}
	var dummyError = errors.New("some error message")

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		assert.Equal(t, dummySessionID.Body, r)
		return nil, dummyError
	}

	// SUT + act
	var result = GetRequestBody(
		dummySessionID,
	)

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestGetRequestBody_Success(t *testing.T) {
	// arrange
	var bodyContent = "some body content"
	var dummySessionID = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
		Body:       ioutil.NopCloser(strings.NewReader(bodyContent)),
	}
	var dummyBuffer = &bytes.Buffer{}
	var dummyReadCloser = ioutil.NopCloser(nil)

	// mock
	createMock(t)

	// expect
	ioutilReadAllExpected = 1
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		return ioutil.ReadAll(r)
	}
	bytesNewBufferExpected = 1
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		assert.Equal(t, []byte(bodyContent), buf)
		return dummyBuffer
	}
	ioutilNopCloserExpected = 1
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		assert.Equal(t, dummyBuffer, r)
		return dummyReadCloser
	}

	// SUT + act
	var result = GetRequestBody(
		dummySessionID,
	)

	// assert
	assert.Equal(t, bodyContent, result)
	assert.Equal(t, dummyReadCloser, dummySessionID.Body)

	// verify
	verifyAll(t)
}

func TestFullDump_DumpError(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{}
	var dummyRequestBytes = []byte("some request bytes")
	var dummyDumpError = errors.New("some dump error")
	var dummyFormat = "FullDump Failed: %v\r\nSimpleDump: %v\r\n"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	httputilDumpRequestExpected = 1
	httputilDumpRequest = func(req *http.Request, body bool) ([]byte, error) {
		httputilDumpRequestCalled++
		assert.Equal(t, dummyHTTPRequest, req)
		assert.True(t, body)
		return dummyRequestBytes, dummyDumpError
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, dummyFormat, format)
		assert.Equal(t, 2, len(a))
		assert.Equal(t, dummyDumpError, a[0])
		assert.Equal(t, dummyHTTPRequest, a[1])
		return dummyResult
	}

	// SUT + act
	var result = FullDump(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}

func TestFullDump_Success(t *testing.T) {
	// arrange
	var dummyRemoteAddress = "some remote address"
	var dummyHTTPRequest = &http.Request{
		RemoteAddr: dummyRemoteAddress,
	}
	var dummyRequestBytes = []byte("some request bytes")
	var dummyFormat = "%vRemote Address: %v\r\n"
	var dummyResult = "some result"

	// mock
	createMock(t)

	// expect
	httputilDumpRequestExpected = 1
	httputilDumpRequest = func(req *http.Request, body bool) ([]byte, error) {
		httputilDumpRequestCalled++
		assert.Equal(t, dummyHTTPRequest, req)
		assert.True(t, body)
		return dummyRequestBytes, nil
	}
	fmtSprintfExpected = 1
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		assert.Equal(t, dummyFormat, format)
		assert.Equal(t, 2, len(a))
		assert.Equal(t, string(dummyRequestBytes), a[0])
		assert.Equal(t, dummyRemoteAddress, a[1])
		return dummyResult
	}

	// SUT + act
	var result = FullDump(
		dummyHTTPRequest,
	)

	// assert
	assert.Equal(t, dummyResult, result)

	// verify
	verifyAll(t)
}
