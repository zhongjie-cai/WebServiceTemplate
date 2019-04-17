package health

import (
	"net/http"

	"github.com/google/uuid"
)

func handleHealthLogic(
	responseWriter http.ResponseWriter,
	request *http.Request,
	sessionID uuid.UUID,
) {
	switch request.Method {
	case http.MethodGet:
		responseOk(
			sessionID,
			configAppVersion(),
			responseWriter,
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
		"Health",
		handleHealthLogicFunc,
	)
}

// HostEntry hosts the service entry for "/health"
func HostEntry() {
	httpHandleFunc(
		"/health",
		handlerFunc)
	httpHandleFunc(
		"/health/",
		handlerFunc)
}
