package loglevel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_NonSupportedLogLevels(t *testing.T) {
	// arrange
	var unsupportedValue = maxLogLevel

	// mock
	createMock(t)

	// SUT
	var sut = LogLevel(unsupportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, debugName, result)

	// verify
	verifyAll(t)
}

func TestString_SupportedLogLevel(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var sut = Error

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, errorName, result)

	// verify
	verifyAll(t)
}

func TestFromString_NoMatchFound(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// mock
	createMock(t)

	// SUT + act
	var result = FromString(dummyValue)

	// assert
	assert.Equal(t, Debug, result)

	// tear down
	verifyAll(t)
}

func TestFromString_HappyPath(t *testing.T) {
	for key, value := range logLevelNameMapping {
		// mock
		createMock(t)

		// SUT + act
		var result = FromString(key)

		// assert
		assert.Equal(t, value, result)

		// tear down
		verifyAll(t)
	}
}
