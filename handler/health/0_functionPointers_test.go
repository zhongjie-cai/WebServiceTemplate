package health

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

var (
	configAppVersionExpected int
	configAppVersionCalled   int
	responseOkExpected       int
	responseOkCalled         int
)

func createMock(t *testing.T) {
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	configAppVersion = func() string {
		configAppVersionCalled++
		return ""
	}
	responseOkExpected = 0
	responseOkCalled = 0
	responseOk = func(sessionID uuid.UUID, responseContent interface{}, responseWriter http.ResponseWriter) {
		responseOkCalled++
	}
}

func verifyAll(t *testing.T) {
	configAppVersion = config.AppVersion
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected number of calls to configAppVersion")
	responseOk = response.Ok
	assert.Equal(t, responseOkExpected, responseOkCalled, "Unexpected number of calls to responseOk")
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
