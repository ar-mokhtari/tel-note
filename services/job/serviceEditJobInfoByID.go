package job

import (
	"tel-note/protocol"
)

func EditJobInfoByID(ID uint, NewJob protocol.Job) protocol.ResponseStatus {
	if storage.EditJob(ID, NewJob) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
