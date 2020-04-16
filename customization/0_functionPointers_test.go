package customization

import (
	"testing"
)

func createMock(t *testing.T) {
}

func verifyAll(t *testing.T) {
	PreBootstrapFunc = nil
	PostBootstrapFunc = nil
	AppClosingFunc = nil
	DefaultAllowedLogType = nil
	DefaultAllowedLogLevel = nil
	SessionAllowedLogType = nil
	SessionAllowedLogLevel = nil
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
	InstrumentRouter = nil
	HTTPRoundTripper = nil
	WrapHTTPRequest = nil
	DefaultNetworkRetryDelay = nil
	DefaultNetworkTimeout = nil
	SkipServerCertVerification = nil
	GraceShutdownWaitTime = nil
}
