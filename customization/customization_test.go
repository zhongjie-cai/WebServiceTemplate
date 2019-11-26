package customization

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
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
	PreActionFunc = func(sessionID uuid.UUID) error { return nil }
	PostActionFunc = func(sessionID uuid.UUID) error { return nil }
	CreateErrorResponseFunc = func(err error) (responseMessage string, statusCode int) { return "", 0 }
	Routes = func() []serverModel.Route { return nil }
	Statics = func() []serverModel.Static { return nil }
	Middlewares = func() []serverModel.MiddlewareFunc { return nil }
	AppErrors = func() (map[apperrorEnum.Code]string, map[apperrorEnum.Code]int) { return nil, nil }

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
	assert.Nil(t, PreActionFunc)
	assert.Nil(t, PostActionFunc)
	assert.Nil(t, CreateErrorResponseFunc)
	assert.Nil(t, Routes)
	assert.Nil(t, Statics)
	assert.Nil(t, Middlewares)
	assert.Nil(t, AppErrors)

	// verify
	verifyAll(t)
}
