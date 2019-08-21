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
		statusCode = http.StatusForbidden
	case apperror.CodeOperationLock:
		statusCode = http.StatusLocked
	case apperror.CodeAccessForbidden:
		statusCode = http.StatusForbidden
	case apperror.CodeDataCorruption:
		statusCode = http.StatusConflict
	case apperror.CodeNotImplemented:
		statusCode = http.StatusNotImplemented
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

// Write responds to the consumer with corresponding HTTP status code and response body
func Write(
	sessionID uuid.UUID,
	responseObject interface{},
	responseError apperror.AppError,
) {
	var responseWriter = sessionGetResponseWriter(
		sessionID,
	)
	var responseMessage string
	var statusCode int
	if responseError != nil {
		responseMessage, statusCode = createErrorResponseFunc(
			responseError,
		)
	} else {
		responseMessage, statusCode = createOkResponseFunc(
			responseObject,
		)
	}
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
		"Write",
		"%v",
		statusCode,
	)
}

// Override overrides the default response.Write functionality by the given callback function; consumers must manually deal with response writer accordingly
func Override(
	sessionID uuid.UUID,
	callback func(*http.Request, http.ResponseWriter),
) {
	var httpRequest = sessionGetRequest(
		sessionID,
	)
	var responseWriter = sessionGetResponseWriter(
		sessionID,
	)
	callback(
		httpRequest,
		responseWriter,
	)
	sessionClearResponseWriter(
		sessionID,
	)
}