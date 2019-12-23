package headerutil

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func getHeaderLogStyle(session sessionModel.Session) headerstyle.HeaderStyle {
	var headerLogStyle = headerstyle.DoNotLog
	if customization.SessionHTTPHeaderLogStyle != nil {
		headerLogStyle = customization.SessionHTTPHeaderLogStyle(session)
	} else if customization.DefaultHTTPHeaderLogStyle != nil {
		headerLogStyle = customization.DefaultHTTPHeaderLogStyle()
	}
	return headerLogStyle
}

func logCombinedHTTPHeader(session sessionModel.Session, header http.Header) {
	var content = jsonutilMarshalIgnoreError(header)
	loggerAPIRequest(
		session,
		"Header",
		"",
		content,
	)
}

func logPerNameHTTPHeader(session sessionModel.Session, header http.Header) {
	for name, value := range header {
		loggerAPIRequest(
			session,
			"Header",
			name,
			stringsJoin(
				value,
				",",
			),
		)
	}
}

func logPerValueHTTPHeader(session sessionModel.Session, header http.Header) {
	for name, value := range header {
		for _, item := range value {
			loggerAPIRequest(
				session,
				"Header",
				name,
				item,
			)
		}
	}
}

// LogHTTPHeader helps log of HTTP header object according to customizations
func LogHTTPHeader(session sessionModel.Session, header http.Header) {
	var headerLogStyle = getHeaderLogStyleFunc(session)
	switch headerLogStyle {
	case headerstyle.LogCombined:
		logCombinedHTTPHeaderFunc(session, header)
	case headerstyle.LogPerName:
		logPerNameHTTPHeaderFunc(session, header)
	case headerstyle.LogPerValue:
		logPerValueHTTPHeaderFunc(session, header)
	}
}

// LogHTTPHeaderForName helps log of HTTP header object for a specific name entry according to customizations
func LogHTTPHeaderForName(session sessionModel.Session, name string, values []string) {
	var header = http.Header{
		name: values,
	}
	logHTTPHeaderFunc(
		session,
		header,
	)
}
