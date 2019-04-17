package favicon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/common"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

var (
	httpHandleFuncExpected              int
	httpHandleFuncCalled                int
	httpServeFileExpected               int
	httpServeFileCalled                 int
	responseErrorExpected               int
	responseErrorCalled                 int
	configAppPathExpected               int
	configAppPathCalled                 int
	apperrorGetInvalidOperationExpected int
	apperrorGetInvalidOperationCalled   int
	commonHandleInSessionExpected       int
	commonHandleInSessionCalled         int
	handleFaviconLogicFuncExpected      int
	handleFaviconLogicFuncCalled        int
	handlerFuncExpected                 int
	handlerFuncCalled                   int
)

func createMock(t *testing.T) {
	httpHandleFuncExpected = 0
	httpHandleFuncCalled = 0
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
	}
	httpServeFileExpected = 0
	httpServeFileCalled = 0
	httpServeFile = func(responseWriter http.ResponseWriter, request *http.Request, name string) {
		httpServeFileCalled++
	}
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error, responseWriter http.ResponseWriter) {
		responseErrorCalled++
	}
	configAppPathExpected = 0
	configAppPathCalled = 0
	configAppPath = func() string {
		configAppPathCalled++
		return ""
	}
	apperrorGetInvalidOperationExpected = 0
	apperrorGetInvalidOperationCalled = 0
	apperrorGetInvalidOperation = func(innerError error) apperror.AppError {
		apperrorGetInvalidOperationCalled++
		return nil
	}
	commonHandleInSessionExpected = 0
	commonHandleInSessionCalled = 0
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, endpoint string, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
	}
	handleFaviconLogicFuncExpected = 0
	handleFaviconLogicFuncCalled = 0
	handleFaviconLogicFunc = func(responseWriter http.ResponseWriter, request *http.Request, sessionID uuid.UUID) {
		handleFaviconLogicFuncCalled++
	}
	handlerFuncExpected = 0
	handlerFuncCalled = 0
	handlerFunc = func(responseWriter http.ResponseWriter, request *http.Request) {
		handlerFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	httpHandleFunc = http.HandleFunc
	if httpHandleFuncExpected != httpHandleFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpHandleFunc, expected %v, actual %v", httpHandleFuncExpected, httpHandleFuncCalled))
	}
	httpServeFile = http.ServeFile
	if httpServeFileExpected != httpServeFileCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to httpServeFile, expected %v, actual %v", httpServeFileExpected, httpServeFileCalled))
	}
	responseError = response.Error
	if responseErrorExpected != responseErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to responseError, expected %v, actual %v", responseErrorExpected, responseErrorCalled))
	}
	configAppPath = config.AppPath
	if configAppPathExpected != configAppPathCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppPath, expected %v, actual %v", configAppPathExpected, configAppPathCalled))
	}
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
	if apperrorGetInvalidOperationExpected != apperrorGetInvalidOperationCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorGetInvalidOperation, expected %v, actual %v", apperrorGetInvalidOperationExpected, apperrorGetInvalidOperationCalled))
	}
	commonHandleInSession = common.HandleInSession
	if commonHandleInSessionExpected != commonHandleInSessionCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to commonHandleInSession, expected %v, actual %v", commonHandleInSessionExpected, commonHandleInSessionCalled))
	}
	handleFaviconLogicFunc = handleFaviconLogic
	if handleFaviconLogicFuncExpected != handleFaviconLogicFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to handleFaviconLogicFunc, expected %v, actual %v", handleFaviconLogicFuncExpected, handleFaviconLogicFuncCalled))
	}
	handlerFunc = handler
	if handlerFuncExpected != handlerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to handlerFunc, expected %v, actual %v", handlerFuncExpected, handlerFuncCalled))
	}
}
