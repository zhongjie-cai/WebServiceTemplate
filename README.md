# WebServiceTemplate
This project (for Golang) is provided as a template for quickly create any Golang web services.

Original source: https://github.com/zhongjie-cai/WebServiceTemplate

Library dependencies (must be present in vendor folder or in Go path):
* [UUID](https://github.com/google/uuid): `go get -u github.com/google/uuid`
* [MUX](https://github.com/gorilla/mux): `go get -u github.com/gorilla/mux`
* [Cache](https://github.com/patrickmn/go-cache): `go get -u github.com/patrickmn/go-cache`
* [Testify](https://github.com/stretchr/testify): `go get -u github.com/stretchr/testify` (For tests only)

A sample application is shown below:

# main.go
```golang
package main

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
	
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
	"github.com/zhongjie-cai/WebServiceTemplate/application"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	"github.com/zhongjie-cai/WebServiceTemplate/server/model"
	"github.com/zhongjie-cai/WebServiceTemplate/session"
)

// This is a sample of how to setup application for running the server
func main() {
	customization.AppName = func() string {
		return "WebServiceTemplate"
	}
	customization.AppPort = func() string {
		return "18605"
	}
	customization.AppVersion = func() string {
		return "0.1.0"
	}
	customization.LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		fmt.Printf("<%v|%v> %v\n", category, subcategory, description)
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
	customization.Routes = func() []model.Route {
		return []model.Route{
			model.Route{
				Endpoint:   "Health",
				Method:     http.MethodGet,
				Path:       "/health",
				ActionFunc: getHealth,
			},
			model.Route{
				Endpoint:   "SwaggerRedirect",
				Method:     http.MethodGet,
				Path:       "/docs",
				ActionFunc: swaggerRedirect,
			},
		}
	}
	application.Start()
}

// getHealth is an example of how a normal HTTP handling method is written with this template library
func getHealth(
	sessionID uuid.UUID,
	requestBody string,
	parameters map[string]string,
) (interface{}, apperror.AppError) {
	return "some version number", nil
}

// swaggerRedirect is an example of how a special HTTP handling method, which overrides the default library behavior, is written with this template library
func swaggerRedirect(
	sessionID uuid.UUID,
	requestBody string,
	parameters map[string]string,
) (interface{}, apperror.AppError) {
	response.Override(
		sessionID,
		func(
			httpRequest *http.Request,
			responseWriter http.ResponseWriter,
		) {
			http.Redirect(
				responseWriter,
				httpRequest,
				"/docs/",
				http.StatusPermanentRedirect,
			)
		},
	)
}

// swaggerHandler is an example of how a normal HTTP static content hosting is written with this template library
func swaggerHandler() http.Handler {
	return http.StripPrefix(
		"/docs/",
		http.FileServer(
			http.Dir(
				"./docs",
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
