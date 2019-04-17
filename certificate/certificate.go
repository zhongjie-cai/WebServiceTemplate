package certificate

import (
	"crypto/tls"
	"crypto/x509"
)

var (
	caCertPool        *x509.CertPool
	clientCertificate *tls.Certificate
	serverCertificate *tls.Certificate
)

func loadTLSCertificate(certBytes, keyBytes []byte) (*tls.Certificate, error) {
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

func appendCertsFromPEM(certPool *x509.CertPool, certBytes []byte) bool {
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

// GetClientCertificate returns the client certificate loaded from local storage
func GetClientCertificate() (*tls.Certificate, error) {
	if clientCertificate != nil {
		return clientCertificate, nil
	}
	return nil,
		apperrorWrapSimpleError(
			nil,
			"Client certificate not initialized",
		)
}

// GetServerCertificate returns the server certificate loaded from local storage
func GetServerCertificate() (*tls.Certificate, error) {
	if serverCertificate != nil {
		return serverCertificate, nil
	}
	return nil,
		apperrorWrapSimpleError(
			nil,
			"Server certificate not initialized",
		)
}

// GetClientCertPool returns the client cert pool (CA root) loaded from local storage
func GetClientCertPool() (*x509.CertPool, error) {
	if caCertPool != nil {
		return caCertPool, nil
	}
	return nil,
		apperrorWrapSimpleError(
			nil,
			"CA cert pool not initialized",
		)
}

// Initialize initializes the certificates used by the application
func Initialize(
	clientCertContent string,
	clientKeyContent string,
	serverCertContent string,
	serverKeyContent string,
	caCertContent string,
) error {
	var clientCert, clientCertError = loadTLSCertificateFunc(
		[]byte(clientCertContent),
		[]byte(clientKeyContent),
	)
	if clientCertError != nil {
		return apperrorWrapSimpleError(
			clientCertError,
			"Failed to initialize client certificate",
		)
	}
	clientCertificate = clientCert
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
