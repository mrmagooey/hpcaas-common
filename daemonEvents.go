package common

// DaemonEvent type
//go:generate stringer -type=DaemonEvent
type DaemonEvent int

// DaemonEvents const
const (
	InitializationComplete DaemonEvent = iota

	CommandUpdateCodeInfo
	CommandUpdateRuntimeInfo
	CommandStart
	CommandKill
	CommandUpload

	CodeSuccessfullyStarted
	CodeNotStarted

	FailedToKillCode
	KilledCode

	ResultUploadingError
	ResultUploadingFinished
	CodeFinishesOk
	CodeFinishesError
)
