package session

import (
	"github.com/google/uuid"
)

// func pointers for injection / testing: logger.go
var (
	uuidNew = uuid.New
)
