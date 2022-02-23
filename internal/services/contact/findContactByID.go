package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindContactByID(id uint) (config.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.Contact{}
}
