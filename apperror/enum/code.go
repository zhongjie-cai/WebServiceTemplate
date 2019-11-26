package enum

import "net/http"

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

// HTTPStatusCode translates the built-in error Code to corresponding HTTP status code
func (code Code) HTTPStatusCode() int {
	var statusCode int
	switch code {
	case CodeGeneralFailure:
		statusCode = http.StatusInternalServerError
	case CodeUnauthorized:
		statusCode = http.StatusUnauthorized
	case CodeInvalidOperation:
		statusCode = http.StatusMethodNotAllowed
	case CodeBadRequest:
		statusCode = http.StatusBadRequest
	case CodeNotFound:
		statusCode = http.StatusNotFound
	case CodeCircuitBreak:
		statusCode = http.StatusForbidden
	case CodeOperationLock:
		statusCode = http.StatusLocked
	case CodeAccessForbidden:
		statusCode = http.StatusForbidden
	case CodeDataCorruption:
		statusCode = http.StatusConflict
	case CodeNotImplemented:
		statusCode = http.StatusNotImplemented
	default:
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}
