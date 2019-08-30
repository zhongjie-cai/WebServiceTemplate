package customization

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

func TestReset(t *testing.T) {
	// stub
	PreBootstrapFunc = func() error { return nil }
	PostBootstrapFunc = func() error { return nil }
	AppClosingFunc = func() error { return nil }
	LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {}
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
	CreateErrorResponseFunc = func(appError apperror.AppError) (responseMessage string, statusCode int) { return "", 0 }
	Routes = func() []model.Route { return nil }
	Statics = func() []model.Static { return nil }
	Middlewares = func() []model.MiddlewareFunc { return nil }

	// mock
	createMock(t)

	// SUT + act
	Reset()

	// assert
	assert.Nil(t, PreBootstrapFunc)
	assert.Nil(t, PostBootstrapFunc)
	assert.Nil(t, AppClosingFunc)
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
	assert.Nil(t, CreateErrorResponseFunc)
	assert.Nil(t, Routes)
	assert.Nil(t, Statics)
	assert.Nil(t, Middlewares)

	// verify
	verifyAll(t)
}
