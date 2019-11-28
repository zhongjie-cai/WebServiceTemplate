package model

// AppError is the error wrapper interface for all WebServiceTemplate service generated errors
type AppError interface {
	// Golang internal error interface
	error
	// Code returns the string representation of the error code enum
	Code() string
	// HTTPStatusCode returns the corresponding HTTP status code mapped to the error code value
	HTTPStatusCode() int
	// InnerErrors returns the inner errors array
	InnerErrors() []error
	// Messages returns the string representations of all inner errors
	Messages() []string
	// ExtraData returns the serialized map of all attached extra data
	ExtraData() map[string]string
	// Append adds the given list of inner errors into the current app error object
	Append(innerErrors ...error)
	// Attach adds the given value to the current app error's extra data map by given name
	Attach(name string, value interface{})
}
