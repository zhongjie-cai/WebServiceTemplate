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

	"github.com/zhongjie-cai/WebServiceTemplate/session"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/application"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
	"github.com/zhongjie-cai/WebServiceTemplate/response"
	serverModel "github.com/zhongjie-cai/WebServiceTemplate/server/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
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
	customization.LoggingFunc = func(session sessionModel.Session, logType logtype.LogType, logLevel loglevel.LogLevel, category, subcategory, description string) {
		fmt.Printf("<%v|%v> %v\n", category, subcategory, description)
	}
	customization.Middlewares = func() []serverModel.MiddlewareFunc {
		return []serverModel.MiddlewareFunc{
			loggingRequestURIMiddleware,
		}
	}
	customization.Statics = func() []serverModel.Static {
		return []serverModel.Static{
			serverModel.Static{
				Name:       "SwaggerUI",
				PathPrefix: "/docs/",
				Handler:    swaggerHandler(),
			},
		}
	}
	customization.Routes = func() []serverModel.Route {
		return []serverModel.Route{
			serverModel.Route{
				Endpoint:   "Health",
				Method:     http.MethodGet,
				Path:       "/health",
				ActionFunc: getHealth,
			},
			serverModel.Route{
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
) (interface{}, error) {
	var appVersion = "some application version"
	session.LogMethodLogic(
		sessionID,
		loglevel.Warn,
		"Health",
		"Message",
		"AppVersion = %v",
		appVersion,
	)
	return appVersion, nil
}

// swaggerRedirect is an example of how a special HTTP handling method, which overrides the default library behavior, is written with this template library
func swaggerRedirect(
	sessionID uuid.UUID,
) (interface{}, error) {
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

# Logging

The library allows the user to customize its logging function by setting the variable `LoggingFunc` under the `customization` package. 
The logging is split into two management areas: log type and log level. 

## Log Type

The log type definitions can be found under the `logtype` sub-package under the `logging` package. 
Apart from all `Method`-prefixed log types, all remainig log types are managed by the library internally and should not be worried by the user. 

The application uses `AllowedLogType` concept to filter out any logging that does not fall into the allowed log type list. 
To configure the `AllowedLogType`, the user can:
- set the variable `DefaultAllowedLogType` under the `customization` package, which affects default logging types for all application (if no session level configuration found)
- set the variable `SessionAllowedLogType` under the `customization` package, which affects session logging types for all sessions (if configured)

## Log Level

The log level definitions can be found under the `loglevel` sub-package under the `logging` package. 
Log level only affects all `Method`-prefixed log types, and the log level is per session configurable. 

The application uses `DefaultAllowedLogLevel` concept to filter out any logging that does not fall into the allowed log level list. 
To configure the `AllowedLogType`, the user can:
- set the variable `DefaultAllowedLogLevel` under the `customization` package, which affects default logging levels for all application (if no session level configuration found)
- set the variable `SessionAllowedLogLevel` under the `customization` package, which affects session logging levels for all sessions (if configured)

## Session Logging

The registered session allows the user to add manual logging to its codebase, through several listed methods as
```golang
session.LogMethodEnter(sessionID uuid.UUID)
session.LogMethodParameter(sessionID uuid.UUID, parameters ...interface{})
session.LogMethodLogic(sessionID uuid.UUID, logLevel loglevel.LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})
session.LogMethodReturn(sessionID uuid.UUID, returns ...interface{})
session.LogMethodExit(sessionID uuid.UUID)
```

The `Enter`, `Parameter`, `Return` and `Exit` are limited to the scope of method boundary area loggings. 
The `Logic` is the normal logging that can be used in any place at any level in the codebase to enforce the user's customized logging entries.

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

# External Web Requests

The library provides a way to send out HTTP/HTTPS requests to external web services based on current session. 
Using this provided feature ensures the logging of the web service requests into corresponding log type for the given session. 

```golang
...

var networkRequest = session.CreateNetworkRequest(
	sessionID,
	HTTP.POST,                       // Method
	"https://www.example.com/tests", // URL
	"{\"foo\":\"bar\"}",             // Payload
	map[string]string{               // Headers
		"Content-Type": "application/json",
		"Accept": "application/json",
	},
)
var testSample testSampleStruct
var statusCode, responseHeader, responseError = networkRequest.Process(
	&testSample,
)

...
```

Network requests would send out client certificate for mTLS communications if the following customization is in place.

```golang
customization.SendClientCert = func() bool { return true }
customization.ClientCertContent = func() string { return "your client certificate content" }
customization.ClientKeyContent = func() string { return "your client key content" }
```

Network requests could also be customized for：

## HTTP Client's HTTP Transport (http.RoundTripper)

This is to enable the 3rd party monitoring libraries, e.g. new relic, to wrap the HTTP transport for better handling of network communications. 

```golang
customization.HTTPRoundTripper = func(originalTransport http.RoundTripper) http.RoundTripper {
	return ... 
}
```

# HTTP Request (http.Request)

This is to enable the 3rd party monitoring libraries, e.g. new relic, to wrap individual HTTP request for better handling of web requests.

```golang
customization.WrapHTTPRequest = func(session sessionModel.Session, httpRequest *http.Request) *http.Request {
	return ...
}
```

# Network Timeout

This is to provide the default HTTP request timeouts for HTTP Client over all network communications.

```golang
customization.DefaultNetworkTimeout = func() time.Duration {
	return ...
}
```
