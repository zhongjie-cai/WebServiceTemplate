package favicon

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

func TestHandleGetFavicon(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppPath = "some app path"

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
	handleGetFavicon(
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
	routeHandleFuncExpected = 1
	routeHandleFunc = func(router *mux.Router, endpoint string, method string, path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
		routeHandleFuncCalled++
		var expectedPointer = fmt.Sprintf("%v", reflect.ValueOf(handleGetFaviconFunc))
		assert.Equal(t, dummyRouter, router)
		assert.Equal(t, "Favicon", endpoint)
		assert.Equal(t, http.MethodGet, method)
		assert.Equal(t, "/favicon.ico", path)
		assert.Equal(t, expectedPointer, fmt.Sprintf("%v", reflect.ValueOf(handler)))
		return nil
	}

	// SUT + act
	HostEntry(
		dummyRouter,
	)

	// verify
	verifyAll(t)
}
