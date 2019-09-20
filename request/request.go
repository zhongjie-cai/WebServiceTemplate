package request

import (
	"crypto/x509"
	"net/http"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
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
func GetLoginID(httpRequest *http.Request) uuid.UUID {
	if httpRequest == nil {
		return uuidNew()
	}
	return getUUIDFromHeaderFunc(
		httpRequest.Header,
		"login-id",
	)
}

// GetCorrelationID parses and returns the correlation ID in request header
func GetCorrelationID(httpRequest *http.Request) uuid.UUID {
	if httpRequest == nil {
		return uuidNew()
	}
	return getUUIDFromHeaderFunc(
		httpRequest.Header,
		"correlation-id",
	)
}

// GetAllowedLogType parses and returns the allowed log type in request header
func GetAllowedLogType(httpRequest *http.Request) logtype.LogType {
	var headerValue = httpRequest.Header.Get("log-type")
	return logtypeFromString(headerValue)
}

// GetClientCertificates parses and returns the client certificates in request header
func GetClientCertificates(httpRequest *http.Request) ([]*x509.Certificate, error) {
	if httpRequest == nil ||
		httpRequest.TLS == nil {
		return nil,
			apperrorWrapSimpleError(
				nil,
				"Invalid request or insecure communication channel",
			)
	}
	return httpRequest.TLS.PeerCertificates, nil
}

// GetRequestBody parses and returns the content of the httpRequest body in string representation
func GetRequestBody(
	httpRequest *http.Request,
) string {
	var bodyBytes []byte
	var bodyError error
	if httpRequest.Body != nil {
		defer httpRequest.Body.Close()
		bodyBytes, bodyError = ioutilReadAll(
			httpRequest.Body,
		)
		if bodyError != nil {
			return ""
		}
		httpRequest.Body = ioutilNopCloser(
			bytesNewBuffer(
				bodyBytes,
			),
		)
	}
	var bodyContent = string(bodyBytes)
	return bodyContent
}
