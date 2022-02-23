package sex

import (
	"tel-note/internal/config"
)

func NewSex(SexName string) config.ResponseStatus {
	if storage.NewSex(SexName) {
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
