package sex

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

func NewSex(SexName string) (status config.ResponseStatus) {
	if status.State, globalVars.AllSex = storage.NewSex(SexName); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditSexByID(ID uint8, NewSexName string) config.ResponseStatus {
	if storage.EditSex(ID, NewSexName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteSexByID(ID uint8) config.ResponseStatus {
	if storage.DeleteSex(ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func FindSexByID(ID uint8) (protocol.Sex, config.ResponseStatus) {
	if starte, data := storage.FindSexByID(ID); starte {
		return data, config.ResponseStatus{State: true}
	}
	return protocol.Sex{}, config.ResponseStatus{State: false}
}
