package logtype

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stringsJoinExpected int
	stringsJoinCalled   int
)

func createMock(t *testing.T) {
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
}

func verifyAll(t *testing.T) {
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected method call to stringsJoin")
}
