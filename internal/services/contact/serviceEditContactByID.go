package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func EditContactByID(newData protocol.Contact, ID uint) *config.ResponseStatus {
	if storage.EditContactByID(newData, ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: true, String: "not found"}
}
