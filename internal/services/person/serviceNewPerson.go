package person

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

func NewPerson(person protocol.Person) (status config.ResponseStatus) {
	if status.State, globalVars.AllPerson = storage.NewPerson(person); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
