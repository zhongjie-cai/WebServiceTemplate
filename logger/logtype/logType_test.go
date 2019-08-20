package logtype

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_AppRoot(t *testing.T) {
	// arrange
	var appRootValue = 0

	// mock
	createMock(t)

	// SUT
	var sut = LogType(appRootValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, AppRoot, sut)
	assert.Equal(t, AppRootName, result)

	// verify
	verifyAll(t)
}

func TestString_NonSupportedLogCategories(t *testing.T) {
	// arrange
	var unsupportedValue = 1 << 31

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var sut = LogType(unsupportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestString_SingleSupportedLogType(t *testing.T) {
	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var sut = MethodLogic

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, MethodLogicName, result)

	// verify
	verifyAll(t)
}

func TestString_MultipleSupportedLogCategories(t *testing.T) {
	// arrange
	var supportedValue = APIEnter | APIRequest | MethodLogic | APIResponse | APIExit

	// mock
	createMock(t)

	// expect
	stringsJoinExpected = 1
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return strings.Join(a, sep)
	}

	// SUT
	var sut = LogType(supportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, GeneralLogging, sut)
	assert.True(t, strings.Contains(result, APIEnterName))
	assert.True(t, strings.Contains(result, APIRequestName))
	assert.True(t, strings.Contains(result, MethodLogicName))
	assert.True(t, strings.Contains(result, APIResponseName))
	assert.True(t, strings.Contains(result, APIExitName))

	// verify
	verifyAll(t)
}

func TestHasFlag_FlagMatch_AppRoot(t *testing.T) {
	// arrange
	var flag = AppRoot

	// mock
	createMock(t)

	// SUT
	var sut = AppRoot

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestHasFlag_FlagNoMatch_AppRoot(t *testing.T) {
	// arrange
	var flag = AppRoot

	// mock
	createMock(t)

	// SUT
	var sut = APIEnter | APIExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestHasFlag_FlagMatch_NotAppRoot(t *testing.T) {
	// arrange
	var flag = MethodLogic

	// mock
	createMock(t)

	// SUT
	var sut = APIEnter | MethodLogic | APIExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestHasFlag_FlagNoMatch_NotAppRoot(t *testing.T) {
	// arrange
	var flag = MethodLogic

	// mock
	createMock(t)

	// SUT
	var sut = APIEnter | APIExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.False(t, result)

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
	assert.Equal(t, AppRoot, result)

	// tear down
	verifyAll(t)
}

func TestFromString_AppRoot(t *testing.T) {
	// arrange
	var dummyValue = AppRootName

	// mock
	createMock(t)

	// SUT + act
	var result = FromString(dummyValue)

	// assert
	assert.Equal(t, AppRoot, result)

	// tear down
	verifyAll(t)
}

func TestFromString_APIEnter(t *testing.T) {
	for key, value := range logTypeNameMapping {
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
