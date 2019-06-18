package model

// ParameterType defines the type specification of a route parameter
type ParameterType string

// These are constants for parameter types and their corresponding replacement RegExp statements
const (
	ParameterTypeString  ParameterType = "string"
	regexpForString      string        = "\\w+"
	ParameterTypeInteger ParameterType = "int"
	regexpForInteger     string        = "\\d+"
	ParameterTypeUUID    ParameterType = "uuid"
	regexpForUUID        string        = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
)

// Parameter holds the information of parameters in route path definitions
type Parameter struct {
	Name string
	Type ParameterType
}

// ParameterTypeMap exposes the replacement mapping for parameters
var ParameterTypeMap = map[ParameterType]string{
	ParameterTypeString:  regexpForString,
	ParameterTypeInteger: regexpForInteger,
	ParameterTypeUUID:    regexpForUUID,
}
