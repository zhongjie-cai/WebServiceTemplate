package apperror

import "errors"

// Code are codes returned by service indicating operation results; it is an integer value of the enum that corresponds to a given error
type Code int

// These are the integer values of the enum that corresponds to a given error
const (
	CodeGeneralFailure Code = iota
	CodeUnauthorized
	CodeInvalidOperation
	CodeBadRequest
	CodeNotFound
	CodeCircuitBreak
	CodeOperationLock
	CodeAccessForbidden
	CodeDataCorruption
	CodeNotImplemented
	codeMaxValue

	ErrorPrintFormat   string = "(%v)%v" // (Code)Message
	ErrorPointer       string = " -> "
	ErrorHolderLeft    string = "[ "
	ErrorHolderRight   string = " ]"
	ErrorSeparator     string = " | "
	ErrorMessageIndent string = "  "
)

// String translates the enum
func (code Code) String() string {
	var names = [...]string{
		"GeneralFailure",
		"Unauthorized",
		"InvalidOperation",
		"BadRequest",
		"NotFound",
		"CircuitBreak",
		"OperationLock",
		"AccessForbidden",
		"DataCorruption",
		"NotImplemented",
	}
	if code < 0 || code >= codeMaxValue {
		return "Unknown"
	}
	return names[code]
}

// AppError is the error wrapper interface for all WebServiceTemplate service generated errors
type AppError interface {
	error
	Code() Code
	InnerErrors() []error
	Messages() []string
}

type appError struct {
	error
	code        Code
	innerErrors []error
}

func (appError appError) Code() Code {
	return appError.code
}

func (appError appError) Error() string {
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

func (appError appError) InnerErrors() []error {
	return appError.innerErrors
}

func (appError appError) Messages() []string {
	var messages = []string{
		fmtSprintf(
			ErrorPrintFormat,
			appError.code,
			appError.error.Error(),
		),
	}
	for _, innerError := range appError.innerErrors {
		var typedError, isTyped = innerError.(AppError)
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

// GetGeneralFailureError creates a generic error based on GeneralFailure
func GetGeneralFailureError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeGeneralFailure,
		"An error occurred during execution",
	)
}

// GetUnauthorized creates an error related to Unauthorized
func GetUnauthorized(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeUnauthorized,
		"Access denied due to authorization error",
	)
}

// GetInvalidOperation creates an error related to InvalidOperation
func GetInvalidOperation(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeInvalidOperation,
		"Operation (method) not allowed",
	)
}

// GetBadRequestError creates an error related to BadRequest
func GetBadRequestError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeBadRequest,
		"Request URI or body is invalid",
	)
}

// GetNotFoundError creates an error related to NotFound
func GetNotFoundError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeNotFound,
		"Requested resource is not found in the storage",
	)
}

// GetCircuitBreakError creates an error related to CircuitBreak
func GetCircuitBreakError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeCircuitBreak,
		"Operation refused due to internal circuit break on correlation ID",
	)
}

// GetOperationLockError creates an error related to OperationLock
func GetOperationLockError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeOperationLock,
		"Operation refused due to mutex lock on correlation ID or trip ID",
	)
}

// GetAccessForbiddenError creates an error related to AccessForbidden
func GetAccessForbiddenError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeAccessForbidden,
		"Operation failed due to access forbidden",
	)
}

// GetDataCorruptionError creates an error related to DataCorruption
func GetDataCorruptionError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeDataCorruption,
		"Operation failed due to internal storage data corruption",
	)
}

// GetNotImplementedError creates an error related to NotImplemented
func GetNotImplementedError(innerError error) AppError {
	return wrapErrorFunc(
		innerError,
		CodeNotImplemented,
		"Operation failed due to internal business logic not implemented",
	)
}

// ConsolidateAllErrors adds up all errors in param list and generate a single error if the list is not empty
func ConsolidateAllErrors(
	baseErrorMessage string,
	allErrors ...error,
) AppError {
	if allErrors == nil || len(allErrors) == 0 {
		return nil
	}
	var allErrorMessages []string
	for _, err := range allErrors {
		if err != nil {
			var errorMessage = err.Error()
			if errorMessage == "" {
				allErrorMessages = append(
					allErrorMessages,
					"Unknown Error",
				)
			} else {
				allErrorMessages = append(
					allErrorMessages,
					err.Error(),
				)
			}
		}
	}
	if len(allErrorMessages) == 0 {
		return nil
	}
	var consilidatedErrorMessage = stringsJoin(
		allErrorMessages,
		ErrorSeparator,
	)
	return wrapSimpleErrorFunc(
		errors.New(consilidatedErrorMessage),
		baseErrorMessage,
	)
}

// WrapError wraps an inner error with given message as a new error with given error code
func WrapError(innerError error, errorCode Code, messageFormat string, parameters ...interface{}) AppError {
	var innerErrors []error
	if innerError != nil {
		innerErrors = []error{innerError}
	}
	return appError{
		fmtErrorf(messageFormat, parameters...),
		errorCode,
		innerErrors,
	}
}

// WrapSimpleError wraps an inner error with given message as a new general failure error
func WrapSimpleError(innerError error, messageFormat string, parameters ...interface{}) AppError {
	return wrapErrorFunc(
		innerError,
		CodeGeneralFailure,
		messageFormat,
		parameters...,
	)
}

// GetInnermostErrors finds the innermost error wrapped within the given error
func GetInnermostErrors(err error) []error {
	var innermostErrors []error
	var appError, ok = err.(AppError)
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
