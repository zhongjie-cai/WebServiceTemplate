package server

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
)

// func pointers for injection / testing: server.go
var (
	configAppPort                   = config.AppPort
	certificateGetServerCertificate = certificate.GetServerCertificate
	certificateGetClientCertPool    = certificate.GetClientCertPool
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	faviconHostEntry                = favicon.HostEntry
	swaggerHostEntry                = swagger.HostEntry
	healthHostEntry                 = health.HostEntry
	createServerFunc                = createServer
	listenAndServeTLSFunc           = listenAndServeTLS
	hostEntriesFunc                 = hostEntries
	runServerFunc                   = runServer
)
