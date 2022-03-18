package person

import (
	"tel-note/protocol"
)

func EditPerson(ID uint, newPerson protocol.Person) protocol.ResponseStatus {
	if storage.EditPerson(ID, newPerson) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
