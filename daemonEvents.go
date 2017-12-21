package common

// Event type
//go:generate stringer -type=Event
type Event int

// Events const
const (
	InitializationComplete Event = iota

	ReceiveCodeInfo
	ReceiveRuntimeInfo

	// orchestrator initiated events
	ReadyToStart
	StartCommandReceived
	KillCommandReceived
	UploadCommandReceived

	CodeSuccessfullyStarted
	CodeNotStarted

	FailedToKillCode
	KilledCode

	ResultUploadingError
	ResultUploadingFinished
	CodeFinishesOk
	CodeFinishesError
)
