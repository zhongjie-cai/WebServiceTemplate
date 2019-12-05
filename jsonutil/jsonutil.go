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

// TryUnmarshal tries to unmarshal given value to dataTemplate
func TryUnmarshal(value string, dataTemplate interface{}) error {
	if value == "" {
		return nil
	}
	var noQuoteJSONError = jsonUnmarshal(
		[]byte(value),
		dataTemplate,
	)
	if noQuoteJSONError == nil {
		return nil
	}
	var withQuoteJSONError = jsonUnmarshal(
		[]byte("\""+value+"\""),
		dataTemplate,
	)
	if withQuoteJSONError == nil {
		return nil
	}
	return fmtErrorf(
		"Unable to unmarshal value [%v] into data template",
		value,
	)
}
