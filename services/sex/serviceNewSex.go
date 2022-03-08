package sex

import (
	"tel-note/protocol"
)

func NewSex(SexName protocol.Sex) (status protocol.ResponseStatus) {
	if status.State = storage.NewSex(SexName); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
