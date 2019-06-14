package common

import (
	"net/http"

	"github.com/google/uuid"
)

const (
	// RegexpForUUID is used when registering routes with path variables that use UUIDs
	RegexpForUUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
)

// HandleInSession wraps the HTTP handler with session related operations
func HandleInSession(
	responseWriter http.ResponseWriter,
	request *http.Request,
	action func(http.ResponseWriter, *http.Request, uuid.UUID),
) {
	var endpoint = routeGetEndpointName(
		request,
	)
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
