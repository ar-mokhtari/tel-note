package person

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

var inputTargetData struct{}

func FindPersonByChar(inputChar string) (state config.ResponseStatus, result []uint) {
	if status, res := storage.FindPersonByChar(inputChar); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, []uint{}
}

func FindPersonByID(ID uint) (state config.ResponseStatus, result protocol.Person) {
	if status, res := storage.FindPersonByID(ID); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, protocol.Person{}
}

func NewPerson(person protocol.Person) (status config.ResponseStatus) {
	if status.State, globalVars.AllPerson = storage.NewPerson(person); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func EditPerson(ID uint, newPerson protocol.Person) config.ResponseStatus {
	if storage.EditPerson(ID, newPerson) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeletePerson(IDS []uint) []uint {
	return DeletePerson(IDS)
}
