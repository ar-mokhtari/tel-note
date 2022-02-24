package sex

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
)

func NewSex(SexName string) (status protocol.ResponseStatus) {
	if status.State, globalVars.AllSex = storage.NewSex(SexName); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
