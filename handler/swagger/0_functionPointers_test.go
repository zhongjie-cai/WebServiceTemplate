package swagger

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

var (
	configAppPathExpected   int
	configAppPathCalled     int
	httpRedirectExpected    int
	httpRedirectCalled      int
	httpStripPrefixExpected int
	httpStripPrefixCalled   int
	httpFileServerExpected  int
	httpFileServerCalled    int
)

func createMock(t *testing.T) {
	configAppPathExpected = 0
	configAppPathCalled = 0
	configAppPath = func() string {
		configAppPathCalled++
		return ""
	}
	httpRedirectExpected = 0
	httpRedirectCalled = 0
	httpRedirect = func(responseWriter http.ResponseWriter, httpRequest *http.Request, url string, code int) {
		httpRedirectCalled++
	}
	httpStripPrefixExpected = 0
	httpStripPrefixCalled = 0
	httpStripPrefix = func(prefix string, h http.Handler) http.Handler {
		httpStripPrefixCalled++
		return nil
	}
	httpFileServerExpected = 0
	httpFileServerCalled = 0
	httpFileServer = func(root http.FileSystem) http.Handler {
		httpFileServerCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	configAppPath = config.AppPath
	assert.Equal(t, configAppPathExpected, configAppPathCalled, "Unexpected number of calls to configAppPath")
	httpRedirect = http.Redirect
	assert.Equal(t, httpRedirectExpected, httpRedirectCalled, "Unexpected number of calls to httpRedirect")
	httpStripPrefix = http.StripPrefix
	assert.Equal(t, httpStripPrefixExpected, httpStripPrefixCalled, "Unexpected number of calls to httpStripPrefix")
	httpFileServer = http.FileServer
	assert.Equal(t, httpFileServerExpected, httpFileServerCalled, "Unexpected number of calls to httpFileServer")
}

// mock structs
type dummyHandlerStruct struct {
}

func (dhs *dummyHandlerStruct) ServeHTTP(responseWriter http.ResponseWriter, httpRequest *http.Request) {
}

type dummyResponseWriter struct {
	t *testing.T
}

func (drw *dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Header")
	return nil
}

func (drw *dummyResponseWriter) Write([]byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to ResponseWrite.Write")
	return 0, nil
}

func (drw *dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Equal(drw.t, http.StatusMethodNotAllowed, statusCode)
}
