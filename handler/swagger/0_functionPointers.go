package swagger

import (
	"net/http"
)

// func pointers for injection / testing: swagger.go
var (
	httpRedirect    = http.Redirect
	httpStripPrefix = http.StripPrefix
	httpFileServer  = http.FileServer
)
