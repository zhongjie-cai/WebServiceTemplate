package swagger

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
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
	httpHandleFuncExpected      int
	httpHandleFuncCalled        int
	httpHandleExpected          int
	httpHandleCalled            int
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
	httpHandleFuncExpected = 0
	httpHandleFuncCalled = 0
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
	}
	httpHandleExpected = 0
	httpHandleCalled = 0
	httpHandle = func(pattern string, handler http.Handler) {
		httpHandleCalled++
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
	if configAppPathExpected != configAppPathCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppPath, expected %v, actual %v", configAppPathExpected, configAppPathCalled))
	}
	httpRedirect = http.Redirect
	if httpRedirectExpected != httpRedirectCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpRedirect, expected %v, actual %v", httpRedirectExpected, httpRedirectCalled))
	}
	httpStripPrefix = http.StripPrefix
	if httpStripPrefixExpected != httpStripPrefixCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpStripPrefix, expected %v, actual %v", httpStripPrefixExpected, httpStripPrefixCalled))
	}
	httpFileServer = http.FileServer
	if httpFileServerExpected != httpFileServerCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpFileServer, expected %v, actual %v", httpFileServerExpected, httpFileServerCalled))
	}
	httpHandleFunc = http.HandleFunc
	if httpHandleFuncExpected != httpHandleFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpHandleFunc, expected %v, actual %v", httpHandleFuncExpected, httpHandleFuncCalled))
	}
	httpHandle = http.Handle
	if httpHandleExpected != httpHandleCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpHandle, expected %v, actual %v", httpHandleExpected, httpHandleCalled))
	}
	redirectHandlerFunc = redirectHandler
	if redirectHandlerFuncExpected != redirectHandlerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to redirectHandlerFunc, expected %v, actual %v", redirectHandlerFuncExpected, redirectHandlerFuncCalled))
	}
	contentHandlerFunc = contentHandler
	if contentHandlerFuncExpected != contentHandlerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to contentHandlerFunc, expected %v, actual %v", contentHandlerFuncExpected, contentHandlerFuncCalled))
	}
}
