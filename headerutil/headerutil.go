package headerutil

import (
	"net/http"

	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
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

func logCombinedHTTPHeader(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
	var content = jsonutilMarshalIgnoreError(header)
	logFunc(
		session,
		"Header",
		"",
		content,
	)
}

func logPerNameHTTPHeader(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
	for name, value := range header {
		logFunc(
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

func logPerValueHTTPHeader(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
	for name, value := range header {
		for _, item := range value {
			logFunc(
				session,
				"Header",
				name,
				item,
			)
		}
	}
}

// LogHTTPHeader helps log of HTTP header object according to customizations
func LogHTTPHeader(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
	var headerLogStyle = getHeaderLogStyleFunc(session)
	switch headerLogStyle {
	case headerstyle.LogCombined:
		logCombinedHTTPHeaderFunc(session, header, logFunc)
	case headerstyle.LogPerName:
		logPerNameHTTPHeaderFunc(session, header, logFunc)
	case headerstyle.LogPerValue:
		logPerValueHTTPHeaderFunc(session, header, logFunc)
	}
}

// LogHTTPHeaderForName helps log of HTTP header object for a specific name entry according to customizations
func LogHTTPHeaderForName(session sessionModel.Session, name string, values []string, logFunc logger.LogFunc) {
	var header = http.Header{
		name: values,
	}
	logHTTPHeaderFunc(
		session,
		header,
		logFunc,
	)
}
