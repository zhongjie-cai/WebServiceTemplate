package logtype

import "strings"

// LogType is the entry type of logging
type LogType int

// These are the enum definitions of log types and presets
const (
	AppRoot  LogType = 0
	APIEnter LogType = 1 << iota
	APIRequest
	MethodEnter
	MethodParameter
	MethodLogic
	NetworkCall
	NetworkRequest
	NetworkResponse
	NetworkFinish
	MethodReturn
	MethodExit
	APIResponse
	APIExit

	BasicTracing   LogType = MethodLogic
	GeneralTracing LogType = BasicTracing | APIEnter | APIExit
	VerboseTracing LogType = GeneralTracing | MethodEnter | MethodExit
	FullTracing    LogType = VerboseTracing | NetworkCall | NetworkFinish

	BasicDebugging   LogType = MethodLogic
	GeneralDebugging LogType = BasicDebugging | APIRequest | APIResponse
	VerboseDebugging LogType = GeneralDebugging | MethodParameter | MethodReturn
	FullDebugging    LogType = VerboseDebugging | NetworkRequest | NetworkResponse

	BasicLogging   LogType = BasicTracing | BasicDebugging
	GeneralLogging LogType = BasicLogging | GeneralTracing | GeneralDebugging
	VerboseLogging LogType = GeneralLogging | VerboseTracing | VerboseDebugging
	FullLogging    LogType = VerboseLogging | FullTracing | FullDebugging
)

// These are the string representations of log category and preset names
const (
	AppRootName         string = "AppRoot"
	APIEnterName        string = "APIEnter"
	APIRequestName      string = "APIRequest"
	MethodEnterName     string = "MethodEnter"
	MethodParameterName string = "MethodParameter"
	MethodLogicName     string = "MethodLogic"
	NetworkCallName     string = "NetworkCall"
	NetworkRequestName  string = "NetworkRequest"
	NetworkResponseName string = "NetworkResponse"
	NetworkFinishName   string = "NetworkFinish"
	MethodReturnName    string = "MethodReturn"
	MethodExitName      string = "MethodExit"
	APIResponseName     string = "APIResponse"
	APIExitName         string = "APIExit"

	BasicTracingName   string = "BasicTracing"
	GeneralTracingName string = "GeneralTracing"
	VerboseTracingName string = "VerboseTracing"
	FullTracingName    string = "FullTracing"

	BasicDebuggingName   string = "BasicDebugging"
	GeneralDebuggingName string = "GeneralDebugging"
	VerboseDebuggingName string = "VerboseDebugging"
	FullDebuggingName    string = "FullDebugging"

	BasicLoggingName   string = "BasicLogging"
	GeneralLoggingName string = "GeneralLogging"
	VerboseLoggingName string = "VerboseLogging"
	FullLoggingName    string = "FullLogging"
)

var supportedLogTypes = map[LogType]string{
	APIEnter:        APIEnterName,
	APIRequest:      APIRequestName,
	MethodEnter:     MethodEnterName,
	MethodParameter: MethodParameterName,
	MethodLogic:     MethodLogicName,
	NetworkCall:     NetworkCallName,
	NetworkRequest:  NetworkRequestName,
	NetworkResponse: NetworkResponseName,
	NetworkFinish:   NetworkFinishName,
	MethodReturn:    MethodReturnName,
	MethodExit:      MethodExitName,
	APIResponse:     APIResponseName,
	APIExit:         APIExitName,
}

var logTypeNameMapping = map[string]LogType{
	AppRootName:          AppRoot,
	APIEnterName:         APIEnter,
	APIRequestName:       APIRequest,
	MethodEnterName:      MethodEnter,
	MethodParameterName:  MethodParameter,
	MethodLogicName:      MethodLogic,
	NetworkCallName:      NetworkCall,
	NetworkRequestName:   NetworkRequest,
	NetworkResponseName:  NetworkResponse,
	NetworkFinishName:    NetworkFinish,
	MethodReturnName:     MethodReturn,
	MethodExitName:       MethodExit,
	APIResponseName:      APIResponse,
	APIExitName:          APIExit,
	BasicTracingName:     BasicTracing,
	GeneralTracingName:   GeneralTracing,
	VerboseTracingName:   VerboseTracing,
	FullTracingName:      FullTracing,
	BasicDebuggingName:   BasicDebugging,
	GeneralDebuggingName: GeneralDebugging,
	VerboseDebuggingName: VerboseDebugging,
	FullDebuggingName:    FullDebugging,
	BasicLoggingName:     BasicLogging,
	GeneralLoggingName:   GeneralLogging,
	VerboseLoggingName:   VerboseLogging,
	FullLoggingName:      FullLogging,
}

// FromString converts a LogType flag instance to its string representation
func (logtype LogType) String() string {
	if logtype == AppRoot {
		return AppRootName
	}
	var result []string
	for key, value := range supportedLogTypes {
		if logtype&key == key {
			result = append(result, value)
		}
	}
	return stringsJoin(result, "|")
}

// HasFlag checks whether this log category has the flag set or not
func (logtype LogType) HasFlag(flag LogType) bool {
	if flag == AppRoot {
		return true
	}
	if logtype&flag == flag {
		return true
	}
	return false
}

// FromString converts a string representation of LogType flag to its strongly typed instance
func FromString(value string) LogType {
	var splitValues = strings.Split(
		value,
		"|",
	)
	var combinedLogType LogType
	for _, splitValue := range splitValues {
		var logType, found = logTypeNameMapping[splitValue]
		if found {
			combinedLogType = combinedLogType | logType
		}
	}
	return combinedLogType
}
