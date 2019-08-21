package panic

import (
	"net/http"
	"runtime/debug"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"

	"github.com/google/uuid"
)

func getRecoverError(recoverResult interface{}) apperror.AppError {
	var err, ok = recoverResult.(error)
	if !ok {
		err = fmtErrorf("%v", recoverResult)
	}
	return apperrorGetGeneralFailureError(err)
}

func getDebugStack() string {
	return string(debug.Stack())
}

// Handle prevents the application from halting when service handler panics unexpectedly
func Handle(endpointName string, sessionID uuid.UUID, recoverResult interface{}, responseWriter http.ResponseWriter) {
	if recoverResult != nil {
		var appError = getRecoverErrorFunc(
			recoverResult,
		)
		responseWrite(
			sessionID,
			nil,
			appError,
		)
		loggerAppRoot(
			"panic",
			"Handle",
			"%v\n%v",
			appError,
			getDebugStackFunc(),
		)
	}
}
