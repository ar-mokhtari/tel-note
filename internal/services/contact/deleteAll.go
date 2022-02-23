package contact

import (
	"tel-note/internal/config"
)

func DeleteAll() *config.ResponseStatus {
	if storage.DeleteAll() {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}
