package sex

import (
	"tel-note/internal/config"
)

func NewSex(SexName string) config.ResponseStatus {
	if Storage.NewSex(SexName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditSexByID(ID uint8, NewSexName string) config.ResponseStatus {
	if Storage.EditSex(ID, NewSexName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteSexByID(ID uint8) config.ResponseStatus {
	if Storage.DeleteSex(ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
