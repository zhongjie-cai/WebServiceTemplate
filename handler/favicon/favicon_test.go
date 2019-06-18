package favicon

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetFavicon(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppPath = "some app path"

	// mock
	createMock(t)

	// expect
	configAppPathExpected = 1
	configAppPath = func() string {
		configAppPathCalled++
		return dummyAppPath
	}
	httpServeFileExpected = 1
	httpServeFile = func(responseWriter http.ResponseWriter, httpRequest *http.Request, name string) {
		httpServeFileCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, dummyAppPath+"/favicon.ico", name)
	}

	// SUT + act
	GetFavicon(
		dummyResponseWriter,
		dummyHTTPRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}
