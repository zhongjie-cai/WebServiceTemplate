package customization

import (
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// PreBootstrapFunc is to customize the pre-processing logic before bootstrapping
var PreBootstrapFunc func() error

// PostBootstrapFunc is to customize the post-processing logic after bootstrapping
var PostBootstrapFunc func() error

// AppClosingFunc is to customize the application closing logic after server shutdown
var AppClosingFunc func() error

// LoggingFunc is to customize the logging backend for the whole application
var LoggingFunc func(session *session.Session, logType logtype.LogType, category, subcategory, description string)

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

// Routes is to customize the routes registration
var Routes func() []model.Route

// Statics is to customize the static contents registration
var Statics func() []model.Static

// Reset clears all customization of functions for the whole application
func Reset() {
	PreBootstrapFunc = nil
	PostBootstrapFunc = nil
	AppClosingFunc = nil
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
	Routes = nil
	Statics = nil
}