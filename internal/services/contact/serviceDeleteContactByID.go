package contact

import (
	"tel-note/internal/config"
)

func DeleteContactByID(ID uint) *config.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}
