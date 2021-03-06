package apperror

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

// These are print formatting related constants
const (
	ErrorPrintFormat   string = "(%v)%v" // (Code)Message
	ErrorPointer       string = " -> "
	ErrorHolderLeft    string = "[ "
	ErrorHolderRight   string = " ]"
	ErrorSeparator     string = " | "
	ErrorMessageIndent string = "  "
)

type appError struct {
	error
	code        enum.Code
	innerErrors []error
	extraData   map[string]interface{}
}

func (appError *appError) Code() string {
	if customization.AppErrors != nil {
		var codeNames, _ = customization.AppErrors()
		var codeName, found = codeNames[appError.code]
		if found {
			return codeName
		}
	}
	return appError.code.String()
}

func (appError *appError) HTTPStatusCode() int {
	if customization.AppErrors != nil {
		var _, httpStatusCodes = customization.AppErrors()
		var statusCode, found = httpStatusCodes[appError.code]
		if found {
			return statusCode
		}
	}
	return appError.code.HTTPStatusCode()
}

func (appError *appError) Error() string {
	var fullMessage = fmtSprintf(
		ErrorPrintFormat,
		appError.code,
		appError.error.Error(),
	)
	var innerMessages []string
	for _, innerError := range appError.innerErrors {
		innerMessages = append(
			innerMessages,
			ErrorHolderLeft+innerError.Error()+ErrorHolderRight,
		)
	}
	var innerMessage = stringsJoin(innerMessages, ErrorSeparator)
	if innerMessage != "" {
		fullMessage += ErrorPointer + ErrorHolderLeft + innerMessage + ErrorHolderRight
	}
	return fullMessage
}

func (appError *appError) InnerErrors() []error {
	return appError.innerErrors
}

func (appError *appError) Messages() []string {
	var messages = []string{
		fmtSprintf(
			ErrorPrintFormat,
			appError.Code(),
			appError.error.Error(),
		),
	}
	for _, innerError := range appError.innerErrors {
		var typedError, isTyped = innerError.(model.AppError)
		if isTyped {
			var innerMessages = typedError.Messages()
			for _, innerMessage := range innerMessages {
				messages = append(
					messages,
					ErrorMessageIndent+innerMessage,
				)
			}
		} else {
			messages = append(
				messages,
				ErrorMessageIndent+innerError.Error(),
			)
		}
	}
	return messages
}

func (appError *appError) ExtraData() map[string]string {
	var result = map[string]string{}
	for name, value := range appError.extraData {
		var data = jsonutilMarshalIgnoreError(value)
		result[name] = data
	}
	return result
}

func (appError *appError) Append(innerErrors ...error) {
	var cleanedInnerErrors = cleanupInnerErrorsFunc(innerErrors)
	if len(cleanedInnerErrors) == 0 {
		return
	}
	appError.innerErrors = append(
		appError.innerErrors,
		cleanedInnerErrors...,
	)
}

func (appError *appError) Attach(name string, value interface{}) {
	if appError.extraData == nil {
		appError.extraData = map[string]interface{}{}
	}
	appError.extraData[name] = value
}

// Initialize checks and validates the customization of AppErrors from consumer code
func Initialize() error {
	if customization.AppErrors == nil {
		return nil
	}
	var innerErrors = []error{}
	var codeNames, httpStatusCodes = customization.AppErrors()
	for code, codeName := range codeNames {
		if code < enum.CodeReservedCount {
			innerErrors = append(
				innerErrors,
				fmtErrorf(
					"AppError code [%v] configured for code name [%v] is conflicting with reserved error codes",
					code,
					codeName,
				),
			)
		}
	}
	for code, httpStatusCode := range httpStatusCodes {
		if code < enum.CodeReservedCount {
			innerErrors = append(
				innerErrors,
				fmtErrorf(
					"AppError code [%v] configured for HTTP status code [%v] is conflicting with reserved error codes",
					code,
					httpStatusCode,
				),
			)
		}
	}
	return wrapSimpleErrorFunc(
		innerErrors,
		"Failed to initialize AppError customization",
	)
}

// GetGeneralFailureError creates a generic error based on GeneralFailure
func GetGeneralFailureError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeGeneralFailure,
		"An error occurred during execution",
	)
}

// GetUnauthorized creates an error related to Unauthorized
func GetUnauthorized(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeUnauthorized,
		"Access denied due to authorization error",
	)
}

// GetInvalidOperation creates an error related to InvalidOperation
func GetInvalidOperation(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeInvalidOperation,
		"Operation (method) not allowed",
	)
}

// GetBadRequestError creates an error related to BadRequest
func GetBadRequestError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeBadRequest,
		"Request URI or body is invalid",
	)
}

// GetNotFoundError creates an error related to NotFound
func GetNotFoundError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeNotFound,
		"Requested resource is not found in the storage",
	)
}

// GetCircuitBreakError creates an error related to CircuitBreak
func GetCircuitBreakError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeCircuitBreak,
		"Operation refused due to internal circuit break on correlation ID",
	)
}

// GetOperationLockError creates an error related to OperationLock
func GetOperationLockError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeOperationLock,
		"Operation refused due to mutex lock on correlation ID or trip ID",
	)
}

// GetAccessForbiddenError creates an error related to AccessForbidden
func GetAccessForbiddenError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeAccessForbidden,
		"Operation failed due to access forbidden",
	)
}

// GetDataCorruptionError creates an error related to DataCorruption
func GetDataCorruptionError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeDataCorruption,
		"Operation failed due to internal storage data corruption",
	)
}

// GetNotImplementedError creates an error related to NotImplemented
func GetNotImplementedError(innerErrors ...error) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeNotImplemented,
		"Operation failed due to internal business logic not implemented",
	)
}

// GetCustomError creates a customized error with given code and formatted message
func GetCustomError(errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
	return &appError{
		fmtErrorf(messageFormat, parameters...),
		errorCode,
		nil,
		nil,
	}
}

func cleanupInnerErrors(innerErrors []error) []error {
	var cleanedInnerErrors = []error{}
	for _, innerError := range innerErrors {
		if innerError != nil {
			cleanedInnerErrors = append(
				cleanedInnerErrors,
				innerError,
			)
		}
	}
	return cleanedInnerErrors
}

// WrapError wraps an inner error with given message as a new error with given error code
func WrapError(innerErrors []error, errorCode enum.Code, messageFormat string, parameters ...interface{}) model.AppError {
	var cleanedInnerErrors = cleanupInnerErrorsFunc(
		innerErrors,
	)
	if len(cleanedInnerErrors) == 0 {
		return nil
	}
	return &appError{
		fmtErrorf(messageFormat, parameters...),
		errorCode,
		cleanedInnerErrors,
		nil,
	}
}

// WrapSimpleError wraps an inner error with given message as a new general failure error
func WrapSimpleError(innerErrors []error, messageFormat string, parameters ...interface{}) model.AppError {
	return wrapErrorFunc(
		innerErrors,
		enum.CodeGeneralFailure,
		messageFormat,
		parameters...,
	)
}

// GetInnermostErrors finds the innermost error wrapped within the given error
func GetInnermostErrors(err error) []error {
	var innermostErrors []error
	var appError, ok = err.(model.AppError)
	if ok {
		for _, innerError := range appError.InnerErrors() {
			innermostErrors = append(
				innermostErrors,
				GetInnermostErrors(innerError)...,
			)
		}
	} else {
		innermostErrors = append(
			innermostErrors,
			err,
		)
	}
	return innermostErrors
}
