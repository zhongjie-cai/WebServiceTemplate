package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
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

func shutDown(
	runtimeContext context.Context,
	server *http.Server,
) error {
	return server.Shutdown(
		runtimeContext,
	)
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

	var signalInterrupt = make(chan os.Signal, 1)
	signalNotify(
		signalInterrupt,
		os.Interrupt,
	)

	var hostError error
	go func() {
		hostError = listenAndServeFunc(
			server,
			serveHTTPS,
		)
		signalInterrupt <- os.Interrupt
	}()

	<-signalInterrupt
	var runtimeContext, cancelCallback = contextWithTimeout(
		contextBackground(),
		15*time.Second,
	)
	defer cancelCallback()

	var shutdownError = shutDownFunc(
		runtimeContext,
		server,
	)

	return apperrorConsolidateAllErrors(
		"One or more errors have occurred during server hosting",
		hostError,
		shutdownError,
	)
}

// Host hosts the service entries and starts HTTPS server
func Host(
	serveHTTPS bool,
	validateClientCert bool,
	appPort string,
) error {
	var router, routerError = registerInstantiate()
	if routerError != nil {
		return apperrorWrapSimpleError(
			routerError,
			"Failed to host entries on port %v",
			appPort,
		)
	}
	loggerAppRoot(
		"server",
		"Host",
		"Targeting port [%v] HTTPS [%v] mTLS [%v]",
		appPort,
		serveHTTPS,
		validateClientCert,
	)
	var hostError = runServerFunc(
		serveHTTPS,
		validateClientCert,
		appPort,
		router,
	)
	if hostError != nil {
		return apperrorWrapSimpleError(
			hostError,
			"Failed to run server on port %v",
			appPort,
		)
	}
	return nil
}
