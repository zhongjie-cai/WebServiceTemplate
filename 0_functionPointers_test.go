package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
)

var (
	configAppPortExpected            int
	configAppPortCalled              int
	configAppVersionExpected         int
	configAppVersionCalled           int
	configInitializeExpected         int
	configInitializeCalled           int
	configSendClientCertExpected     int
	configSendClientCertCalled       int
	configClientCertContentExpected  int
	configClientCertContentCalled    int
	configClientKeyContentExpected   int
	configClientKeyContentCalled     int
	configServeHTTPSExpected         int
	configServeHTTPSCalled           int
	configServerCertContentExpected  int
	configServerCertContentCalled    int
	configServerKeyContentExpected   int
	configServerKeyContentCalled     int
	configValidateClientCertExpected int
	configValidateClientCertCalled   int
	configCaCertContentExpected      int
	configCaCertContentCalled        int
	certificateInitializeExpected    int
	certificateInitializeCalled      int
	loggerAppRootExpected            int
	loggerAppRootCalled              int
	bootstrapApplicationFuncExpected int
	bootstrapApplicationFuncCalled   int
	connectStoragesFuncExpected      int
	connectStoragesFuncCalled        int
	disconnectStoragesFuncExpected   int
	disconnectStoragesFuncCalled     int
	serverHostExpected               int
	serverHostCalled                 int
	apperrorWrapSimpleErrorExpected  int
	apperrorWrapSimpleErrorCalled    int
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
	configSendClientCertExpected = 0
	configSendClientCertCalled = 0
	configSendClientCert = func() bool {
		configSendClientCertCalled++
		return false
	}
	configClientCertContentExpected = 0
	configClientCertContentCalled = 0
	configClientCertContent = func() string {
		configClientCertContentCalled++
		return ""
	}
	configClientKeyContentExpected = 0
	configClientKeyContentCalled = 0
	configClientKeyContent = func() string {
		configClientKeyContentCalled++
		return ""
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
	certificateInitialize = func(sendClientCert bool, clientCertContent string, clientKeyContent string, serveHTTPS bool, serverCertContent string, serverKeyContent string, validateClientCert bool, caCertContent string) error {
		certificateInitializeCalled++
		return nil
	}
	loggerAppRootExpected = 0
	loggerAppRootCalled = 0
	loggerAppRoot = func(sessionID uuid.UUID, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerAppRootCalled++
	}
	bootstrapApplicationFuncExpected = 0
	bootstrapApplicationFuncCalled = 0
	bootstrapApplicationFunc = func() error {
		bootstrapApplicationFuncCalled++
		return nil
	}
	connectStoragesFuncExpected = 0
	connectStoragesFuncCalled = 0
	connectStoragesFunc = func() error {
		connectStoragesFuncCalled++
		return nil
	}
	disconnectStoragesFuncExpected = 0
	disconnectStoragesFuncCalled = 0
	disconnectStoragesFunc = func() error {
		disconnectStoragesFuncCalled++
		return nil
	}
	serverHostExpected = 0
	serverHostCalled = 0
	serverHost = func(serveHTTPS bool, validateClientCert bool, appPort string) error {
		serverHostCalled++
		return nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	configAppPort = config.AppPort
	assert.Equal(t, configAppPortExpected, configAppPortCalled, "Unexpected method call to configAppPort")
	configAppVersion = config.AppVersion
	assert.Equal(t, configAppVersionExpected, configAppVersionCalled, "Unexpected method call to configAppVersion")
	configInitialize = config.Initialize
	assert.Equal(t, configInitializeExpected, configInitializeCalled, "Unexpected method call to configInitialize")
	configSendClientCert = config.SendClientCert
	assert.Equal(t, configSendClientCertExpected, configSendClientCertCalled, "Unexpected method call to configSendClientCert")
	configClientCertContent = config.ClientCertContent
	assert.Equal(t, configClientCertContentExpected, configClientCertContentCalled, "Unexpected method call to configClientCertContent")
	configClientKeyContent = config.ClientKeyContent
	assert.Equal(t, configClientKeyContentExpected, configClientKeyContentCalled, "Unexpected method call to configClientKeyContent")
	configServeHTTPS = config.ServeHTTPS
	assert.Equal(t, configServeHTTPSExpected, configServeHTTPSCalled, "Unexpected method call to configServeHTTPS")
	configServerCertContent = config.ServerCertContent
	assert.Equal(t, configServerCertContentExpected, configServerCertContentCalled, "Unexpected method call to configServerCertContent")
	configServerKeyContent = config.ServerKeyContent
	assert.Equal(t, configServerKeyContentExpected, configServerKeyContentCalled, "Unexpected method call to configServerKeyContent")
	configValidateClientCert = config.ValidateClientCert
	assert.Equal(t, configValidateClientCertExpected, configValidateClientCertCalled, "Unexpected method call to configValidateClientCert")
	configCaCertContent = config.CaCertContent
	assert.Equal(t, configCaCertContentExpected, configCaCertContentCalled, "Unexpected method call to configCaCertContent")
	certificateInitialize = certificate.Initialize
	assert.Equal(t, certificateInitializeExpected, certificateInitializeCalled, "Unexpected method call to certificateInitialize")
	loggerAppRoot = logger.AppRoot
	assert.Equal(t, loggerAppRootExpected, loggerAppRootCalled, "Unexpected method call to loggerAppRoot")
	bootstrapApplicationFunc = bootstrapApplication
	assert.Equal(t, bootstrapApplicationFuncExpected, bootstrapApplicationFuncCalled, "Unexpected method call to bootstrapApplicationFunc")
	connectStoragesFunc = connectStorages
	assert.Equal(t, connectStoragesFuncExpected, connectStoragesFuncCalled, "Unexpected method call to connectStoragesFunc")
	disconnectStoragesFunc = disconnectStorages
	assert.Equal(t, disconnectStoragesFuncExpected, disconnectStoragesFuncCalled, "Unexpected method call to disconnectStoragesFunc")
	serverHost = server.Host
	assert.Equal(t, serverHostExpected, serverHostCalled, "Unexpected method call to serverHost")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
}
