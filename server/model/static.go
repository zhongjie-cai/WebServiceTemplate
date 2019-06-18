package model

import "net/http"

// Static holds the registration information of a static content hosting
type Static struct {
	Name       string
	PathPrefix string
	Handler    http.Handler
}
