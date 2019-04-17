package request

import (
	"crypto/x509"
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"

	"github.com/google/uuid"
)

func getUUIDFromHeader(
	header http.Header,
	name string,
) uuid.UUID {
	var parsedUUID uuid.UUID
	var parseError error
	var headerValue = header.Get(name)
	parsedUUID, parseError = uuidParse(headerValue)
	if parseError != nil {
		parsedUUID = uuidNew()
	}
	return parsedUUID
}

// GetLoginID parses and returns the login ID in request header
func GetLoginID(request *http.Request) uuid.UUID {
	if request == nil {
		return uuidNew()
	}
	return getUUIDFromHeaderFunc(
		request.Header,
		"login-id",
	)
}

// GetCorrelationID parses and returns the correlation ID in request header
func GetCorrelationID(request *http.Request) uuid.UUID {
	if request == nil {
		return uuidNew()
	}
	return getUUIDFromHeaderFunc(
		request.Header,
		"correlation-id",
	)
}

// GetAllowedLogType parses and returns the allowed log type in request header
func GetAllowedLogType(request *http.Request) logtype.LogType {
	var headerValue = request.Header.Get("log-type")
	return logtypeFromString(headerValue)
}

// GetClientCertificates parses and returns the client certificates in request header
func GetClientCertificates(request *http.Request) ([]*x509.Certificate, error) {
	if request == nil ||
		request.TLS == nil {
		return nil,
			apperrorWrapSimpleError(
				nil,
				"Invalid request or insecure communication channel",
			)
	}
	return request.TLS.PeerCertificates, nil
}

// GetRequestBody parses and returns the content of the request body in string representation
func GetRequestBody(
	request *http.Request,
	sessionID uuid.UUID,
) string {
	var bodyBytes []byte
	var bodyError error
	if request.Body != nil {
		defer request.Body.Close()
		bodyBytes, bodyError = ioutilReadAll(
			request.Body,
		)
		if bodyError != nil {
			loggerAPIRequest(
				sessionID,
				"request",
				"GetRequestBody",
				"Error getting request body: %v",
				bodyError,
			)
			return ""
		}
	}
	var bodyContent = string(bodyBytes)
	loggerAPIRequest(
		sessionID,
		"request",
		"GetRequestBody",
		bodyContent,
	)
	return bodyContent
}
