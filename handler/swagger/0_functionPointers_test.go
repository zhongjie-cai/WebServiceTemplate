package swagger

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

var (
	configAppPathExpected       int
	configAppPathCalled         int
	httpRedirectExpected        int
	httpRedirectCalled          int
	httpStripPrefixExpected     int
	httpStripPrefixCalled       int
	httpFileServerExpected      int
	httpFileServerCalled        int
	routeHandleFuncExpected     int
	routeHandleFuncCalled       int
	routeHostStaticExpected     int
	routeHostStaticCalled       int
	redirectHandlerFuncExpected int
	redirectHandlerFuncCalled   int
	contentHandlerFuncExpected  int
	contentHandlerFuncCalled    int
)

func createMock(t *testing.T) {
	configAppPathExpected = 0
	configAppPathCalled = 0
	configAppPath = func() string {
		configAppPathCalled++
		return ""
	}
	httpRedirectExpected = 0
	httpRedirectCalled = 0
	httpRedirect = func(responseWriter http.ResponseWriter, request *http.Request, url string, code int) {
		httpRedirectCalled++
	}
	httpStripPrefixExpected = 0
	httpStripPrefixCalled = 0
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		return nil
	}
	httpFileServerExpected = 0
	httpFileServerCalled = 0
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		return nil
	}
	routeHandleFuncExpected = 0
	routeHandleFuncCalled = 0
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		return nil
	}
	routeHostStaticExpected = 0
	routeHostStaticCalled = 0
	routeHostStatic = func(router *mux.Router, name string, path string, handler http.Handler) *mux.Route {
		routeHostStaticCalled++
		return nil
	}
	redirectHandlerFuncExpected = 0
	redirectHandlerFuncCalled = 0
	redirectHandlerFunc = func(responseWriter http.ResponseWriter, request *http.Request) {
		redirectHandlerFuncCalled++
	}
	contentHandlerFuncExpected = 0
	contentHandlerFuncCalled = 0
	contentHandlerFunc = func() http.Handler {
		contentHandlerFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	configAppPath = config.AppPath
	assert.Equal(t, configAppPathExpected, configAppPathCalled, "Unexpected method call to configAppPath")
	httpRedirect = http.Redirect
	assert.Equal(t, httpRedirectExpected, httpRedirectCalled, "Unexpected method call to httpRedirect")
	httpStripPrefix = http.StripPrefix
	assert.Equal(t, httpStripPrefixExpected, httpStripPrefixCalled, "Unexpected method call to httpStripPrefix")
	httpFileServer = http.FileServer
	assert.Equal(t, httpFileServerExpected, httpFileServerCalled, "Unexpected method call to httpFileServer")
	routeHandleFunc = route.HandleFunc
	assert.Equal(t, routeHandleFuncExpected, routeHandleFuncCalled, "Unexpected method call to routeHandleFunc")
	routeHostStatic = route.HostStatic
	assert.Equal(t, routeHostStaticExpected, routeHostStaticCalled, "Unexpected method call to routeHostStatic")
	redirectHandlerFunc = redirectHandler
	assert.Equal(t, redirectHandlerFuncExpected, redirectHandlerFuncCalled, "Unexpected method call to redirectHandlerFunc")
	contentHandlerFunc = contentHandler
	assert.Equal(t, contentHandlerFuncExpected, contentHandlerFuncCalled, "Unexpected method call to contentHandlerFunc")
}
