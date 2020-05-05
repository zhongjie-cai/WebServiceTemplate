package request

import (
	"net/http"
)

// GetRequestBody parses and returns the content of the httpRequest body in string representation
func GetRequestBody(
	httpRequest *http.Request,
) string {
	var bodyBytes []byte
	var bodyError error
	if httpRequest != nil &&
		httpRequest.Body != nil {
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

// FullDump dumps the complete content of a given HTTP request, including its method, URL, body, headers and caller address
func FullDump(
	httpRequest *http.Request,
) string {
	var requestBytes, dumpError = httputilDumpRequest(
		httpRequest,
		true,
	)
	if dumpError != nil {
		return fmtSprintf(
			"FullDump Failed: %v\r\nSimpleDump: %v\r\n",
			dumpError,
			httpRequest,
		)
	}
	return fmtSprintf(
		"%vRemote Address: %v\r\n",
		string(requestBytes),
		httpRequest.RemoteAddr,
	)
}
