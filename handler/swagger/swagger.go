package swagger

import (
	"net/http"
)

func redirectHandler(responseWriter http.ResponseWriter, request *http.Request) {
	httpRedirect(
		responseWriter,
		request,
		"/docs/index.html",
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
func HostEntry() {
	httpHandleFunc(
		"/docs",
		redirectHandlerFunc)
	httpHandle(
		"/docs/",
		contentHandlerFunc())
}
