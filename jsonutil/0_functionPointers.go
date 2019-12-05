package jsonutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// func pointers for injection / testing: logger.go
var (
	jsonNewEncoder                 = json.NewEncoder
	stringsTrimRight               = strings.TrimRight
	jsonUnmarshal                  = json.Unmarshal
	fmtErrorf                      = fmt.Errorf
	reflectTypeOf                  = reflect.TypeOf
	stringsToLower                 = strings.ToLower
	strconvAtoi                    = strconv.Atoi
	strconvParseBool               = strconv.ParseBool
	strconvParseInt                = strconv.ParseInt
	strconvParseFloat              = strconv.ParseFloat
	strconvParseUint               = strconv.ParseUint
	tryUnmarshalPrimitiveTypesFunc = tryUnmarshalPrimitiveTypes
)
