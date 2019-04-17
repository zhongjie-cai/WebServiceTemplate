package jsonutil

import (
	"encoding/json"
	"strings"
)

// func pointers for injection / testing: logger.go
var (
	jsonNewEncoder   = json.NewEncoder
	stringsTrimRight = strings.TrimRight
)
