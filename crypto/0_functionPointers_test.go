package crypto

import (
	"crypto/aes"
	"crypto/cipher"
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
	assert.Equal(t, aesNewCipherExpected, aesNewCipherCalled, "Unexpected method call to aesNewCipher")
	cipherNewGCM = cipher.NewGCM
	assert.Equal(t, cipherNewGCMExpected, cipherNewGCMCalled, "Unexpected method call to cipherNewGCM")
	ioReadFull = io.ReadFull
	assert.Equal(t, ioReadFullExpected, ioReadFullCalled, "Unexpected method call to ioReadFull")
	apperrorWrapSimpleError = apperror.WrapSimpleError
	assert.Equal(t, apperrorWrapSimpleErrorExpected, apperrorWrapSimpleErrorCalled, "Unexpected method call to apperrorWrapSimpleError")
}
