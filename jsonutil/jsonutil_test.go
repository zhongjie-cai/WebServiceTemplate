package jsonutil

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalIgnoreError_Empty(t *testing.T) {
	// arrange
	var dummyObject *struct {
		Foo  string
		Test int
	}
	var expectedResult = "null"

	// mock
	createMock(t)

	// expect
	jsonNewEncoderExpected = 1
	jsonNewEncoder = func(w io.Writer) *json.Encoder {
		jsonNewEncoderCalled++
		return json.NewEncoder(w)
	}
	stringsTrimRightExpected = 1
	stringsTrimRight = func(s string, cutset string) string {
		stringsTrimRightCalled++
		return strings.TrimRight(s, cutset)
	}

	// SUT + act
	var result = MarshalIgnoreError(
		dummyObject,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}

func TestMarshalIgnoreError_Success(t *testing.T) {
	// arrange
	var dummyObject = struct {
		Foo  string
		Test int
	}{
		"<bar />",
		123,
	}
	var expectedResult = "{\"Foo\":\"<bar />\",\"Test\":123}"

	// mock
	createMock(t)

	// expect
	jsonNewEncoderExpected = 1
	jsonNewEncoder = func(w io.Writer) *json.Encoder {
		jsonNewEncoderCalled++
		return json.NewEncoder(w)
	}
	stringsTrimRightExpected = 1
	stringsTrimRight = func(s string, cutset string) string {
		stringsTrimRightCalled++
		return strings.TrimRight(s, cutset)
	}

	// SUT + act
	var result = MarshalIgnoreError(
		dummyObject,
	)

	// assert
	assert.Equal(t, expectedResult, result)

	// verify
	verifyAll(t)
}
