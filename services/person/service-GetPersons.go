package person

import "tel-note/protocol"

func GetPersons() []*protocol.Person {
	return storage.GetPersons()
}
