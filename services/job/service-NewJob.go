package job

import (
	"tel-note/protocol"
)

func NewJob(job protocol.Job) (status protocol.ResponseStatus) {
	if status.State = storage.NewJob(job); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
