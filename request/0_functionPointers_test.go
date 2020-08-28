package request

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
)

var (
	uuidParseExpected              int
	uuidParseCalled                int
	uuidNewExpected                int
	uuidNewCalled                  int
	apperrorGetCustomErrorExpected int
	apperrorGetCustomErrorCalled   int
	ioutilReadAllExpected          int
	ioutilReadAllCalled            int
	ioutilNopCloserExpected        int
	ioutilNopCloserCalled          int
	bytesNewBufferExpected         int
	bytesNewBufferCalled           int
	httputilDumpRequestExpected    int
	httputilDumpRequestCalled      int
	fmtSprintfExpected             int
	fmtSprintfCalled               int
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
	httputilDumpRequestExpected = 0
	httputilDumpRequestCalled = 0
	httputilDumpRequest = func(req *http.Request, body bool) ([]byte, error) {
		httputilDumpRequestCalled++
		return nil, nil
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
}

func verifyAll(t *testing.T) {
	uuidParse = uuid.Parse
	assert.Equal(t, uuidParseExpected, uuidParseCalled, "Unexpected number of calls to uuidParse")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	apperrorGetCustomError = apperror.GetCustomError
	assert.Equal(t, apperrorGetCustomErrorExpected, apperrorGetCustomErrorCalled, "Unexpected number of calls to apperrorGetCustomError")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected number of calls to ioutilReadAll")
	ioutilNopCloser = ioutil.NopCloser
	assert.Equal(t, ioutilNopCloserExpected, ioutilNopCloserCalled, "Unexpected number of calls to ioutilNopCloser")
	bytesNewBuffer = bytes.NewBuffer
	assert.Equal(t, bytesNewBufferExpected, bytesNewBufferCalled, "Unexpected number of calls to bytesNewBuffer")
	httputilDumpRequest = httputil.DumpRequest
	assert.Equal(t, httputilDumpRequestExpected, httputilDumpRequestCalled, "Unexpected number of calls to httputilDumpRequest")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to fmtSprintf")
}
