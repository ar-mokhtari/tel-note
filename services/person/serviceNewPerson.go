package person

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
)

func NewPerson(person protocol.Person) (status protocol.ResponseStatus) {
	if status.State, globalVars.AllPerson = storage.NewPerson(person); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
