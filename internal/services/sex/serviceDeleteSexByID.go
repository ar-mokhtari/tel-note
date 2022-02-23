package sex

import (
	"tel-note/internal/config"
)

func DeleteSexByID(ID uint8) config.ResponseStatus {
	if storage.DeleteSex(ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
