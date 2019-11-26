package panic

import (
	"net/http"
	"runtime/debug"

	"github.com/google/uuid"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
)

func getRecoverError(recoverResult interface{}) apperrorModel.AppError {
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
