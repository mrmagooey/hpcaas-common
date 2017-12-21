package common

// DaemonEvent type
//go:generate stringer -type=DaemonEvent
type DaemonEvent int

// DaemonEvents const
const (
	InitializationComplete DaemonEvent = iota

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
