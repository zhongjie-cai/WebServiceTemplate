package session

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	uuidNewExpected int
	uuidNewCalled   int
	getFuncExpected int
	getFuncCalled   int
)

func createMock(t *testing.T) {
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	getFuncExpected = 0
	getFuncCalled = 0
	getFunc = func(sessionID uuid.UUID) *Session {
		getFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to uuidNew")
	getFunc = Get
	assert.Equal(t, getFuncExpected, getFuncCalled, "Unexpected number of calls to getFunc")
}

// mock structs
type dummyResponseWriter struct {
	t *testing.T
}

func (drw dummyResponseWriter) Header() http.Header {
	assert.Fail(drw.t, "Unexpected number of calls to Header")
	return nil
}

func (drw dummyResponseWriter) Write(bytes []byte) (int, error) {
	assert.Fail(drw.t, "Unexpected number of calls to Write")
	return 0, nil
}

func (drw dummyResponseWriter) WriteHeader(statusCode int) {
	assert.Fail(drw.t, "Unexpected number of calls to WriteHeader")
}
