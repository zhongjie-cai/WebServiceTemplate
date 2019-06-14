package config

// These variables are injected to application during compile time
var (
	// appVersion is for injecting the version information from compiler
	appVersion = "0.0.0.0"
	// appPort is for injecting the port information from compiler
	appPort = "18605"
	// appName is for injecting the name of application from compiler
	appName = "WebServiceTemplate"
	// appPath is for injecting the application executable path from compiler
	appPath = "."
)

// These variables are intialized during application startup from system info and environment variables directly
var (
	// bootTime is the time when the host has started execution
	bootTime = ""
	// isLocalhost is the control switch for whether or not the current running environment is a localhost (testing) environment; for localhost, no encryption/decryption is needed, and logs print all details
	isLocalhost = false
	// sendClientCert is the control switch for whether or not sending client certificate over external HTTP/S communications
	sendClientCert = false
	// serveHTTPS is the control switch for whether or not hosting the web-service with HTTPS
	serveHTTPS = false
	// validateClientCert is the control switch for whether or not validating the client certificate of incoming HTTP/S requests
	validateClientCert = false
	// cryptoKey is the key phrase that is used to encrypt & decrypt secrets; consists of two parts, one in code and another configured in environment variable
	cryptoKey = ""
)

// These variables are initialized during application startup from decrypting certain environment variables
var (
	// clientCertContent is the client certificate cert string
	clientCertContent = ""
	// clientKeyContent is the client certificate key string
	clientKeyContent = ""
	// serverCertContent is the server certificate cert string
	serverCertContent = ""
	// serverKeyContent is the server certificate key string
	serverKeyContent = ""
	// caCertContent is the CA certificate cert string
	caCertContent = ""
)

const (
	// CryptoKeyPartial is the second half of the encryption key that should be combined with the injected value from compiler
	CryptoKeyPartial string = "UEvaxQGW6YC9aeCs"
)

func initializeBootTime() {
	var timeNowUTC = timeutilGetTimeNowUTC()
	bootTime = timeutilFormatDateTime(timeNowUTC)
}

func initializeGeneralEnvironmentVariables() error {
	isLocalhost = stringsEqualFold(getEnvironmentVariable("IsLocalhost"), "true")
	sendClientCert = stringsEqualFold(getEnvironmentVariable("SendClientCert"), "true")
	serveHTTPS = stringsEqualFold(getEnvironmentVariable("ServeHTTPS"), "true")
	validateClientCert = stringsEqualFold(getEnvironmentVariable("ValidateClientCert"), "true")
	return nil
}

func initializeCryptoKey() error {
	cryptoKey = CryptoKeyPartial + getEnvironmentVariable("CryptoKey")
	if len(cryptoKey) != 32 {
		cryptoKey = ""
		if isLocalhost {
			return nil
		}
		return apperrorWrapSimpleError(
			nil,
			"Invalid crypto key length: make sure environment variable is set properly",
		)
	}
	return nil
}

func decryptFromEnvironmentVariable(name string) (string, error) {
	var value = getEnvironmentVariable(name)
	if value == "" {
		return "", nil
	}
	var result, err = cryptoDecrypt(
		value,
		cryptoKey,
	)
	if err != nil {
		if isLocalhost {
			return value, nil
		}
		return "",
			apperrorWrapSimpleError(
				err,
				"Failed to decrypt environment variable [%v]",
				name,
			)
	}
	return result, nil
}

func initializeEncryptedEnvironmentVariables() error {
	var clientCertError error
	var clientKeyError error
	var serverCertError error
	var serverKeyError error
	var caCertError error
	clientCertContent, clientCertError = decryptFromEnvironmentVariableFunc("ClientCertContent")
	clientKeyContent, clientKeyError = decryptFromEnvironmentVariableFunc("ClientKeyContent")
	serverCertContent, serverCertError = decryptFromEnvironmentVariableFunc("ServerCertContent")
	serverKeyContent, serverKeyError = decryptFromEnvironmentVariableFunc("ServerKeyContent")
	caCertContent, caCertError = decryptFromEnvironmentVariableFunc("CaCertContent")
	return apperrorConsolidateAllErrors(
		"Failed to decrypt environment variables",
		clientCertError,
		clientKeyError,
		serverCertError,
		serverKeyError,
		caCertError,
	)
}

// Initialize initiates all application related configuration properties
func Initialize() error {
	initializeBootTimeFunc()
	var environmentVariableError = initializeGeneralEnvironmentVariablesFunc()
	if environmentVariableError != nil {
		return apperrorWrapSimpleError(
			environmentVariableError,
			"Failed to load general environment variables",
		)
	}
	var cryptoKeyError = initializeCryptoKeyFunc()
	if cryptoKeyError != nil {
		return apperrorWrapSimpleError(
			cryptoKeyError,
			"Failed to load crypto key from environment variables",
		)
	}
	environmentVariableError = initializeEncryptedEnvironmentVariablesFunc()
	if environmentVariableError != nil {
		return apperrorWrapSimpleError(
			environmentVariableError,
			"Failed to load encrypted environment variables",
		)
	}
	return nil
}

// AppVersion returns the version information of the application
func AppVersion() string {
	return appVersion
}

// AppPort returns the hosting port of the application
func AppPort() string {
	return appPort
}

// AppName returns the name of the application
func AppName() string {
	return appName
}

// AppPath returns the execution path of the application
func AppPath() string {
	return appPath
}

// IsLocalhost returns the control switch for whether or not the current running environment is a localhost (testing) environment; for localhost, no encryption/decryption is needed, and logs print all details
func IsLocalhost() bool {
	return isLocalhost
}

// CryptoKey returns the encryption/decryption key of the application
func CryptoKey() string {
	return cryptoKey
}

// SendClientCert returns the control switch for whether or not sending client certificate over external HTTP/S communications
func SendClientCert() bool {
	return sendClientCert
}

// ClientCertContent returns the client certificate cert content of the application
func ClientCertContent() string {
	return clientCertContent
}

// ClientKeyContent returns the client certificate key content of the application
func ClientKeyContent() string {
	return clientKeyContent
}

// ServeHTTPS returns the control switch for whether or not hosting the web-service with HTTPS
func ServeHTTPS() bool {
	return serveHTTPS
}

// ServerCertContent returns the server certificate cert content of the application
func ServerCertContent() string {
	return serverCertContent
}

// ServerKeyContent returns the server certificate key content of the application
func ServerKeyContent() string {
	return serverKeyContent
}

// ValidateClientCert returns the control switch for whether or not validating the client certificate of incoming HTTP/S requests
func ValidateClientCert() bool {
	return validateClientCert
}

// CaCertContent returns the CA certificate cert content of the application
func CaCertContent() string {
	return caCertContent
}
