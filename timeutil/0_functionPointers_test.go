package timeutil

import (
	"fmt"
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
	if timeNowExpected != timeNowCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to timeNow, expected %v, actual %v", timeNowExpected, timeNowCalled))
	}
}
