package jsonutil

import (
	"bytes"
)

// These are constants to configure the JSON encoder
const (
	EscapeHTML bool = false
)

// MarshalIgnoreError returns the string representation of the given object; returns empty string in case of error
func MarshalIgnoreError(v interface{}) string {
	var buffer = &bytes.Buffer{}
	var encoder = jsonNewEncoder(buffer)
	encoder.SetEscapeHTML(EscapeHTML)
	encoder.Encode(v)
	var result = string(buffer.Bytes())
	return stringsTrimRight(result, "\n")
}
