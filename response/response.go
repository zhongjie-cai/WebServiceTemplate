package response

import (
	"net/http"

	"github.com/google/uuid"
	apperrorModel "github.com/zhongjie-cai/WebServiceTemplate/apperror/model"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

// These are the constants used by the HTTP modules
const (
	ContentTypeJSON = "application/json; charset=utf-8"
)

// errorResponseModel defines the response object that is written back to consumer of the API
type errorResponseModel struct {
	Code      string            `json:"code"`
	Messages  []string          `json:"messages"`
	ExtraData map[string]string `json:"extraData,omitempty"`
}

// overrideResponse defines a dummy response returned by override to suppress logging
type overrideResponse struct{}

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

func getAppError(
	err error,
) apperrorModel.AppError {
	var appError, isAppError = err.(apperrorModel.AppError)
	if isAppError {
		return appError
	}
	return apperrorGetGeneralFailureError(
		err,
	)
}

func generateErrorResponse(
	appError apperrorModel.AppError,
) errorResponseModel {
	var code = appError.Code()
	var messages = appError.Messages()
	var extraData = appError.ExtraData()
	var response = errorResponseModel{
		Code:      code,
		Messages:  messages,
		ExtraData: extraData,
	}
	return response
}

func createErrorResponse(
	err error,
) (string, int) {
	var appError = getAppErrorFunc(err)
	var response = generateErrorResponseFunc(appError)
	var responseMessage = jsonutilMarshalIgnoreError(response)
	var statusCode = appError.HTTPStatusCode()
	return responseMessage, statusCode
}

func writeResponse(
	sessionID uuid.UUID,
	responseWriter http.ResponseWriter,
	statusCode int,
	responseMessage string,
) {
	loggerAPIResponse(
		sessionID,
		"response",
		strconvItoa(statusCode),
		responseMessage,
	)
	responseWriter.Header().Set("Content-Type", ContentTypeJSON)
	responseWriter.WriteHeader(statusCode)
	responseWriter.Write([]byte(responseMessage))
	loggerAPIExit(
		sessionID,
		"response",
		"Write",
		"%v",
		statusCode,
	)
}

func constructResponse(
	responseObject interface{},
	responseError error,
) (string, int) {
	if responseError != nil {
		if customization.CreateErrorResponseFunc != nil {
			return customization.CreateErrorResponseFunc(
				responseError,
			)
		}
		return createErrorResponseFunc(
			responseError,
		)
	}
	return createOkResponseFunc(
		responseObject,
	)
}

// Write responds to the consumer with corresponding HTTP status code and response body
func Write(
	sessionID uuid.UUID,
	responseObject interface{},
	responseError error,
) {
	var responseWriter = sessionGetResponseWriter(
		sessionID,
	)
	var responseMessage, statusCode = constructResponseFunc(
		responseObject,
		responseError,
	)
	var _, isOverrided = responseObject.(overrideResponse)
	if !isOverrided {
		writeResponseFunc(
			sessionID,
			responseWriter,
			statusCode,
			responseMessage,
		)
	} else {
		loggerAPIExit(
			sessionID,
			"response",
			"Write",
			"Overrided",
		)
	}
}

// Override overrides the default response.Write functionality by the given callback function; consumers must manually deal with response writer accordingly
func Override(
	sessionID uuid.UUID,
	callback func(*http.Request, http.ResponseWriter),
) (interface{}, error) {
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
	return overrideResponse{}, nil
}
