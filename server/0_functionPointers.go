package server

import (
	"context"
	"os/signal"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/register"
)

// func pointers for injection / testing: server.go
var (
	certificateGetServerCertificate = certificate.GetServerCertificate
	certificateGetCaCertPool        = certificate.GetCaCertPool
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	registerInstantiate             = register.Instantiate
	loggerAppRoot                   = logger.AppRoot
	signalNotify                    = signal.Notify
	contextWithTimeout              = context.WithTimeout
	contextBackground               = context.Background
	configGraceShutdownWaitTime     = config.GraceShutdownWaitTime
	createServerFunc                = createServer
	listenAndServeFunc              = listenAndServe
	shutDownFunc                    = shutDown
	consolidateErrorFunc            = consolidateError
	runServerFunc                   = runServer
)
