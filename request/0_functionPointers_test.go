package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

var (
	uuidParseExpected               int
	uuidParseCalled                 int
	uuidNewExpected                 int
	uuidNewCalled                   int
	logtypeFromStringExpected       int
	logtypeFromStringCalled         int
	apperrorWrapSimpleErrorExpected int
	apperrorWrapSimpleErrorCalled   int
	ioutilReadAllExpected           int
	ioutilReadAllCalled             int
	loggerAPIRequestExpected        int
	loggerAPIRequestCalled          int
	getUUIDFromHeaderFuncExpected   int
	getUUIDFromHeaderFuncCalled     int
)

func createMock(t *testing.T) {
	uuidParseExpected = 0
	uuidParseCalled = 0
	uuidParse = func(s string) (uuid.UUID, error) {
		uuidParseCalled++
		return uuid.Nil, nil
	}
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	logtypeFromStringExpected = 0
	logtypeFromStringCalled = 0
	logtypeFromString = func(value string) logtype.LogType {
		logtypeFromStringCalled++
		return 0
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	ioutilReadAllExpected = 0
	ioutilReadAllCalled = 0
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		return nil, nil
	}
	loggerAPIRequestExpected = 0
	loggerAPIRequestCalled = 0
	loggerAPIRequest = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAPIRequestCalled++
	}
	getUUIDFromHeaderFuncExpected = 0
	getUUIDFromHeaderFuncCalled = 0
	getUUIDFromHeaderFunc = func(header http.Header, name string) uuid.UUID {
		getUUIDFromHeaderFuncCalled++
		return uuid.Nil
	}
}

func verifyAll(t *testing.T) {
	uuidParse = uuid.Parse
	if uuidParseExpected != uuidParseCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to uuidParse, expected %v, actual %v", uuidParseExpected, uuidParseCalled))
	}
	uuidNew = uuid.New
	if uuidNewExpected != uuidNewCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to uuidNew, expected %v, actual %v", uuidNewExpected, uuidNewCalled))
	}
	logtypeFromString = logtype.FromString
	if logtypeFromStringExpected != logtypeFromStringCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to logtypeFromString, expected %v, actual %v", logtypeFromStringExpected, logtypeFromStringCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
	ioutilReadAll = ioutil.ReadAll
	if ioutilReadAllExpected != ioutilReadAllCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to ioutilReadAll, expected %v, actual %v", ioutilReadAllExpected, ioutilReadAllCalled))
	}
	loggerAPIRequest = logger.APIRequest
	if loggerAPIRequestExpected != loggerAPIRequestCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAPIRequest, expected %v, actual %v", loggerAPIRequestExpected, loggerAPIRequestCalled))
	}
	getUUIDFromHeaderFunc = getUUIDFromHeader
	if getUUIDFromHeaderFuncExpected != getUUIDFromHeaderFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to getUUIDFromHeaderFunc, expected %v, actual %v", getUUIDFromHeaderFuncExpected, getUUIDFromHeaderFuncCalled))
	}
}
