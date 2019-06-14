package health

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

var (
	routeHandleFuncExpected           int
	routeHandleFuncCalled             int
	configAppVersionExpected          int
	configAppVersionCalled            int
	responseOkExpected                int
	responseOkCalled                  int
	commonHandleInSessionExpected     int
	commonHandleInSessionCalled       int
	handleGetHealthFuncExpected       int
	handleGetHealthFuncCalled         int
	handleGetHealthReportFuncExpected int
	handleGetHealthReportFuncCalled   int
)

func createMock(t *testing.T) {
	routeHandleFuncExpected = 0
	routeHandleFuncCalled = 0
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		return nil
	}
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	configAppVersion = func() string {
		configAppVersionCalled++
		return ""
	}
	responseOkExpected = 0
	responseOkCalled = 0
	responseOk = func(sessionID uuid.UUID, responseContent interface{}, responseWriter http.ResponseWriter) {
		responseOkCalled++
	}
	commonHandleInSessionExpected = 0
	commonHandleInSessionCalled = 0
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
	}
	handleGetHealthFuncExpected = 0
	handleGetHealthFuncCalled = 0
	handleGetHealthFunc = func(responseWriter http.ResponseWriter, request *http.Request) {
		handleGetHealthFuncCalled++
	}
	handleGetHealthReportFuncExpected = 0
	handleGetHealthReportFuncCalled = 0
	handleGetHealthReportFunc = func(responseWriter http.ResponseWriter, request *http.Request) {
		handleGetHealthReportFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	routeHandleFunc = route.HandleFunc
	assert.Equal(t, routeHandleFuncExpected, routeHandleFuncCalled, "Unexpected method call to routeHandleFunc")
	configAppVersion = config.AppVersion
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected method call to configAppVersion")
	responseOk = response.Ok
	assert.Equal(t, responseOkExpected, responseOkCalled, "Unexpected method call to responseOk")
	commonHandleInSession = common.HandleInSession
	assert.Equal(t, commonHandleInSessionExpected, commonHandleInSessionCalled, "Unexpected method call to commonHandleInSession")
	handleGetHealthFunc = handleGetHealth
	assert.Equal(t, handleGetHealthFuncExpected, handleGetHealthFuncCalled, "Unexpected method call to handleGetHealthFunc")
	handleGetHealthReportFunc = handleGetHealthReport
	assert.Equal(t, handleGetHealthReportFuncExpected, handleGetHealthReportFuncCalled, "Unexpected method call to handleGetHealthReportFunc")
}
