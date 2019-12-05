package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	jsonNewEncoderExpected                 int
	jsonNewEncoderCalled                   int
	stringsTrimRightExpected               int
	stringsTrimRightCalled                 int
	jsonUnmarshalExpected                  int
	jsonUnmarshalCalled                    int
	fmtErrorfExpected                      int
	fmtErrorfCalled                        int
	reflectTypeOfExpected                  int
	reflectTypeOfCalled                    int
	stringsToLowerExpected                 int
	stringsToLowerCalled                   int
	strconvAtoiExpected                    int
	strconvAtoiCalled                      int
	strconvParseBoolExpected               int
	strconvParseBoolCalled                 int
	strconvParseIntExpected                int
	strconvParseIntCalled                  int
	strconvParseFloatExpected              int
	strconvParseFloatCalled                int
	strconvParseUintExpected               int
	strconvParseUintCalled                 int
	tryUnmarshalPrimitiveTypesFuncExpected int
	tryUnmarshalPrimitiveTypesFuncCalled   int
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
	reflectTypeOfExpected = 0
	reflectTypeOfCalled = 0
	reflectTypeOf = func(i interface{}) reflect.Type {
		reflectTypeOfCalled++
		return nil
	}
	stringsToLowerExpected = 0
	stringsToLowerCalled = 0
	stringsToLower = func(s string) string {
		stringsToLowerCalled++
		return ""
	}
	strconvAtoiExpected = 0
	strconvAtoiCalled = 0
	strconvAtoi = func(s string) (int, error) {
		strconvAtoiCalled++
		return 0, nil
	}
	strconvParseBoolExpected = 0
	strconvParseBoolCalled = 0
	strconvParseBool = func(str string) (bool, error) {
		strconvParseBoolCalled++
		return false, nil
	}
	strconvParseIntExpected = 0
	strconvParseIntCalled = 0
	strconvParseInt = func(s string, base int, bitSize int) (int64, error) {
		strconvParseIntCalled++
		return 0, nil
	}
	strconvParseFloatExpected = 0
	strconvParseFloatCalled = 0
	strconvParseFloat = func(s string, bitSize int) (float64, error) {
		strconvParseFloatCalled++
		return 0, nil
	}
	strconvParseUintExpected = 0
	strconvParseUintCalled = 0
	strconvParseUint = func(s string, base int, bitSize int) (uint64, error) {
		strconvParseUintCalled++
		return 0, nil
	}
	tryUnmarshalPrimitiveTypesFuncExpected = 0
	tryUnmarshalPrimitiveTypesFuncCalled = 0
	tryUnmarshalPrimitiveTypesFunc = func(value string, dataTemplate interface{}) bool {
		tryUnmarshalPrimitiveTypesFuncCalled++
		return false
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
	reflectTypeOf = reflect.TypeOf
	assert.Equal(t, reflectTypeOfExpected, reflectTypeOfCalled, "Unexpected number of calls to reflectTypeOf")
	stringsToLower = strings.ToLower
	assert.Equal(t, stringsToLowerExpected, stringsToLowerCalled, "Unexpected number of calls to stringsToLower")
	strconvAtoi = strconv.Atoi
	assert.Equal(t, strconvAtoiExpected, strconvAtoiCalled, "Unexpected number of calls to strconvAtoi")
	strconvParseBool = strconv.ParseBool
	assert.Equal(t, strconvParseBoolExpected, strconvParseBoolCalled, "Unexpected number of calls to strconvParseBool")
	strconvParseInt = strconv.ParseInt
	assert.Equal(t, strconvParseIntExpected, strconvParseIntCalled, "Unexpected number of calls to strconvParseInt")
	strconvParseFloat = strconv.ParseFloat
	assert.Equal(t, strconvParseFloatExpected, strconvParseFloatCalled, "Unexpected number of calls to strconvParseFloat")
	strconvParseUint = strconv.ParseUint
	assert.Equal(t, strconvParseUintExpected, strconvParseUintCalled, "Unexpected number of calls to strconvParseUint")
	tryUnmarshalPrimitiveTypesFunc = tryUnmarshalPrimitiveTypes
	assert.Equal(t, tryUnmarshalPrimitiveTypesFuncExpected, tryUnmarshalPrimitiveTypesFuncCalled, "Unexpected number of calls to tryUnmarshalPrimitiveTypesFunc")
}
