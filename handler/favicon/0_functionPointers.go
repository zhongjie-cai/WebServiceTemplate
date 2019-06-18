package favicon

import (
	"net/http"
)

// func pointers for injection / testing: favicon.go
var (
	httpServeFile = http.ServeFile
)
