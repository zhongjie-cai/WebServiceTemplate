package apperror

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fmtSprintfExpected          int
	fmtSprintfCalled            int
	fmtErrorfExpected           int
	fmtErrorfCalled             int
	stringsJoinExpected         int
	stringsJoinCalled           int
	wrapErrorFuncExpected       int
	wrapErrorFuncCalled         int
	wrapSimpleErrorFuncExpected int
	wrapSimpleErrorFuncCalled   int
)

func createMock(t *testing.T) {
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
	wrapErrorFuncExpected = 0
	wrapErrorFuncCalled = 0
	wrapErrorFunc = func(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
		wrapErrorFuncCalled++
		return nil
	}
	wrapSimpleErrorFuncExpected = 0
	wrapSimpleErrorFuncCalled = 0
	wrapSimpleErrorFunc = func(innerError error, messageFormat string, parameters ...interface{}) AppError {
		wrapSimpleErrorFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	fmtSprintf = fmt.Sprintf
	if fmtSprintfExpected != fmtSprintfCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to fmtSprintf, expected %v, actual %v", fmtSprintfExpected, fmtSprintfCalled))
	}
	fmtErrorf = fmt.Errorf
	if fmtErrorfExpected != fmtErrorfCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to fmtErrorf, expected %v, actual %v", fmtErrorfExpected, fmtErrorfCalled))
	}
	stringsJoin = strings.Join
	if stringsJoinExpected != stringsJoinCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to stringsJoin, expected %v, actual %v", stringsJoinExpected, stringsJoinCalled))
	}
	wrapErrorFunc = WrapError
	if wrapErrorFuncExpected != wrapErrorFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to wrapErrorFunc, expected %v, actual %v", wrapErrorFuncExpected, wrapErrorFuncCalled))
	}
	wrapSimpleErrorFunc = WrapSimpleError
	if wrapSimpleErrorFuncExpected != wrapSimpleErrorFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to wrapSimpleErrorFunc, expected %v, actual %v", wrapSimpleErrorFuncExpected, wrapSimpleErrorFuncCalled))
	}
}
