package contact

import (
	"tel-note/protocol"
)

func NewContact(inputContact protocol.Contact) (status protocol.ResponseStatus, resData protocol.ContactStorage) {
	if status.State = storage.AddContact(inputContact); status.State {
		return protocol.ResponseStatus{State: true}, resData
	}
	return protocol.ResponseStatus{State: false}, protocol.ContactStorage{}

}
