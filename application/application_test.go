package application

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func TestDoPreBootstraping_NoPreBoostrapFunc(t *testing.T) {
	// arrange
	var preBootstrapFuncExpected int
	var preBootstrapFuncCalled int

	// stub
	customization.PreBootstrapFunc = nil

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPreBootstraping", subcategory)
		assert.Equal(t, "customization.PreBootstrapFunc is not configured; skipped execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	result := doPreBootstraping()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, preBootstrapFuncExpected, preBootstrapFuncCalled, "Unexpected number of calls to PreBootstrapFunc")
}

func TestDoPreBootstraping_PreBoostrapFuncError(t *testing.T) {
	// arrange
	var preBootstrapFuncExpected int
	var preBootstrapFuncCalled int
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	preBootstrapFuncExpected = 1
	customization.PreBootstrapFunc = func() error {
		preBootstrapFuncCalled++
		return dummyError
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPreBootstraping", subcategory)
		assert.Equal(t, "Failed to execute customization.PreBootstrapFunc. Error: %v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	result := doPreBootstraping()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, preBootstrapFuncExpected, preBootstrapFuncCalled, "Unexpected number of calls to PreBootstrapFunc")
}

func TestDoPreBootstraping_PreBoostrapFuncSuccess(t *testing.T) {
	// arrange
	var preBootstrapFuncExpected int
	var preBootstrapFuncCalled int

	// mock
	createMock(t)

	// expect
	preBootstrapFuncExpected = 1
	customization.PreBootstrapFunc = func() error {
		preBootstrapFuncCalled++
		return nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPreBootstraping", subcategory)
		assert.Equal(t, "customization.PreBootstrapFunc executed successfully", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	result := doPreBootstraping()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, preBootstrapFuncExpected, preBootstrapFuncCalled, "Unexpected number of calls to PreBootstrapFunc")
}

func TestBootstrapApplication_WithError(t *testing.T) {
	// arrange
	var dummyLoggerError = errors.New("some logger error")
	var dummyConfigError = errors.New("some config error")
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"
	var dummyCertError = errors.New("some cert error")

	// mock
	createMock(t)

	// expect
	loggerInitializeExpected = 1
	loggerInitialize = func() error {
		loggerInitializeCalled++
		return dummyLoggerError
	}
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return dummyConfigError
	}
	configServeHTTPSExpected = 1
	config.ServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configServerCertContentExpected = 1
	config.ServerCertContent = func() string {
		configServerCertContentCalled++
		return dummyServerCertContent
	}
	configServerKeyContentExpected = 1
	config.ServerKeyContent = func() string {
		configServerKeyContentCalled++
		return dummyServerKeyContent
	}
	configValidateClientCertExpected = 1
	config.ValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	configCaCertContentExpected = 1
	config.CaCertContent = func() string {
		configCaCertContentCalled++
		return dummyCaCertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyCaCertContent, caCertContent)
		return dummyCertError
	}
	loggerAppRootExpected = 3
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "bootstrapApplication", subcategory)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Application logger not initialized cleanly. Potential error: %v", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyLoggerError, parameters[0])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Application configuration not initialized cleanly. Potential error: %v", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyConfigError, parameters[0])
		} else if loggerAppRootCalled == 3 {
			assert.Equal(t, "Failed to bootstrap server application. Error: %v", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyCertError, parameters[0])
		}
	}

	// SUT + act
	result := bootstrapApplication()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestBootstrapApplication_NoError(t *testing.T) {
	// arrange
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyServerCertContent = "some server cert content"
	var dummyServerKeyContent = "some server key content"
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyCaCertContent = "some CA cert content"

	// mock
	createMock(t)

	// expect
	loggerInitializeExpected = 1
	loggerInitialize = func() error {
		loggerInitializeCalled++
		return nil
	}
	configInitializeExpected = 1
	configInitialize = func() error {
		configInitializeCalled++
		return nil
	}
	configServeHTTPSExpected = 1
	config.ServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configServerCertContentExpected = 1
	config.ServerCertContent = func() string {
		configServerCertContentCalled++
		return dummyServerCertContent
	}
	configServerKeyContentExpected = 1
	config.ServerKeyContent = func() string {
		configServerKeyContentCalled++
		return dummyServerKeyContent
	}
	configValidateClientCertExpected = 1
	config.ValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	configCaCertContentExpected = 1
	config.CaCertContent = func() string {
		configCaCertContentCalled++
		return dummyCaCertContent
	}
	certificateInitializeExpected = 1
	certificateInitialize = func(serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyServerCertContent, serverCertContent)
		assert.Equal(t, dummyServerKeyContent, serverKeyContent)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyCaCertContent, caCertContent)
		return nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "bootstrapApplication", subcategory)
		assert.Equal(t, "Application bootstrapped successfully", messageFormat)
	}

	// SUT + act
	result := bootstrapApplication()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestDoPostBootstraping_NoPostBoostrapFunc(t *testing.T) {
	// arrange
	var postBootstrapFuncExpected int
	var postBootstrapFuncCalled int

	// stub
	customization.PostBootstrapFunc = nil

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPostBootstraping", subcategory)
		assert.Equal(t, "customization.PostBootstrapFunc is not configured; skipped execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	result := doPostBootstraping()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, postBootstrapFuncExpected, postBootstrapFuncCalled, "Unexpected number of calls to PostBootstrapFunc")
}

func TestDoPostBootstraping_PostBoostrapFuncError(t *testing.T) {
	// arrange
	var postBootstrapFuncExpected int
	var postBootstrapFuncCalled int
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	postBootstrapFuncExpected = 1
	customization.PostBootstrapFunc = func() error {
		postBootstrapFuncCalled++
		return dummyError
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPostBootstraping", subcategory)
		assert.Equal(t, "Failed to execute customization.PostBootstrapFunc. Error: %v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	result := doPostBootstraping()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, postBootstrapFuncExpected, postBootstrapFuncCalled, "Unexpected number of calls to PostBootstrapFunc")
}

func TestDoPostBootstraping_PostBoostrapFuncSuccess(t *testing.T) {
	// arrange
	var postBootstrapFuncExpected int
	var postBootstrapFuncCalled int

	// mock
	createMock(t)

	// expect
	postBootstrapFuncExpected = 1
	customization.PostBootstrapFunc = func() error {
		postBootstrapFuncCalled++
		return nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doPostBootstraping", subcategory)
		assert.Equal(t, "customization.PostBootstrapFunc executed successfully", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	result := doPostBootstraping()

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
	assert.Equal(t, postBootstrapFuncExpected, postBootstrapFuncCalled, "Unexpected number of calls to PostBootstrapFunc")
}

func TestDoApplicationStarting_HostError(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50
	var dummyHostError = errors.New("some host error")

	// mock
	createMock(t)

	// expect
	configAppPortExpected = 1
	config.AppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	config.AppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	config.ServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	config.ValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return dummyHostError
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doApplicationStarting", subcategory)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Trying to start server (v-%v)", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Failed to host server. Error: %v", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyHostError, parameters[0])
		}
	}

	// SUT + act
	doApplicationStarting()

	// verify
	verifyAll(t)
}

func TestDoApplicationStarting_HostSuccess(t *testing.T) {
	// arrange
	var dummyAppVersion = "dummyAppVersion"
	var dummyAppPort = "dummyAppPort"
	var dummyServeHTTPS = rand.Intn(100) < 50
	var dummyValidateClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// expect
	configAppPortExpected = 1
	config.AppPort = func() string {
		configAppPortCalled++
		return dummyAppPort
	}
	configAppVersionExpected = 1
	config.AppVersion = func() string {
		configAppVersionCalled++
		return dummyAppVersion
	}
	configServeHTTPSExpected = 1
	config.ServeHTTPS = func() bool {
		configServeHTTPSCalled++
		return dummyServeHTTPS
	}
	configValidateClientCertExpected = 1
	config.ValidateClientCert = func() bool {
		configValidateClientCertCalled++
		return dummyValidateClientCert
	}
	serverHostExpected = 1
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		assert.Equal(t, dummyServeHTTPS, serveHTTPS)
		assert.Equal(t, dummyValidateClientCert, validateClientCert)
		assert.Equal(t, dummyAppPort, appPort)
		return nil
	}
	loggerAppRootExpected = 2
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doApplicationStarting", subcategory)
		if loggerAppRootCalled == 1 {
			assert.Equal(t, "Trying to start server (v-%v)", messageFormat)
			assert.Equal(t, 1, len(parameters))
			assert.Equal(t, dummyAppVersion, parameters[0])
		} else if loggerAppRootCalled == 2 {
			assert.Equal(t, "Server hosting terminated", messageFormat)
			assert.Equal(t, 0, len(parameters))
		}
	}

	// SUT + act
	doApplicationStarting()

	// verify
	verifyAll(t)
}

func TestDoApplicationClosing_NilAppClosingFunc(t *testing.T) {
	// arrange
	var appClosingFuncExpected int
	var appClosingFuncCalled int

	// stub
	customization.AppClosingFunc = nil

	// mock
	createMock(t)

	// expect
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doApplicationClosing", subcategory)
		assert.Equal(t, "customization.AppClosingFunc is not configured; skipped execution", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	doApplicationClosing()

	// verify
	verifyAll(t)
	assert.Equal(t, appClosingFuncExpected, appClosingFuncCalled, "Unexpected number of calls to AppClosingFunc")
}

func TestDoApplicationClosing_AppClosingError(t *testing.T) {
	// arrange
	var appClosingFuncExpected int
	var appClosingFuncCalled int
	var dummyClosingError = errors.New("some closing error")

	// mock
	createMock(t)

	// expect
	appClosingFuncExpected = 1
	customization.AppClosingFunc = func() error {
		appClosingFuncCalled++
		return dummyClosingError
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doApplicationClosing", subcategory)
		assert.Equal(t, "Failed to execute customization.AppClosingFunc. Error: %v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyClosingError, parameters[0])
	}

	// SUT + act
	doApplicationClosing()

	// verify
	verifyAll(t)
	assert.Equal(t, appClosingFuncExpected, appClosingFuncCalled, "Unexpected number of calls to AppClosingFunc")
}

func TestDoApplicationClosing_AppClosingSuccess(t *testing.T) {
	// arrange
	var appClosingFuncExpected int
	var appClosingFuncCalled int

	// mock
	createMock(t)

	// expect
	appClosingFuncExpected = 1
	customization.AppClosingFunc = func() error {
		appClosingFuncCalled++
		return nil
	}
	loggerAppRootExpected = 1
	loggerAppRoot = func(category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
		assert.Equal(t, "application", category)
		assert.Equal(t, "doApplicationClosing", subcategory)
		assert.Equal(t, "customization.AppClosingFunc executed successfully", messageFormat)
		assert.Equal(t, 0, len(parameters))
	}

	// SUT + act
	doApplicationClosing()

	// verify
	verifyAll(t)
	assert.Equal(t, appClosingFuncExpected, appClosingFuncCalled, "Unexpected number of calls to AppClosingFunc")
}

func TestStart_PreBoostrapExit(t *testing.T) {
	// mock
	createMock(t)

	// expect
	doPreBootstrapingFuncExpected = 1
	doPreBootstrapingFunc = func() bool {
		doPreBootstrapingFuncCalled++
		return false
	}

	// SUT + act
	Start()

	// verify
	verifyAll(t)
}

func TestStart_BoostrappingExit(t *testing.T) {
	// mock
	createMock(t)

	// expect
	doPreBootstrapingFuncExpected = 1
	doPreBootstrapingFunc = func() bool {
		doPreBootstrapingFuncCalled++
		return true
	}
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() bool {
		bootstrapApplicationFuncCalled++
		return false
	}

	// SUT + act
	Start()

	// verify
	verifyAll(t)
}

func TestStart_PostBoostrapExit(t *testing.T) {
	// mock
	createMock(t)

	// expect
	doPreBootstrapingFuncExpected = 1
	doPreBootstrapingFunc = func() bool {
		doPreBootstrapingFuncCalled++
		return true
	}
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() bool {
		bootstrapApplicationFuncCalled++
		return true
	}
	doPostBootstrapingFuncExpected = 1
	doPostBootstrapingFunc = func() bool {
		doPostBootstrapingFuncCalled++
		return false
	}

	// SUT + act
	Start()

	// verify
	verifyAll(t)
}

func TestStart_RunApplication(t *testing.T) {
	// mock
	createMock(t)

	// expect
	doPreBootstrapingFuncExpected = 1
	doPreBootstrapingFunc = func() bool {
		doPreBootstrapingFuncCalled++
		return true
	}
	bootstrapApplicationFuncExpected = 1
	bootstrapApplicationFunc = func() bool {
		bootstrapApplicationFuncCalled++
		return true
	}
	doPostBootstrapingFuncExpected = 1
	doPostBootstrapingFunc = func() bool {
		doPostBootstrapingFuncCalled++
		return true
	}
	doApplicationStartingFuncExpected = 1
	doApplicationStartingFunc = func() {
		doApplicationStartingFuncCalled++
	}
	doApplicationClosingFuncExpected = 1
	doApplicationClosingFunc = func() {
		doApplicationClosingFuncCalled++
	}

	// SUT + act
	Start()

	// verify
	verifyAll(t)
}
