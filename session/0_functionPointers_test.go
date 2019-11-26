package session

import (
	"encoding/json"
	"net/http"
	"net/textproto"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/request"
	"github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	uuidNewExpected                         int
	uuidNewCalled                           int
	jsonMarshalExpected                     int
	jsonMarshalCalled                       int
	jsonUnmarshalExpected                   int
	jsonUnmarshalCalled                     int
	fmtErrorfExpected                       int
	fmtErrorfCalled                         int
	muxVarsExpected                         int
	muxVarsCalled                           int
	requestGetRequestBodyExpected           int
	requestGetRequestBodyCalled             int
	apperrorGetBadRequestErrorExpected      int
	apperrorGetBadRequestErrorCalled        int
	apperrorWrapSimpleErrorExpected         int
	apperrorWrapSimpleErrorCalled           int
	textprotoCanonicalMIMEHeaderKeyExpected int
	textprotoCanonicalMIMEHeaderKeyCalled   int
	getFuncExpected                         int
	getFuncCalled                           int
	tryUnmarshalFuncExpected                int
	tryUnmarshalFuncCalled                  int
	getRequestFuncExpected                  int
	getRequestFuncCalled                    int
	getAllQueriesFuncExpected               int
	getAllQueriesFuncCalled                 int
	getAllHeadersFuncExpected               int
	getAllHeadersFuncCalled                 int
	configIsLocalhostExpected               int
	configIsLocalhostCalled                 int
)

func createMock(t *testing.T) {
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	jsonMarshalExpected = 0
	jsonMarshalCalled = 0
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		return nil, nil
	}
	jsonUnmarshalExpected = 0
	jsonUnmarshalCalled = 0
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return nil
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	muxVarsExpected = 0
	muxVarsCalled = 0
	muxVars = func(r *http.Request) map[string]string {
		muxVarsCalled++
		return nil
	}
	requestGetRequestBodyExpected = 0
	requestGetRequestBodyCalled = 0
	requestGetRequestBody = func(httpRequest *http.Request) string {
		requestGetRequestBodyCalled++
		return ""
	}
	apperrorGetBadRequestErrorExpected = 0
	apperrorGetBadRequestErrorCalled = 0
	apperrorGetBadRequestError = func(innerErrors ...error) apperrorModel.AppError {
		apperrorGetBadRequestErrorCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerErrors []error, messageFormat string, parameters ...interface{}) apperrorModel.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
	textprotoCanonicalMIMEHeaderKeyExpected = 0
	textprotoCanonicalMIMEHeaderKeyCalled = 0
	textprotoCanonicalMIMEHeaderKey = func(s string) string {
		textprotoCanonicalMIMEHeaderKeyCalled++
		return ""
	}
	getFuncExpected = 0
	getFuncCalled = 0
	getFunc = func(sessionID uuid.UUID) model.Session {
		getFuncCalled++
		return nil
	}
	tryUnmarshalFuncExpected = 0
	tryUnmarshalFuncCalled = 0
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) apperrorModel.AppError {
		tryUnmarshalFuncCalled++
		return nil
	}
	getRequestFuncExpected = 0
	getRequestFuncCalled = 0
	getRequestFunc = func(sessionID uuid.UUID) *http.Request {
		getRequestFuncCalled++
		return nil
	}
	getAllQueriesFuncExpected = 0
	getAllQueriesFuncCalled = 0
	getAllQueriesFunc = func(session *session, name string) []string {
		getAllQueriesFuncCalled++
		return nil
	}
	getAllHeadersFuncExpected = 0
	getAllHeadersFuncCalled = 0
	getAllHeadersFunc = func(session *session, name string) []string {
		getAllHeadersFuncCalled++
		return nil
	}
	configIsLocalhostExpected = 0
	configIsLocalhostCalled = 0
	config.IsLocalhost = func() bool {
		configIsLocalhostCalled++
		return false
	}
}

func verifyAll(t *testing.T) {
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	jsonMarshal = json.Marshal
	assert.Equal(t, jsonMarshalExpected, jsonMarshalCalled, "Unexpected number of calls to jsonMarshal")
	jsonUnmarshal = json.Unmarshal
	assert.Equal(t, jsonUnmarshalExpected, jsonUnmarshalCalled, "Unexpected number of calls to jsonUnmarshal")
	muxVars = mux.Vars
	assert.Equal(t, muxVarsExpected, muxVarsCalled, "Unexpected number of calls to muxVars")
	requestGetRequestBody = request.GetRequestBody
	assert.Equal(t, requestGetRequestBodyExpected, requestGetRequestBodyCalled, "Unexpected number of calls to requestGetRequestBody")
	apperrorGetBadRequestError = apperror.GetBadRequestError
	assert.Equal(t, apperrorGetBadRequestErrorExpected, apperrorGetBadRequestErrorCalled, "Unexpected number of calls to apperrorGetBadRequestError")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected number of calls to apperrorWrapSimpleError")
	textprotoCanonicalMIMEHeaderKey = textproto.CanonicalMIMEHeaderKey
	assert.Equal(t, textprotoCanonicalMIMEHeaderKeyExpected, textprotoCanonicalMIMEHeaderKeyCalled, "Unexpected number of calls to textprotoCanonicalMIMEHeaderKey")
	getFunc = Get
	assert.Equal(t, getFuncExpected, getFuncCalled, "Unexpected number of calls to getFunc")
	tryUnmarshalFunc = tryUnmarshal
	assert.Equal(t, tryUnmarshalFuncExpected, tryUnmarshalFuncCalled, "Unexpected number of calls to tryUnmarshalFunc")
	getRequestFunc = GetRequest
	assert.Equal(t, getRequestFuncExpected, getRequestFuncCalled, "Unexpected number of calls to getRequestFunc")
	getAllQueriesFunc = getAllQueries
	assert.Equal(t, getAllQueriesFuncExpected, getAllQueriesFuncCalled, "Unexpected number of calls to getAllQueriesFunc")
	getAllHeadersFunc = getAllHeaders
	assert.Equal(t, getAllHeadersFuncExpected, getAllHeadersFuncCalled, "Unexpected number of calls to getAllHeadersFunc")
	config.IsLocalhost = func() bool { return false }
	assert.Equal(t, configIsLocalhostExpected, configIsLocalhostCalled, "Unexpected number of calls to configIsLocalhost")
}

// mock structs
type dummyResponseWriter struct {
	t *testing.T
}

func (drw dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to Header")
	return nil
}

func (drw dummyResponseWriter) Write(bytes []byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to Write")
	return 0, nil
}

func (drw dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to WriteHeader")
}

type dummyAttachment struct {
	ID   uuid.UUID
	Foo  string
	Test int
}
