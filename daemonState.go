package common

// DaemonState type
//go:generate stringer -type=DaemonState
type DaemonState int

// Daemon States
const (
	DaemonStateInitializing DaemonState = iota

	DaemonStateListeningForInfo
	DaemonStateCodeInfoReceived
	// RuntimeInfoReceived

	DaemonStateReadyToStartCode
	DaemonStateStartingCode
	DaemonStateCodeRunning
	DaemonStateCodeBeingKilled
	DaemonStateCodeFailedToKill
	DaemonStateCodeKilled
	DaemonStateCodeStopped

	DaemonStateCodeError
	DaemonStateCodeErrorNotStarted

	DaemonStateResultUploadingState
	ResultUploadingErrorState
	ResultUploadingCompleteState
)
