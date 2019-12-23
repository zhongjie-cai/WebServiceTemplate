package headerstyle

// HeaderStyle is the style of logging header
type HeaderStyle int

// These are the enum definitions of log types and presets
const (
	DoNotLog    HeaderStyle = 0
	LogCombined HeaderStyle = iota
	LogPerName
	LogPerValue
	maxHeaderStyle
)

// These are the string representations of log category and preset names
const (
	doNotLogName    string = "DoNotLog"
	logCombinedName string = "LogCombined"
	logPerNameName  string = "LogPerName"
	logPerValueName string = "LogPerValue"
)

var supportedHeaderStyles = map[HeaderStyle]string{
	DoNotLog:    doNotLogName,
	LogCombined: logCombinedName,
	LogPerName:  logPerNameName,
	LogPerValue: logPerValueName,
}

var headerStyleNameMapping = map[string]HeaderStyle{
	doNotLogName:    DoNotLog,
	logCombinedName: LogCombined,
	logPerNameName:  LogPerName,
	logPerValueName: LogPerValue,
}

// FromString converts a HeaderStyle flag instance to its string representation
func (headerStyle HeaderStyle) String() string {
	for key, value := range supportedHeaderStyles {
		if headerStyle == key {
			return value
		}
	}
	return doNotLogName
}

// FromString converts a string representation of HeaderStyle flag to its strongly typed instance
func FromString(value string) HeaderStyle {
	var headerStyle, found = headerStyleNameMapping[value]
	if !found {
		return DoNotLog
	}
	return headerStyle
}
