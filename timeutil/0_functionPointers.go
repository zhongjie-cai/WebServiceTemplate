package timeutil

import (
	"time"
)

// func pointers for injection / testing: timeutil.go
var (
	timeNow = time.Now
)
