package enum

import (
	"math"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeEnumString_UnknownNegative(t *testing.T) {
	// arrange
	var testCode Code

	// mock
	createMock(t)

	// SUT
	testCode = -1

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unknown", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GeneralFailure(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeGeneralFailure

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "GeneralFailure", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_Unauthorized(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeUnauthorized

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unauthorized", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_InvalidOperation(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeInvalidOperation

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "InvalidOperation", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_BadRequest(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeBadRequest

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "BadRequest", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_NotFound(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeNotFound

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "NotFound", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_CircuitBreak(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeCircuitBreak

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "CircuitBreak", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_OperationLock(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeOperationLock

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "OperationLock", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_AccessForbidden(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeAccessForbidden

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "AccessForbidden", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GetDataCorruption(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeDataCorruption

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "DataCorruption", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_GetNotImplemented(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var testCode = CodeNotImplemented

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "NotImplemented", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumString_UnknownTooBig(t *testing.T) {
	// arrange
	var testCode Code

	// mock
	createMock(t)

	// SUT
	testCode = 999

	// act
	var convertedString = testCode.String()

	// assert
	assert.Equal(t, "Unknown", convertedString)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_GeneralFailure(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeGeneralFailure

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusInternalServerError, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_Unauthorized(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeUnauthorized

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusUnauthorized, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_InvalidOperation(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeInvalidOperation

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusMethodNotAllowed, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_BadRequest(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeBadRequest

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusBadRequest, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_NotFound(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeNotFound

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusNotFound, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_CircuitBreak(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeCircuitBreak

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusForbidden, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_OperationLock(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeOperationLock

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusLocked, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_AccessForbidden(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeAccessForbidden

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusForbidden, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_DataCorruption(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeDataCorruption

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusConflict, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_NotImplemented(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = CodeNotImplemented

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusNotImplemented, result)

	// verify
	verifyAll(t)
}

func TestCodeEnumHTTPStatusCode_OtherCode(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummyCode = Code(rand.Intn(math.MaxInt8) + int(CodeReservedCount))

	// act
	var result = dummyCode.HTTPStatusCode()

	// assert
	assert.Equal(t, http.StatusInternalServerError, result)

	// verify
	verifyAll(t)
}
