package globalVars

import "tel-note/protocol"

//init global variable
var (
	SexStore     protocol.SexStorage
	ContactStore protocol.ContactStorage
	PersonStore  protocol.PersonStorage
	CityStore    protocol.CityStorage
	JobStore     protocol.JobStorage
)

func Init() {
	SexStore,
		ContactStore,
		PersonStore,
		CityStore,
		JobStore =
		protocol.SexStorage{},
		protocol.ContactStorage{},
		protocol.PersonStorage{},
		protocol.CityStorage{},
		protocol.JobStorage{}
}
