package model

import (
	"github.com/google/uuid"
)

// Route holds the registration information of a dynamic route hosting
type Route struct {
	Endpoint   string
	Method     string
	Path       string
	Parameters map[string]Parameter
	ActionFunc func(uuid.UUID, string)
}
