package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/application"
	"github.com/zhongjie-cai/WebServiceTemplate/handler/swagger"
)

var (
	fmtPrintfExpected        int
	fmtPrintfCalled          int
	swaggerHandlerExpected   int
	swaggerHandlerCalled     int
	applicationStartExpected int
	applicationStartCalled   int
)

func createMock(t *testing.T) {
	fmtPrintfExpected = 0
	fmtPrintfCalled = 0
	fmtPrintf = func(format string, a ...interface{}) (n int, err error) {
		fmtPrintfCalled++
		return 0, nil
	}
	swaggerHandlerExpected = 0
	swaggerHandlerCalled = 0
	swaggerHandler = func() http.Handler {
		swaggerHandlerCalled++
		return nil
	}
	applicationStartExpected = 0
	applicationStartCalled = 0
	applicationStart = func() {
		applicationStartCalled++
	}
}

func verifyAll(t *testing.T) {
	fmtPrintf = fmt.Printf
	assert.Equal(t, fmtPrintfExpected, fmtPrintfCalled, "Unexpected number of calls to fmtPrintf")
	swaggerHandler = swagger.Handler
	assert.Equal(t, swaggerHandlerExpected, swaggerHandlerCalled, "Unexpected number of calls to swaggerHandler")
	applicationStart = application.Start
	assert.Equal(t, applicationStartExpected, applicationStartCalled, "Unexpected number of calls to applicationStart")
}
