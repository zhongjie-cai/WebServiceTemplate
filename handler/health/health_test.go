package health

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
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

func TestHandleHealthLogic_MethodGet(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppVersion = "some app version"

	// mock + assert
	configAppVersionExpected = 1
	configAppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	responseOkExpected = 1
	responseOk = func(sessionID uuid.UUID, responseContent interface{}, responseWriter http.ResponseWriter) {
		responseOkCalled++
		assert.Equal(t, dummySessionID, sessionID)
		assert.Equal(t, dummyAppVersion, responseContent)
		assert.Equal(t, dummyResponseWriter, responseWriter)
	}

	// SUT + act
	handleHealthLogic(
		dummyResponseWriter,
		dummyRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}

func TestHandleHealthLogic_MethodNotGet(t *testing.T) {
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
	handleHealthLogic(
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
		assert.Equal(t, "Health", endpoint)
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handleHealthLogicFunc))
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
	httpHandleFuncExpected = 2
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		httpHandleFuncCalled++
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handlerFunc))
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
		if httpHandleFuncCalled == 1 {
			assert.Equal(t, "/health", pattern)
		} else if httpHandleFuncCalled == 2 {
			assert.Equal(t, "/health/", pattern)
		}
	}

	// SUT + act
	HostEntry()

	// verify
	verifyAll(t)
}
