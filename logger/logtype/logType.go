package logtype

// LogType is the category of logging
type LogType int

// These are the enum definitions of log categories and presets
const (
	AppRoot  LogType = 0
	APIEnter LogType = 1 << iota
	APIRequest
	MethodEnter
	MethodParameter
	MethodLogic
	DependencyCall
	DependencyRequest
	DependencyResponse
	DependencyFinish
	MethodReturn
	MethodExit
	APIResponse
	APIExit
	GeneralTracing   LogType = APIEnter | APIExit
	VerboseTracing   LogType = GeneralTracing | MethodEnter | MethodExit
	FullTracing      LogType = VerboseTracing | DependencyCall | DependencyFinish
	BasicDebugging   LogType = MethodLogic
	GeneralDebugging LogType = BasicDebugging | APIRequest | APIResponse
	VerboseDebugging LogType = GeneralDebugging | MethodReturn | MethodExit
	FullDebugging    LogType = VerboseDebugging | DependencyRequest | DependencyResponse
	BasicLogging     LogType = BasicDebugging
	GeneralLogging   LogType = BasicLogging | GeneralTracing | GeneralDebugging
	VerboseLogging   LogType = GeneralLogging | VerboseTracing | VerboseDebugging
	FullLogging      LogType = VerboseLogging | FullTracing | FullDebugging
)

// These are the string representations of log category and preset names
const (
	AppRootName            string = "AppRoot"
	APIEnterName           string = "APIEnter"
	APIRequestName         string = "APIRequest"
	MethodEnterName        string = "MethodEnter"
	MethodParameterName    string = "MethodParameter"
	MethodLogicName        string = "MethodLogic"
	DependencyCallName     string = "DependencyCall"
	DependencyRequestName  string = "DependencyRequest"
	DependencyResponseName string = "DependencyResponse"
	DependencyFinishName   string = "DependencyFinish"
	MethodReturnName       string = "MethodReturn"
	MethodExitName         string = "MethodExit"
	APIResponseName        string = "APIResponse"
	APIExitName            string = "APIExit"
	GeneralTracingName     string = "GeneralTracing"
	VerboseTracingName     string = "VerboseTracing"
	FullTracingName        string = "FullTracing"
	BasicDebuggingName     string = "BasicDebugging"
	GeneralDebuggingName   string = "GeneralDebugging"
	VerboseDebuggingName   string = "VerboseDebugging"
	FullDebuggingName      string = "FullDebugging"
	BasicLoggingName       string = "BasicLogging"
	GeneralLoggingName     string = "GeneralLogging"
	VerboseLoggingName     string = "VerboseLogging"
	FullLoggingName        string = "FullLogging"
)

var supportedLogCategories = map[LogType]string{
	APIEnter:           APIEnterName,
	APIRequest:         APIRequestName,
	MethodEnter:        MethodEnterName,
	MethodParameter:    MethodParameterName,
	MethodLogic:        MethodLogicName,
	DependencyCall:     DependencyCallName,
	DependencyRequest:  DependencyRequestName,
	DependencyResponse: DependencyResponseName,
	DependencyFinish:   DependencyFinishName,
	MethodReturn:       MethodReturnName,
	MethodExit:         MethodExitName,
	APIResponse:        APIResponseName,
	APIExit:            APIExitName,
}

var logTypeNameMapping = map[string]LogType{
	AppRootName:            AppRoot,
	APIEnterName:           APIEnter,
	APIRequestName:         APIRequest,
	MethodEnterName:        MethodEnter,
	MethodParameterName:    MethodParameter,
	MethodLogicName:        MethodLogic,
	DependencyCallName:     DependencyCall,
	DependencyRequestName:  DependencyRequest,
	DependencyResponseName: DependencyResponse,
	DependencyFinishName:   DependencyFinish,
	MethodReturnName:       MethodReturn,
	MethodExitName:         MethodExit,
	APIResponseName:        APIResponse,
	APIExitName:            APIExit,
	GeneralTracingName:     GeneralTracing,
	VerboseTracingName:     VerboseTracing,
	FullTracingName:        FullTracing,
	BasicDebuggingName:     BasicDebugging,
	GeneralDebuggingName:   GeneralDebugging,
	VerboseDebuggingName:   VerboseDebugging,
	FullDebuggingName:      FullDebugging,
	BasicLoggingName:       BasicLogging,
	GeneralLoggingName:     GeneralLogging,
	VerboseLoggingName:     VerboseLogging,
	FullLoggingName:        FullLogging,
}

// FromString converts a LogType flag instance to its string representation
func (logCategory LogType) String() string {
	if logCategory == AppRoot {
		return AppRootName
	}
	var result []string
	for key, value := range supportedLogCategories {
		if logCategory&key == key {
			result = append(result, value)
		}
	}
	return stringsJoin(result, "|")
}

// HasFlag checks whether this log category has the flag set or not
func (logCategory LogType) HasFlag(flag LogType) bool {
	if flag == AppRoot {
		return true
	}
	if logCategory&flag == flag {
		return true
	}
	return false
}

// FromString converts a string representation of LogType flag to its strongly typed instance
func FromString(value string) LogType {
	var logType, found = logTypeNameMapping[value]
	if !found {
		return AppRoot
	}
	return logType
}
