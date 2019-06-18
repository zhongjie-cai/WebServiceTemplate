package swagger

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

// Redirect handles HTTP redirection for swagger UI requests
func Redirect(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
	sessionID uuid.UUID,
) {
	httpRedirect(
		responseWriter,
		httpRequest,
		"/docs/",
		http.StatusPermanentRedirect,
	)
}

// Handler handles the hosting of the swagger UI static content
func Handler() http.Handler {
	return httpStripPrefix(
		"/docs/",
		httpFileServer(
			http.Dir(
				config.AppPath()+"/docs",
			),
		),
	)
}
