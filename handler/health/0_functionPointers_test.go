package health

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
	configAppVersionExpected            int
	configAppVersionCalled              int
	responseOkExpected                  int
	responseOkCalled                    int
	responseErrorExpected               int
	responseErrorCalled                 int
	apperrorGetInvalidOperationExpected int
	apperrorGetInvalidOperationCalled   int
	commonHandleInSessionExpected       int
	commonHandleInSessionCalled         int
	handleHealthLogicFuncExpected       int
	handleHealthLogicFuncCalled         int
	handlerFuncExpected                 int
	handlerFuncCalled                   int
)

func createMock(t *testing.T) {
	httpHandleFuncExpected = 0
	httpHandleFuncCalled = 0
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
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
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error, responseWriter http.ResponseWriter) {
		responseErrorCalled++
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
	handleHealthLogicFuncExpected = 0
	handleHealthLogicFuncCalled = 0
	handleHealthLogicFunc = func(responseWriter http.ResponseWriter, request *http.Request, sessionID uuid.UUID) {
		handleHealthLogicFuncCalled++
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
	configAppVersion = config.AppVersion
	if configAppVersionExpected != configAppVersionCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppVersion, expected %v, actual %v", configAppVersionExpected, configAppVersionCalled))
	}
	responseOk = response.Ok
	if responseOkExpected != responseOkCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to responseOk, expected %v, actual %v", responseOkExpected, responseOkCalled))
	}
	responseError = response.Error
	if responseErrorExpected != responseErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to responseError, expected %v, actual %v", responseErrorExpected, responseErrorCalled))
	}
	apperrorGetInvalidOperation = apperror.GetInvalidOperation
	if apperrorGetInvalidOperationExpected != apperrorGetInvalidOperationCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorGetInvalidOperation, expected %v, actual %v", apperrorGetInvalidOperationExpected, apperrorGetInvalidOperationCalled))
	}
	commonHandleInSession = common.HandleInSession
	if commonHandleInSessionExpected != commonHandleInSessionCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to commonHandleInSession, expected %v, actual %v", commonHandleInSessionExpected, commonHandleInSessionCalled))
	}
	handleHealthLogicFunc = handleHealthLogic
	if handleHealthLogicFuncExpected != handleHealthLogicFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to handleHealthLogicFunc, expected %v, actual %v", handleHealthLogicFuncExpected, handleHealthLogicFuncCalled))
	}
	handlerFunc = handler
	if handlerFuncExpected != handlerFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to handlerFunc, expected %v, actual %v", handlerFuncExpected, handlerFuncCalled))
	}
}
