package globalVars

import "tel-note/internal/protocol"

//init global variable
var (
	AllContact protocol.ContactStorage
	AllPerson  protocol.PersonStorage
)

func Init() {
	AllContact, AllPerson =
		protocol.ContactStorage{},
		protocol.PersonStorage{}
}
