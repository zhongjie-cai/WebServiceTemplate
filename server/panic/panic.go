package panic

import (
	"runtime/debug"

	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
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
func Handle(session sessionModel.Session, recoverResult interface{}) {
	if recoverResult != nil {
		var appError = getRecoverErrorFunc(
			recoverResult,
		)
		responseWrite(
			session,
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
