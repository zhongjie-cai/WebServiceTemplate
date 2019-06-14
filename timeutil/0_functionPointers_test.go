package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	timeNowExpected int
	timeNowCalled   int
)

func createMock(t *testing.T) {
	timeNowExpected = 0
	timeNowCalled = 0
	timeNow = func() time.Time {
		timeNowCalled++
		return time.Time{}
	}
}

func verifyAll(t *testing.T) {
	timeNow = time.Now
	assert.Equal(t, timeNowExpected, timeNowCalled, "Unexpected method call to timeNow")
}
