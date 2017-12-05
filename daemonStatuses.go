package common

// DaemonStatus new type so that we can add our methods
type DaemonStatus int

// The set of possible daemon statuses
const (
	DaemonStartedStatus DaemonStatus = iota
	DaemonRunningStatus
	DaemonErrorStatus
)

// DaemonStatuses A slice containing all DaemonStatus strings
var DaemonStatuses = []string{
	"DaemonStartedStatus",
	"DaemonRunningStatus",
	"DaemonErrorStatus",
}

func (cs DaemonStatus) String() string {
	return DaemonStatuses[int(cs)]
}

// CodeStatus new type so that we can add our methods
type CodeStatus int

// The set of user code statuss
// these are all mutually exclusive
const (
	CodeWaitingStatus CodeStatus = iota
	CodeMissingStatus
	CodeFailedToStartStatus
	CodeRunningStatus
	CodeStoppedStatus
	CodeKilledStatus
	CodeFailedToKillStatus
	CodeErrorStatus
)

// CodeStatuses A slice containing all CodeStatus strings
var CodeStatuses = []string{
	"CodeWaitingStatus",
	"CodeMissingStatus",
	"CodeFailedToStartStatus",
	"CodeRunningStatus",
	"CodeStoppedStatus",
	"CodeKilledStatus",
	"CodeFailedToKillStatus",
	"CodeErrorStatus",
}

func (cs CodeStatus) String() string {
	return CodeStatuses[int(cs)]
}

// ResultStatus represents the status of the results coming from the container
type ResultStatus int

// The set of result statuss
const (
	ResultWaitingStatus ResultStatus = iota
	ResultUploadingStatus
	ResultsUploadingFinishedStatus
	ResultErrorStatus
)

// ResultStatuses the names of the statuss
var ResultStatuses = []string{
	"ResultWaitingStatus",
	"ResultUploadingStatus",
	"ResultsUploadingFinishedStatus",
	"ResultErrorStatus",
}

func (rs ResultStatus) String() string {
	return ResultStatuses[int(rs)]
}

// StartedStatus new type so that we can add our methods
type StartedStatus int

// The set of user code start statuss
const (
	StartedByDaemonStatus StartedStatus = iota
	StartedExternallyStatus
)

// HowStartedStatuses names
var HowStartedStatuses = []string{
	"StartedByDaemonStatus",
	"StartedExternallyStatus",
}

func (rs StartedStatus) String() string {
	return HowStartedStatuses[int(rs)]
}
