package model

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	regexpMatchStringExpected int
	regexpMatchStringCalled   int
)

func createMock(t *testing.T) {
	regexpMatchStringExpected = 0
	regexpMatchStringCalled = 0
	regexpMatchString = func(pattern string, s string) (bool, error) {
		regexpMatchStringCalled++
		return false, nil
	}
}

func verifyAll(t *testing.T) {
	regexpMatchString = regexp.MatchString
	assert.Equal(t, regexpMatchStringExpected, regexpMatchStringCalled, "Unexpected number of calls to regexpMatchString")
}
