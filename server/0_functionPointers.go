package server

import (
	"context"
	"os/signal"

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
	apperrorConsolidateAllErrors    = apperror.ConsolidateAllErrors
	registerInstantiate             = register.Instantiate
	loggerAppRoot                   = logger.AppRoot
	signalNotify                    = signal.Notify
	contextWithTimeout              = context.WithTimeout
	contextBackground               = context.Background
	createServerFunc                = createServer
	listenAndServeFunc              = listenAndServe
	shutDownFunc                    = shutDown
	runServerFunc                   = runServer
)
