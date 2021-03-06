package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP_MethodNotAllowed(t *testing.T) {
	// arrange
	var dummyHTTPRequest = &http.Request{
		Method:     http.MethodGet,
		RequestURI: "http://localhost/",
		Header:     map[string][]string{},
	}
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyRequestString = "some request string"

	// mock
	createMock(t)

	// expect
	requestFullDumpExpected = 1
	requestFullDump = func(httpRequest *http.Request) string {
		requestFullDumpCalled++
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		return dummyRequestString
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "RouteError", category)
		assert.Equal(t, "MethodNotAllowed", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyRequestString, parameters[0])
	}
	httpErrorExpected = 1
	httpError = func(w http.ResponseWriter, error string, code int) {
		httpErrorCalled++
		assert.Equal(t, dummyResponseWriter, w)
		assert.Equal(t, "405 - resource URI action not allowed", error)
		assert.Equal(t, http.StatusMethodNotAllowed, code)
	}

	// SUT
	var sut = &MethodNotAllowedHandler{}

	// act
	sut.ServeHTTP(
		dummyResponseWriter,
		dummyHTTPRequest,
	)

	// verify
	verifyAll(t)
}
