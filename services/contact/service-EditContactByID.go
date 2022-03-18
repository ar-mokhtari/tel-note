package contact

import (
	"tel-note/protocol"
)

func EditContactByID(newData protocol.Contact, ID uint) *protocol.ResponseStatus {
	if storage.EditContactByID(newData, ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: true, String: "not found"}
}
