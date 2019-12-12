package logtype

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sortStringsExpected  int
	sortStringsCalled    int
	stringsJoinExpected  int
	stringsJoinCalled    int
	stringsSplitExpected int
	stringsSplitCalled   int
)

func createMock(t *testing.T) {
	sortStringsExpected = 0
	sortStringsCalled = 0
	sortStrings = func(a []string) {
		sortStringsCalled++
	}
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
	sortStrings = sort.Strings
	assert.Equal(t, sortStringsExpected, sortStringsCalled, "Unexpected number of calls to sortStrings")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	stringsSplit = strings.Split
	assert.Equal(t, stringsSplitExpected, stringsSplitCalled, "Unexpected number of calls to stringsSplit")
}
