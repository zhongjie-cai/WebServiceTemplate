package main

import (
	"github.com/google/uuid"
)

func bootstrapApplication() error {
	var configError = configInitialize()
	if configError != nil {
		return apperrorWrapSimpleError(
			configError,
			"Failed to bootstrap application for configuration",
		)
	}
	var certError = certificateInitialize(
		configSendClientCert(),
		configClientCertContent(),
		configClientKeyContent(),
		configServeHTTPS(),
		configServerCertContent(),
		configServerKeyContent(),
		configValidateClientCert(),
		configCaCertContent(),
	)
	if certError != nil {
		return apperrorWrapSimpleError(
			certError,
			"Failed to bootstrap application for certificates",
		)
	}
	return nil
}

func connectStorages() error {
	// TODO: connect the application to its dedicated storage
	return nil
}

func disconnectStorages() error {
	// TODO: disconnect the application from its dedicated storage
	return nil
}

func main() {
	err := bootstrapApplicationFunc()
	if err != nil {
		loggerAppRoot(
			uuid.Nil,
			"main",
			"bootstrapApplicationFunc",
			"Failed to initialize server due to %v.",
			err,
		)
		return
	}
	var appVersion = configAppVersion()
	var appPort = configAppPort()
	loggerAppRoot(
		uuid.Nil,
		"main",
		"applicationStart",
		"Started server (v-%v) on port %v.",
		appVersion,
		appPort,
	)
	err = connectStoragesFunc()
	if err != nil {
		loggerAppRoot(
			uuid.Nil,
			"main",
			"connectStorages",
			"Failed to initialize server due to %v.",
			err,
		)
		return
	}
	defer func() {
		err = disconnectStoragesFunc()
		if err != nil {
			loggerAppRoot(
				uuid.Nil,
				"main",
				"disconnectStorages",
				"Failed to terminate server cleanly due to %v.",
				err,
			)
			return
		}
	}()
	err = serverHost(
		configServeHTTPS(),
		configValidateClientCert(),
		appPort,
	)
	if err != nil {
		loggerAppRoot(
			uuid.Nil,
			"main",
			"applicationStop",
			"Stopped server due to %v.",
			err,
		)
	} else {
		loggerAppRoot(
			uuid.Nil,
			"main",
			"applicationStop",
			"Stopped server peacefully.",
		)
	}
}
