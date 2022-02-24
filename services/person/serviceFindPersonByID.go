package person

import (
	"tel-note/protocol"
)

func FindPersonByID(ID uint) (state protocol.ResponseStatus, result protocol.Person) {
	if status, res := storage.FindPersonByID(ID); status {
		return protocol.ResponseStatus{State: true}, res
	}
	return protocol.ResponseStatus{State: false}, protocol.Person{}
}
