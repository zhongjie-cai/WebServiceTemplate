package model

// AppError is the error wrapper interface for all WebServiceTemplate service generated errors
type AppError interface {
	error
	Code() string
	HTTPStatusCode() int
	InnerErrors() []error
	Messages() []string
	Append(innerErrors ...error)
}
