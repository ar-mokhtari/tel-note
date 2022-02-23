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

func EditJobInfoByID(ID uint, NewJob protocol.Job) config.ResponseStatus {
	if storage.EditJob(ID, NewJob) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteJobByID(ID []uint) []uint {
	return storage.DeleteJob(ID)
}

func FindJobByID(inputID uint) (status config.ResponseStatus, job protocol.Job) {
	if state, data := storage.FindJobByID(inputID); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.Job{}
}
