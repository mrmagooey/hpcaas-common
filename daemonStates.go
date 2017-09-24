package common

// The set of possible daemon states
const (
	DaemonStartedState uint8 = iota
	DaemonRunningState
	DaemonErrorState
)

// CodeState new type so that we can add our methods
type CodeState int

// The set of user code states
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

// The set of user code start states
const (
	StartedByDaemonState int = iota
	StartedExternallyState
)
