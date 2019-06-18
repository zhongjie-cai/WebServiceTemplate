package logger

import (
	"time"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
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

// Initialize initiates and checks all application logging related function injections
func Initialize() error {
	if customization.LoggingFunc == nil {
		return apperrorWrapSimpleError(
			nil,
			"customization.LoggingFunc is not configured; fallback to default logging function.",
		)
	}
	return nil
}

func defaultLogging(
	session *session.Session,
	logType logtype.LogType,
	category,
	subcategory,
	description string,
) {
	var logEntryString = jsonutilMarshalIgnoreError(
		logEntry{
			Application: config.AppName(),
			Version:     config.AppVersion(),
			Timestamp:   timeutilGetTimeNowUTC(),
			Session:     session.ID,
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

func prepareLogging(
	sessionID uuid.UUID,
	logType logtype.LogType,
	category,
	subcategory,
	description string,
) {
	var session = sessionGet(
		sessionID,
	)
	if !session.AllowedLogType.HasFlag(logType) &&
		!config.IsLocalhost() {
		return
	}
	if customization.LoggingFunc == nil {
		defaultLoggingFunc(
			session,
			logType,
			category,
			subcategory,
			description,
		)
	} else {
		customization.LoggingFunc(
			session,
			logType,
			category,
			subcategory,
			description,
		)
	}
}

// AppRoot logs the given message as AppRoot category
func AppRoot(category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		uuid.Nil,
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
	prepareLoggingFunc(
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
