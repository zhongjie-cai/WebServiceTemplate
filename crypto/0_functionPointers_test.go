package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

var (
	aesNewCipherExpected            int
	aesNewCipherCalled              int
	cipherNewGCMExpected            int
	cipherNewGCMCalled              int
	ioReadFullExpected              int
	ioReadFullCalled                int
	apperrorWrapSimpleErrorExpected int
	apperrorWrapSimpleErrorCalled   int
)

func createMock(t *testing.T) {
	aesNewCipherExpected = 0
	aesNewCipherCalled = 0
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return nil, nil
	}
	cipherNewGCMExpected = 0
	cipherNewGCMCalled = 0
	cipherNewGCM = func(cipher cipher.Block) (cipher.AEAD, error) {
		cipherNewGCMCalled++
		return nil, nil
	}
	ioReadFullExpected = 0
	ioReadFullCalled = 0
	ioReadFull = func(r io.Reader, buf []byte) (n int, err error) {
		ioReadFullCalled++
		return 0, nil
	}
	apperrorWrapSimpleErrorExpected = 0
	apperrorWrapSimpleErrorCalled = 0
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	aesNewCipher = aes.NewCipher
	if aesNewCipherExpected != aesNewCipherCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to aesNewCipher, expected %v, actual %v", aesNewCipherExpected, aesNewCipherCalled))
	}
	cipherNewGCM = cipher.NewGCM
	if cipherNewGCMExpected != cipherNewGCMCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to cipherNewGCM, expected %v, actual %v", cipherNewGCMExpected, cipherNewGCMCalled))
	}
	ioReadFull = io.ReadFull
	if ioReadFullExpected != ioReadFullCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to ioReadFull, expected %v, actual %v", ioReadFullExpected, ioReadFullCalled))
	}
	apperrorWrapSimpleError = apperror.WrapSimpleError
	if apperrorWrapSimpleErrorExpected != apperrorWrapSimpleErrorCalled {
		assert.Fail(t, fmt.Sprintf("Unexpected method call to apperrorWrapSimpleError, expected %v, actual %v", apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled))
	}
}
