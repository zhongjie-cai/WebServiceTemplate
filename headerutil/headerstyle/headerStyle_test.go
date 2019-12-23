package headerstyle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_NonSupportedHeaderStyles(t *testing.T) {
	// arrange
	var unsupportedValue = maxHeaderStyle

	// mock
	createMock(t)

	// SUT
	var sut = HeaderStyle(unsupportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, doNotLogName, result)

	// verify
	verifyAll(t)
}

func TestString_SupportedHeaderStyle(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var sut = LogCombined

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, logCombinedName, result)

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
	assert.Equal(t, DoNotLog, result)

	// tear down
	verifyAll(t)
}

func TestFromString_HappyPath(t *testing.T) {
	for key, value := range headerStyleNameMapping {
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
