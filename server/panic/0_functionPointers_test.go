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
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	responseErrorExpected = 0
	responseErrorCalled = 0
	responseError = func(sessionID uuid.UUID, err error) {
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
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	getRecoverErrorFunc = getRecoverError
	assert.Equal(t, getRecoverErrorFuncExpected, getRecoverErrorFuncCalled, "Unexpected number of calls to getRecoverErrorFunc")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	responseError = response.Error
	assert.Equal(t, responseErrorExpected, responseErrorCalled, "Unexpected number of calls to responseError")
	apperrorGetGeneralFailureError = apperror.GetGeneralFailureError
	assert.Equal(t, apperrorGetGeneralFailureErrorExpected, apperrorGetGeneralFailureErrorCalled, "Unexpected number of calls to apperrorGetGeneralFailureError")
	getDebugStackFunc = getDebugStack
	assert.Equal(t, getDebugStackFuncExpected, getDebugStackFuncCalled, "Unexpected number of calls to getDebugStackFunc")
}

// mock structs
type dummyPanicResponseWriter struct {
	t *testing.T
}

func (drw *dummyPanicResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Header")
	return nil
}

func (drw *dummyPanicResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyPanicResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.WriteHeader")
}
