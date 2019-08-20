# WebServiceTemplate
This project (for Golang) is provided as a template for quickly create any Golang web services.

Original source: https://github.com/zhongjie-cai/WebServiceTemplate

Library dependencies (must be present in vendor folder or in Go path):
* [UUID](https://github.com/google/uuid): `go get github.com/google/uuid`
* [MUX](https://github.com/gorilla/mux): `go get github.com/gorilla/mux`

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
			model.Route{
				Endpoint:   "SwaggerRedirect",
				Method:     http.MethodGet,
				Path:       "/docs",
				ActionFunc: swagger.Redirect,
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
	sessionID uuid.UUID,
	requestBody string,
	parameters map[string]string,
) {
	response.Ok(
		sessionID,
		"some version number",
	)
}
```

# handler/swagger/swagger.go

```golang
package swagger

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// Redirect handles HTTP redirection for swagger UI requests
func Redirect(
	sessionID uuid.UUID,
	requestBody string,
	parameters map[string]string,
) {
	var httpRequest = session.GetRequest(
		sessionID,
	)
	var responseWriter = session.GetResponseWriter(
		sessionID,
	)
	http.Redirect(
		responseWriter,
		httpRequest,
		"/docs/",
		http.StatusPermanentRedirect,
	)
}

// Handler handles the hosting of the swagger UI static content
func Handler() http.Handler {
	return http.StripPrefix(
		"/docs/",
		http.FileServer(
			http.Dir(
				config.AppPath()+"/docs",
			),
		),
	)
}
```

# Request & Response

The registered handler receives the request body as a string, thus it is normally not necessary to load request from session.
However, if specific data is needed from request, one could always retrieve request from session through following function call using sessionID:

```golang
var httpRequest = session.GetRequest(sessionID)
```

The response functions accept the session ID and internally load the response writer accordingly, thus it is normally not necessary to load response writer from session.
However, if specific operation is needed for response, one could always retrieve response writer through following function call using sessionID:

```golang
var responseWriter = session.GetResponseWriter(sessionID)
```

# Swagger UI

Copy the swagger UI folder "/docs/" from this library to your repository root path.  
The "openapi.json" is the swagger definition (in OpenAPI v3 format).  
