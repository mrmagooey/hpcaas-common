package common

// DaemonState type
//go:generate stringer -type=DaemonState
type DaemonState int

// States
const (
	Initializing DaemonState = iota
	Preparation
	UpdateInProgress
	CodeRunning
	CodeBeingKilled
	CodeFailedToKill
	CodeKilled
	CodeStopped
	CodeError
	ResultUploadingState
	ResultUploadingErrorState
	ResultUploadingCompleteState
)
