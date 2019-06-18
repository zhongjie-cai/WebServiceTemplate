package application

// PreBootstrapFunc can be replaced by a customized function of your own to preprocess before bootstrapping
var PreBootstrapFunc func() error

// PostBootstrapFunc can be replaced by a customized function of your own to postprocess after bootstrapping
var PostBootstrapFunc func() error

// AppClosingFunc can be replaced by a customized function of your own to finalize the closing of the application
var AppClosingFunc func() error

func doPreBootstraping() bool {
	if PreBootstrapFunc != nil {
		var preBootstrapError = PreBootstrapFunc()
		if preBootstrapError != nil {
			loggerAppRoot(
				"application",
				"doPreBootstraping",
				"Failed to execute pre-bootstrap function. Error: %v",
				preBootstrapError,
			)
			return false
		}
		loggerAppRoot(
			"application",
			"doPreBootstraping",
			"PreBootstrapFunc executed successfully",
		)
	} else {
		loggerAppRoot(
			"application",
			"doPreBootstraping",
			"PreBootstrapFunc is not configured; skipped execution",
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
			"Logger not initialized cleanly. Potential error: %v",
			loggerError,
		)
	}
	var configError = configInitialize()
	if configError != nil {
		loggerAppRoot(
			"application",
			"bootstrapApplication",
			"Configuration not initialized cleanly. Potential error: %v",
			configError,
		)
	}
	var certError = certificateInitialize(
		configServeHTTPS(),
		configServerCertContent(),
		configServerKeyContent(),
		configValidateClientCert(),
		configCaCertContent(),
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
	loggerAppRoot(
		"application",
		"bootstrapApplication",
		"Application bootstrapped successfully",
	)
	return true
}

func doPostBootstraping() bool {
	if PostBootstrapFunc != nil {
		var postBootstrapError = PostBootstrapFunc()
		if postBootstrapError != nil {
			loggerAppRoot(
				"application",
				"doPostBootstraping",
				"Failed to execute post-bootstrap function. Error: %v",
				postBootstrapError,
			)
			return false
		}
		loggerAppRoot(
			"application",
			"doPostBootstraping",
			"PostBootstrapFunc executed successfully",
		)
	} else {
		loggerAppRoot(
			"application",
			"doPostBootstraping",
			"PostBootstrapFunc is not configured; skipped execution",
		)
	}
	return true
}

func doApplicationStarting() {
	loggerAppRoot(
		"application",
		"doApplicationStarting",
		"Trying to start server (v-%v)",
		configAppVersion(),
	)
	var serverHostError = serverHost(
		configServeHTTPS(),
		configValidateClientCert(),
		configAppPort(),
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
	if AppClosingFunc == nil {
		return
	}
	var appClosingError = AppClosingFunc()
	if appClosingError != nil {
		loggerAppRoot(
			"application",
			"doApplicationClosing",
			"Failed to execute application closing function. Error: %v",
			appClosingError,
		)
	} else {
		loggerAppRoot(
			"application",
			"doApplicationClosing",
			"Application closed successfully",
		)
	}
}

// Start bootstraps and starts the application web server according to configured function values
func Start() {
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
