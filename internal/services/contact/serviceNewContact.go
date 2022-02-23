package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

func NewContact(inputContact protocol.Contact) (status config.ResponseStatus, resData protocol.ContactStorage) {
	if status.State, globalVars.AllContact = storage.AddContact(inputContact); status.State {
		return config.ResponseStatus{State: true}, resData
	}
	return config.ResponseStatus{State: false}, protocol.ContactStorage{}

}
