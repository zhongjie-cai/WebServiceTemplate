package loglevel

// LogLevel is the severity level of logging
type LogLevel int

// These are the enum definitions of log types and presets
const (
	Debug LogLevel = 0
	Info  LogLevel = iota
	Warn
	Error
	Fatal
	maxLogLevel
)

// These are the string representations of log category and preset names
const (
	DebugName string = "Debug"
	InfoName  string = "Info"
	WarnName  string = "Warn"
	ErrorName string = "Error"
	FatalName string = "Fatal"
)

var supportedLogLevels = map[LogLevel]string{
	Debug: DebugName,
	Info:  InfoName,
	Warn:  WarnName,
	Error: ErrorName,
	Fatal: FatalName,
}

var logLevelNameMapping = map[string]LogLevel{
	DebugName: Debug,
	InfoName:  Info,
	WarnName:  Warn,
	ErrorName: Error,
	FatalName: Fatal,
}

// FromString converts a LogLevel flag instance to its string representation
func (logLevel LogLevel) String() string {
	for key, value := range supportedLogLevels {
		if logLevel == key {
			return value
		}
	}
	return DebugName
}

// FromString converts a string representation of LogLevel flag to its strongly typed instance
func FromString(value string) LogLevel {
	var logLevel, found = logLevelNameMapping[value]
	if !found {
		return Debug
	}
	return logLevel
}
