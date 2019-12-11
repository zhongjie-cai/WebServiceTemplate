package application

import (
	"github.com/zhongjie-cai/WebServiceTemplate/config"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
)

func doPreBootstraping() bool {
	if customization.PreBootstrapFunc != nil {
		var preBootstrapError = customization.PreBootstrapFunc()
		if preBootstrapError != nil {
			loggerAppRoot(
				"application",
				"doPreBootstraping",
				"Failed to execute customization.PreBootstrapFunc. Error: %v",
				preBootstrapError,
			)
			return false
		}
		loggerAppRoot(
			"application",
			"doPreBootstraping",
			"customization.PreBootstrapFunc executed successfully",
		)
	} else {
		loggerAppRoot(
			"application",
			"doPreBootstraping",
			"customization.PreBootstrapFunc is not configured; skipped execution",
		)
	}
	return true
}

func bootstrapApplication() bool {
	var loggerError = loggerInitialize()
	if loggerError != nil {
		loggerAppRoot(
			"application",
			"bootstrapApplication",
			"Application logger not initialized cleanly. Potential error: %v",
			loggerError,
		)
	}
	var configError = configInitialize()
	if configError != nil {
		loggerAppRoot(
			"application",
			"bootstrapApplication",
			"Application configuration not initialized cleanly. Potential error: %v",
			configError,
		)
	}
	var certError = certificateInitialize(
		config.ServeHTTPS(),
		config.ServerCertContent(),
		config.ServerKeyContent(),
		config.ValidateClientCert(),
		config.CaCertContent(),
		config.ClientCertContent(),
		config.ClientKeyContent(),
	)
	if certError != nil {
		loggerAppRoot(
			"application",
			"bootstrapApplication",
			"Failed to bootstrap server application. Error: %v",
			certError,
		)
		return false
	}
	var apperrorError = apperrorInitialize()
	if apperrorError != nil {
		loggerAppRoot(
			"application",
			"bootstrapApplication",
			"Failed to bootstrap server application. Error: %v",
			apperrorError,
		)
		return false
	}
	networkInitialize(
		config.DefaultNetworkTimeout(),
	)
	loggerAppRoot(
		"application",
		"bootstrapApplication",
		"Application bootstrapped successfully",
	)
	return true
}

func doPostBootstraping() bool {
	if customization.PostBootstrapFunc != nil {
		var postBootstrapError = customization.PostBootstrapFunc()
		if postBootstrapError != nil {
			loggerAppRoot(
				"application",
				"doPostBootstraping",
				"Failed to execute customization.PostBootstrapFunc. Error: %v",
				postBootstrapError,
			)
			return false
		}
		loggerAppRoot(
			"application",
			"doPostBootstraping",
			"customization.PostBootstrapFunc executed successfully",
		)
	} else {
		loggerAppRoot(
			"application",
			"doPostBootstraping",
			"customization.PostBootstrapFunc is not configured; skipped execution",
		)
	}
	return true
}

func doApplicationStarting() {
	loggerAppRoot(
		"application",
		"doApplicationStarting",
		"Trying to start server (v-%v)",
		config.AppVersion(),
	)
	var serverHostError = serverHost(
		config.ServeHTTPS(),
		config.ValidateClientCert(),
		config.AppPort(),
	)
	if serverHostError != nil {
		loggerAppRoot(
			"application",
			"doApplicationStarting",
			"Failed to host server. Error: %v",
			serverHostError,
		)
	} else {
		loggerAppRoot(
			"application",
			"doApplicationStarting",
			"Server hosting terminated",
		)
	}
}

func doApplicationClosing() {
	if customization.AppClosingFunc == nil {
		loggerAppRoot(
			"application",
			"doApplicationClosing",
			"customization.AppClosingFunc is not configured; skipped execution",
		)
		return
	}
	var appClosingError = customization.AppClosingFunc()
	if appClosingError != nil {
		loggerAppRoot(
			"application",
			"doApplicationClosing",
			"Failed to execute customization.AppClosingFunc. Error: %v",
			appClosingError,
		)
	} else {
		loggerAppRoot(
			"application",
			"doApplicationClosing",
			"customization.AppClosingFunc executed successfully",
		)
	}
}

// Start bootstraps and starts the application web server according to configured function values
func Start() {
	sessionInitialize()
	if !doPreBootstrapingFunc() {
		return
	}
	if !bootstrapApplicationFunc() {
		return
	}
	if !doPostBootstrapingFunc() {
		return
	}
	defer doApplicationClosingFunc()
	doApplicationStartingFunc()
}
