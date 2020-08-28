package customization

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	serverModel "github.com/zhongjie-cai/WebServiceTemplate/server/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

// PreBootstrapFunc is to customize the pre-processing logic before bootstrapping
var PreBootstrapFunc func() error

// PostBootstrapFunc is to customize the post-processing logic after bootstrapping
var PostBootstrapFunc func() error

// AppClosingFunc is to customize the application closing logic after server shutdown
var AppClosingFunc func() error

// DefaultAllowedLogType is to customize the default allowed log type loading logic for the whole application
var DefaultAllowedLogType func() logtype.LogType

// DefaultAllowedLogLevel is to customize the default allowed log type loading logic for the whole application
var DefaultAllowedLogLevel func() loglevel.LogLevel

// DefaultHTTPHeaderLogStyle is to customize the log style for HTTP headers for the whole application
var DefaultHTTPHeaderLogStyle func() headerstyle.HeaderStyle

// SessionAllowedLogType is to customize the allowed log type determination logic for every HTTP session
var SessionAllowedLogType func(session sessionModel.Session) logtype.LogType

// SessionAllowedLogLevel is to customize the allowed log level determination logic for every HTTP session
var SessionAllowedLogLevel func(session sessionModel.Session) loglevel.LogLevel

// SessionHTTPHeaderLogStyle is to customize the log style for HTTP headers for every HTTP session
var SessionHTTPHeaderLogStyle func(session sessionModel.Session) headerstyle.HeaderStyle

// LoggingFunc is to customize the logging backend for the whole application
var LoggingFunc func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string)

// AppVersion is to customize the application version string
var AppVersion func() string

// AppPort is to customize the application port number
var AppPort func() string

// AppName is to customize the application name string
var AppName func() string

// AppPath is to customize the application startup system path
var AppPath func() string

// IsLocalhost is to customize the check for localhost
var IsLocalhost func() bool

// ServeHTTPS is to customize the server hosting security option (HTTP v.s. HTTPS)
var ServeHTTPS func() bool

// ServerCertContent is to customize the loading logic for server certificate content
var ServerCertContent func() string

// ServerKeyContent is to customize the loading logic for server key content
var ServerKeyContent func() string

// ValidateClientCert is to customize the server hosting security option (mTLS v.s. none)
var ValidateClientCert func() bool

// CaCertContent is to customize the loading logic for CA certificate content
var CaCertContent func() string

// SendClientCert is to customize the HTTP request option to the external web services (mTLS v.s. none)
var SendClientCert func(url string) bool

// ClientCertContent is to customize the loading logic for client certificate content
var ClientCertContent func() string

// ClientKeyContent is to customize the loading logic for client key content
var ClientKeyContent func() string

// PreActionFunc is to customize the pre-action function used before each route action takes place, e.g. authorization, etc.
var PreActionFunc func(session sessionModel.Session) error

// PostActionFunc is to customize the post-action function used after each route action takes place, e.g. finalization, etc.
var PostActionFunc func(session sessionModel.Session) error

// CreateErrorResponseFunc is to customize the generation of HTTP error response
var CreateErrorResponseFunc func(err error) (responseMessage string, statusCode int)

// Routes is to customize the routes registration
var Routes func() []serverModel.Route

// Statics is to customize the static contents registration
var Statics func() []serverModel.Static

// Middlewares is to customize the middlewares registration
var Middlewares func() []serverModel.MiddlewareFunc

// NotFoundHandler is to customize the handler for routes that are not found in router
var NotFoundHandler func() http.Handler

// MethodNotAllowedHandler is to customize the handler for routes with methods that are not allowed in router
var MethodNotAllowedHandler func() http.Handler

// InstrumentRouter is to customize the instrumentation on top of a fully configured router; usually useful for 3rd party monitoring tools such as new relic, etc.
var InstrumentRouter func(router *mux.Router) *mux.Router

// AppErrors is to append customized AppErrors with their string representations and corresponding HTTP status codes; customized enum must be after apperrorEnum.CodeReservedCount
var AppErrors func() (map[apperrorEnum.Code]string, map[apperrorEnum.Code]int)

// HTTPRoundTripper is to customize the creation of the HTTP transport for any network communications through HTTP/HTTPS by session
var HTTPRoundTripper func(originalTransport http.RoundTripper) http.RoundTripper

// WrapHTTPRequest is to customize the creation of the HTTP request for any network communications through HTTP/HTTPS by session; utilize this method if needed for new relic wrapping, etc.
var WrapHTTPRequest func(session sessionModel.Session, httpRequest *http.Request) *http.Request

// DefaultNetworkRetryDelay is to customize the default retry delay for any network communications through HTTP/HTTPS by session
var DefaultNetworkRetryDelay func() time.Duration

// DefaultNetworkTimeout is to customize the default timeout for any network communications through HTTP/HTTPS by session
var DefaultNetworkTimeout func() time.Duration

// SkipServerCertVerification is to customize the skip of server certificate verification for any network communications through HTTP/HTTPS by session
var SkipServerCertVerification func() bool

// GraceShutdownWaitTime is to customize the graceful shutdown wait time for the application
var GraceShutdownWaitTime func() time.Duration

// Reset clears all customization of functions for the whole application
func Reset() {
	PreBootstrapFunc = nil
	PostBootstrapFunc = nil
	AppClosingFunc = nil
	DefaultAllowedLogType = nil
	DefaultAllowedLogLevel = nil
	DefaultHTTPHeaderLogStyle = nil
	SessionAllowedLogType = nil
	SessionAllowedLogLevel = nil
	SessionHTTPHeaderLogStyle = nil
	LoggingFunc = nil
	AppVersion = nil
	AppPort = nil
	AppName = nil
	AppPath = nil
	IsLocalhost = nil
	ServeHTTPS = nil
	ServerCertContent = nil
	ServerKeyContent = nil
	ValidateClientCert = nil
	CaCertContent = nil
	SendClientCert = nil
	ClientCertContent = nil
	ClientKeyContent = nil
	PreActionFunc = nil
	PostActionFunc = nil
	CreateErrorResponseFunc = nil
	Routes = nil
	Statics = nil
	Middlewares = nil
	NotFoundHandler = nil
	MethodNotAllowedHandler = nil
	InstrumentRouter = nil
	AppErrors = nil
	HTTPRoundTripper = nil
	WrapHTTPRequest = nil
	DefaultNetworkRetryDelay = nil
	DefaultNetworkTimeout = nil
	SkipServerCertVerification = nil
	GraceShutdownWaitTime = nil
}
