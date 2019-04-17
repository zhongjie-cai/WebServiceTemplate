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
}

func verifyAll(t *testing.T) {
	jsonNewEncoder = json.NewEncoder
	if jsonNewEncoderExpected != jsonNewEncoderCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to jsonNewEncoder, expected %v, actual %v", jsonNewEncoderExpected, jsonNewEncoderCalled))
	}
	stringsTrimRight = strings.TrimRight
	if stringsTrimRightExpected != stringsTrimRightCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to stringsTrimRight, expected %v, actual %v", stringsTrimRightExpected, stringsTrimRightCalled))
	}
}
