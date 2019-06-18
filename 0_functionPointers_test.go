package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
)

var (
	configAppPortExpected             int
	configAppPortCalled               int
	configAppVersionExpected          int
	configAppVersionCalled            int
	configInitializeExpected          int
	configInitializeCalled            int
	configServeHTTPSExpected          int
	configServeHTTPSCalled            int
	configServerCertContentExpected   int
	configServerCertContentCalled     int
	configServerKeyContentExpected    int
	configServerKeyContentCalled      int
	configValidateClientCertExpected  int
	configValidateClientCertCalled    int
	configCaCertContentExpected       int
	configCaCertContentCalled         int
	certificateInitializeExpected     int
	certificateInitializeCalled       int
	loggerInitializeExpected          int
	loggerInitializeCalled            int
	loggerAppRootExpected             int
	loggerAppRootCalled               int
	serverHostExpected                int
	serverHostCalled                  int
	doPreBootstrapingFuncExpected     int
	doPreBootstrapingFuncCalled       int
	bootstrapApplicationFuncExpected  int
	bootstrapApplicationFuncCalled    int
	doPostBootstrapingFuncExpected    int
	doPostBootstrapingFuncCalled      int
	doApplicationStartingFuncExpected int
	doApplicationStartingFuncCalled   int
	doApplicationClosingFuncExpected  int
	doApplicationClosingFuncCalled    int
)

func createMock(t *testing.T) {
	configAppPortExpected = 0
	configAppPortCalled = 0
	configAppPort = func() string {
		configAppPortCalled++
		return ""
	}
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	configAppVersion = func() string {
		configAppVersionCalled++
		return ""
	}
	configInitializeExpected = 0
	configInitializeCalled = 0
	configInitialize = func() error {
		configInitializeCalled++
		return nil
	}
	configServeHTTPSExpected = 0
	configServeHTTPSCalled = 0
	configServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return false
	}
	configServerCertContentExpected = 0
	configServerCertContentCalled = 0
	configServerCertContent = func() string {
		configServerCertContentCalled++
		return ""
	}
	configServerKeyContentExpected = 0
	configServerKeyContentCalled = 0
	configServerKeyContent = func() string {
		configServerKeyContentCalled++
		return ""
	}
	configValidateClientCertExpected = 0
	configValidateClientCertCalled = 0
	configValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return false
	}
	configCaCertContentExpected = 0
	configCaCertContentCalled = 0
	configCaCertContent = func() string {
		configCaCertContentCalled++
		return ""
	}
	certificateInitializeExpected = 0
	certificateInitializeCalled = 0
	certificateInitialize = func(serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		return nil
	}
	loggerInitializeExpected = 0
	loggerInitializeCalled = 0
	loggerInitialize = func() error {
		loggerInitializeCalled++
		return nil
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	serverHostExpected = 0
	serverHostCalled = 0
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		return nil
	}
	doPreBootstrapingFuncExpected = 0
	doPreBootstrapingFuncCalled = 0
	doPreBootstrapingFunc = func() bool {
		doPreBootstrapingFuncCalled++
		return false
	}
	bootstrapApplicationFuncExpected = 0
	bootstrapApplicationFuncCalled = 0
	bootstrapApplicationFunc = func() bool {
		bootstrapApplicationFuncCalled++
		return false
	}
	doPostBootstrapingFuncExpected = 0
	doPostBootstrapingFuncCalled = 0
	doPostBootstrapingFunc = func() bool {
		doPostBootstrapingFuncCalled++
		return false
	}
	doApplicationStartingFuncExpected = 0
	doApplicationStartingFuncCalled = 0
	doApplicationStartingFunc = func() {
		doApplicationStartingFuncCalled++
	}
	doApplicationClosingFuncExpected = 0
	doApplicationClosingFuncCalled = 0
	doApplicationClosingFunc = func() {
		doApplicationClosingFuncCalled++
	}
}

func verifyAll(t *testing.T) {
	configAppPort = config.AppPort
	assert.Equal(t, configAppPortExpected, configAppPortCalled, "Unexpected number of calls to configAppPort")
	configAppVersion = config.AppVersion
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected number of calls to configAppVersion")
	configInitialize = config.Initialize
	assert.Equal(t, configInitializeExpected, configInitializeCalled, "Unexpected number of calls to configInitialize")
	configServeHTTPS = config.ServeHTTPS
	assert.Equal(t, configServeHTTPSExpected, configServeHTTPSCalled, "Unexpected number of calls to configServeHTTPS")
	configServerCertContent = config.ServerCertContent
	assert.Equal(t, configServerCertContentExpected, configServerCertContentCalled, "Unexpected number of calls to configServerCertContent")
	configServerKeyContent = config.ServerKeyContent
	assert.Equal(t, configServerKeyContentExpected, configServerKeyContentCalled, "Unexpected number of calls to configServerKeyContent")
	configValidateClientCert = config.ValidateClientCert
	assert.Equal(t, configValidateClientCertExpected, configValidateClientCertCalled, "Unexpected number of calls to configValidateClientCert")
	configCaCertContent = config.CaCertContent
	assert.Equal(t, configCaCertContentExpected, configCaCertContentCalled, "Unexpected number of calls to configCaCertContent")
	certificateInitialize = certificate.Initialize
	assert.Equal(t, certificateInitializeExpected, certificateInitializeCalled, "Unexpected number of calls to certificateInitialize")
	loggerInitialize = logger.Initialize
	assert.Equal(t, loggerInitializeExpected, loggerInitializeCalled, "Unexpected number of calls to loggerInitialize")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	serverHost = server.Host
	assert.Equal(t, serverHostExpected, serverHostCalled, "Unexpected number of calls to serverHost")
	doPreBootstrapingFunc = doPreBootstraping
	assert.Equal(t, doPreBootstrapingFuncExpected, doPreBootstrapingFuncCalled, "Unexpected number of calls to doPreBootstrapingFunc")
	bootstrapApplicationFunc = bootstrapApplication
	assert.Equal(t, bootstrapApplicationFuncExpected, bootstrapApplicationFuncCalled, "Unexpected number of calls to bootstrapApplicationFunc")
	doPostBootstrapingFunc = doPostBootstraping
	assert.Equal(t, doPostBootstrapingFuncExpected, doPostBootstrapingFuncCalled, "Unexpected number of calls to doPostBootstrapingFunc")
	doApplicationStartingFunc = doApplicationStarting
	assert.Equal(t, doApplicationStartingFuncExpected, doApplicationStartingFuncCalled, "Unexpected number of calls to doApplicationStartingFunc")
	doApplicationClosingFunc = doApplicationClosing
	assert.Equal(t, doApplicationClosingFuncExpected, doApplicationClosingFuncCalled, "Unexpected number of calls to doApplicationClosingFunc")

	PreBootstrapFunc = nil
	PostBootstrapFunc = nil
	AppClosingFunc = nil
}
