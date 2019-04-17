package session

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	uuidNewExpected int
	uuidNewCalled   int
)

func createMock(t *testing.T) {
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
}

func verifyAll(t *testing.T) {
	uuidNew = uuid.New
	if uuidNewExpected != uuidNewCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to uuidNew, expected %v, actual %v", uuidNewExpected, uuidNewCalled))
	}
}
