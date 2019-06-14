package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func createServer(
	serveHTTPS bool,
	validateClientCert bool,
	appPort string,
	router *mux.Router,
) *http.Server {
	var tlsConfig = &tls.Config{
		// PFS because we can but this will reject client with RSA certificates
		CipherSuites: []uint16{tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
		// Force it server side
		PreferServerCipherSuites: true,
		// TLS 1.2 as minimum requirement
		MinVersion: tls.VersionTLS12,
	}
	if serveHTTPS {
		var serverCert = certificateGetServerCertificate()
		tlsConfig.Certificates = []tls.Certificate{
			*serverCert,
		}
		if validateClientCert {
			var clientCertPool = certificateGetClientCertPool()
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
			tlsConfig.ClientCAs = clientCertPool
		}
	}
	return &http.Server{
		Addr:         ":" + appPort,
		TLSConfig:    tlsConfig,
		Handler:      router,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 180,
	}
}

func listenAndServe(
	server *http.Server,
	serveHTTPS bool,
) error {
	if serveHTTPS {
		return server.ListenAndServeTLS("", "")
	}
	return server.ListenAndServe()
}

func runServer(
	serveHTTPS bool,
	validateClientCert bool,
	appPort string,
	router *mux.Router,
) error {
	var server = createServerFunc(
		serveHTTPS,
		validateClientCert,
		appPort,
		router,
	)
	var serverError = listenAndServeFunc(
		server,
		serveHTTPS,
	)
	if serverError != nil {
		return apperrorWrapSimpleError(
			serverError,
			"Failed to host service on port %v",
			appPort,
		)
	}
	return nil
}

// Host hosts the service entries and starts HTTPS server
func Host(
	serveHTTPS bool,
	validateClientCert bool,
	appPort string,
) error {
	var router, entryError = routeRegisterEntries(
		healthHostEntry,
		faviconHostEntry,
		swaggerHostEntry,
	)
	if entryError != nil {
		return apperrorWrapSimpleError(
			entryError,
			"Failed to host entries on port %v",
			appPort,
		)
	}
	var serverError = runServerFunc(
		serveHTTPS,
		validateClientCert,
		appPort,
		router,
	)
	if serverError != nil {
		return apperrorWrapSimpleError(
			serverError,
			"Failed to run server on port %v",
			appPort,
		)
	}
	return nil
}
