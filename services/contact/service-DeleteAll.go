package contact

import "tel-note/protocol"

func DeleteAll() *protocol.ResponseStatus {
	if storage.DeleteAll() {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}
