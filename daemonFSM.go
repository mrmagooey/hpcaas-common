package common

import "github.com/looplab/fsm"

// DaemonFSM The finite state machine that the daemon will follow
// only an fsm around the code lifecycle
var DaemonFSM = fsm.Events{
	{
		Name: InitializationComplete.String(),
		Src:  []string{Initializing.String()},
		Dst:  ListeningForInfo.String(),
	},
	{
		Name: ReceiveCodeInfo.String(),
		Src:  []string{ListeningForInfo.String()},
		Dst:  CodeInfoReceived.String(),
	},
	{
		Name: ReceiveRuntimeInfo.String(),
		Src:  []string{CodeInfoReceived.String()},
		Dst:  ReadyToStartCode.String(),
	},
	// {
	// 	Name: ReadyToStart.String(),
	// 	Src:  []string{RuntimeInfoReceived.String()},
	// 	Dst:  ReadyToStartCode.String(),
	// },
	{
		Name: StartCommandReceived.String(),
		Src:  []string{ReadyToStartCode.String()},
		Dst:  StartingCode.String(),
	},
	{
		Name: CodeSuccessfullyStarted.String(),
		Src:  []string{StartingCode.String()},
		Dst:  CodeRunning.String(),
	},
	{
		Name: CodeNotStarted.String(),
		Src:  []string{StartingCode.String()},
		Dst:  CodeErrorNotStarted.String(),
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
		Name: FailedToKillCode.String(),
		Src:  []string{CodeBeingKilled.String()},
		Dst:  CodeFailedToKill.String(),
	},
	{
		Name: KilledCode.String(),
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
		Name: ResultUploadingFinished.String(),
		Src: []string{
			ResultUploadingState.String(),
		},
		Dst: ResultUploadingCompleteState.String(),
	},
	{
		Name: ResultUploadingError.String(),
		Src: []string{
			ResultUploadingState.String(),
		},
		Dst: ResultUploadingErrorState.String(),
	},
}
