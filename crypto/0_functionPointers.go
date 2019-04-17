package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// func pointers for injection / testing: crypto.go
var (
	aesNewCipher            = aes.NewCipher
	cipherNewGCM            = cipher.NewGCM
	ioReadFull              = io.ReadFull
	apperrorWrapSimpleError = apperror.WrapSimpleError
)
