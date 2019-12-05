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

func tryUnmarshalPrimitiveTypes(value string, dataTemplate interface{}) bool {
	if value == "" {
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*string)(nil)) {
		(*(dataTemplate).(*string)) = value
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*bool)(nil)) {
		var parsedValue, parseError = strconvParseBool(
			stringsToLower(
				value,
			),
		)
		if parseError != nil {
			return false
		}
		(*(dataTemplate).(*bool)) = parsedValue
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*int)(nil)) {
		var parsedValue, parseError = strconvAtoi(value)
		if parseError != nil {
			return false
		}
		(*(dataTemplate).(*int)) = parsedValue
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*int64)(nil)) {
		var parsedValue, parseError = strconvParseInt(value, 0, 64)
		if parseError != nil {
			return false
		}
		(*(dataTemplate).(*int64)) = parsedValue
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*float64)(nil)) {
		var parsedValue, parseError = strconvParseFloat(value, 64)
		if parseError != nil {
			return false
		}
		(*(dataTemplate).(*float64)) = parsedValue
		return true
	}
	if reflectTypeOf(dataTemplate) == reflectTypeOf((*byte)(nil)) {
		var parsedValue, parseError = strconvParseUint(value, 0, 8)
		if parseError != nil {
			return false
		}
		(*(dataTemplate).(*byte)) = byte(parsedValue)
		return true
	}
	return false
}

// TryUnmarshal tries to unmarshal given value to dataTemplate
func TryUnmarshal(value string, dataTemplate interface{}) error {
	if tryUnmarshalPrimitiveTypesFunc(
		value,
		dataTemplate,
	) {
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
