package contact

import "tel-note/protocol"

func GetContacts() protocol.ContactStorage {
	return storage.GetContacts()
}
