package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

// Encrypt would encrypt the plain text using given key via AES-GCM crypto suite
func Encrypt(plainText string, key string) (string, error) {
	var cipher, cipherError = aesNewCipher([]byte(key))
	if cipherError != nil {
		return "",
			apperrorWrapSimpleError(
				cipherError,
				"Failed to create new cipher",
			)
	}
	var gcm, _ = cipherNewGCM(cipher)
	var nonce = make([]byte, gcm.NonceSize())
	ioReadFull(rand.Reader, nonce)
	var resultBytes = gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(resultBytes), nil
}

// Decrypt would decrypt the cipher text using given key via AES-GCM crypto suite
func Decrypt(cipherText string, key string) (string, error) {
	var cipher, cipherError = aesNewCipher([]byte(key))
	if cipherError != nil {
		return "",
			apperrorWrapSimpleError(
				cipherError,
				"Failed to create new cipher",
			)
	}
	var gcm, _ = cipherNewGCM(cipher)
	var cipherBytes, decodeError = base64.StdEncoding.DecodeString(cipherText)
	if decodeError != nil {
		return "",
			apperrorWrapSimpleError(
				decodeError,
				"Failed to decode cipher text",
			)
	}
	var nonceSize = gcm.NonceSize()
	if len(cipherBytes) < nonceSize {
		return "",
			apperrorWrapSimpleError(
				nil,
				"Failed to decode cipher text: cipherText too short",
			)
	}
	var nonce []byte
	nonce, cipherBytes = cipherBytes[:nonceSize], cipherBytes[nonceSize:]
	var resultBytes, gcmError = gcm.Open(nil, nonce, cipherBytes, nil)
	if gcmError != nil {
		return "",
			apperrorWrapSimpleError(
				gcmError,
				"Failed to decode using cipher text",
			)
	}
	return string(resultBytes), nil
}
