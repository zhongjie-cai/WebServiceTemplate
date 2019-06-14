package server

import (
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/certificate"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
	"github.com/zhongjie-cai/WebServiceTemplate/server/route"
)

// func pointers for injection / testing: server.go
var (
	certificateGetServerCertificate = certificate.GetServerCertificate
	certificateGetClientCertPool    = certificate.GetClientCertPool
	apperrorWrapSimpleError         = apperror.WrapSimpleError
	faviconHostEntry                = favicon.HostEntry
	swaggerHostEntry                = swagger.HostEntry
	healthHostEntry                 = health.HostEntry
	routeRegisterEntries            = route.RegisterEntries
	createServerFunc                = createServer
	listenAndServeFunc              = listenAndServe
	runServerFunc                   = runServer
)
