package jsonutil

import (
	"encoding/json"
	"fmt"
	"strings"
)

// func pointers for injection / testing: logger.go
var (
	jsonNewEncoder   = json.NewEncoder
	stringsTrimRight = strings.TrimRight
	jsonUnmarshal    = json.Unmarshal
	fmtErrorf        = fmt.Errorf
)
