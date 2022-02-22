package globalVars

import "tel-note/internal/protocol"

//init global variable
var (
	AllSex     protocol.SexStorage
	AllContact protocol.ContactStorage
	AllPerson  protocol.PersonStorage
	AllCity    protocol.CityStorage
	AllJob     protocol.JobStorage
)

func Init() {
	AllSex,
		AllContact,
		AllPerson,
		AllCity,
		AllJob =
		protocol.SexStorage{},
		protocol.ContactStorage{},
		protocol.PersonStorage{},
		protocol.CityStorage{},
		protocol.JobStorage{}
}
