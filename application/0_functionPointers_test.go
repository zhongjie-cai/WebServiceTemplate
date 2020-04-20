package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/network"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
	sessionInitializeExpected                int
	sessionInitializeCalled                  int
	configAppPortExpected                    int
	configAppPortCalled                      int
	configAppVersionExpected                 int
	configAppVersionCalled                   int
	configInitializeExpected                 int
	configInitializeCalled                   int
	configServeHTTPSExpected                 int
	configServeHTTPSCalled                   int
	configServerCertContentExpected          int
	configServerCertContentCalled            int
	configServerKeyContentExpected           int
	configServerKeyContentCalled             int
	configValidateClientCertExpected         int
	configValidateClientCertCalled           int
	configCaCertContentExpected              int
	configCaCertContentCalled                int
	configClientCertContentExpected          int
	configClientCertContentCalled            int
	configClientKeyContentExpected           int
	configClientKeyContentCalled             int
	configDefaultNetworkTimeoutExpected      int
	configDefaultNetworkTimeoutCalled        int
	configSkipServerCertVerificationExpected int
	configSkipServerCertVerificationCalled   int
	certificateInitializeExpected            int
	certificateInitializeCalled              int
	apperrorInitializeExpected               int
	apperrorInitializeCalled                 int
	networkInitializeExpected                int
	networkInitializeCalled                  int
	loggerInitializeExpected                 int
	loggerInitializeCalled                   int
	loggerAppRootExpected                    int
	loggerAppRootCalled                      int
	serverHostExpected                       int
	serverHostCalled                         int
	serverHaltExpected                       int
	serverHaltCalled                         int
	doPreBootstrapingFuncExpected            int
	doPreBootstrapingFuncCalled              int
	bootstrapApplicationFuncExpected         int
	bootstrapApplicationFuncCalled           int
	doPostBootstrapingFuncExpected           int
	doPostBootstrapingFuncCalled             int
	doApplicationStartingFuncExpected        int
	doApplicationStartingFuncCalled          int
	doApplicationClosingFuncExpected         int
	doApplicationClosingFuncCalled           int
)

func createMock(t *testing.T) {
	sessionInitializeExpected = 0
	sessionInitializeCalled = 0
	sessionInitialize = func() {
		sessionInitializeCalled++
	}
	configAppPortExpected = 0
	configAppPortCalled = 0
	config.AppPort = func() string {
		configAppPortCalled++
		return ""
	}
	configAppVersionExpected = 0
	configAppVersionCalled = 0
	config.AppVersion = func() string {
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
	config.ServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return false
	}
	configServerCertContentExpected = 0
	configServerCertContentCalled = 0
	config.ServerCertContent = func() string {
		configServerCertContentCalled++
		return ""
	}
	configServerKeyContentExpected = 0
	configServerKeyContentCalled = 0
	config.ServerKeyContent = func() string {
		configServerKeyContentCalled++
		return ""
	}
	configValidateClientCertExpected = 0
	configValidateClientCertCalled = 0
	config.ValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return false
	}
	configCaCertContentExpected = 0
	configCaCertContentCalled = 0
	config.CaCertContent = func() string {
		configCaCertContentCalled++
		return ""
	}
	configClientCertContentExpected = 0
	configClientCertContentCalled = 0
	config.ClientCertContent = func() string {
		configClientCertContentCalled++
		return ""
	}
	configClientKeyContentExpected = 0
	configClientKeyContentCalled = 0
	config.ClientKeyContent = func() string {
		configClientKeyContentCalled++
		return ""
	}
	configDefaultNetworkTimeoutExpected = 0
	configDefaultNetworkTimeoutCalled = 0
	config.DefaultNetworkTimeout = func() time.Duration {
		configDefaultNetworkTimeoutCalled++
		return 0
	}
	configSkipServerCertVerificationExpected = 0
	configSkipServerCertVerificationCalled = 0
	config.SkipServerCertVerification = func() bool {
		configSkipServerCertVerificationCalled++
		return false
	}
	certificateInitializeExpected = 0
	certificateInitializeCalled = 0
	certificateInitialize = func(serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string, clientCertContent string, clientKeyContent string) error {
		certificateInitializeCalled++
		return nil
	}
	apperrorInitializeExpected = 0
	apperrorInitializeCalled = 0
	apperrorInitialize = func() error {
		apperrorInitializeCalled++
		return nil
	}
	networkInitializeExpected = 0
	networkInitializeCalled = 0
	networkInitialize = func(networkTimeout time.Duration, skipServerCertVerification bool) {
		networkInitializeCalled++
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
	serverHaltExpected = 0
	serverHaltCalled = 0
	serverHalt = func() {
		serverHaltCalled++
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
	config.AppPort = func() string { return "" }
	assert.Equal(t, configAppPortExpected, configAppPortCalled, "Unexpected number of calls to configAppPort")
	config.AppVersion = func() string { return "" }
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected number of calls to configAppVersion")
	configInitialize = config.Initialize
	assert.Equal(t, configInitializeExpected, configInitializeCalled, "Unexpected number of calls to configInitialize")
	config.ServeHTTPS = func() bool { return false }
	assert.Equal(t, configServeHTTPSExpected, configServeHTTPSCalled, "Unexpected number of calls to configServeHTTPS")
	config.ServerCertContent = func() string { return "" }
	assert.Equal(t, configServerCertContentExpected, configServerCertContentCalled, "Unexpected number of calls to configServerCertContent")
	config.ServerKeyContent = func() string { return "" }
	assert.Equal(t, configServerKeyContentExpected, configServerKeyContentCalled, "Unexpected number of calls to configServerKeyContent")
	config.ValidateClientCert = func() bool { return false }
	assert.Equal(t, configValidateClientCertExpected, configValidateClientCertCalled, "Unexpected number of calls to configValidateClientCert")
	config.CaCertContent = func() string { return "" }
	assert.Equal(t, configCaCertContentExpected, configCaCertContentCalled, "Unexpected number of calls to configCaCertContent")
	config.ClientCertContent = func() string { return "" }
	assert.Equal(t, configClientCertContentExpected, configClientCertContentCalled, "Unexpected number of calls to configClientCertContent")
	config.ClientKeyContent = func() string { return "" }
	assert.Equal(t, configClientKeyContentExpected, configClientKeyContentCalled, "Unexpected number of calls to configClientKeyContent")
	config.DefaultNetworkTimeout = func() time.Duration { return 0 }
	assert.Equal(t, configDefaultNetworkTimeoutExpected, configDefaultNetworkTimeoutCalled, "Unexpected number of calls to configDefaultNetworkTimeout")
	config.SkipServerCertVerification = func() bool { return false }
	assert.Equal(t, configSkipServerCertVerificationExpected, configSkipServerCertVerificationCalled, "Unexpected number of calls to configSkipServerCertVerification")
	sessionInitialize = session.Initialize
	assert.Equal(t, sessionInitializeExpected, sessionInitializeCalled, "Unexpected number of calls to sessionInitialize")
	certificateInitialize = certificate.Initialize
	assert.Equal(t, certificateInitializeExpected, certificateInitializeCalled, "Unexpected number of calls to certificateInitialize")
	apperrorInitialize = apperror.Initialize
	assert.Equal(t, apperrorInitializeExpected, apperrorInitializeCalled, "Unexpected number of calls to apperrorInitialize")
	networkInitialize = network.Initialize
	assert.Equal(t, networkInitializeExpected, networkInitializeCalled, "Unexpected number of calls to networkInitialize")
	loggerInitialize = logger.Initialize
	assert.Equal(t, loggerInitializeExpected, loggerInitializeCalled, "Unexpected number of calls to loggerInitialize")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected number of calls to loggerAppRoot")
	serverHost = server.Host
	assert.Equal(t, serverHostExpected, serverHostCalled, "Unexpected number of calls to serverHost")
	serverHalt = server.Halt
	assert.Equal(t, serverHaltExpected, serverHaltCalled, "Unexpected number of calls to serverHalt")
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

	customization.PreBootstrapFunc = nil
	customization.PostBootstrapFunc = nil
	customization.AppClosingFunc = nil
}
