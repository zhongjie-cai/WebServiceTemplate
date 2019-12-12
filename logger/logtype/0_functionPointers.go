package logtype

import (
	"sort"
	"strings"
)

// func pointers for injection / testing: logType.go
var (
	sortStrings  = sort.Strings
	stringsJoin  = strings.Join
	stringsSplit = strings.Split
)
