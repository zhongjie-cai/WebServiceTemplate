package model

import (
	"net/http"

	"github.com/google/uuid"
)

// Route holds the registration information of a dynamic route hosting
type Route struct {
	Endpoint   string
	Method     string
	Path       string
	Parameters map[string]Parameter
	ActionFunc func(http.ResponseWriter, *http.Request, uuid.UUID)
}
