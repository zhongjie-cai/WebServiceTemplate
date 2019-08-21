package model

import (
	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// ActionFunc defines the action function to be called for route processing logic
type ActionFunc func(
	sessionID uuid.UUID,
	requestBody string,
	parameters map[string]string,
) (
	responseObject interface{},
	responseError apperror.AppError,
)
