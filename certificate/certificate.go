package certificate

import (
	"crypto/tls"
	"crypto/x509"
)

var (
	caCertPool        *x509.CertPool
	serverCertificate *tls.Certificate
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
				tlsError,
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
			apperrorWrapSimpleError(
				nil,
				"Failed to parse certificate bytes",
			)
	}
	return certPool, nil
}

func initializeServerCert(
	serveHTTPS bool,
	serverCertContent string,
	serverKeyContent string,
) error {
	serverCertificate = nil
	if !serveHTTPS {
		return nil
	}
	var serverCert, serverCertError = loadTLSCertificateFunc(
		[]byte(serverCertContent),
		[]byte(serverKeyContent),
	)
	if serverCertError != nil {
		return apperrorWrapSimpleError(
			serverCertError,
			"Failed to initialize server certificate",
		)
	}
	serverCertificate = serverCert
	return nil
}

func initializeCaCertPool(
	validateClientCert bool,
	caCertContent string,
) error {
	caCertPool = nil
	if !validateClientCert {
		return nil
	}
	var certPool, poolError = loadX509CertPoolFunc(
		[]byte(caCertContent),
	)
	if poolError != nil {
		return apperrorWrapSimpleError(
			poolError,
			"Failed to initialize CA cert pool",
		)
	}
	caCertPool = certPool
	return nil
}

// Initialize initializes the certificates used by the application
func Initialize(
	serveHTTPS bool,
	serverCertContent string,
	serverKeyContent string,
	validateClientCert bool,
	caCertContent string,
) error {
	var (
		serverCertError = initializeServerCertFunc(
			serveHTTPS,
			serverCertContent,
			serverKeyContent,
		)
		caCertPoolError = initializeCaCertPoolFunc(
			validateClientCert,
			caCertContent,
		)
	)
	return apperrorConsolidateAllErrors(
		"Failed to initialize certificates for application",
		serverCertError,
		caCertPoolError,
	)
}

// GetServerCertificate returns the server certificate loaded from local storage
func GetServerCertificate() *tls.Certificate {
	return serverCertificate
}

// GetClientCertPool returns the client cert pool (CA root) loaded from local storage
func GetClientCertPool() *x509.CertPool {
	return caCertPool
}
