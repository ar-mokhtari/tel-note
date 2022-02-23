package person

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func EditPerson(ID uint, newPerson protocol.Person) config.ResponseStatus {
	if storage.EditPerson(ID, newPerson) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
