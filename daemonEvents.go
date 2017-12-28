package common

// DaemonEvent type
//go:generate stringer -type=DaemonEvent
type DaemonEvent int

// DaemonEvents const
const (
	InitializationComplete DaemonEvent = iota

	DaemonEventCommandUpdateCodeInfo
	DaemonEventCommandUpdateRuntimeInfo
	DaemonEventCommandStart
	DaemonEventCommandKill
	DaemonEventCommandUpload

	DaemonEventCodeSuccessfullyStarted
	DaemonEventCodeNotStarted

	DaemonEventFailedToKillCode
	DaemonEventKilledCode

	DaemonEventResultUploadingError
	DaemonEventResultUploadingFinished
	DaemonEventCodeFinishesOk
	DaemonEventCodeFinishesError
)
