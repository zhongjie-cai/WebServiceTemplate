package health

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetHealth(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyHTTPRequest, _ = http.NewRequest(
		http.MethodGet,
		"http://localhost",
		nil,
	)
	var dummyResponseWriter = &dummyResponseWriter{t}
	var dummyAppVersion = "some app version"

	// mock
	createMock(t)

	// expect
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
	GetHealth(
		dummyResponseWriter,
		dummyHTTPRequest,
		dummySessionID,
	)

	// verify
	verifyAll(t)
}
