package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	shutdownSignal = make(chan os.Signal, 1)
)

func createServer(
	serveHTTPS bool,
	validateClientCert bool,
	appPort string,
	router *mux.Router,
) *http.Server {
	var tlsConfig = &tls.Config{
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
			var clientCertPool = certificateGetCaCertPool()
			if clientCertPool != nil {
				tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
				tlsConfig.ClientCAs = clientCertPool
			} else {
				tlsConfig.ClientAuth = tls.RequireAnyClientCert
			}
		} else {
			tlsConfig.ClientAuth = tls.RequestClientCert
		}
	}
	return &http.Server{
		Addr:      ":" + appPort,
		TLSConfig: tlsConfig,
		Handler:   router,
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

func consolidateError(
	hostError error,
	shutdownError error,
) error {
	if hostError == http.ErrServerClosed {
		hostError = nil
	}
	if shutdownError == http.ErrServerClosed {
		shutdownError = nil
	}
	return apperrorWrapSimpleError(
		[]error{
			hostError,
			shutdownError,
		},
		"One or more errors have occurred during server hosting",
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

	signalNotify(
		shutdownSignal,
		os.Interrupt,
		os.Kill,
	)

	var hostError error
	go func() {
		hostError = listenAndServeFunc(
			server,
			serveHTTPS,
		)
		haltFunc()
	}()

	<-shutdownSignal

	loggerAppRoot(
		"server",
		"Host",
		"Interrupt signal received: Terminating server",
	)

	var runtimeContext, cancelCallback = contextWithTimeout(
		contextBackground(),
		configGraceShutdownWaitTime(),
	)
	defer cancelCallback()

	var shutdownError = shutDownFunc(
		runtimeContext,
		server,
	)

	return consolidateErrorFunc(
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
			[]error{routerError},
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
	loggerAppRoot(
		"server",
		"Host",
		"Server terminated",
	)
	if hostError != nil {
		return apperrorWrapSimpleError(
			[]error{hostError},
			"Failed to run server on port %v",
			appPort,
		)
	}
	return nil
}

// Halt emits the signal interrupt to the server to conduct a graceful shutdown
func Halt() {
	shutdownSignal <- os.Interrupt
}
