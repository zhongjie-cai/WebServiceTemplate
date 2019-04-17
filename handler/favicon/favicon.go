package favicon

import (
	"net/http"

	"github.com/google/uuid"
)

func handleFaviconLogic(
	responseWriter http.ResponseWriter,
	request *http.Request,
	sessionID uuid.UUID,
) {
	switch request.Method {
	case http.MethodGet:
		httpServeFile(
			responseWriter,
			request,
			configAppPath()+"/favicon.ico",
		)
	default:
		responseError(
			sessionID,
			apperrorGetInvalidOperation(nil),
			responseWriter,
		)
	}
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	commonHandleInSession(
		responseWriter,
		request,
		"Favicon",
		handleFaviconLogicFunc,
	)
}

// HostEntry hosts the service entry for "/favicon.ico"
func HostEntry() {
	httpHandleFunc(
		"/favicon.ico",
		handlerFunc,
	)
}
