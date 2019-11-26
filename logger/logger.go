package logger

import (
	"time"

	"github.com/google/uuid"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

type logEntry struct {
	Application string            `json:"application"`
	Version     string            `json:"version"`
	Timestamp   time.Time         `json:"timestamp"`
	Session     uuid.UUID         `json:"session"`
	Name        string            `json:"name"`
	Type        logtype.LogType   `json:"type"`
	Level       loglevel.LogLevel `json:"level"`
	Category    string            `json:"category"`
	Subcategory string            `json:"subcategory"`
	Description string            `json:"description"`
}

// Initialize initiates and checks all application logging related function injections
func Initialize() error {
	if customization.LoggingFunc == nil {
		return apperrorGetCustomError(
			apperrorEnum.CodeGeneralFailure,
			"customization.LoggingFunc is not configured; fallback to default logging function.",
		)
	}
	return nil
}

func defaultLogging(
	session sessionModel.Session,
	logType logtype.LogType,
	logLevel loglevel.LogLevel,
	category,
	subcategory,
	description string,
) {
	var logEntryString = jsonutilMarshalIgnoreError(
		logEntry{
			Application: config.AppName(),
			Version:     config.AppVersion(),
			Timestamp:   timeutilGetTimeNowUTC(),
			Session:     session.GetID(),
			Name:        session.GetName(),
			Type:        logType,
			Level:       logLevel,
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
	logLevel loglevel.LogLevel,
	category,
	subcategory,
	description string,
) {
	var session = sessionGet(
		sessionID,
	)
	if !session.IsLogAllowed(logType, logLevel) {
		return
	}
	if customization.LoggingFunc == nil {
		defaultLoggingFunc(
			session,
			logType,
			logLevel,
			category,
			subcategory,
			description,
		)
	} else {
		customization.LoggingFunc(
			session,
			logType,
			logLevel,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// MethodLogic logs the given message as MethodLogic category
func MethodLogic(sessionID uuid.UUID, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		sessionID,
		logtype.MethodLogic,
		logLevel,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
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
		loglevel.Info,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}
