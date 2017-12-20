package common

// Event type
//go:generate stringer -type=Event
type Event int

// Events const
const (
	InitializationCompleteEvent Event = iota
	StartCommandReceived
	KillCommandReceived
	DaemonFailedToKillCode
	DaemonKilledCode
	UploadCommandReceived
	UpdateReceived
	UpdateFinished
	ResultUploadingErrorEvent
	ResultUploadingFinishedEvent
	CodeFinishesOk
	CodeFinishesError
)
