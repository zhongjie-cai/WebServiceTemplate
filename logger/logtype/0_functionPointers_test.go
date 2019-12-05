package logtype

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stringsJoinExpected  int
	stringsJoinCalled    int
	stringsSplitExpected int
	stringsSplitCalled   int
)

func createMock(t *testing.T) {
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
	stringsSplitExpected = 0
	stringsSplitCalled = 0
	stringsSplit = func(s, sep string) []string {
		stringsSplitCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	stringsSplit = strings.Split
	assert.Equal(t, stringsSplitExpected, stringsSplitCalled, "Unexpected number of calls to stringsSplit")
}
