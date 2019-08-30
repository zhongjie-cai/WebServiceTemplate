package model

// Route holds the registration information of a dynamic route hosting
type Route struct {
	Endpoint   string
	Method     string
	Path       string
	Parameters map[string]ParameterType
	Queries    map[string]ParameterType
	ActionFunc ActionFunc
}
