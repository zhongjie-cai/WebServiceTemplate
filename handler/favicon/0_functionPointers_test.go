package favicon

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

var (
	routeHandleFuncExpected       int
	routeHandleFuncCalled         int
	httpServeFileExpected         int
	httpServeFileCalled           int
	configAppPathExpected         int
	configAppPathCalled           int
	commonHandleInSessionExpected int
	commonHandleInSessionCalled   int
	handleGetFaviconFuncExpected  int
	handleGetFaviconFuncCalled    int
)

func createMock(t *testing.T) {
	routeHandleFuncExpected = 0
	routeHandleFuncCalled = 0
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		return nil
	}
	httpServeFileExpected = 0
	httpServeFileCalled = 0
	httpServeFile = func(responseWriter http.ResponseWriter, request *http.Request, name string) {
		httpServeFileCalled++
	}
	configAppPathExpected = 0
	configAppPathCalled = 0
	configAppPath = func() string {
		configAppPathCalled++
		return ""
	}
	commonHandleInSessionExpected = 0
	commonHandleInSessionCalled = 0
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
	}
	handleGetFaviconFuncExpected = 0
	handleGetFaviconFuncCalled = 0
	handleGetFaviconFunc = func(responseWriter http.ResponseWriter, request *http.Request) {
		handleGetFaviconFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	routeHandleFunc = route.HandleFunc
	assert.Equal(t, routeHandleFuncExpected, routeHandleFuncCalled, "Unexpected method call to routeHandleFunc")
	httpServeFile = http.ServeFile
	assert.Equal(t, httpServeFileExpected, httpServeFileCalled, "Unexpected method call to httpServeFile")
	configAppPath = config.AppPath
	assert.Equal(t, configAppPathExpected, configAppPathCalled, "Unexpected method call to configAppPath")
	commonHandleInSession = common.HandleInSession
	assert.Equal(t, commonHandleInSessionExpected, commonHandleInSessionCalled, "Unexpected method call to commonHandleInSession")
	handleGetFaviconFunc = handleGetFavicon
	assert.Equal(t, handleGetFaviconFuncExpected, handleGetFaviconFuncCalled, "Unexpected method call to handleGetFaviconFunc")
}
