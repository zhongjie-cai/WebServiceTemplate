package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

var (
	uuidParseExpected                           int
	uuidParseCalled                             int
	uuidNewExpected                             int
	uuidNewCalled                               int
	configDefaultAllowedLogTypeExpected         int
	configDefaultAllowedLogTypeCalled           int
	configDefaultAllowedLogLevelExpected        int
	configDefaultAllowedLogLevelCalled          int
	customizationSessionAllowedLogTypeExpected  int
	customizationSessionAllowedLogTypeCalled    int
	customizationSessionAllowedLogLevelExpected int
	customizationSessionAllowedLogLevelCalled   int
	apperrorGetCustomErrorExpected              int
	apperrorGetCustomErrorCalled                int
	ioutilReadAllExpected                       int
	ioutilReadAllCalled                         int
	ioutilNopCloserExpected                     int
	ioutilNopCloserCalled                       int
	bytesNewBufferExpected                      int
	bytesNewBufferCalled                        int
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
	configDefaultAllowedLogTypeExpected = 0
	configDefaultAllowedLogTypeCalled = 0
	config.DefaultAllowedLogType = nil
	configDefaultAllowedLogLevelExpected = 0
	configDefaultAllowedLogLevelCalled = 0
	config.DefaultAllowedLogLevel = nil
	customizationSessionAllowedLogTypeExpected = 0
	customizationSessionAllowedLogTypeCalled = 0
	customization.SessionAllowedLogType = nil
	customizationSessionAllowedLogLevelExpected = 0
	customizationSessionAllowedLogLevelCalled = 0
	customization.SessionAllowedLogLevel = nil
	apperrorGetCustomErrorExpected = 0
	apperrorGetCustomErrorCalled = 0
	apperrorGetCustomError = func(errorCode apperrorEnum.Code, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorGetCustomErrorCalled++
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
	config.DefaultAllowedLogType = nil
	assert.Equal(t, configDefaultAllowedLogTypeExpected, configDefaultAllowedLogTypeCalled, "Unexpected number of calls to configDefaultAllowedLogType")
	config.DefaultAllowedLogLevel = nil
	assert.Equal(t, configDefaultAllowedLogLevelExpected, configDefaultAllowedLogLevelCalled, "Unexpected number of calls to configDefaultAllowedLogLevel")
	customization.SessionAllowedLogType = nil
	assert.Equal(t, customizationSessionAllowedLogTypeExpected, customizationSessionAllowedLogTypeCalled, "Unexpected number of calls to customizationSessionAllowedLogType")
	customization.SessionAllowedLogLevel = nil
	assert.Equal(t, customizationSessionAllowedLogLevelExpected, customizationSessionAllowedLogLevelCalled, "Unexpected number of calls to customizationSessionAllowedLogLevel")
	apperrorGetCustomError = apperror.GetCustomError
	assert.Equal(t, apperrorGetCustomErrorExpected, apperrorGetCustomErrorCalled, "Unexpected number of calls to apperrorGetCustomError")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected number of calls to ioutilReadAll")
	ioutilNopCloser = ioutil.NopCloser
	assert.Equal(t, ioutilNopCloserExpected, ioutilNopCloserCalled, "Unexpected number of calls to ioutilNopCloser")
	bytesNewBuffer = bytes.NewBuffer
	assert.Equal(t, bytesNewBufferExpected, bytesNewBufferCalled, "Unexpected number of calls to bytesNewBuffer")
}
