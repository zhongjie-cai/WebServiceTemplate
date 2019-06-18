package health

import (
	"net/http"

	"github.com/google/uuid"
)

// GetHealth handles the HTTP request for getting health report
func GetHealth(
	responseWriter http.ResponseWriter,
	httpRequest *http.Request,
	sessionID uuid.UUID,
) {
	responseOk(
		sessionID,
		configAppVersion(),
		responseWriter,
	)
}
