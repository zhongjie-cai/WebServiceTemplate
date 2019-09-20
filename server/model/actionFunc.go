package model

import (
	"github.com/google/uuid"
)

// ActionFunc defines the action function to be called for route processing logic
type ActionFunc func(
	sessionID uuid.UUID,
) (
	responseObject interface{},
	responseError error,
)
