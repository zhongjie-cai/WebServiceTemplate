package request

import (
	"crypto/x509"
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

// GetAllowedLogType parses and returns the allowed log type in request header
func GetAllowedLogType(httpRequest *http.Request) logtype.LogType {
	if httpRequest == nil {
		return logtype.GeneralTracing
	}
	var headerValues, headerValuesFound = httpRequest.Header["Log-Type"]
	if !headerValuesFound || len(headerValues) == 0 {
		return logtype.GeneralTracing
	}
	var logType logtype.LogType
	for _, headerValue := range headerValues {
		logType = logType | logtypeFromString(
			headerValue,
		)
	}
	if logType == logtype.AppRoot {
		return logtype.GeneralTracing
	}
	return logType
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
