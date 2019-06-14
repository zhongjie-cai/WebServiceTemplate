package request

import (
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
	assert.Equal(t, uuidParseExpected, uuidParseCalled, "Unexpected method call to uuidParse")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected method call to uuidNew")
	logtypeFromString = logtype.FromString
	assert.Equal(t, logtypeFromStringExpected, logtypeFromStringCalled, "Unexpected method call to logtypeFromString")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected method call to ioutilReadAll")
	loggerAPIRequest = logger.APIRequest
	assert.Equal(t, loggerAPIRequestExpected, loggerAPIRequestCalled, "Unexpected method call to loggerAPIRequest")
	getUUIDFromHeaderFunc = getUUIDFromHeader
	assert.Equal(t, getUUIDFromHeaderFuncExpected, getUUIDFromHeaderFuncCalled, "Unexpected method call to getUUIDFromHeaderFunc")
}
