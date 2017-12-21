package common

// DaemonState type
//go:generate stringer -type=DaemonState
type DaemonState int

// States
const (
	Initializing DaemonState = iota

	ListeningForInfo
	CodeInfoReceived
	// RuntimeInfoReceived

	ReadyToStartCode
	StartingCode
	CodeRunning
	CodeBeingKilled
	CodeFailedToKill
	CodeKilled
	CodeStopped

	CodeError
	CodeErrorNotStarted

	ResultUploadingState
	ResultUploadingErrorState
	ResultUploadingCompleteState
)
