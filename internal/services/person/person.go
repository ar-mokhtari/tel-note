package person

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindPersonByChar(inputChar string) (state config.ResponseStatus, result []uint) {
	if status, res := Storage.FindPersonByChar(inputChar); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, []uint{}
}

func FindPersonByID(ID uint) (state config.ResponseStatus, result protocol.Person) {
	if status, res := Storage.FindPersonByID(ID); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, protocol.Person{}
}

func NewPerson(person protocol.Person) config.ResponseStatus {
	if Storage.NewPerson(person) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func EditPerson(ID uint, newPerson protocol.Person) config.ResponseStatus {
	if Storage.EditPerson(ID, newPerson) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeletePerson(IDS []uint) []uint {
	return DeletePerson(IDS)
}
