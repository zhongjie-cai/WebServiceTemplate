package jsonutil

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
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

func TestTryUnmarshal_NoQuoteJSONEmpty(t *testing.T) {
	// arrange
	var dummyValue string
	var dummyDataTemplate int

	// mock
	createMock(t)

	// SUT + act
	var err = TryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Primitive(t *testing.T) {
	// arrange
	var dummyValue = rand.Int()
	var dummyValueString = strconv.Itoa(dummyValue)
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValueString), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = TryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Struct(t *testing.T) {
	// arrange
	var dummyValueString = "{\"foo\":\"bar\",\"test\":123}"
	var dummyDataTemplate struct {
		Foo  string `json:"foo"`
		Test int    `json:"test"`
	}

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		assert.Equal(t, []byte(dummyValueString), data)
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = TryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "bar", dummyDataTemplate.Foo)
	assert.Equal(t, 123, dummyDataTemplate.Test)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_WithQuoteJSONSuccess(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate string

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 2
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		if jsonUnmarshalCalled == 1 {
			assert.Equal(t, []byte(dummyValue), data)
		} else if jsonUnmarshalCalled == 2 {
			assert.Equal(t, []byte("\""+dummyValue+"\""), data)
		}
		return json.Unmarshal(data, v)
	}

	// SUT + act
	var err = TryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestTryUnmarshal_Failure(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate uuid.UUID
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	jsonUnmarshalExpected = 2
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		if jsonUnmarshalCalled == 1 {
			assert.Equal(t, []byte(dummyValue), data)
		} else if jsonUnmarshalCalled == 2 {
			assert.Equal(t, []byte("\""+dummyValue+"\""), data)
		}
		return json.Unmarshal(data, v)
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "Unable to unmarshal value [%v] into data template", format)
		assert.Equal(t, 1, len(a))
		assert.Equal(t, dummyValue, a[0])
		return dummyError
	}

	// SUT + act
	var err = TryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}
