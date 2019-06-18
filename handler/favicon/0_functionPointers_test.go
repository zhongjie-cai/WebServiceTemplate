package favicon

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

var (
	httpServeFileExpected int
	httpServeFileCalled   int
	configAppPathExpected int
	configAppPathCalled   int
)

func createMock(t *testing.T) {
	httpServeFileExpected = 0
	httpServeFileCalled = 0
	httpServeFile = func(responseWriter http.ResponseWriter, httpRequest *http.Request, name string) {
		httpServeFileCalled++
	}
	configAppPathExpected = 0
	configAppPathCalled = 0
	config.AppPath = func() string {
		configAppPathCalled++
		return ""
	}
}

func verifyAll(t *testing.T) {
	httpServeFile = http.ServeFile
	assert.Equal(t, httpServeFileExpected, httpServeFileCalled, "Unexpected number of calls to httpServeFile")
	config.AppPath = func() string { return "" }
	assert.Equal(t, configAppPathExpected, configAppPathCalled, "Unexpected number of calls to configAppPath")
}

// mock structs
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
