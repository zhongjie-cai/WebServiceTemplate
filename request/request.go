package request

import (
	"crypto/x509"
	"net/http"

	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

// GetAllowedLogType parses and returns the allowed log type in request header
func GetAllowedLogType(httpRequest *http.Request) logtype.LogType {
	if httpRequest == nil ||
		customization.SessionAllowedLogType == nil {
		return config.DefaultAllowedLogType()
	}
	return customization.SessionAllowedLogType(httpRequest)
}

// GetAllowedLogLevel parses and returns the allowed log level in request header
func GetAllowedLogLevel(httpRequest *http.Request) loglevel.LogLevel {
	if httpRequest == nil ||
		customization.SessionAllowedLogLevel == nil {
		return config.DefaultAllowedLogLevel()
	}
	return customization.SessionAllowedLogLevel(httpRequest)
}

// GetClientCertificates parses and returns the client certificates in request header
func GetClientCertificates(httpRequest *http.Request) ([]*x509.Certificate, error) {
	if httpRequest == nil ||
		httpRequest.TLS == nil {
		return nil,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
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
