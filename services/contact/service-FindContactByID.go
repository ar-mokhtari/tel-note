package contact

import (
	"tel-note/protocol"
)

func FindContactByID(id uint) (protocol.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.Contact{}
}
