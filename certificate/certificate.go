package certificate

import (
	"crypto/tls"
	"crypto/x509"

	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
)

var (
	caCertPool        *x509.CertPool
	serverCertificate *tls.Certificate
	clientCertificate *tls.Certificate
)

func loadTLSCertificate(
	certBytes, keyBytes []byte,
) (*tls.Certificate, error) {
	var tlsCert, tlsError = tlsX509KeyPair(
		certBytes,
		keyBytes,
	)
	if tlsError != nil {
		return nil,
			apperrorWrapSimpleError(
				[]error{tlsError},
				"Failed to load certificate content",
			)
	}
	return &tlsCert, nil
}

func appendCertsFromPEM(
	certPool *x509.CertPool,
	certBytes []byte,
) bool {
	return certPool.AppendCertsFromPEM(certBytes)
}

func loadX509CertPool(certBytes []byte) (*x509.CertPool, error) {
	var certPool = x509NewCertPool()
	var appendSuccess = appendCertsFromPEMFunc(certPool, certBytes)
	if !appendSuccess {
		return nil,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"Failed to parse certificate bytes",
			)
	}
	return certPool, nil
}

func initializeTLSCertiticate(
	shouldLoadCert bool,
	certContent string,
	keyContent string,
) (*tls.Certificate, error) {
	if !shouldLoadCert || certContent == "" || keyContent == "" {
		return nil, nil
	}
	var cert, certError = loadTLSCertificateFunc(
		[]byte(certContent),
		[]byte(keyContent),
	)
	if certError != nil {
		return nil,
			apperrorWrapSimpleError(
				[]error{certError},
				"Failed to initialize certificate by key-cert pair",
			)
	}
	return cert, nil
}

func initializeX509CertPool(
	shouldLoadCert bool,
	certContent string,
) (*x509.CertPool, error) {
	if !shouldLoadCert || certContent == "" {
		return nil, nil
	}
	var certPool, poolError = loadX509CertPoolFunc(
		[]byte(certContent),
	)
	if poolError != nil {
		return nil,
			apperrorWrapSimpleError(
				[]error{poolError},
				"Failed to initialize cert pool by cert content",
			)
	}
	return certPool, nil
}

// Initialize initializes the certificates used by the application
func Initialize(
	serveHTTPS bool,
	serverCertContent string,
	serverKeyContent string,
	validateClientCert bool,
	caCertContent string,
	clientCertContent string,
	clientKeyContent string,
) error {
	var (
		serverCert, serverCertError = initializeTLSCertiticateFunc(
			serveHTTPS,
			serverCertContent,
			serverKeyContent,
		)
		certPool, caCertPoolError = initializeX509CertPoolFunc(
			validateClientCert,
			caCertContent,
		)
		clientCert, clientCertError = initializeTLSCertiticateFunc(
			true,
			clientCertContent,
			clientKeyContent,
		)
	)
	serverCertificate = serverCert
	caCertPool = certPool
	clientCertificate = clientCert
	return apperrorWrapSimpleError(
		[]error{
			serverCertError,
			caCertPoolError,
			clientCertError,
		},
		"Failed to initialize certificates for application",
	)
}

// GetServerCertificate returns the server certificate loaded from local storage
func GetServerCertificate() *tls.Certificate {
	return serverCertificate
}

// GetCaCertPool returns the client cert pool (CA root) loaded from local storage
func GetCaCertPool() *x509.CertPool {
	return caCertPool
}

// GetClientCertificate returns the client certificate loaded from local storage
func GetClientCertificate() *tls.Certificate {
	return clientCertificate
}

// HasClientCert returns whether a client certificate is available for runtime or not
func HasClientCert() bool {
	return clientCertificate != nil
}
