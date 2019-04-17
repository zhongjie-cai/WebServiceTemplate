package panic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

var (
	fmtErrorfExpected                      int
	fmtErrorfCalled                        int
	getRecoverErrorFuncExpected            int
	getRecoverErrorFuncCalled              int
	loggerAppRootExpected                  int
	loggerAppRootCalled                    int
	responseErrorExpected                  int
	responseErrorCalled                    int
	apperrorGetGeneralFailureErrorExpected int
	apperrorGetGeneralFailureErrorCalled   int
	getDebugStackFuncExpected              int
	getDebugStackFuncCalled                int
)

func createMock(t *testing.T) {
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	getRecoverErrorFuncExpected = 0
	getRecoverErrorFuncCalled = 0
	getRecoverErrorFunc = func(recoverResult interface{}) apperror.AppError {
		getRecoverErrorFuncCalled++
		return nil
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error, responseWriter http.ResponseWriter) {
		responseErrorCalled++
	}
	apperrorGetGeneralFailureErrorExpected = 0
	apperrorGetGeneralFailureErrorCalled = 0
	apperrorGetGeneralFailureError = func(innerError error) apperror.AppError {
		apperrorGetGeneralFailureErrorCalled++
		return nil
	}
	getDebugStackFuncExpected = 0
	getDebugStackFuncCalled = 0
	getDebugStackFunc = func() string {
		getDebugStackFuncCalled++
		return ""
	}
}

func verifyAll(t *testing.T) {
	fmtErrorf = fmt.Errorf
	if fmtErrorfExpected != fmtErrorfCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to fmtErrorf, expected %v, actual %v", fmtErrorfExpected, fmtErrorfCalled))
	}
	getRecoverErrorFunc = getRecoverError
	if getRecoverErrorFuncExpected != getRecoverErrorFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getRecoverErrorFunc, expected %v, actual %v", getRecoverErrorFuncExpected, getRecoverErrorFuncCalled))
	}
	loggerAppRoot = logger.AppRoot
	if loggerAppRootExpected != loggerAppRootCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAppRoot, expected %v, actual %v", loggerAppRootExpected, loggerAppRootCalled))
	}
	responseError = response.Error
	if responseErrorExpected != responseErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to responseError, expected %v, actual %v", responseErrorExpected, responseErrorCalled))
	}
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	if apperrorGetGeneralFailureErrorExpected != apperrorGetGeneralFailureErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorGetGeneralFailureError, expected %v, actual %v", apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled))
	}
	getDebugStackFunc = getDebugStack
	if getDebugStackFuncExpected != getDebugStackFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getDebugStackFunc, expected %v, actual %v", getDebugStackFuncExpected, getDebugStackFuncCalled))
	}
}
