package sex

import (
	"tel-note/internal/config"
)

func EditSexByID(ID uint8, NewSexName string) config.ResponseStatus {
	if storage.EditSex(ID, NewSexName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
