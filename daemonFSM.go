package common

import "github.com/looplab/fsm"

// DaemonFSM The finite state machine that the daemon will follow
// only an fsm around the code lifecycle
var DaemonFSM = fsm.Events{
	{
		Name: InitializationComplete.String(),
		Src:  []string{DaemonStateInitializing.String()},
		Dst:  DaemonStateListeningForInfo.String(),
	},
	{
		Name: DaemonEventCommandUpdateCodeInfo.String(),
		Src:  []string{DaemonStateListeningForInfo.String()},
		Dst:  DaemonStateCodeInfoReceived.String(),
	},
	{
		Name: DaemonEventCommandUpdateRuntimeInfo.String(),
		Src:  []string{DaemonStateCodeInfoReceived.String()},
		Dst:  DaemonStateReadyToStartCode.String(),
	},
	{
		Name: DaemonEventCommandStart.String(),
		Src:  []string{DaemonStateReadyToStartCode.String()},
		Dst:  DaemonStateStartingCode.String(),
	},
	{
		Name: DaemonEventCodeSuccessfullyStarted.String(),
		Src:  []string{DaemonStateStartingCode.String()},
		Dst:  DaemonStateCodeRunning.String(),
	},
	{
		Name: DaemonEventCodeNotStarted.String(),
		Src:  []string{DaemonStateStartingCode.String()},
		Dst:  DaemonStateCodeErrorNotStarted.String(),
	},
	{
		Name: DaemonEventCodeFinishesOk.String(),
		Src:  []string{DaemonStateCodeRunning.String()},
		Dst:  DaemonStateCodeStopped.String(),
	},
	{
		Name: DaemonEventCodeFinishesError.String(),
		Src:  []string{DaemonStateCodeRunning.String()},
		Dst:  DaemonStateCodeError.String(),
	},
	{
		Name: DaemonEventCommandKill.String(),
		Src:  []string{DaemonStateCodeRunning.String()},
		Dst:  DaemonStateCodeBeingKilled.String(),
	},
	{
		Name: DaemonEventFailedToKillCode.String(),
		Src:  []string{DaemonStateCodeBeingKilled.String()},
		Dst:  DaemonStateCodeFailedToKill.String(),
	},
	{
		Name: DaemonEventKilledCode.String(),
		Src:  []string{DaemonStateCodeBeingKilled.String()},
		Dst:  DaemonStateCodeKilled.String(),
	},
	{
		Name: DaemonEventCommandUpload.String(),
		Src: []string{
			DaemonStateCodeKilled.String(),
			DaemonStateCodeError.String(),
			DaemonStateCodeStopped.String(),
			DaemonStateCodeFailedToKill.String(),
		},
		Dst: DaemonStateResultUploadingState.String(),
	},
	{
		Name: DaemonEventResultUploadingFinished.String(),
		Src: []string{
			DaemonStateResultUploadingState.String(),
		},
		Dst: ResultUploadingCompleteState.String(),
	},
	{
		Name: DaemonEventResultUploadingError.String(),
		Src: []string{
			DaemonStateResultUploadingState.String(),
		},
		Dst: ResultUploadingErrorState.String(),
	},
}
