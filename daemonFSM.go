package common

import "github.com/looplab/fsm"

// DaemonFSM The finite state machine that the daemon will follow
// only an fsm around the code lifecycle
var DaemonFSM = fsm.Events{
	{
		Name: InitializationCompleteEvent.String(),
		Src:  []string{Initializing.String()},
		Dst:  Preparation.String(),
	},
	{
		Name: UpdateReceived.String(),
		Src:  []string{Preparation.String()},
		Dst:  UpdateInProgress.String(),
	},
	{
		Name: UpdateFinished.String(),
		Src:  []string{UpdateInProgress.String()},
		Dst:  Preparation.String(),
	},
	{
		Name: StartCommandReceived.String(),
		Src:  []string{Preparation.String()},
		Dst:  CodeRunning.String(),
	},
	{
		Name: CodeFinishesOk.String(),
		Src:  []string{CodeRunning.String()},
		Dst:  CodeStopped.String(),
	},
	{
		Name: CodeFinishesError.String(),
		Src:  []string{CodeRunning.String()},
		Dst:  CodeError.String(),
	},
	{
		Name: KillCommandReceived.String(),
		Src:  []string{CodeRunning.String()},
		Dst:  CodeBeingKilled.String(),
	},
	{
		Name: DaemonFailedToKillCode.String(),
		Src:  []string{CodeBeingKilled.String()},
		Dst:  CodeFailedToKill.String(),
	},
	{
		Name: DaemonKilledCode.String(),
		Src:  []string{CodeBeingKilled.String()},
		Dst:  CodeKilled.String(),
	},
	{
		Name: UploadCommandReceived.String(),
		Src: []string{
			CodeKilled.String(),
			CodeError.String(),
			CodeStopped.String(),
			CodeFailedToKill.String(),
		},
		Dst: ResultUploadingState.String(),
	},
	{
		Name: ResultUploadingFinishedEvent.String(),
		Src: []string{
			ResultUploadingState.String(),
		},
		Dst: ResultUploadingCompleteState.String(),
	},
	{
		Name: ResultUploadingErrorEvent.String(),
		Src: []string{
			ResultUploadingState.String(),
		},
		Dst: ResultUploadingErrorState.String(),
	},
}
