package application

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// func pointers for injection / testing: main.go
var (
	sessionInit               = session.Init
	configInitialize          = config.Initialize
	certificateInitialize     = certificate.Initialize
	apperrorInitialize        = apperror.Initialize
	loggerInitialize          = logger.Initialize
	loggerAppRoot             = logger.AppRoot
	serverHost                = server.Host
	doPreBootstrapingFunc     = doPreBootstraping
	bootstrapApplicationFunc  = bootstrapApplication
	doPostBootstrapingFunc    = doPostBootstraping
	doApplicationStartingFunc = doApplicationStarting
	doApplicationClosingFunc  = doApplicationClosing
)
