package apperror

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/jsonutil"
)

var (
	fmtSprintfExpected                 int
	fmtSprintfCalled                   int
	fmtErrorfExpected                  int
	fmtErrorfCalled                    int
	stringsJoinExpected                int
	stringsJoinCalled                  int
	jsonutilMarshalIgnoreErrorExpected int
	jsonutilMarshalIgnoreErrorCalled   int
	cleanupInnerErrorsFuncExpected     int
	cleanupInnerErrorsFuncCalled       int
	wrapErrorFuncExpected              int
	wrapErrorFuncCalled                int
	wrapSimpleErrorFuncExpected        int
	wrapSimpleErrorFuncCalled          int
	customizationAppErrorsExpected     int
	customizationAppErrorsCalled       int
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
	jsonutilMarshalIgnoreErrorExpected = 0
	jsonutilMarshalIgnoreErrorCalled = 0
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		return ""
	}
	cleanupInnerErrorsFuncExpected = 0
	cleanupInnerErrorsFuncCalled = 0
	cleanupInnerErrorsFunc = func(innerErrors []error) []error {
		cleanupInnerErrorsFuncCalled++
		return nil
	}
	wrapErrorFuncExpected = 0
	wrapErrorFuncCalled = 0
	wrapErrorFunc = func(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
		wrapErrorFuncCalled++
		return nil
	}
	wrapSimpleErrorFuncExpected = 0
	wrapSimpleErrorFuncCalled = 0
	wrapSimpleErrorFunc = func(innerErrors []error, messageFormat string, parameters ...interface{}) model.AppError {
		wrapSimpleErrorFuncCalled++
		return nil
	}
	customizationAppErrorsExpected = 0
	customizationAppErrorsCalled = 0
	customization.AppErrors = nil
}

func verifyAll(t *testing.T) {
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	jsonutilMarshalIgnoreError = jsonutil.MarshalIgnoreError
	assert.Equal(t, jsonutilMarshalIgnoreErrorExpected, jsonutilMarshalIgnoreErrorCalled, "Unexpected number of calls to jsonutilMarshalIgnoreError")
	cleanupInnerErrorsFunc = cleanupInnerErrors
	assert.Equal(t, cleanupInnerErrorsFuncExpected, cleanupInnerErrorsFuncCalled, "Unexpected number of calls to cleanupInnerErrorsFunc")
	wrapErrorFunc = WrapError
	assert.Equal(t, wrapErrorFuncExpected, wrapErrorFuncCalled, "Unexpected number of calls to wrapErrorFunc")
	wrapSimpleErrorFunc = WrapSimpleError
	assert.Equal(t, wrapSimpleErrorFuncExpected, wrapSimpleErrorFuncCalled, "Unexpected number of calls to wrapSimpleErrorFunc")
	customization.AppErrors = nil
	assert.Equal(t, customizationAppErrorsExpected, customizationAppErrorsCalled, "Unexpected number of calls to customization.AppErrors")
}
