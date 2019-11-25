package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

var (
	uuidParseExpected               int
	uuidParseCalled                 int
	uuidNewExpected                 int
	uuidNewCalled                   int
	logtypeFromStringExpected       int
	logtypeFromStringCalled         int
	loglevelFromStringExpected      int
	loglevelFromStringCalled        int
	apperrorWrapSimpleErrorExpected int
	apperrorWrapSimpleErrorCalled   int
	ioutilReadAllExpected           int
	ioutilReadAllCalled             int
	ioutilNopCloserExpected         int
	ioutilNopCloserCalled           int
	bytesNewBufferExpected          int
	bytesNewBufferCalled            int
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
	loglevelFromStringExpected = 0
	loglevelFromStringCalled = 0
	loglevelFromString = func(value string) loglevel.LogLevel {
		loglevelFromStringCalled++
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
	ioutilNopCloserExpected = 0
	ioutilNopCloserCalled = 0
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		return nil
	}
	bytesNewBufferExpected = 0
	bytesNewBufferCalled = 0
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	uuidParse = uuid.Parse
	assert.Equal(t, uuidParseExpected, uuidParseCalled, "Unexpected number of calls to uuidParse")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	logtypeFromString = logtype.FromString
	assert.Equal(t, logtypeFromStringExpected, logtypeFromStringCalled, "Unexpected number of calls to logtypeFromString")
	loglevelFromString = loglevel.FromString
	assert.Equal(t, loglevelFromStringExpected, loglevelFromStringCalled, "Unexpected number of calls to loglevelFromString")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected number of calls to ioutilReadAll")
	ioutilNopCloser = ioutil.NopCloser
	assert.Equal(t, ioutilNopCloserExpected, ioutilNopCloserCalled, "Unexpected number of calls to ioutilNopCloser")
	bytesNewBuffer = bytes.NewBuffer
	assert.Equal(t, bytesNewBufferExpected, bytesNewBufferCalled, "Unexpected number of calls to bytesNewBuffer")
}
