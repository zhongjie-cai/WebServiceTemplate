package health

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

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

func TestHandleGetHealth(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppVersion = "some app version"

	// mock
	createMock(t)

	// expect
	commonHandleInSessionExpected = 1
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyRequest, request)
		action(
			dummyResponseWriter,
			dummyRequest,
			dummySessionID,
		)
	}
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
	handleGetHealth(
		dummyResponseWriter,
		dummyRequest,
	)

	// verify
	verifyAll(t)
}

func TestHandleGetHealthReport(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppVersion = "some app version"

	// mock
	createMock(t)

	// expect
	commonHandleInSessionExpected = 1
	commonHandleInSession = func(responseWriter http.ResponseWriter, request *http.Request, action func(http.ResponseWriter, *http.Request, uuid.UUID)) {
		commonHandleInSessionCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyRequest, request)
		action(
			dummyResponseWriter,
			dummyRequest,
			dummySessionID,
		)
	}
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
	handleGetHealthReport(
		dummyResponseWriter,
		dummyRequest,
	)

	// verify
	verifyAll(t)
}

func TestHostEntry(t *testing.T) {
	// arrange
	var dummyRouter = &mux.Router{}

	// mock
	createMock(t)

	// expect
	routeHandleFuncExpected = 2
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		assert.Equal(t, dummyRouter, router)
		assert.Equal(t, http.MethodGet, method)
		if routeHandleFuncCalled == 1 {
			assert.Equal(t, "Health", endpoint)
			assert.Equal(t, "/health", path)
			var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handleGetHealthFunc))
			assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
		} else if routeHandleFuncCalled == 2 {
			assert.Equal(t, "HealthReport", endpoint)
			assert.Equal(t, "/health/report", path)
			var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handleGetHealthReportFunc))
			assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
		}
		return nil
	}

	// SUT + act
	HostEntry(
		dummyRouter,
	)

	// verify
	verifyAll(t)
}
