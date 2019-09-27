package customization

import (
	"testing"
)

var ()

func createMock(t *testing.T) {
}

func verifyAll(t *testing.T) {
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
	AuthorizationFunc = nil
	CreateErrorResponseFunc = nil
	Routes = nil
	Statics = nil
	Middlewares = nil
}
