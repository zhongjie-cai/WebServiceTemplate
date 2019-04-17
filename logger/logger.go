package logger

import (
	"time"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

type logEntry struct {
	Application string          `json:"application"`
	Version     string          `json:"version"`
	Timestamp   time.Time       `json:"timestamp"`
	Session     uuid.UUID       `json:"session"`
	Login       uuid.UUID       `json:"login"`
	Endpoint    string          `json:"endpoint"`
	Level       logtype.LogType `json:"level"`
	Category    string          `json:"category"`
	Subcategory string          `json:"subcategory"`
	Description string          `json:"description"`
}

func doLogging(sessionID uuid.UUID, logType logtype.LogType, category, subcategory, description string) {
	var session = sessionGet(
		sessionID,
	)
	if !session.AllowedLogType.HasFlag(logType) {
		return
	}
	var logEntryString = jsonutilMarshalIgnoreError(
		logEntry{
			Application: configAppName(),
			Version:     configAppVersion(),
			Timestamp:   timeutilGetTimeNowUTC(),
			Session:     sessionID,
			Login:       session.LoginID,
			Endpoint:    session.Endpoint,
			Level:       logType,
			Category:    category,
			Subcategory: subcategory,
			Description: description,
		},
	)
	fmtPrintln(
		logEntryString,
	)
}

// AppRoot logs the given message as AppRoot category
func AppRoot(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.AppRoot,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// APIEnter logs the given message as APIEnter category
func APIEnter(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.APIEnter,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// APIRequest logs the given message as APIRequest category
func APIRequest(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.APIRequest,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodEnter logs the given message as MethodEnter category
func MethodEnter(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.MethodEnter,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodParameter logs the given message as MethodParameter category
func MethodParameter(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.MethodParameter,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodLogic logs the given message as MethodLogic category
func MethodLogic(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.MethodLogic,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// DependencyCall logs the given message as DependencyCall category
func DependencyCall(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.DependencyCall,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// DependencyRequest logs the given message as DependencyRequest category
func DependencyRequest(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.DependencyRequest,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// DependencyResponse logs the given message as DependencyResponse category
func DependencyResponse(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.DependencyResponse,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// DependencyFinish logs the given message as DependencyFinish category
func DependencyFinish(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.DependencyFinish,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodReturn logs the given message as MethodReturn category
func MethodReturn(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.MethodReturn,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodExit logs the given message as MethodExit category
func MethodExit(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.MethodExit,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// APIResponse logs the given message as APIResponse category
func APIResponse(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.APIResponse,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// APIExit logs the given message as APIExit category
func APIExit(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	doLoggingFunc(
		sessionID,
		logtype.APIExit,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}
