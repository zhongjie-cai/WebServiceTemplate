package application

import (
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
)

// func pointers for injection / testing: main.go
var (
	configAppPort             = config.AppPort
	configAppVersion          = config.AppVersion
	configInitialize          = config.Initialize
	configServeHTTPS          = config.ServeHTTPS
	configServerCertContent   = config.ServerCertContent
	configServerKeyContent    = config.ServerKeyContent
	configValidateClientCert  = config.ValidateClientCert
	configCaCertContent       = config.CaCertContent
	certificateInitialize     = certificate.Initialize
	loggerInitialize          = logger.Initialize
	loggerAppRoot             = logger.AppRoot
	serverHost                = server.Host
	doPreBootstrapingFunc     = doPreBootstraping
	bootstrapApplicationFunc  = bootstrapApplication
	doPostBootstrapingFunc    = doPostBootstraping
	doApplicationStartingFunc = doApplicationStarting
	doApplicationClosingFunc  = doApplicationClosing
)
