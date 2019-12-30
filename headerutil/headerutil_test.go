package headerutil

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/headerutil/headerstyle"
	"github.com/zhongjie-cai/WebServiceTemplate/logger"
	sessionModel "github.com/zhongjie-cai/WebServiceTemplate/session/model"
)

func TestGetHeaderLogStyle_NoCustomization(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}

	// mock
	createMock(t)

	// SUT + act
	var result = getHeaderLogStyle(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, headerstyle.DoNotLog, result)

	// verify
	verifyAll(t)
}

func TestGetHeaderLogStyle_SessionCustomized(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeaderStyle = headerstyle.HeaderStyle(rand.Int())

	// mock
	createMock(t)

	// expect
	customizationSessionHTTPHeaderLogStyleExpected = 1
	customization.SessionHTTPHeaderLogStyle = func(session sessionModel.Session) headerstyle.HeaderStyle {
		customizationSessionHTTPHeaderLogStyleCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderStyle
	}

	// SUT + act
	var result = getHeaderLogStyle(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyHeaderStyle, result)

	// verify
	verifyAll(t)
}

func TestGetHeaderLogStyle_DefaultCustomized(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeaderStyle = headerstyle.HeaderStyle(rand.Int())

	// mock
	createMock(t)

	// expect
	customizationDefaultHTTPHeaderLogStyleExpected = 1
	customization.DefaultHTTPHeaderLogStyle = func() headerstyle.HeaderStyle {
		customizationDefaultHTTPHeaderLogStyleCalled++
		return dummyHeaderStyle
	}

	// SUT + act
	var result = getHeaderLogStyle(
		dummySessionObject,
	)

	// assert
	assert.Equal(t, dummyHeaderStyle, result)

	// verify
	verifyAll(t)
}

func TestLogCombinedHTTPHeader(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyContent = "some content"

	// mock
	createMock(t)
	var loggerLogFunc func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{})

	// expect
	jsonutilMarshalIgnoreErrorExpected = 1
	jsonutilMarshalIgnoreError = func(v interface{}) string {
		jsonutilMarshalIgnoreErrorCalled++
		assert.Equal(t, dummyHeader, v)
		return dummyContent
	}
	loggerLogFuncExpected = 1
	loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Header", category)
		assert.Zero(t, subcategory)
		assert.Equal(t, dummyContent, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	logCombinedHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogPerNameHTTPHeader(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}

	// mock
	createMock(t)
	var loggerLogFunc func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{})

	// expect
	stringsJoinExpected = 2
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		assert.Equal(t, ",", sep)
		return strings.Join(a, sep)
	}
	loggerLogFuncExpected = 2
	loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Header", category)
		if subcategory == "foo" {
			assert.Equal(t, "bar1,bar2", messageFormat)
		} else if subcategory == "test" {
			assert.Equal(t, "123", messageFormat)
		}
		assert.Empty(t, parameters)
	}

	// SUT + act
	logPerNameHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogPerValueHTTPHeader(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}

	// mock
	createMock(t)
	var loggerLogFunc func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{})

	// expect
	loggerLogFuncExpected = 3
	loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, "Header", category)
		if messageFormat == "bar1" {
			assert.Equal(t, "foo", subcategory)
		} else if messageFormat == "bar2" {
			assert.Equal(t, "foo", subcategory)
		} else if messageFormat == "123" {
			assert.Equal(t, "test", subcategory)
		}
		assert.Empty(t, parameters)
	}

	// SUT + act
	logPerValueHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeader_DoNotLog(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyHeaderLogStyle = headerstyle.DoNotLog
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	getHeaderLogStyleFuncExpected = 1
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderLogStyle
	}

	// SUT + act
	LogHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeader_LogCombined(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyHeaderLogStyle = headerstyle.LogCombined
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	getHeaderLogStyleFuncExpected = 1
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderLogStyle
	}
	logCombinedHTTPHeaderFuncExpected = 1
	logCombinedHTTPHeaderFunc = func(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
		logCombinedHTTPHeaderFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHeader, header)
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(loggerLogFunc)), fmt.Sprintf("%v", reflect.ValueOf(logFunc)))
	}

	// SUT + act
	LogHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeader_LogPerName(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyHeaderLogStyle = headerstyle.LogPerName
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	getHeaderLogStyleFuncExpected = 1
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderLogStyle
	}
	logPerNameHTTPHeaderFuncExpected = 1
	logPerNameHTTPHeaderFunc = func(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
		logPerNameHTTPHeaderFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHeader, header)
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(loggerLogFunc)), fmt.Sprintf("%v", reflect.ValueOf(logFunc)))
	}

	// SUT + act
	LogHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeader_LogPerValue(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyHeaderLogStyle = headerstyle.LogPerValue
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	getHeaderLogStyleFuncExpected = 1
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderLogStyle
	}
	logPerValueHTTPHeaderFuncExpected = 1
	logPerValueHTTPHeaderFunc = func(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
		logPerValueHTTPHeaderFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyHeader, header)
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(loggerLogFunc)), fmt.Sprintf("%v", reflect.ValueOf(logFunc)))
	}

	// SUT + act
	LogHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeader_Other(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyHeader = http.Header{
		"foo":  []string{"bar1", "bar2"},
		"test": []string{"123"},
	}
	var dummyHeaderLogStyle = headerstyle.HeaderStyle(100 + rand.Intn(100))
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	getHeaderLogStyleFuncExpected = 1
	getHeaderLogStyleFunc = func(session sessionModel.Session) headerstyle.HeaderStyle {
		getHeaderLogStyleFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		return dummyHeaderLogStyle
	}

	// SUT + act
	LogHTTPHeader(
		dummySessionObject,
		dummyHeader,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}

func TestLogHTTPHeaderForName(t *testing.T) {
	// arrange
	var dummySessionObject = &dummySession{t}
	var dummyName = "some name"
	var dummyValues = []string{"some value 1", "some value 2"}
	var loggerLogFunc = func(session sessionModel.Session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		loggerLogFuncCalled++
	}

	// mock
	createMock(t)

	// expect
	logHTTPHeaderFuncExpected = 1
	logHTTPHeaderFunc = func(session sessionModel.Session, header http.Header, logFunc logger.LogFunc) {
		logHTTPHeaderFuncCalled++
		assert.Equal(t, dummySessionObject, session)
		assert.Equal(t, dummyValues, header[dummyName])
		assert.Equal(t, fmt.Sprintf("%v", reflect.ValueOf(loggerLogFunc)), fmt.Sprintf("%v", reflect.ValueOf(logFunc)))
	}

	// SUT + act
	LogHTTPHeaderForName(
		dummySessionObject,
		dummyName,
		dummyValues,
		loggerLogFunc,
	)

	// verify
	verifyAll(t)
}
