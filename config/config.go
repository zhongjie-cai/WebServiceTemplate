package config

var (
	// appVersion is for injecting the version information from compiler
	appVersion = "0.0.0.0"
	// appPort is for injecting the port information from compiler
	appPort = "443"
	// appName is for injecting the name of application from compiler
	appName = "WebServiceTemplate"
	// appPath is for injecting the application executable path from compiler
	appPath = "."
	// cryptoKey is the key phrase that is used to encrypt & decrypt secrets; consists of two parts, one in code and another configured in environment variable
	cryptoKey = ""
	// bootTime is the time when the host has started execution
	bootTime = ""
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

func initializeCryptoKey() error {
	cryptoKey = CryptoKeyPartial + getEnvironmentVariable("CryptoKey")
	if len(cryptoKey) != 32 {
		cryptoKey = ""
		return apperrorWrapSimpleError(
			nil,
			"Invalid crypto key length: make sure environment variable is set properly",
		)
	}
	return nil
}

func initializeEnvironmentVariables() error {
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
		return value,
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
	caCertContent, caCertError = decryptFromEnvironmentVariableFunc("CACertContent")
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
	var cryptoKeyError = initializeCryptoKeyFunc()
	if cryptoKeyError != nil {
		return apperrorWrapSimpleError(
			cryptoKeyError,
			"Failed to initialize crypto key",
		)
	}
	var environmentVariableError = initializeEnvironmentVariablesFunc()
	if environmentVariableError != nil {
		return apperrorWrapSimpleError(
			environmentVariableError,
			"Failed to load environment variables",
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

// CryptoKey returns the encryption/decryption key of the application
func CryptoKey() string {
	return cryptoKey
}

// ClientCertContent returns the client certificate cert content of the application
func ClientCertContent() string {
	return clientCertContent
}

// ClientKeyContent returns the client certificate key content of the application
func ClientKeyContent() string {
	return clientKeyContent
}

// ServerCertContent returns the server certificate cert content of the application
func ServerCertContent() string {
	return serverCertContent
}

// ServerKeyContent returns the server certificate key content of the application
func ServerKeyContent() string {
	return serverKeyContent
}

// CACertContent returns the CA certificate cert content of the application
func CACertContent() string {
	return caCertContent
}
