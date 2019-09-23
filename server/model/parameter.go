package model

// ParameterType defines the type specification of a route parameter
type ParameterType string

// These are constants for parameter types and their corresponding replacement RegExp statements
const (
	ParameterTypeAnything ParameterType = "\\.*"
	ParameterTypeString   ParameterType = "\\w+"
	ParameterTypeInteger  ParameterType = "\\d+"
	ParameterTypeUUID     ParameterType = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	ParameterTypeDateTime ParameterType = "[\\d]{4}-[\\d]{2}-[\\d]{2}T[\\d]{2}:[\\d]{2}:[\\d]{2}(\\.[\\d]{3})?Z([\\d]{2}:[\\d]{2})?"
	ParameterTypeDate     ParameterType = "[\\d]{4}-[\\d]{2}-[\\d]{2}"
	ParameterTypeTime     ParameterType = "[\\d]{2}:[\\d]{2}:[\\d]{2}(\\.[\\d]{3})?"
	ParameterTypeBoolean  ParameterType = "(?i)(true|false)"
	ParameterTypeFloat    ParameterType = "\\d+(\\.\\d+)?"
)
