package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP_NotFound(t *testing.T) {
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
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyHTTPRequest, v)
		return dummyRequestString
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "RouteError", category)
		assert.Equal(t, "NotFound", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyRequestString, parameters[0])
	}
	httpErrorExpected = 1
	httpError = func(w http.ResponseWriter, error string, code int) {
		httpErrorCalled++
		assert.Equal(t, dummyResponseWriter, w)
		assert.Equal(t, "404 - resource URI not found", error)
		assert.Equal(t, http.StatusNotFound, code)
	}

	// SUT
	var sut = &NotFoundHandler{}

	// act
	sut.ServeHTTP(
		dummyResponseWriter,
		dummyHTTPRequest,
	)

	// verify
	verifyAll(t)
}
