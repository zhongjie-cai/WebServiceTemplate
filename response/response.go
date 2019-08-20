package response

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// These are the constants used by the HTTP modules
const (
	ContentTypeJSON = "application/json; charset=utf-8"
	ContentTypeXML  = "application/xml; charset=utf-8"
)

// errorResponseModel defines the response object that is written back to consumer of the API
type errorResponseModel struct {
	Code     int      `json:"code"`
	Type     string   `json:"type"`
	Messages []string `json:"messages"`
}

func getAppError(err error) apperror.AppError {
	var appError, isAppError = err.(apperror.AppError)
	if !isAppError {
		appError = apperrorGetGeneralFailureError(err)
	}
	return appError
}

func getStatusCode(appError apperror.AppError) int {
	var statusCode int
	switch appError.Code() {
	case apperror.CodeGeneralFailure:
		statusCode = http.StatusInternalServerError
	case apperror.CodeInvalidOperation:
		statusCode = http.StatusMethodNotAllowed
	case apperror.CodeBadRequest:
		statusCode = http.StatusBadRequest
	case apperror.CodeCircuitBreak:
		statusCode = http.StatusBadRequest
	case apperror.CodeOperationLock:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}

func createOkResponse(
	responseContent interface{},
) (string, int) {
	if responseContent == nil {
		return "", http.StatusNoContent
	}
	var responseMessage = jsonutilMarshalIgnoreError(responseContent)
	if responseMessage == "" {
		return "", http.StatusNoContent
	}
	return responseMessage, http.StatusOK
}

func generateErrorResponse(
	appError apperror.AppError,
) errorResponseModel {
	var code = appError.Code()
	var messages = appError.Messages()
	var response = errorResponseModel{
		Code:     int(code),
		Type:     code.String(),
		Messages: messages,
	}
	return response
}

func createErrorResponse(
	appError apperror.AppError,
) (string, int) {
	var response = generateErrorResponseFunc(appError)
	var responseMessage = jsonutilMarshalIgnoreError(response)
	var statusCode = getStatusCodeFunc(appError)
	return responseMessage, statusCode
}

func writeResponse(
	responseWriter http.ResponseWriter,
	statusCode int,
	responseMessage string,
) {
	responseWriter.Header().Set("Content-Type", ContentTypeJSON)
	responseWriter.WriteHeader(statusCode)
	responseWriter.Write([]byte(responseMessage))
}

// Ok responds to the consumer with HTTP-2xx status code and status message
func Ok(
	sessionID uuid.UUID,
	responseContent interface{},
) {
	var responseWriter = sessionGetResponseWriter(
		sessionID,
	)
	var responseMessage, statusCode = createOkResponseFunc(
		responseContent,
	)
	loggerAPIResponse(
		sessionID,
		"response",
		strconvItoa(statusCode),
		responseMessage,
	)
	writeResponseFunc(
		responseWriter,
		statusCode,
		responseMessage,
	)
	loggerAPIExit(
		sessionID,
		"response",
		"Ok",
		"",
	)
}

// Error responds to the consumer with HTTP-4xx or HTTP-5xx status code and status message
func Error(
	sessionID uuid.UUID,
	err error,
) {
	var responseWriter = sessionGetResponseWriter(
		sessionID,
	)
	var appError = getAppErrorFunc(
		err,
	)
	var responseMessage, statusCode = createErrorResponseFunc(
		appError,
	)
	loggerAPIResponse(
		sessionID,
		"response",
		strconvItoa(statusCode),
		responseMessage,
	)
	writeResponseFunc(
		responseWriter,
		statusCode,
		responseMessage,
	)
	loggerAPIExit(
		sessionID,
		"response",
		"Error",
		"",
	)
}
