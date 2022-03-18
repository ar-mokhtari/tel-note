package contact

import (
	"tel-note/protocol"
)

func NewContact(inputContact protocol.Contact) (status protocol.ResponseStatus, resData []*protocol.Contact) {
	if status.State = storage.AddContact(inputContact); status.State {
		return protocol.ResponseStatus{State: true}, resData
	}
	return protocol.ResponseStatus{State: false}, []*protocol.Contact{}

}
