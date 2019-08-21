package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	cache "github.com/patrickmn/go-cache"
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
	defaultSession        = &Session{
		ID:             uuid.Nil,
		Endpoint:       "",
		LoginID:        uuid.Nil,
		AllowedLogType: logtype.BasicLogging,
		Request:        defaultRequest,
		ResponseWriter: defaultResponseWriter,
	}
)

// Session is the storage for the current HTTP request session, containing information needed for logging, monitoring, etc.
type Session struct {
	ID             uuid.UUID
	Endpoint       string
	LoginID        uuid.UUID
	CorrelationID  uuid.UUID
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
	endpoint string,
	loginID uuid.UUID,
	correlationID uuid.UUID,
	allowedLogType logtype.LogType,
	httpRequest *http.Request,
	responseWriter http.ResponseWriter,
) uuid.UUID {
	if loginID == uuid.Nil {
		return uuid.Nil
	}
	var sessionID = uuidNew()
	sessionCache.SetDefault(
		sessionID.String(),
		&Session{
			ID:             sessionID,
			Endpoint:       endpoint,
			LoginID:        loginID,
			CorrelationID:  correlationID,
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

// ClearResponseWriter clears the response writer instance and sets it to nil writer that takes no actions at all
func ClearResponseWriter(sessionID uuid.UUID) {
	var sessionObject = getFunc(sessionID)
	if sessionObject == nil {
		return
	}
	sessionObject.ResponseWriter = defaultResponseWriter
	sessionCache.SetDefault(
		sessionID.String(),
		sessionObject,
	)
}
