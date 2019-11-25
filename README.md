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
		return "1.2.3"
	}
	customization.LoggingFunc = func(session *session.Session, logType logtype.LogType, category, subcategory, description string) {
		fmt.Printf("<%v|%v> %v\n", category, subcategory, description)
	}
	customization.Middlewares = func() []model.MiddlewareFunc {
		return []model.MiddlewareFunc{
			loggingRequestURIMiddleware,
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
) (interface{}, apperror.AppError) {
	return "some version number", nil
}

// swaggerRedirect is an example of how a special HTTP handling method, which overrides the default library behavior, is written with this template library
func swaggerRedirect(
	sessionID uuid.UUID,
) (interface{}, apperror.AppError) {
	return response.Override(
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

// loggingRequestURIMiddleware is an example of how a middleware function is written with this template library
func loggingRequestURIMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			responseWriter http.ResponseWriter,
			httpRequest *http.Request,
		) {
			// middleware logic & processing
			fmt.Println(
				httpRequest.RequestURI,
			)
			// hand over to next handler in the chain
			nextHandler.ServeHTTP(
				responseWriter,
				httpRequest,
			)
		},
	)
}
```

# Request & Response

The registered handler could retrieve request body, parameters and query strings through session methods, thus it is normally not necessary to load request from session:

```golang
// request body: {"foo":"bar","test":123}
var body struct {
	Foo string `json:"foo"`
	Test int `json:"test"`
}
var bodyError = session.GetRequestBody(sessionID, &body)

// parameters: "id"=456
var id int
var idError = session.GetRequestParameter(sessionID, "id", &id)

// query strigns: "uuid"="123456-1234-1234-1234-123456789abc"
var uuid uuid.UUID
var uuidError = session.GetRequestQueryString(sessionID, "uuid", &uuid)
```

However, if specific data is needed from request, one could always retrieve request from session through following function call using sessionID:

```golang
var httpRequest = session.GetRequest(sessionID)
```

The response functions accept the session ID and internally load the response writer accordingly, thus it is normally not necessary to load response writer from session.

However, if specific operation is needed for response, one could always retrieve response writer through following function call using sessionID:

```golang
var responseWriter = session.GetResponseWriter(sessionID)
```

# Error Handling

To simplify the error handling, one could utilize the built-in error type `apperror.AppError` interface, which provides support to many basic types of errors that are mapped to corresponding HTTP status codes:

* GeneralFailure => InternalServerError (500)
* InvalidOperation => MethodNotAllowed (405)
* BadRequest => BadRequest (400)
* NotFound => NotFound (404)
* CircuitBreak => Forbidden (403)
* OperationLock => Locked (423)
* AccessForbidden => Forbidden (403)
* DataCorruption => Conflict (409)
* NotImplemented => NotImplemented (501)

However, if specific operation is needed for response, one could always customize the error response creation by setting the `customization.CreateErrorResponseFunc` function:

```golang
customization.CreateErrorResponseFunc = func(err error) (responseMessage string, statusCode int) {
	return err.Error(), 500
}
```

# Session Attachment

The registered session contains an attachment dictionary, which allows the user to attach any object which is JSON serializable into the given session associated to a session ID.

```golang
var myAttachmentName = "my attachment name"
var myAttachmentObject = anyJSONSerializableStruct {
	...
}
var success = session.Attach(sessionID, myAttachmentName, myAttachmentObject)
if !success {
	// failed to attach an object: add your customized logic here if needed
} else {
	// succeeded to attach an object: add your customized logic here if needed
}
```

To retrieve a previously attached object from session, simply use the following sample logic.

```golang
var myAttachmentName = "my attachment name"
var retrievedAttachment anyJSONSerializableStruct
var success = session.GetAttachment(sessionID, myAttachmentName, &retrievedAttachment)
if !success {
	// failed to retrieve an attachment: add your customized logic here if needed
} else {
	// succeeded to retrieve an attachment: add your customized logic here if needed
}
```

In some situations, it is good to detach a certain attachment, especially if it is a big object consuming large memory, which can be done as following.

```golang
var myAttachmentName = "my attachment name"
var success = session.Detach(sessionID, myAttachmentName)
if !success {
	// failed to detach an attachment: add your customized logic here if needed
} else {
	// succeeded to detach an attachment: add your customized logic here if needed
}
```

# Swagger UI

Copy the swagger UI folder "/docs/" from this library to your repository root path.  
The "openapi.json" is the swagger definition (in OpenAPI v3 format).  
