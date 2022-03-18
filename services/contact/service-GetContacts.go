package contact

import "tel-note/protocol"

func GetContacts() []*protocol.Contact {
	return storage.GetContacts()
}
