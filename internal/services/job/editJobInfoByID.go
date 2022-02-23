package job

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func EditJobInfoByID(ID uint, NewJob protocol.Job) config.ResponseStatus {
	if storage.EditJob(ID, NewJob) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
