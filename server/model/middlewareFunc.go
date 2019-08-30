package model

import (
	"github.com/gorilla/mux"
)

// MiddlewareFunc warps around mux.MiddlewareFunc, which receives an http.Handler and returns another http.Handler.
// Typically, the returned handler is a closure which does something with the http.ResponseWriter and http.Request passed
// to it, and then calls the handler passed as parameter to the MiddlewareFunc.
type MiddlewareFunc mux.MiddlewareFunc
