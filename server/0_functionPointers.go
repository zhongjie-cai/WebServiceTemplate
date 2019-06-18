package server

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/register"
)

// func pointers for injection / testing: server.go
var (
	certificateGetServerCertificate = certificate.GetServerCertificate
	certificateGetClientCertPool    = certificate.GetClientCertPool
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	registerInstantiate             = register.Instantiate
	loggerAppRoot                   = logger.AppRoot
	createServerFunc                = createServer
	listenAndServeFunc              = listenAndServe
	runServerFunc                   = runServer
)
