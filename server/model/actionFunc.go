package model

import (
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

// ActionFunc defines the action function to be called for route processing logic
type ActionFunc func(
	session sessionModel.Session,
) (
	responseObject interface{},
	responseError error,
)
