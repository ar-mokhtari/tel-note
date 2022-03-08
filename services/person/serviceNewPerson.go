package person

import (
	"tel-note/protocol"
)

func NewPerson(person protocol.Person) (status protocol.ResponseStatus) {
	if status.State = storage.NewPerson(person); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
