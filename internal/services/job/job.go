package job

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func NewJob(job protocol.Job) config.ResponseStatus {
	if Storage.NewJob(job) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditJobInfoByID(ID uint, NewJob protocol.Job) config.ResponseStatus {
	if Storage.EditJob(ID, NewJob) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteJobByID(ID []uint) []uint {
	return Storage.DeleteJob(ID)
}
