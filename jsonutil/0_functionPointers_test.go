package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	jsonNewEncoderExpected   int
	jsonNewEncoderCalled     int
	stringsTrimRightExpected int
	stringsTrimRightCalled   int
	jsonUnmarshalExpected    int
	jsonUnmarshalCalled      int
	fmtErrorfExpected        int
	fmtErrorfCalled          int
)

func createMock(t *testing.T) {
	jsonNewEncoderExpected = 0
	jsonNewEncoderCalled = 0
	jsonNewEncoder = func(w io.Writer) *json.Encoder {
		jsonNewEncoderCalled++
		return nil
	}
	stringsTrimRightExpected = 0
	stringsTrimRightCalled = 0
	stringsTrimRight = func(s string, cutset string) string {
		stringsTrimRightCalled++
		return ""
	}
	jsonUnmarshalExpected = 0
	jsonUnmarshalCalled = 0
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return nil
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	jsonNewEncoder = json.NewEncoder
	assert.Equal(t, jsonNewEncoderExpected, jsonNewEncoderCalled, "Unexpected number of calls to jsonNewEncoder")
	stringsTrimRight = strings.TrimRight
	assert.Equal(t, stringsTrimRightExpected, stringsTrimRightCalled, "Unexpected number of calls to stringsTrimRight")
	jsonUnmarshal = json.Unmarshal
	assert.Equal(t, jsonUnmarshalExpected, jsonUnmarshalCalled, "Unexpected number of calls to jsonUnmarshal")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
}
