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

//TODO:::How about to use map
var CityMapStore = make(map[uint][]*protocol.City)
