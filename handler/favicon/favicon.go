package favicon

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleGetFavicon(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	commonHandleInSession(
		responseWriter,
		request,
		func(
			responseWriter http.ResponseWriter,
			request *http.Request,
			sessionID uuid.UUID,
		) {
			httpServeFile(
				responseWriter,
				request,
				configAppPath()+"/favicon.ico",
			)
		},
	)
}

// HostEntry hosts the service entry for "/favicon.ico"
func HostEntry(router *mux.Router) {
	routeHandleFunc(
		router,
		"Favicon",
		http.MethodGet,
		"/favicon.ico",
		handleGetFaviconFunc,
	)
}
