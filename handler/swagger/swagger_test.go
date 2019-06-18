package swagger

import (
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestRedirectHandler(t *testing.T) {
	// arrange
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// expect
	httpRedirectExpected = 1
	httpRedirect = func(responseWriter http.ResponseWriter, httpRequest *http.Request, url string, code int) {
		httpRedirectCalled++
		assert.Equal(t, dummyResponseWriter, responseWriter)
		assert.Equal(t, dummyHTTPRequest, httpRequest)
		assert.Equal(t, "/docs/", url)
		assert.Equal(t, http.StatusPermanentRedirect, code)
	}

	// SUT + act
	Redirect(
		dummyResponseWriter,
		dummyHTTPRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}

func TestContentHandler(t *testing.T) {
	// arrange
	var dummyAppPath = "some app path"
	var dummyFileHandler = &dummyHandlerStruct{}
	var dummyForwardedHandler = &dummyHandlerStruct{}

	// mock
	createMock(t)

	// expect
	configAppPathExpected = 1
	configAppPath = func() string {
		configAppPathCalled++
		return dummyAppPath
	}
	httpFileServerExpected = 1
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		assert.Equal(t, http.Dir(dummyAppPath+"/docs"), root)
		return dummyFileHandler
	}
	httpStripPrefixExpected = 1
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		assert.Equal(t, "/docs/", prefix)
		assert.Equal(t, dummyFileHandler, h)
		return dummyForwardedHandler
	}

	// SUT + act
	var result = Handler()

	// assert
	assert.Equal(t, dummyForwardedHandler, result)

	// verify
	verifyAll(t)
}
