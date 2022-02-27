package person

import "tel-note/protocol"

func GetPersons() protocol.PersonStorage {
	return storage.GetPersons()
}
