package sex

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
)

func NewSex(SexName protocol.Sex) (status protocol.ResponseStatus) {
	if status.State, globalVars.SexStore = storage.NewSex(SexName); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
