package swagger

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
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
		assert.Equal(t, "/docs/", url)
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
	var dummyRouter = &mux.Router{}
	var dummyHandler = &dummyHandlerStruct{}

	// mock
	createMock(t)

	// expect
	routeHandleFuncExpected = 1
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		assert.Equal(t, dummyRouter, router)
		assert.Equal(t, "SwaggerUI", endpoint)
		assert.Equal(t, http.MethodGet, method)
		assert.Equal(t, "/docs", path)
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(redirectHandlerFunc))
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
		return nil
	}
	contentHandlerFuncExpected = 1
	contentHandlerFunc = func() http.Handler {
		contentHandlerFuncCalled++
		return dummyHandler
	}
	routeHostStaticExpected = 1
	routeHostStatic = func(router *mux.Router, name string, path string, handler http.Handler) *mux.Route {
		routeHostStaticCalled++
		assert.Equal(t, dummyRouter, router)
		assert.Equal(t, "SwaggerUI", name)
		assert.Equal(t, "/docs/", path)
		assert.Equal(t, dummyHandler, handler)
		return nil
	}

	// SUT + act
	HostEntry(
		dummyRouter,
	)

	// verify
	verifyAll(t)
}
