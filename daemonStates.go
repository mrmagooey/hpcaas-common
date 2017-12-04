package common

// DaemonState new type so that we can add our methods
type DaemonState int

// The set of possible daemon states
const (
	DaemonStartedState DaemonState = iota
	DaemonRunningState
	DaemonErrorState
)

// DaemonStates A slice containing all DaemonState strings
var DaemonStates = []string{
	"DaemonStartedState",
	"DaemonRunningState",
	"DaemonErrorState",
}

func (cs DaemonState) String() string {
	return DaemonStates[int(cs)]
}

// CodeState new type so that we can add our methods
type CodeState int

// The set of user code states
// these are all mutually exclusive
const (
	CodeWaitingState CodeState = iota
	CodeMissingState
	CodeFailedToStartState
	CodeRunningState
	CodeStoppedState
	CodeKilledState
	CodeFailedToKillState
	CodeErrorState
)

// CodeStates A slice containing all CodeState strings
var CodeStates = []string{
	"CodeWaitingState",
	"CodeMissingState",
	"CodeFailedToStartState",
	"CodeRunningState",
	"CodeStoppedState",
	"CodeKilledState",
	"CodeFailedToKillState",
	"CodeErrorState",
}

func (cs CodeState) String() string {
	return CodeStates[int(cs)]
}

// ResultState represents the state of the results coming from the container
type ResultState int

// The set of result states
const (
	ResultWaitingState ResultState = iota
	ResultUploadingState
	ResultsUploadingFinishedState
	ResultErrorState
)

// ResultStates the names of the states
var ResultStates = []string{
	"ResultWaitingState",
	"ResultUploadingState",
	"ResultsUploadingFinishedState",
	"ResultErrorState",
}

func (rs ResultState) String() string {
	return ResultStates[int(rs)]
}

// StartedState new type so that we can add our methods
type StartedState int

// The set of user code start states
const (
	StartedByDaemonState StartedState = iota
	StartedExternallyState
)

// HowStartedStates names
var HowStartedStates = []string{
	"StartedByDaemonState",
	"StartedExternallyState",
}

func (rs StartedState) String() string {
	return HowStartedStates[int(rs)]
}
