// Code generated by "stringer -type=DaemonState"; DO NOT EDIT.

package common

import "strconv"

const _DaemonState_name = "InitializingListeningForInfoCodeInfoReceivedReadyToStartCodeStartingCodeCodeRunningCodeBeingKilledCodeFailedToKillCodeKilledCodeStoppedCodeErrorCodeErrorNotStartedResultUploadingStateResultUploadingErrorStateResultUploadingCompleteState"

var _DaemonState_index = [...]uint8{0, 12, 28, 44, 60, 72, 83, 98, 114, 124, 135, 144, 163, 183, 208, 236}

func (i DaemonState) String() string {
	if i < 0 || i >= DaemonState(len(_DaemonState_index)-1) {
		return "DaemonState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DaemonState_name[_DaemonState_index[i]:_DaemonState_index[i+1]]
}
