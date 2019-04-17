package common

import (
	"net/http"

	"github.com/google/uuid"
)

// HandleInSession wraps the HTTP handler with session related operations
func HandleInSession(
	responseWriter http.ResponseWriter,
	request *http.Request,
	endpoint string,
	action func(http.ResponseWriter, *http.Request, uuid.UUID),
) {
	var sessionID = sessionRegister(
		endpoint,
		requestGetLoginID(
			request,
		),
		requestGetCorrelationID(
			request,
		),
		requestGetAllowedLogType(
			request,
		),
		request,
		responseWriter,
	)
	defer func() {
		panicHandle(
			endpoint,
			sessionID,
			recover(),
			responseWriter,
		)
		sessionUnregister(
			sessionID,
		)
	}()
	loggerAPIEnter(
		sessionID,
		"handler",
		endpoint,
		request.Method,
	)
	action(
		responseWriter,
		request,
		sessionID,
	)
	loggerAPIExit(
		sessionID,
		"handler",
		endpoint,
		request.Method,
	)
}
