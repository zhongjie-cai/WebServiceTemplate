package health

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
)

// GetHealth handles the HTTP request for getting health report
func GetHealth(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
	sessionID uuid.UUID,
) {
	responseOk(
		sessionID,
		config.AppVersion(),
		responseWriter,
	)
}
