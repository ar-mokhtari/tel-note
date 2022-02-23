package job

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

func NewJob(job protocol.Job) (status config.ResponseStatus) {
	if status.State, globalVars.AllJob = storage.NewJob(job); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
