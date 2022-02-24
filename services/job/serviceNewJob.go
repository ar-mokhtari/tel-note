package job

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
)

func NewJob(job protocol.Job) (status protocol.ResponseStatus) {
	if status.State, globalVars.AllJob = storage.NewJob(job); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
