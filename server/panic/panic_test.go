package panic

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestGetRecoverError_Error(t *testing.T) {
	// arrange
	var dummyRecoverResult = errors.New("some error")
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	apperrorGetGeneralFailureErrorExpected = 1
	apperrorGetGeneralFailureError = func(innerError error) apperror.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, dummyRecoverResult, innerError)
		return dummyAppError
	}

	// SUT + act
	var result = getRecoverError(
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyAppError, result)

	// verify
	verifyAll(t)
}

func TestGetRecoverError_NonError(t *testing.T) {
	// arrange
	var dummyRecoverResult = "some recovery result"
	var dummyError = errors.New("some error")
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "%v", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyRecoverResult, a[0])
		return dummyError
	}
	apperrorGetGeneralFailureErrorExpected = 1
	apperrorGetGeneralFailureError = func(innerError error) apperror.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, dummyError, innerError)
		return dummyAppError
	}

	// SUT + act
	var result = getRecoverError(
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyAppError, result)

	// verify
	verifyAll(t)
}

func TestGetDebugStack(t *testing.T) {
	// SUT + act
	var result = getDebugStack()

	// assert
	assert.NotZero(t, result)

	// verify
	verifyAll(t)
}

func TestHandlePanic(t *testing.T) {
	// arrange
	var dummyEndpointName = "some endpoint name"
	var dummySessionID = uuid.New()
	var dummyError = errors.New("some error")
	var dummyRecoverResult = dummyError.(interface{})
	var dummyAppError = apperror.GetGeneralFailureError(dummyError)
	var dummyDebugStack = "some debug stack"
	var dummyResponseWriter = &dummyPanicResponseWriter{t}

	// mock
	createMock(t)

	// expect
	getRecoverErrorFuncExpected = 1
	getRecoverErrorFunc = func(recoverResult interface{}) apperror.AppError {
		getRecoverErrorFuncCalled++
		assert.Equal(t, dummyRecoverResult, recoverResult)
		return dummyAppError
	}
	responseWriteExpected = 1
	responseWrite = func(sessionID uuid.UUID, responseObject interface{}, responseError apperror.AppError) {
		responseWriteCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyAppError, responseError)
	}
	getDebugStackFuncExpected = 1
	getDebugStackFunc = func() string {
		getDebugStackFuncCalled++
		return dummyDebugStack
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "panic", category)
		assert.Equal(t, "Handle", subcategory)
		assert.Equal(t, "%v\n%v", messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyAppError, parameters[0])
		assert.Equal(t, dummyDebugStack, parameters[1])
	}

	// SUT + act
	Handle(
		dummyEndpointName,
		dummySessionID,
		dummyRecoverResult,
		dummyResponseWriter,
	)

	// verify
	verifyAll(t)
}
