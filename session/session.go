package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	cache "github.com/patrickmn/go-cache"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

type nilResponseWriter struct{}

func (r *nilResponseWriter) Header() http.Header {
	return http.Header{}
}

func (r *nilResponseWriter) Write(body []byte) (int, error) {
	return 0, nil
}

func (r *nilResponseWriter) WriteHeader(status int) {
}

var (
	sessionCache          = cache.New(15*time.Minute, 30*time.Minute)
	defaultRequest, _     = http.NewRequest("", "", nil)
	defaultResponseWriter = &nilResponseWriter{}
	defaultName           = "AppRoot"
	defaultSession        = &Session{
		ID:             uuid.Nil,
		Name:           defaultName,
		AllowedLogType: logtype.BasicLogging,
		Request:        defaultRequest,
		ResponseWriter: defaultResponseWriter,
	}
)

// Session is the storage for the current HTTP request session, containing information needed for logging, monitoring, etc.
type Session struct {
	ID             uuid.UUID
	Name           string
	AllowedLogType logtype.LogType
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Init initialize the default session for the application
func Init(appName string, roleType string, hostName string, version string, buildTime string) {
	// Initialize default session entry
	sessionCache.Set(
		uuid.Nil.String(),
		defaultSession,
		cache.NoExpiration,
	)
}

// Register registers the information of a session for given session ID
func Register(
	name string,
	allowedLogType logtype.LogType,
	httpRequest *http.Request,
	responseWriter http.ResponseWriter,
) uuid.UUID {
	var sessionID = uuidNew()
	sessionCache.SetDefault(
		sessionID.String(),
		&Session{
			ID:             sessionID,
			Name:           name,
			AllowedLogType: allowedLogType,
			Request:        httpRequest,
			ResponseWriter: responseWriter,
		},
	)
	return sessionID
}

// Unregister unregisters the information of a session for given session ID
func Unregister(sessionID uuid.UUID) {
	sessionCache.Delete(
		sessionID.String(),
	)
}

// Get retrieves a registered session for given session ID
func Get(sessionID uuid.UUID) *Session {
	var cacheItem, sessionLoaded = sessionCache.Get(sessionID.String())
	if !sessionLoaded {
		return defaultSession
	}
	var session, ok = cacheItem.(*Session)
	if !ok {
		return defaultSession
	}
	return session
}

// GetName returns the name registered to session object for given session ID
func GetName(sessionID uuid.UUID) string {
	var sessionObject = getFunc(sessionID)
	if sessionObject == nil {
		return defaultName
	}
	return sessionObject.Name
}

// GetRequest returns the HTTP request object from session object for given session ID
func GetRequest(sessionID uuid.UUID) *http.Request {
	var sessionObject = getFunc(sessionID)
	if sessionObject == nil {
		return defaultRequest
	}
	return sessionObject.Request
}

// GetResponseWriter returns the HTTP response writer object from session object for given session ID
func GetResponseWriter(sessionID uuid.UUID) http.ResponseWriter {
	var sessionObject = getFunc(sessionID)
	if sessionObject == nil {
		return defaultResponseWriter
	}
	return sessionObject.ResponseWriter
}

func tryUnmarshal(value string, dataTemplate interface{}) apperror.AppError {
	var noQuoteJSONError = jsonUnmarshal(
		[]byte(value),
		dataTemplate,
	)
	if noQuoteJSONError == nil {
		return nil
	}
	if value == "" {
		return nil
	}
	var withQuoteJSONError = jsonUnmarshal(
		[]byte("\""+value+"\""),
		dataTemplate,
	)
	if withQuoteJSONError == nil {
		return nil
	}
	return apperrorGetBadRequestError(
		fmtErrorf(
			"Unable to unmarshal value [%v] into data template",
			value,
		),
	)
}

// GetRequestBody loads HTTP request body associated to session and unmarshals the content JSON to given data template
func GetRequestBody(sessionID uuid.UUID, dataTemplate interface{}) apperror.AppError {
	var httpRequest = getRequestFunc(
		sessionID,
	)
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
	return tryUnmarshalFunc(
		requestBody,
		dataTemplate,
	)
}

// GetRequestParameter loads HTTP request parameter associated to session for given name and unmarshals the content to given data template
func GetRequestParameter(sessionID uuid.UUID, name string, dataTemplate interface{}) apperror.AppError {
	var httpRequest = getRequestFunc(
		sessionID,
	)
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
	return tryUnmarshalFunc(
		value,
		dataTemplate,
	)
}

func getAllQueries(sessionID uuid.UUID, name string) []string {
	var httpRequest = getRequestFunc(
		sessionID,
	)
	var queries, found = httpRequest.URL.Query()[name]
	if !found {
		return nil
	}
	return queries
}

// GetRequestQuery loads HTTP request single query string associated to session for given name and unmarshals the content to given data template
func GetRequestQuery(sessionID uuid.UUID, name string, dataTemplate interface{}) apperror.AppError {
	var queries = getAllQueriesFunc(
		sessionID,
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
	return tryUnmarshalFunc(
		queries[0],
		dataTemplate,
	)
}

// GetRequestQueries loads HTTP request query strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func GetRequestQueries(sessionID uuid.UUID, name string, dataTemplate interface{}, fillCallback func()) apperror.AppError {
	var queries = getAllQueriesFunc(
		sessionID,
		name,
	)
	var unmarshalErrors = []error{}
	for _, query := range queries {
		var unmarshalError = tryUnmarshalFunc(
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
	return apperrorConsolidateAllErrors(
		"Failed to get request query strings",
		unmarshalErrors...,
	)
}

func getAllHeaders(sessionID uuid.UUID, name string) []string {
	var httpRequest = getRequestFunc(
		sessionID,
	)
	var canonicalName = textprotoCanonicalMIMEHeaderKey(name)
	var headers, found = httpRequest.Header[canonicalName]
	if !found {
		return nil
	}
	return headers
}

// GetRequestHeader loads HTTP request single header string associated to session for given name and unmarshals the content to given data template
func GetRequestHeader(sessionID uuid.UUID, name string, dataTemplate interface{}) apperror.AppError {
	var headers = getAllHeadersFunc(
		sessionID,
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
	return tryUnmarshalFunc(
		headers[0],
		dataTemplate,
	)
}

// GetRequestHeaders loads HTTP request header strings associated to session for given name and unmarshals the content to given data template; the fillCallback is called when each unmarshal operation succeeds, so consumer could fill in external arrays using data template during the process
func GetRequestHeaders(sessionID uuid.UUID, name string, dataTemplate interface{}, fillCallback func()) apperror.AppError {
	var headers = getAllHeadersFunc(
		sessionID,
		name,
	)
	var unmarshalErrors = []error{}
	for _, header := range headers {
		var unmarshalError = tryUnmarshalFunc(
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
	return apperrorConsolidateAllErrors(
		"Failed to get request header strings",
		unmarshalErrors...,
	)
}
