package panic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestGetRecoverError_Error(t *testing.T) {
	// arrange
	var dummyRecoverResult = errors.New("some error")
	var dummyAppError = apperror.GetCustomError(0, "some app error")

	// mock
	createMock(t)

	// expect
	apperrorGetGeneralFailureErrorExpected = 1
	apperrorGetGeneralFailureError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyRecoverResult, innerErrors[0])
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
	var dummyAppError = apperror.GetCustomError(0, "some app error")

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
	apperrorGetGeneralFailureError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetGeneralFailureErrorCalled++
		assert.Equal(t, 1, len(innerErrors))
		assert.Equal(t, dummyError, innerErrors[0])
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
	var dummySessionObject = &dummySession{}
	var dummyError = errors.New("some error")
	var dummyRecoverResult = dummyError.(interface{})
	var dummyAppError = apperror.GetGeneralFailureError(dummyError)
	var dummyDebugStack = "some debug stack"

	// mock
	createMock(t)

	// expect
	getRecoverErrorFuncExpected = 1
	getRecoverErrorFunc = func(recoverResult interface{}) apperrorModel.AppError {
		getRecoverErrorFuncCalled++
		assert.Equal(t, dummyRecoverResult, recoverResult)
		return dummyAppError
	}
	responseWriteExpected = 1
	responseWrite = func(session sessionModel.Session, responseObject interface{}, responseError error) {
		responseWriteCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Nil(t, responseObject)
		assert.Equal(t, dummyAppError, responseError)
	}
	getDebugStackFuncExpected = 1
	getDebugStackFunc = func() string {
		getDebugStackFuncCalled++
		return dummyDebugStack
	}
	loggerMethodLogicExpected = 1
	loggerMethodLogic = func(session sessionModel.Session, logLevel loglevel.LogLevel, category, subcategory, messageFormat string, parameters ...interface{}) {
		loggerMethodLogicCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, loglevel.Fatal, logLevel)
		assert.Equal(t, "panic", category)
		assert.Equal(t, "Handle", subcategory)
		assert.Equal(t, "%v\n%v", messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyAppError, parameters[0])
		assert.Equal(t, dummyDebugStack, parameters[1])
	}

	// SUT + act
	Handle(
		dummySessionObject,
		dummyRecoverResult,
	)

	// verify
	verifyAll(t)
}
