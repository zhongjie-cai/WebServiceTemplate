package customization

import (
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	serverModel "github.com/zhongjie-cai/WebServiceTemplate/server/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestReset(t *testing.T) {
	// stub
	PreBootstrapFunc = func() error { return nil }
	PostBootstrapFunc = func() error { return nil }
	AppClosingFunc = func() error { return nil }
	DefaultAllowedLogType = func() logtype.LogType { return logtype.LogType(0) }
	DefaultAllowedLogLevel = func() loglevel.LogLevel { return loglevel.LogLevel(0) }
	DefaultHTTPHeaderLogStyle = func() headerstyle.HeaderStyle { return headerstyle.HeaderStyle(0) }
	SessionAllowedLogType = func(session sessionModel.Session) logtype.LogType { return logtype.LogType(0) }
	SessionAllowedLogLevel = func(session sessionModel.Session) loglevel.LogLevel { return loglevel.LogLevel(0) }
	SessionHTTPHeaderLogStyle = func(session sessionModel.Session) headerstyle.HeaderStyle { return headerstyle.HeaderStyle(0) }
	LoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
	}
	AppVersion = func() string { return "" }
	AppPort = func() string { return "" }
	AppName = func() string { return "" }
	AppPath = func() string { return "" }
	IsLocalhost = func() bool { return false }
	ServeHTTPS = func() bool { return false }
	ServerCertContent = func() string { return "" }
	ServerKeyContent = func() string { return "" }
	ValidateClientCert = func() bool { return false }
	CaCertContent = func() string { return "" }
	SendClientCert = func(url string) bool { return false }
	ClientCertContent = func() string { return "" }
	ClientKeyContent = func() string { return "" }
	PreActionFunc = func(session sessionModel.Session) error { return nil }
	PostActionFunc = func(session sessionModel.Session) error { return nil }
	CreateErrorResponseFunc = func(err error) (responseMessage string, statusCode int) { return "", 0 }
	Routes = func() []serverModel.Route { return nil }
	Statics = func() []serverModel.Static { return nil }
	Middlewares = func() []serverModel.MiddlewareFunc { return nil }
	InstrumentRouter = func(router *mux.Router) *mux.Router { return nil }
	AppErrors = func() (map[apperrorEnum.Code]string, map[apperrorEnum.Code]int) { return nil, nil }
	HTTPRoundTripper = func(originalTransport http.RoundTripper) http.RoundTripper { return nil }
	WrapHTTPRequest = func(session sessionModel.Session, httpRequest *http.Request) *http.Request { return nil }
	DefaultNetworkRetryDelay = func() time.Duration { return 0 }
	DefaultNetworkTimeout = func() time.Duration { return 0 }
	SkipServerCertVerification = func() bool { return false }

	// mock
	createMock(t)

	// SUT + act
	Reset()

	// assert
	assert.Nil(t, PreBootstrapFunc)
	assert.Nil(t, PostBootstrapFunc)
	assert.Nil(t, AppClosingFunc)
	assert.Nil(t, DefaultAllowedLogType)
	assert.Nil(t, DefaultAllowedLogLevel)
	assert.Nil(t, DefaultHTTPHeaderLogStyle)
	assert.Nil(t, SessionAllowedLogType)
	assert.Nil(t, SessionAllowedLogLevel)
	assert.Nil(t, SessionHTTPHeaderLogStyle)
	assert.Nil(t, LoggingFunc)
	assert.Nil(t, AppVersion)
	assert.Nil(t, AppPort)
	assert.Nil(t, AppName)
	assert.Nil(t, AppPath)
	assert.Nil(t, IsLocalhost)
	assert.Nil(t, ServeHTTPS)
	assert.Nil(t, ServerCertContent)
	assert.Nil(t, ServerKeyContent)
	assert.Nil(t, ValidateClientCert)
	assert.Nil(t, CaCertContent)
	assert.Nil(t, SendClientCert)
	assert.Nil(t, ClientCertContent)
	assert.Nil(t, ClientKeyContent)
	assert.Nil(t, PreActionFunc)
	assert.Nil(t, PostActionFunc)
	assert.Nil(t, CreateErrorResponseFunc)
	assert.Nil(t, Routes)
	assert.Nil(t, Statics)
	assert.Nil(t, Middlewares)
	assert.Nil(t, InstrumentRouter)
	assert.Nil(t, AppErrors)
	assert.Nil(t, HTTPRoundTripper)
	assert.Nil(t, WrapHTTPRequest)
	assert.Nil(t, DefaultNetworkRetryDelay)
	assert.Nil(t, DefaultNetworkTimeout)
	assert.Nil(t, SkipServerCertVerification)

	// verify
	verifyAll(t)
}
