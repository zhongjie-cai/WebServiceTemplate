package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	cache "github.com/patrickmn/go-cache"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

var (
	sessionCache          = cache.New(15*time.Minute, 30*time.Minute)
	defaultRequest        = &http.Request{}
	defaultResponseWriter = &nilResponseWriter{}
	defaultName           = "AppRoot"
	defaultSession        *session
)

// Initialize ensures the NilSession is set for the AppRoot function
func Initialize() {
	model.NilSession = defaultSession
}

type session struct {
	ID              uuid.UUID
	Name            string
	AllowedLogType  logtype.LogType
	AllowedLogLevel loglevel.LogLevel
	Request         *http.Request
	ResponseWriter  http.ResponseWriter
	attachment      map[string]interface{}
}

// GetID returns the ID of this registered session object
func (session *session) GetID() uuid.UUID {
	if session == nil {
		return uuid.Nil
	}
	return session.ID
}

// GetName returns the name registered to session object for given session ID
func (session *session) GetName() string {
	if session == nil {
		return defaultName
	}
	return session.Name
}

// GetRequest returns the HTTP request object from session object for given session ID
func (session *session) GetRequest() *http.Request {
	if session == nil {
		return defaultRequest
	}
	return session.Request
}

// GetResponseWriter returns the HTTP response writer object from session object for given session ID
func (session *session) GetResponseWriter() http.ResponseWriter {
	if session == nil {
		return defaultResponseWriter
	}
	return session.ResponseWriter
}

// GetRequestBody loads HTTP request body associated to session and unmarshals the content JSON to given data template
func (session *session) GetRequestBody(dataTemplate interface{}) apperrorModel.AppError {
	var httpRequest = session.GetRequest()
	var requestBody = requestGetRequestBody(
		httpRequest,
	)
	if requestBody == "" {
		return apperrorGetBadRequestError(
			fmtErrorf(
				"The request body is empty",
			),
		)
	}
	loggerAPIRequest(
		session,
		"Body",
		"",
		requestBody,
	)
	return apperrorGetBadRequestError(
		jsonutilTryUnmarshal(
			requestBody,
			dataTemplate,
		),
	)
}

// GetRequestParameter loads HTTP request parameter associated to session for given name and unmarshals the content to given data template
func (session *session) GetRequestParameter(name string, dataTemplate interface{}) apperrorModel.AppError {
	var httpRequest = session.GetRequest()
	var parameters = muxVars(
		httpRequest,
	)
	var value, found = parameters[name]
	if !found {
		return apperrorGetBadRequestError(
			fmtErrorf(
				"The expected parameter [%v] is not found in request",
				name,
			),
		)
	}
	loggerAPIRequest(
		session,
		"Parameter",
		name,
		value,
	)
	return apperrorGetBadRequestError(
		jsonutilTryUnmarshal(
			value,
			dataTemplate,
		),
	)
}

func getAllQueries(session *session, name string) []string {
	var httpRequest = session.GetRequest()
	var queries, found = httpRequest.URL.Query()[name]
	if !found {
		return nil
	}
	return queries
}

// GetRequestQuery loads HTTP request single query string associated to session for given name and unmarshals the content to given data template
func (session *session) GetRequestQuery(name string, dataTemplate interface{}) apperrorModel.AppError {
	var queries = getAllQueriesFunc(
		session,
		name,
	)
	if len(queries) == 0 {
		return apperrorGetBadRequestError(
			fmtErrorf(
				"The expected query string [%v] is not found in request",
				name,
			),
		)
	}
	var value = queries[0]
	loggerAPIRequest(
		session,
		"Query",
		name,
		value,
	)
	return apperrorGetBadRequestError(
		jsonutilTryUnmarshal(
			value,
			dataTemplate,
		),
	)
}

// GetRequestQueries loads HTTP request query strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func (session *session) GetRequestQueries(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	var queries = getAllQueriesFunc(
		session,
		name,
	)
	var unmarshalErrors = []error{}
	for _, query := range queries {
		loggerAPIRequest(
			session,
			"Query",
			name,
			query,
		)
		var unmarshalError = jsonutilTryUnmarshal(
			query,
			dataTemplate,
		)
		if unmarshalError != nil {
			unmarshalErrors = append(
				unmarshalErrors,
				unmarshalError,
			)
		} else {
			fillCallback()
		}
	}
	return apperrorGetBadRequestError(
		unmarshalErrors...,
	)
}

func getAllHeaders(session *session, name string) []string {
	var httpRequest = session.GetRequest()
	var canonicalName = textprotoCanonicalMIMEHeaderKey(name)
	var headers, found = httpRequest.Header[canonicalName]
	if !found {
		return nil
	}
	return headers
}

// GetRequestHeader loads HTTP request single header string associated to session for given name and unmarshals the content to given data template
func (session *session) GetRequestHeader(name string, dataTemplate interface{}) apperrorModel.AppError {
	var headers = getAllHeadersFunc(
		session,
		name,
	)
	if len(headers) == 0 {
		return apperrorGetBadRequestError(
			fmtErrorf(
				"The expected header string [%v] is not found in request",
				name,
			),
		)
	}
	var value = headers[0]
	loggerAPIRequest(
		session,
		"Header",
		name,
		value,
	)
	return apperrorGetBadRequestError(
		jsonutilTryUnmarshal(
			value,
			dataTemplate,
		),
	)
}

// GetRequestHeaders loads HTTP request header strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func (session *session) GetRequestHeaders(name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	var headers = getAllHeadersFunc(
		session,
		name,
	)
	var unmarshalErrors = []error{}
	for _, header := range headers {
		loggerAPIRequest(
			session,
			"Header",
			name,
			header,
		)
		var unmarshalError = jsonutilTryUnmarshal(
			header,
			dataTemplate,
		)
		if unmarshalError != nil {
			unmarshalErrors = append(
				unmarshalErrors,
				unmarshalError,
			)
		} else {
			fillCallback()
		}
	}
	return apperrorGetBadRequestError(
		unmarshalErrors...,
	)
}

// Attach attaches any value object into the given session associated to the session ID
func (session *session) Attach(name string, value interface{}) bool {
	if session == nil {
		return false
	}
	if session.attachment == nil {
		session.attachment = map[string]interface{}{}
	}
	session.attachment[name] = value
	return true
}

// Detach detaches any value object from the given session associated to the session ID
func (session *session) Detach(name string) bool {
	if session == nil {
		return false
	}
	if session.attachment != nil {
		delete(session.attachment, name)
	}
	return true
}

// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template
func (session *session) GetAttachment(name string, dataTemplate interface{}) bool {
	if session == nil {
		return false
	}
	var attachment, found = session.attachment[name]
	if !found {
		return false
	}
	var bytes, marshalError = jsonMarshal(attachment)
	if marshalError != nil {
		return false
	}
	var unmarshalError = jsonUnmarshal(
		bytes,
		dataTemplate,
	)
	return unmarshalError == nil
}

func isLoggingTypeMatch(session *session, logType logtype.LogType) bool {
	var allowedLogType logtype.LogType
	if session == nil {
		allowedLogType = config.DefaultAllowedLogType()
	} else {
		allowedLogType = session.AllowedLogType
	}
	return allowedLogType.HasFlag(logType)
}

func isLoggingLevelMatch(session *session, logLevel loglevel.LogLevel) bool {
	var allowedLogLevel loglevel.LogLevel
	if session == nil {
		allowedLogLevel = config.DefaultAllowedLogLevel()
	} else {
		allowedLogLevel = session.AllowedLogLevel
	}
	return allowedLogLevel <= logLevel
}

// IsLoggingAllowed checks the passed in log type and level and determines whether they match the session log criteria or not
func (session *session) IsLoggingAllowed(logType logtype.LogType, logLevel loglevel.LogLevel) bool {
	if !config.IsLocalhost() {
		var loggingTypeMatched = isLoggingTypeMatchFunc(
			session,
			logType,
		)
		var loggingLevelMatched = isLoggingLevelMatchFunc(
			session,
			logLevel,
		)
		if !loggingTypeMatched ||
			(logType == logtype.MethodLogic && !loggingLevelMatched) {
			return false
		}
	}
	return true
}

// Register registers the information of a session for given session ID
func Register(
	name string,
	allowedLogType logtype.LogType,
	allowedLogLevel loglevel.LogLevel,
	httpRequest *http.Request,
	responseWriter http.ResponseWriter,
) model.Session {
	var sessionID = uuidNew()
	var session = &session{
		ID:              sessionID,
		Name:            name,
		AllowedLogType:  allowedLogType,
		AllowedLogLevel: allowedLogLevel,
		Request:         httpRequest,
		ResponseWriter:  responseWriter,
		attachment:      map[string]interface{}{},
	}
	sessionCache.SetDefault(
		sessionID.String(),
		session,
	)
	return session
}

// Unregister unregisters the information of a session for given session ID
func Unregister(session model.Session) {
	var sessionID = session.GetID()
	sessionCache.Delete(
		sessionID.String(),
	)
}

// Get retrieves a registered session for given session ID
func Get(sessionID uuid.UUID) model.Session {
	var cacheItem, sessionLoaded = sessionCache.Get(sessionID.String())
	if !sessionLoaded {
		return defaultSession
	}
	var session, ok = cacheItem.(model.Session)
	if !ok {
		return defaultSession
	}
	return session
}

// GetName returns the name registered to session object for given session ID
func GetName(sessionID uuid.UUID) string {
	var session = getFunc(sessionID)
	return session.GetName()
}

// GetRequest returns the HTTP request object from session object for given session ID
func GetRequest(sessionID uuid.UUID) *http.Request {
	var session = getFunc(sessionID)
	return session.GetRequest()
}

// GetResponseWriter returns the HTTP response writer object from session object for given session ID
func GetResponseWriter(sessionID uuid.UUID) http.ResponseWriter {
	var session = getFunc(sessionID)
	return session.GetResponseWriter()
}

// GetRequestBody loads HTTP request body associated to session and unmarshals the content JSON to given data template
func GetRequestBody(sessionID uuid.UUID, dataTemplate interface{}) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestBody(dataTemplate)
}

// GetRequestParameter loads HTTP request parameter associated to session for given name and unmarshals the content to given data template
func GetRequestParameter(sessionID uuid.UUID, name string, dataTemplate interface{}) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestParameter(name, dataTemplate)
}

// GetRequestQuery loads HTTP request single query string associated to session for given name and unmarshals the content to given data template
func GetRequestQuery(sessionID uuid.UUID, name string, dataTemplate interface{}) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestQuery(name, dataTemplate)
}

// GetRequestQueries loads HTTP request query strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func GetRequestQueries(sessionID uuid.UUID, name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestQueries(name, dataTemplate, fillCallback)
}

// GetRequestHeader loads HTTP request single header string associated to session for given name and unmarshals the content to given data template
func GetRequestHeader(sessionID uuid.UUID, name string, dataTemplate interface{}) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestHeader(name, dataTemplate)
}

// GetRequestHeaders loads HTTP request header strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func GetRequestHeaders(sessionID uuid.UUID, name string, dataTemplate interface{}, fillCallback func()) apperrorModel.AppError {
	var session = getFunc(sessionID)
	return session.GetRequestHeaders(name, dataTemplate, fillCallback)
}

// Attach attaches any value object into the given session associated to the session ID
func Attach(sessionID uuid.UUID, name string, value interface{}) bool {
	var session = getFunc(sessionID)
	return session.Attach(name, value)
}

// Detach detaches any value object from the given session associated to the session ID
func Detach(sessionID uuid.UUID, name string) bool {
	var session = getFunc(sessionID)
	return session.Detach(name)
}

// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template
func GetAttachment(sessionID uuid.UUID, name string, dataTemplate interface{}) bool {
	var session = getFunc(sessionID)
	return session.GetAttachment(
		name,
		dataTemplate,
	)
}

func getMethodName() string {
	var pc, _, _, ok = runtimeCaller(3)
	if !ok {
		return "?"
	}
	var fn = runtimeFuncForPC(pc)
	return fn.Name()
}

// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
func LogMethodEnter(sessionID uuid.UUID) {
	var session = getFunc(sessionID)
	var methodName = getMethodNameFunc()
	loggerMethodEnter(
		session,
		methodName,
		"",
		"",
	)
}

// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
func LogMethodParameter(sessionID uuid.UUID, parameters ...interface{}) {
	var session = getFunc(sessionID)
	var methodName = getMethodNameFunc()
	for index, parameter := range parameters {
		loggerMethodParameter(
			session,
			methodName,
			strconvItoa(index),
			"%v",
			parameter,
		)
	}
}

// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
func LogMethodLogic(sessionID uuid.UUID, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	var session = getFunc(sessionID)
	loggerMethodLogic(
		session,
		logLevel,
		category,
		subcategory,
		messageFormat,
		parameters...,
	)
}

// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
func LogMethodReturn(sessionID uuid.UUID, returns ...interface{}) {
	var session = getFunc(sessionID)
	var methodName = getMethodNameFunc()
	for index, returnValue := range returns {
		loggerMethodReturn(
			session,
			methodName,
			strconvItoa(index),
			"%v",
			returnValue,
		)
	}
}

// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
func LogMethodExit(sessionID uuid.UUID) {
	var session = getFunc(sessionID)
	var methodName = getMethodNameFunc()
	loggerMethodExit(
		session,
		methodName,
		"",
		"",
	)
}

// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
func CreateNetworkRequest(sessionID uuid.UUID, method string, url string, payload string, header map[string]string) networkModel.NetworkRequest {
	var session = getFunc(sessionID)
	return networkNewNetworkRequest(
		session,
		method,
		url,
		payload,
		header,
	)
}
