package swagger

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock struct
type dummyHandlerStruct struct {
}

func (dhs *dummyHandlerStruct) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
}

type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Equal(drw.t, http.StatusMethodNotAllowed, statusCode)
}

func TestRedirectHandler(t *testing.T) {
	// arrange
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}

	// mock
	createMock(t)

	// expect
	httpRedirectExpected = 1
	httpRedirect = func(responseWriter http.ResponseWriter, request *http.Request, url string, code int) {
		httpRedirectCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyRequest, request)
		assert.Equal(t, "/docs/index.html", url)
		assert.Equal(t, http.StatusPermanentRedirect, code)
	}

	// SUT + act
	redirectHandler(
		dummyResponseWriter,
		dummyRequest,
	)

	// verify
	verifyAll(t)
}

func TestContentHandler(t *testing.T) {
	// arrange
	var dummyAppPath = "some app path"
	var dummyFileHandler = &dummyHandlerStruct{}
	var dummyForwardedHandler = &dummyHandlerStruct{}

	// mock
	createMock(t)

	// expect
	configAppPathExpected = 1
	configAppPath = func() string {
		configAppPathCalled++
		return dummyAppPath
	}
	httpFileServerExpected = 1
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		assert.Equal(t, http.Dir(dummyAppPath+"/docs"), root)
		return dummyFileHandler
	}
	httpStripPrefixExpected = 1
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		assert.Equal(t, "/docs/", prefix)
		assert.Equal(t, dummyFileHandler, h)
		return dummyForwardedHandler
	}

	// SUT + act
	var result = contentHandler()

	// assert
	assert.Equal(t, dummyForwardedHandler, result)

	// verify
	verifyAll(t)
}

func TestHostEntry(t *testing.T) {
	// arrange
	var dummyHandler = &dummyHandlerStruct{}

	// mock
	createMock(t)

	// expect
	httpHandleFuncExpected = 1
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
		assert.Equal(t, "/docs", pattern)
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(redirectHandlerFunc))
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
	}
	contentHandlerFuncExpected = 1
	contentHandlerFunc = func() http.Handler {
		contentHandlerFuncCalled++
		return dummyHandler
	}
	httpHandleExpected = 1
	httpHandle = func(pattern string, handler http.Handler) {
		httpHandleCalled++
		assert.Equal(t, "/docs/", pattern)
		assert.Equal(t, dummyHandler, handler)
	}

	// SUT + act
	HostEntry()

	// verify
	verifyAll(t)
}
