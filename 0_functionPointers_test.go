package main

import (
	"fmt"
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
	configClientCertContentExpected  int
	configClientCertContentCalled    int
	configClientKeyContentExpected   int
	configClientKeyContentCalled     int
	configServerCertContentExpected  int
	configServerCertContentCalled    int
	configServerKeyContentExpected   int
	configServerKeyContentCalled     int
	configCACertContentExpected      int
	configCACertContentCalled        int
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
	configCACertContentExpected = 0
	configCACertContentCalled = 0
	configCACertContent = func() string {
		configCACertContentCalled++
		return ""
	}
	certificateInitializeExpected = 0
	certificateInitializeCalled = 0
	certificateInitialize = func(clientCertContent string, clientKeyContent string, serverCertContent string, serverKeyContent string, caCertContent string) error {
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
	serverHost = func() error {
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
	if configAppPortExpected != configAppPortCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppPort, expected %v, actual %v", configAppPortExpected, configAppPortCalled))
	}
	configAppVersion = config.AppVersion
	if configAppVersionExpected != configAppVersionCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configAppVersion, expected %v, actual %v", configAppVersionExpected, configAppVersionCalled))
	}
	configInitialize = config.Initialize
	if configInitializeExpected != configInitializeCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configInitialize, expected %v, actual %v", configInitializeExpected, configInitializeCalled))
	}

	configClientCertContent = config.ClientCertContent
	if configClientCertContentExpected != configClientCertContentCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configClientCertContent, expected %v, actual %v", configClientCertContentExpected, configClientCertContentCalled))
	}
	configClientKeyContent = config.ClientKeyContent
	if configClientKeyContentExpected != configClientKeyContentCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configClientKeyContent, expected %v, actual %v", configClientKeyContentExpected, configClientKeyContentCalled))
	}
	configServerCertContent = config.ServerCertContent
	if configServerCertContentExpected != configServerCertContentCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configServerCertContent, expected %v, actual %v", configServerCertContentExpected, configServerCertContentCalled))
	}
	configServerKeyContent = config.ServerKeyContent
	if configServerKeyContentExpected != configServerKeyContentCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configServerKeyContent, expected %v, actual %v", configServerKeyContentExpected, configServerKeyContentCalled))
	}
	configCACertContent = config.CACertContent
	if configCACertContentExpected != configCACertContentCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to configCACertContent, expected %v, actual %v", configCACertContentExpected, configCACertContentCalled))
	}

	certificateInitialize = certificate.Initialize
	if certificateInitializeExpected != certificateInitializeCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to certificateInitialize, expected %v, actual %v", certificateInitializeExpected, certificateInitializeCalled))
	}
	loggerAppRoot = logger.AppRoot
	if loggerAppRootExpected != loggerAppRootCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to loggerAppRoot, expected %v, actual %v", loggerAppRootExpected, loggerAppRootCalled))
	}
	bootstrapApplicationFunc = bootstrapApplication
	if bootstrapApplicationFuncExpected != bootstrapApplicationFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to bootstrapApplicationFunc, expected %v, actual %v", bootstrapApplicationFuncExpected, bootstrapApplicationFuncCalled))
	}
	connectStoragesFunc = connectStorages
	if connectStoragesFuncExpected != connectStoragesFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to connectStoragesFunc, expected %v, actual %v", connectStoragesFuncExpected, connectStoragesFuncCalled))
	}
	disconnectStoragesFunc = disconnectStorages
	if disconnectStoragesFuncExpected != disconnectStoragesFuncCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to disconnectStoragesFunc, expected %v, actual %v", disconnectStoragesFuncExpected, disconnectStoragesFuncCalled))
	}
	serverHost = server.Host
	if serverHostExpected != serverHostCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to serverHost, expected %v, actual %v", serverHostExpected, serverHostCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
}
