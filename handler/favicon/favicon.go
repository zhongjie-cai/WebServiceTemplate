package favicon

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

// GetFavicon handles the HTTP request for getting favicon
func GetFavicon(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
	sessionID uuid.UUID,
) {
	httpServeFile(
		responseWriter,
		httpRequest,
		config.AppPath()+"/favicon.ico",
	)
}
