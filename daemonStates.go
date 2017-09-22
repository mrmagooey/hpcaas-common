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

// The set of result states
const (
	ResultWaitingState uint8 = iota + 1
	ResultUploadingState
	ResultsUploadingFinishedState
	ResultErrorState
)

// The set of user code start states
const (
	StartedByDaemonState uint8 = iota + 1
	StartedExternallyState
)
