# WebServiceTemplate
This project (for Golang) is provided as a template for quickly create any Golang web services.

Original source: https://github.com/zhongjie-cai/WebServiceTemplate

Library dependencies (must be present in vendor folder or in Go path):
* [UUID](https://github.com/google/uuid): `go get github.com/google/uuid`

A sample application is shown below:

# main.go
```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/application"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/favicon"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/health"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
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
		}
	}
	customization.Statics = func() []model.Static {
		return []model.Static{
			model.Static{
				Name:       "SwaggerUI",
				PathPrefix: "/docs/",
				Handler:    swagger.Handler(),
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
		fmt.Printf("<%v|%v> %v\n", category, subcategory, description)
	}
	application.Start()
}
```

# handler/health/health.go

```golang
package health

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
)

// GetHealth handles the HTTP request for getting health report
func GetHealth(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
	sessionID string,
) {
	response.Ok(
		sessionID,
		"some version number",
		responseWriter,
	)
}
```

# Swagger UI

Copy the swagger UI folder "/docs/" from this library to your repository root path.  
The "openapi.json" is the swagger definition (in OpenAPI v3 format).  
