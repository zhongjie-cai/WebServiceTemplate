package main

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
)

// func pointers for injection / testing: main.go
var (
	configAppPort            = config.AppPort
	configAppVersion         = config.AppVersion
	configInitialize         = config.Initialize
	configSendClientCert     = config.SendClientCert
	configClientCertContent  = config.ClientCertContent
	configClientKeyContent   = config.ClientKeyContent
	configServeHTTPS         = config.ServeHTTPS
	configServerCertContent  = config.ServerCertContent
	configServerKeyContent   = config.ServerKeyContent
	configValidateClientCert = config.ValidateClientCert
	configCaCertContent      = config.CaCertContent
	certificateInitialize    = certificate.Initialize
	loggerAppRoot            = logger.AppRoot
	bootstrapApplicationFunc = bootstrapApplication
	connectStoragesFunc      = connectStorages
	disconnectStoragesFunc   = disconnectStorages
	serverHost               = server.Host
	apperrorWrapSimpleError  = apperror.WrapSimpleError
)
