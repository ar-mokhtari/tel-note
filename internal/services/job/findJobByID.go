package job

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindJobByID(inputID uint) (status config.ResponseStatus, job protocol.Job) {
	if state, data := storage.FindJobByID(inputID); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.Job{}
}
