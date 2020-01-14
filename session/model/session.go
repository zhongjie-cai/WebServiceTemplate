package model

import (
	"net/http"

	"github.com/google/uuid"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
)

var (
	// NilSession is for application root logging
	NilSession Session
)

// Session is the storage for the current HTTP request session, containing information needed for logging, monitoring, etc.
type Session interface {
	SessionMeta
	SessionHTTP
	SessionAttachment
	SessionLogging
	SessionNetwork
}

// SessionMeta is a subset of Session interface, containing only meta data related methods
type SessionMeta interface {
	// GetID returns the ID of this registered session object
	GetID() uuid.UUID

	// GetName returns the name registered to session object for given session ID
	GetName() string
}

// SessionHTTP is a subset of Session interface, containing only HTTP request & response related methods
type SessionHTTP interface {
	SessionHTTPRequest
	SessionHTTPResponse
}

// SessionHTTPRequest is a subset of Session interface, containing only HTTP request related methods
type SessionHTTPRequest interface {
	// GetRequest returns the HTTP request object from session object for given session ID
	GetRequest() *http.Request

	// GetRequestBody loads HTTP request body associated to session and unmarshals the content JSON to given data template
	GetRequestBody(dataTemplate interface{}) apperrorModel.AppError

	// GetRequestParameter loads HTTP request parameter associated to session for given name and unmarshals the content to given data template
	GetRequestParameter(name string, dataTemplate interface{}) apperrorModel.AppError

	// GetRequestQuery loads HTTP request single query string associated to session for given name and unmarshals the content to given data template
	GetRequestQuery(name string, dataTemplate interface{}) apperrorModel.AppError

	// GetRequestQueries loads HTTP request query strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
	GetRequestQueries(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError

	// GetRequestHeader loads HTTP request single header string associated to session for given name and unmarshals the content to given data template
	GetRequestHeader(name string, dataTemplate interface{}) apperrorModel.AppError

	// GetRequestHeaders loads HTTP request header strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
	GetRequestHeaders(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError
}

// SessionHTTPResponse is a subset of SessionHTTP interface, containing only HTTP response related methods
type SessionHTTPResponse interface {
	// GetResponseWriter returns the HTTP response writer object from session object for given session ID
	GetResponseWriter() http.ResponseWriter
}

// SessionAttachment is a subset of Session interface, containing only attachment related methods
type SessionAttachment interface {
	// Attach attaches any value object into the given session associated to the session ID
	Attach(name string, value interface{}) bool

	// Detach detaches any value object from the given session associated to the session ID
	Detach(name string) bool

	// GetRawAttachment retrieves any value object from the given session associated to the session ID and returns the raw interface (consumer needs to manually cast, but works for struct with private fields)
	GetRawAttachment(name string) (interface{}, bool)

	// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template (only works for structs with exported fields)
	GetAttachment(name string, dataTemplate interface{}) bool
}

// SessionLogging is a subset of Session interface, containing only logging related methods
type SessionLogging interface {
	// IsLoggingAllowed checks the passed in log type and level and determines whether they match the session log criteria or not
	IsLoggingAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool

	// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
	LogMethodEnter()

	// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
	LogMethodParameter(parameters ...interface{})

	// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
	LogMethodLogic(logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})

	// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
	LogMethodReturn(returns ...interface{})

	// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
	LogMethodExit()
}

// SessionNetwork is a subset of Session interface, containing only network related methods
type SessionNetwork interface {
	// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
	CreateNetworkRequest(method string, url string, payload string, header map[string]string) networkModel.NetworkRequest
}
