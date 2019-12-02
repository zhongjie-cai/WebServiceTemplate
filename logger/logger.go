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
	session sessionModel.Session,
	logType logtype.LogType,
	logLevel loglevel.LogLevel,
	category,
	subcategory,
	description string,
) {
	if !session.IsLoggingAllowed(logType, logLevel) {
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
		nil,
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
func APIEnter(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func APIRequest(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func MethodEnter(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func MethodParameter(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func MethodLogic(session sessionModel.Session, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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

// NetworkCall logs the given message as NetworkCall category
func NetworkCall(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
		logtype.NetworkCall,
		loglevel.Info,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// NetworkRequest logs the given message as NetworkRequest category
func NetworkRequest(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
		logtype.NetworkRequest,
		loglevel.Info,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// NetworkResponse logs the given message as NetworkResponse category
func NetworkResponse(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
		logtype.NetworkResponse,
		loglevel.Info,
		category,
		subcategory,
		fmtSprintf(
			messageFormat,
			parameters...,
		),
	)
}

// NetworkFinish logs the given message as NetworkFinish category
func NetworkFinish(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
		logtype.NetworkFinish,
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
func MethodReturn(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func MethodExit(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func APIResponse(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
func APIExit(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	prepareLoggingFunc(
		session,
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
