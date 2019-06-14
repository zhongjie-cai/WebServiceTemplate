package health

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleGetHealth(
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
			responseOk(
				sessionID,
				configAppVersion(),
				responseWriter,
			)
		},
	)
}

func handleGetHealthReport(
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
			responseOk(
				sessionID,
				configAppVersion(),
				responseWriter,
			)
		},
	)
}

// HostEntry hosts the service entry for "/health"
func HostEntry(router *mux.Router) {
	routeHandleFunc(
		router,
		"Health",
		http.MethodGet,
		"/health",
		handleGetHealthFunc,
	)
	routeHandleFunc(
		router,
		"HealthReport",
		http.MethodGet,
		"/health/report",
		handleGetHealthReportFunc,
	)
}
