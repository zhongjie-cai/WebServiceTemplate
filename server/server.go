package server

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func createServer(serverCert *tls.Certificate, clientCertPool *x509.CertPool) *http.Server {
	return &http.Server{
		Addr: ":" + configAppPort(),
		TLSConfig: &tls.Config{
			// Provide server certificates for HTTPS communications
			Certificates: []tls.Certificate{
				*serverCert,
			},
			// Reject any TLS certificate that cannot be validated
			ClientAuth: tls.RequireAndVerifyClientCert,
			// Ensure that we only use our "CA" to validate certificates
			ClientCAs: clientCertPool,
			// PFS because we can but this will reject client with RSA certificates
			CipherSuites: []uint16{tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
			// Force it server side
			PreferServerCipherSuites: true,
			// TLS 1.2 as minimum requirement
			MinVersion: tls.VersionTLS12,
		},
	}
}

func listenAndServeTLS(server *http.Server) error {
	return server.ListenAndServeTLS("", "")
}

func runServer() error {
	var serverCert, certError = certificateGetServerCertificate()
	if certError != nil {
		return apperrorWrapSimpleError(
			certError,
			"Failed to run server due to server cert error",
		)
	}
	var clientCertPool, poolError = certificateGetClientCertPool()
	if poolError != nil {
		return apperrorWrapSimpleError(
			poolError,
			"Failed to run server due to client cert pool error",
		)
	}
	var server = createServerFunc(
		serverCert,
		clientCertPool,
	)
	var serverError = listenAndServeTLSFunc(
		server,
	)
	if serverError != nil {
		return apperrorWrapSimpleError(
			serverError,
			"Failed to host service on port %v",
			configAppPort(),
		)
	}
	return nil
}

func hostEntries(entryFuncs ...func()) error {
	if entryFuncs == nil ||
		len(entryFuncs) == 0 {
		return apperrorWrapSimpleError(
			nil,
			"No host entries found",
		)
	}
	for _, entryFunc := range entryFuncs {
		entryFunc()
	}
	return nil
}

// Host hosts the service entries and starts HTTPS server
func Host() error {
	var entryError = hostEntriesFunc(
		healthHostEntry,
		faviconHostEntry,
		swaggerHostEntry,
	)
	if entryError != nil {
		return apperrorWrapSimpleError(
			entryError,
			"Failed to host entries on port %v",
			configAppPort(),
		)
	}
	var serverError = runServerFunc()
	if serverError != nil {
		return apperrorWrapSimpleError(
			serverError,
			"Failed to run server on port %v",
			configAppPort(),
		)
	}
	return nil
}
