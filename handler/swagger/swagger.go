package swagger

import (
	"net/http"

	"github.com/gorilla/mux"
)

func redirectHandler(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	httpRedirect(
		responseWriter,
		request,
		"/docs/",
		http.StatusPermanentRedirect,
	)
}

func contentHandler() http.Handler {
	return httpStripPrefix(
		"/docs/",
		httpFileServer(
			http.Dir(configAppPath()+"/docs")))
}

// HostEntry hosts the service entry for "/docs"
func HostEntry(router *mux.Router) {
	routeHandleFunc(
		router,
		"SwaggerUI",
		http.MethodGet,
		"/docs",
		redirectHandlerFunc,
	)
	routeHostStatic(
		router,
		"SwaggerUI",
		"/docs/",
		contentHandlerFunc(),
	)
}
