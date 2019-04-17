package favicon

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// mock struct
type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected call to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Equal(drw.t, http.StatusMethodNotAllowed, statusCode)
}

func TestHandleFaviconLogic_MethodGet(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppPath = "some app path"

	// mock + assert
	configAppPathExpected = 1
	configAppPath = func() string {
		configAppPathCalled++
		return dummyAppPath
	}
	httpServeFileExpected = 1
	httpServeFile = func(responseWriter http.ResponseWriter, request *http.Request, name string) {
		httpServeFileCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyRequest, request)
		assert.Equal(t, dummyAppPath+"/favicon.ico", name)
	}

	// SUT + act
	handleFaviconLogic(
		dummyResponseWriter,
		dummyRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}

func TestHandleFaviconLogic_MethodNotGet(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		"whatever",
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock + assert
	apperrorGetInvalidOperationExpected = 1
	apperrorGetInvalidOperation = func(innerError error) apperror.AppError {
		apperrorGetInvalidOperationCalled++
		assert.Nil(t, innerError)
		return dummyAppError
	}
	responseErrorExpected = 1
	responseError = func(sessionID uuid.UUID, err error, responseWriter http.ResponseWriter) {
		responseErrorCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyAppError, err)
		assert.Equal(t, dummyResponseWriter, responseWriter)
	}

	// SUT + act
	handleFaviconLogic(
		dummyResponseWriter,
		dummyRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}

func TestHandler(t *testing.T) {
	// arrange
	var dummyRequest, _ = http.NewRequest(
		"whatever",
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}

	// mock
	createMock(t)

	// expect
	commonHandleInSessionExpected = 1
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, endpoint string, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyRequest, request)
		assert.Equal(t, "Favicon", endpoint)
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handleFaviconLogicFunc))
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(action)))
	}

	// SUT + act
	handler(
		dummyResponseWriter,
		dummyRequest,
	)

	// assert

	// tear down
	verifyAll(t)
}

func TestHostEntry(t *testing.T) {
	// mock
	createMock(t)

	// expect
	httpHandleFuncExpected = 1
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handlerFunc))
		assert.Equal(t, "/favicon.ico", pattern)
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
	}

	// SUT + act
	HostEntry()

	// verify
	verifyAll(t)
}
