package certificate

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/zhongjie-cai/WebServiceTemplate/apperror"
)

// func pointers for injection / testing: certificate.go
var (
	tlsX509KeyPair               = tls.X509KeyPair
	x509NewCertPool              = x509.NewCertPool
	apperrorGetCustomError       = apperror.GetCustomError
	apperrorWrapSimpleError      = apperror.WrapSimpleError
	loadTLSCertificateFunc       = loadTLSCertificate
	appendCertsFromPEMFunc       = appendCertsFromPEM
	loadX509CertPoolFunc         = loadX509CertPool
	initializeTLSCertiticateFunc = initializeTLSCertiticate
	initializeX509CertPoolFunc   = initializeX509CertPool
)
