package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

func TestCryptoLogic_HappyPath(t *testing.T) {
	// arrange
	var expectedPlainText = "some dummy plain text"
	var cryptoKey = "the-key-has-to-be-32-bytes-long!"

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 2
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	cipherNewGCMExpected = 2
	cipherNewGCM = func(c cipher.Block) (cipher.AEAD, error) {
		cipherNewGCMCalled++
		return cipher.NewGCM(c)
	}
	ioReadFullExpected = 1
	ioReadFull = func(r io.Reader, buf []byte) (n int, err error) {
		ioReadFullCalled++
		return io.ReadFull(r, buf)
	}

	// SUT + act
	var cipherText, _ = Encrypt(expectedPlainText, cryptoKey)
	var decryptedText, _ = Decrypt(cipherText, cryptoKey)

	// assert
	assert.Equal(t, expectedPlainText, decryptedText)

	// verify
	verifyAll(t)
}

func TestEncrypt_InvalidKey(t *testing.T) {
	// arrange
	var expectedPlainText = "some dummy plain text"
	var cryptoKey = "invalid key"
	var expectedErrorMessage = "Failed to create new cipher"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 1
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var cipherText, err = Encrypt(expectedPlainText, cryptoKey)

	// assert
	assert.Equal(t, "", cipherText)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestDecrypt_InvalidKey(t *testing.T) {
	// arrange
	var expectedCipherText = "some dummy cipher text"
	var cryptoKey = "invalid key"
	var expectedErrorMessage = "Failed to create new cipher"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 1
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var plainText, err = Decrypt(expectedCipherText, cryptoKey)

	// assert
	assert.Equal(t, "", plainText)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestDecrypt_WrongBase64CipherText(t *testing.T) {
	// arrange
	var expectedCipherText = "some dummy cipher text"
	var cryptoKey = "the-key-has-to-be-32-bytes-long!"
	var expectedErrorMessage = "Failed to decode cipher text"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 1
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	cipherNewGCMExpected = 1
	cipherNewGCM = func(c cipher.Block) (cipher.AEAD, error) {
		cipherNewGCMCalled++
		return cipher.NewGCM(c)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var plainText, err = Decrypt(expectedCipherText, cryptoKey)

	// assert
	assert.Equal(t, "", plainText)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestDecrypt_TooShortCipherText(t *testing.T) {
	// arrange
	var expectedCipherText = base64.StdEncoding.EncodeToString([]byte("foo"))
	var cryptoKey = "the-key-has-to-be-32-bytes-long!"
	var expectedErrorMessage = "Failed to decode cipher text: cipherText too short"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 1
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	cipherNewGCMExpected = 1
	cipherNewGCM = func(c cipher.Block) (cipher.AEAD, error) {
		cipherNewGCMCalled++
		return cipher.NewGCM(c)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.Nil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var plainText, err = Decrypt(expectedCipherText, cryptoKey)

	// assert
	assert.Equal(t, "", plainText)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}

func TestDecrypt_InvalidCipherText(t *testing.T) {
	// arrange
	var expectedCipherText = base64.StdEncoding.EncodeToString([]byte("some dummy cipher text"))
	var cryptoKey = "the-key-has-to-be-32-bytes-long!"
	var expectedErrorMessage = "Failed to decode using cipher text"
	var dummyAppError = apperror.GetGeneralFailureError(nil)

	// mock
	createMock(t)

	// expect
	aesNewCipherExpected = 1
	aesNewCipher = func(key []byte) (cipher.Block, error) {
		aesNewCipherCalled++
		return aes.NewCipher(key)
	}
	cipherNewGCMExpected = 1
	cipherNewGCM = func(c cipher.Block) (cipher.AEAD, error) {
		cipherNewGCMCalled++
		return cipher.NewGCM(c)
	}
	apperrorWrapSimpleErrorExpected = 1
	apperrorWrapSimpleError = func(innerError error, messageFormat string, parameters ...interface{}) apperror.AppError {
		apperrorWrapSimpleErrorCalled++
		assert.NotNil(t, innerError)
		assert.Equal(t, expectedErrorMessage, messageFormat)
		assert.Equal(t, 0, len(parameters))
		return dummyAppError
	}

	// SUT + act
	var plainText, err = Decrypt(expectedCipherText, cryptoKey)

	// assert
	assert.Equal(t, "", plainText)
	assert.Equal(t, dummyAppError, err)

	// verify
	verifyAll(t)
}
