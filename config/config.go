package config

import (
	"time"

	apperrorEnum "github.com/zhongjie-cai/WebServiceTemplate/apperror/enum"
	"github.com/zhongjie-cai/WebServiceTemplate/customization"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/loglevel"
	"github.com/zhongjie-cai/WebServiceTemplate/logger/logtype"
)

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

// SendClientCert returns the control switch for whether or not send client certificate to external web services
var SendClientCert = defaultSendClientCert

// ClientCertContent returns the client certificate cert content of the application
var ClientCertContent = defaultClientCertContent

// ClientKeyContent returns the client certificate key content of the application
var ClientKeyContent = defaultClientKeyContent

// DefaultAllowedLogType returns the default allowed log type of the application
var DefaultAllowedLogType = defaultAllowedLogType

// DefaultAllowedLogLevel returns the default allowed log level of the application
var DefaultAllowedLogLevel = defaultAllowedLogLevel

// DefaultNetworkTimeout returns the default network timeout value of the application
var DefaultNetworkTimeout = defaultNetworkTimeout

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

func defaultSendClientCert() bool {
	return false
}

func defaultClientCertContent() string {
	return ""
}

func defaultClientKeyContent() string {
	return ""
}

func defaultAllowedLogType() logtype.LogType {
	return logtype.BasicLogging
}

func defaultAllowedLogLevel() loglevel.LogLevel {
	return loglevel.Warn
}

func defaultNetworkTimeout() time.Duration {
	return 3 * time.Minute
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
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.%v function is forced to default [%v] due to forceToDefault flag set",
				name,
				defaultFunc(),
			)
	}
	if stringFunc == nil ||
		functionPointerEqualsFunc(stringFunc, defaultFunc) ||
		len(stringFunc()) == 0 {
		return defaultFunc,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
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
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.%v function is forced to default [%v] due to forceToDefault flag set",
				name,
				defaultFunc(),
			)
	}
	if booleanFunc == nil ||
		functionPointerEqualsFunc(booleanFunc, defaultFunc) {
		return defaultFunc,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.%v function is not configured; fallback to default [%v].",
				name,
				defaultFunc(),
			)
	}
	return booleanFunc, nil
}

func validateDefaultAllowedLogType(
	customizedFunc func() logtype.LogType,
	defaultFunc func() logtype.LogType,
) (func() logtype.LogType, error) {
	if customizedFunc == nil {
		return defaultFunc,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.DefaultAllowedLogType function is not configured; fallback to default [%v].",
				defaultFunc(),
			)
	}
	return customizedFunc, nil
}

func validateDefaultAllowedLogLevel(
	customizedFunc func() loglevel.LogLevel,
	defaultFunc func() loglevel.LogLevel,
) (func() loglevel.LogLevel, error) {
	if customizedFunc == nil {
		return defaultFunc,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.DefaultAllowedLogLevel function is not configured; fallback to default [%v].",
				defaultFunc(),
			)
	}
	return customizedFunc, nil
}

func validateDefaultNetworkTimeout(
	customizedFunc func() time.Duration,
	defaultFunc func() time.Duration,
) (func() time.Duration, error) {
	if customizedFunc == nil {
		return defaultFunc,
			apperrorGetCustomError(
				apperrorEnum.CodeGeneralFailure,
				"customization.DefaultNetworkTimeout function is not configured; fallback to default [%v].",
				defaultFunc(),
			)
	}
	return customizedFunc, nil
}

func isServerCertificateAvailable() bool {
	return len(ServerCertContent()) != 0 && len(ServerKeyContent()) != 0
}

func isCaCertificateAvailable() bool {
	return len(CaCertContent()) != 0
}

func isClientCertificateAvailable() bool {
	return len(ClientCertContent()) != 0 && len(ClientKeyContent()) != 0
}

// Initialize initiates and checks all application config related function injections
func Initialize() error {
	const noForceToDefault = false
	var (
		appVersionError             error
		appPortError                error
		appNameError                error
		appPathError                error
		isLocalhostError            error
		serveHTTPSError             error
		serverCertContentError      error
		serverKeyContentError       error
		validateClientCertError     error
		caCertContentError          error
		sendClientCertError         error
		clientCertContentError      error
		clientKeyContentError       error
		defaultAllowedLogTypeError  error
		defaultAllowedLogLevelError error
		defaultNetworkTimeoutError  error
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
	ClientCertContent, clientCertContentError = validateStringFunctionFunc(
		customization.ClientCertContent,
		"ClientCertContent",
		defaultClientCertContent,
		noForceToDefault,
	)
	ClientKeyContent, clientKeyContentError = validateStringFunctionFunc(
		customization.ClientKeyContent,
		"ClientKeyContent",
		defaultClientKeyContent,
		noForceToDefault,
	)
	SendClientCert, sendClientCertError = validateBooleanFunctionFunc(
		customization.SendClientCert,
		"SendClientCert",
		defaultSendClientCert,
		!isClientCertificateAvailableFunc(),
	)
	DefaultAllowedLogType, defaultAllowedLogTypeError = validateDefaultAllowedLogTypeFunc(
		customization.DefaultAllowedLogType,
		defaultAllowedLogType,
	)
	DefaultAllowedLogLevel, defaultAllowedLogLevelError = validateDefaultAllowedLogLevelFunc(
		customization.DefaultAllowedLogLevel,
		defaultAllowedLogLevel,
	)
	DefaultNetworkTimeout, defaultNetworkTimeoutError = validateDefaultNetworkTimeoutFunc(
		customization.DefaultNetworkTimeout,
		defaultNetworkTimeout,
	)
	return apperrorWrapSimpleError(
		[]error{
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
			clientCertContentError,
			clientKeyContentError,
			sendClientCertError,
			defaultAllowedLogTypeError,
			defaultAllowedLogLevelError,
			defaultNetworkTimeoutError,
		},
		"Unexpected errors occur during configuration initialization",
	)
}
