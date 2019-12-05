package logtype

import "strings"

// func pointers for injection / testing: logType.go
var (
	stringsJoin  = strings.Join
	stringsSplit = strings.Split
)
