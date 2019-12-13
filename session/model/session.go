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
	// GetID returns the ID of this registered session object
	GetID() uuid.UUID

	// GetName returns the name registered to session object for given session ID
	GetName() string

	// GetRequest returns the HTTP request object from session object for given session ID
	GetRequest() *http.Request

	// GetResponseWriter returns the HTTP response writer object from session object for given session ID
	GetResponseWriter() http.ResponseWriter

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

	// Attach attaches any value object into the given session associated to the session ID
	Attach(name string, value interface{}) bool

	// Detach detaches any value object from the given session associated to the session ID
	Detach(name string) bool

	// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template
	GetAttachment(name string, dataTemplate interface{}) bool

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

	// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
	CreateNetworkRequest(method string, url string, payload string, header map[string]string) networkModel.NetworkRequest
}
