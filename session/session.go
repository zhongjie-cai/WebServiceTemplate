package session

import (
	"net/http"
	"reflect"

	"github.com/google/uuid"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	networkModel "github.com/zhongjie-cai/WebServiceTemplate/network/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

const (
	defaultName = "AppRoot"
)

var (
	defaultSessionID      uuid.UUID
	defaultRequest        *http.Request
	defaultResponseWriter http.ResponseWriter
	defaultSession        *session
)

// Initialize ensures the NilSession is set for the AppRoot function
func Initialize() {
	defaultSessionID = uuidNew()
	defaultRequest = &http.Request{}
	defaultResponseWriter = &nilResponseWriter{}
	model.NilSession = defaultSession
}

func isInterfaceValueNil(i interface{}) bool {
	if i == nil {
		return true
	}
	var v = reflectValueOf(i)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return !v.IsValid()
}

// Register registers the information of a session for given session ID
func Register(
	name string,
	httpRequest *http.Request,
	responseWriter http.ResponseWriter,
) model.Session {
	var sessionID = uuidNew()
	if httpRequest == nil {
		httpRequest = defaultRequest
	}
	if isInterfaceValueNilFunc(responseWriter) {
		responseWriter = defaultResponseWriter
	}
	var session = &session{
		ID:             sessionID,
		Name:           name,
		Request:        httpRequest,
		ResponseWriter: responseWriter,
		attachment:     map[string]interface{}{},
	}
	session.AllowedLogType = getAllowedLogTypeFunc(session)
	session.AllowedLogLevel = getAllowedLogLevelFunc(session)
	return session
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
		return defaultSessionID
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
	if session == nil ||
		session.Request == nil {
		return defaultRequest
	}
	return session.Request
}

// GetResponseWriter returns the HTTP response writer object from session object for given session ID
func (session *session) GetResponseWriter() http.ResponseWriter {
	if session == nil ||
		isInterfaceValueNilFunc(session.ResponseWriter) {
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
	headerutilLogHTTPHeaderForName(
		session,
		name,
		headers,
		loggerAPIRequest,
	)
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

// GetRawAttachment retrieves any value object from the given session associated to the session ID and returns the raw interface (consumer needs to manually cast, but works for struct with private fields)
func (session *session) GetRawAttachment(name string) (interface{}, bool) {
	if session == nil {
		return nil, false
	}
	var attachment, found = session.attachment[name]
	if !found {
		return nil, false
	}
	return attachment, true
}

// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template
func (session *session) GetAttachment(name string, dataTemplate interface{}) bool {
	var attachment, found = session.GetRawAttachment(name)
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

func getAllowedLogType(session *session) logtype.LogType {
	if customization.SessionAllowedLogType == nil {
		return config.DefaultAllowedLogType()
	}
	return customization.SessionAllowedLogType(session)
}

func getAllowedLogLevel(session *session) loglevel.LogLevel {
	if customization.SessionAllowedLogLevel == nil {
		return config.DefaultAllowedLogLevel()
	}
	return customization.SessionAllowedLogLevel(session)
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

func getMethodName() string {
	var pc, _, _, ok = runtimeCaller(3)
	if !ok {
		return "?"
	}
	var fn = runtimeFuncForPC(pc)
	return fn.Name()
}

// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
func (session *session) LogMethodEnter() {
	var methodName = getMethodNameFunc()
	loggerMethodEnter(
		session,
		methodName,
		"",
		"",
	)
}

// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
func (session *session) LogMethodParameter(parameters ...interface{}) {
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
func (session *session) LogMethodLogic(logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
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
func (session *session) LogMethodReturn(returns ...interface{}) {
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
func (session *session) LogMethodExit() {
	var methodName = getMethodNameFunc()
	loggerMethodExit(
		session,
		methodName,
		"",
		"",
	)
}

func shouldSendClientCert(url string) bool {
	if customization.SendClientCert == nil {
		return certificateHasClientCert()
	}
	return customization.SendClientCert(url)
}

// CreateNetworkRequest generates a network request object to the targeted external web service for the given session associated to the session ID
func (session *session) CreateNetworkRequest(method string, url string, payload string, header map[string]string) networkModel.NetworkRequest {
	var sendClientCert = shouldSendClientCertFunc(url)
	return networkNewNetworkRequest(
		session,
		method,
		url,
		payload,
		header,
		sendClientCert,
	)
}
