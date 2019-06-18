package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func TestMain(t *testing.T) {
	// stub
	appVersion = "some app version"
	appName = "some app name"
	appPath = "some app path"
	appPort = "some app port"

	// mock
	createMock(t)

	// expect
	fmtPrintfExpected = 1
	fmtPrintf = func(format string, a ...interface{}) (n int, err error) {
		fmtPrintfCalled++
		assert.Equal(t, "<%v|%v> %v\n", format)
		assert.Equal(t, 3, len(a))
		return 0, nil
	}
	swaggerHandlerExpected = 1
	swaggerHandler = func() http.Handler {
		swaggerHandlerCalled++
		return nil
	}
	applicationStartExpected = 1
	applicationStart = func() {
		applicationStartCalled++
	}

	// SUT + act
	main()

	// assert
	assert.NotNil(t, customization.Routes)
	assert.NotEmpty(t, customization.Routes())
	assert.NotNil(t, customization.Statics)
	assert.NotEmpty(t, customization.Statics())
	assert.NotNil(t, customization.AppName)
	assert.NotZero(t, customization.AppName())
	assert.NotNil(t, customization.AppPort)
	assert.NotZero(t, customization.AppPort())
	assert.NotNil(t, customization.AppVersion)
	assert.NotZero(t, customization.AppVersion())
	assert.NotNil(t, customization.AppPath)
	assert.NotZero(t, customization.AppPath())
	assert.NotNil(t, customization.IsLocalhost)
	assert.True(t, customization.IsLocalhost())
	assert.NotNil(t, customization.LoggingFunc)
	assert.NotPanics(t, func() { customization.LoggingFunc(nil, 0, "", "", "") })

	// verify
	verifyAll(t)
}
