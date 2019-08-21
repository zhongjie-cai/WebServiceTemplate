package config

import "github.com/zhongjie-cai/WebServiceTemplate/customization"

// AppVersion returns the version information of the application
var AppVersion = defaultAppVersion

// AppPort returns the hosting port of the application
var AppPort = defaultAppPort

// AppName returns the name of the application
var AppName = defaultAppName

// AppPath returns the execution path of the application
var AppPath = defaultAppPath

// IsLocalhost returns the control switch for whether or not the current running environment is a localhost (testing) environment; for localhost, logs print all details
var IsLocalhost = defaultIsLocalhost

// ServeHTTPS returns the control switch for whether or not hosting the web-service with HTTPS
var ServeHTTPS = defaultServeHTTPS

// ServerCertContent returns the server certificate cert content of the application
var ServerCertContent = defaultServerCertContent

// ServerKeyContent returns the server certificate key content of the application
var ServerKeyContent = defaultServerKeyContent

// ValidateClientCert returns the control switch for whether or not validating the client certificate of incoming HTTP/S requests
var ValidateClientCert = defaultValidateClientCert

// CaCertContent returns the CA certificate cert content of the application
var CaCertContent = defaultCaCertContent

func defaultAppVersion() string {
	return "0.0.0.0"
}

func defaultAppPort() string {
	return "18605"
}

func defaultAppName() string {
	return "WebServiceTemplate"
}

func defaultAppPath() string {
	return "."
}

func defaultIsLocalhost() bool {
	return false
}

func defaultServeHTTPS() bool {
	return false
}

func defaultServerCertContent() string {
	return ""
}

func defaultServerKeyContent() string {
	return ""
}

func defaultValidateClientCert() bool {
	return false
}

func defaultCaCertContent() string {
	return ""
}

func functionPointerEquals(left, right interface{}) bool {
	var leftPointer = fmtSprintf("%v", reflectValueOf(left))
	var rightPointer = fmtSprintf("%v", reflectValueOf(right))
	return leftPointer == rightPointer
}

func validateStringFunction(
	stringFunc func() string,
	name string,
	defaultFunc func() string,
	forceToDefault bool,
) (func() string, error) {
	if forceToDefault {
		return defaultFunc,
			apperrorWrapSimpleError(
				nil,
				"customization.%v function is forced to default [%v] due to forceToDefault flag set",
				name,
				defaultFunc(),
			)
	}
	if stringFunc == nil ||
		functionPointerEqualsFunc(stringFunc, defaultFunc) ||
		len(stringFunc()) == 0 {
		return defaultFunc,
			apperrorWrapSimpleError(
				nil,
				"customization.%v function is not configured or is empty; fallback to default [%v]",
				name,
				defaultFunc(),
			)
	}
	return stringFunc, nil
}

func validateBooleanFunction(
	booleanFunc func() bool,
	name string,
	defaultFunc func() bool,
	forceToDefault bool,
) (func() bool, error) {
	if forceToDefault {
		return defaultFunc,
			apperrorWrapSimpleError(
				nil,
				"customization.%v function is forced to default [%v] due to forceToDefault flag set",
				name,
				defaultFunc(),
			)
	}
	if booleanFunc == nil ||
		functionPointerEqualsFunc(booleanFunc, defaultFunc) {
		return defaultFunc,
			apperrorWrapSimpleError(
				nil,
				"customization.%v function is not configured; fallback to default [%v].",
				name,
				defaultFunc(),
			)
	}
	return booleanFunc, nil
}

func isServerCertificateAvailable() bool {
	return len(ServerCertContent()) != 0 && len(ServerKeyContent()) != 0
}

func isCaCertificateAvailable() bool {
	return len(CaCertContent()) != 0
}

// Initialize initiates and checks all application config related function injections
func Initialize() error {
	const noForceToDefault = false
	var (
		appVersionError         error
		appPortError            error
		appNameError            error
		appPathError            error
		isLocalhostError        error
		serveHTTPSError         error
		serverCertContentError  error
		serverKeyContentError   error
		validateClientCertError error
		caCertContentError      error
	)
	AppVersion, appVersionError = validateStringFunctionFunc(
		customization.AppVersion,
		"AppVersion",
		defaultAppVersion,
		noForceToDefault,
	)
	AppPort, appPortError = validateStringFunctionFunc(
		customization.AppPort,
		"AppPort",
		defaultAppPort,
		noForceToDefault,
	)
	AppName, appNameError = validateStringFunctionFunc(
		customization.AppName,
		"AppName",
		defaultAppName,
		noForceToDefault,
	)
	AppPath, appPathError = validateStringFunctionFunc(
		customization.AppPath,
		"AppPath",
		defaultAppPath,
		noForceToDefault,
	)
	IsLocalhost, isLocalhostError = validateBooleanFunctionFunc(
		customization.IsLocalhost,
		"IsLocalhost",
		defaultIsLocalhost,
		noForceToDefault,
	)
	ServerCertContent, serverCertContentError = validateStringFunctionFunc(
		customization.ServerCertContent,
		"ServerCertContent",
		defaultServerCertContent,
		noForceToDefault,
	)
	ServerKeyContent, serverKeyContentError = validateStringFunctionFunc(
		customization.ServerKeyContent,
		"ServerKeyContent",
		defaultServerKeyContent,
		noForceToDefault,
	)
	ServeHTTPS, serveHTTPSError = validateBooleanFunctionFunc(
		customization.ServeHTTPS,
		"ServeHTTPS",
		defaultServeHTTPS,
		!isServerCertificateAvailableFunc(),
	)
	CaCertContent, caCertContentError = validateStringFunctionFunc(
		customization.CaCertContent,
		"CaCertContent",
		defaultCaCertContent,
		noForceToDefault,
	)
	ValidateClientCert, validateClientCertError = validateBooleanFunctionFunc(
		customization.ValidateClientCert,
		"ValidateClientCert",
		defaultValidateClientCert,
		!isCaCertificateAvailableFunc(),
	)
	return apperrorConsolidateAllErrors(
		"Unexpected errors occur during configuration initialization",
		appVersionError,
		appPortError,
		appNameError,
		appPathError,
		isLocalhostError,
		serverCertContentError,
		serverKeyContentError,
		serveHTTPSError,
		caCertContentError,
		validateClientCertError,
	)
}