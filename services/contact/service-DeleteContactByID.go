package contact

import "tel-note/protocol"

func DeleteContactByID(ID uint) *protocol.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}
