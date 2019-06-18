package main

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

var (
	appVersion string
	appName    string
	appPath    string
	appPort    string
)

// This is a sample of how to setup application for running the server
func main() {
	customization.Routes = func() []model.Route {
		return []model.Route{
			model.Route{
				Endpoint:   "Health",
				Method:     http.MethodGet,
				Path:       "/health",
				ActionFunc: health.GetHealth,
			},
			model.Route{
				Endpoint:   "SwaggerRedirect",
				Method:     http.MethodGet,
				Path:       "/docs",
				ActionFunc: swagger.Redirect,
			},
			model.Route{
				Endpoint:   "Favicon",
				Method:     http.MethodGet,
				Path:       "/favicon.ico",
				ActionFunc: favicon.GetFavicon,
			},
		}
	}
	customization.Statics = func() []model.Static {
		return []model.Static{
			model.Static{
				Name:       "SwaggerUI",
				PathPrefix: "/docs/",
				Handler:    swaggerHandler(),
			},
		}
	}
	customization.AppName = func() string {
		return appName
	}
	customization.AppPort = func() string {
		return appPort
	}
	customization.AppVersion = func() string {
		return appVersion
	}
	customization.AppPath = func() string {
		return appPath
	}
	customization.IsLocalhost = func() bool {
		return true
	}
	customization.LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		fmtPrintf("<%v|%v> %v\n", category, subcategory, description)
	}
	applicationStart()
}
