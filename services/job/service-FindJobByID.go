package job

import (
	"tel-note/protocol"
)

func FindJobByID(inputID uint) (status protocol.ResponseStatus, job protocol.Job) {
	if state, data := storage.FindJobByID(inputID); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.Job{}
}
