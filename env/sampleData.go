package env

import (
	"tel-note/protocol"
)

var (
	//  contact variables:
	Alireza  = protocol.Contact{PersonID: 1, Cellphone: "09121234567", Description: "none", JobID: 5}
	Kianoosh = protocol.Contact{PersonID: 2, Cellphone: "0916", Description: "", JobID: 4}
	Omid     = protocol.Contact{PersonID: 3, Cellphone: "0916", Description: "", JobID: 3}
	Khorshid = protocol.Contact{PersonID: 4, Cellphone: "0912", Description: "", JobID: 2}
	Reza     = protocol.Contact{PersonID: 5, Cellphone: "0914", Description: "", JobID: 1}
	Tom      = protocol.Contact{PersonID: 6, Cellphone: "+39", Description: "", JobID: 6}
	Frank    = protocol.Contact{PersonID: 7, Cellphone: "+39", Description: "", JobID: 8}
	Delaram  = protocol.Contact{PersonID: 8, JobID: 4}

	//----------- Test Vars -----------------
	ContactDataTest = protocol.ContactStorage{
		Data: []*protocol.Contact{&Alireza, &Kianoosh, &Omid, &Khorshid, &Reza, &Tom, &Frank, &Delaram},
	}
	JobDataTest = []*protocol.Job{
		AccountManager, Driver, StoreManager, Developer, SecurityMan, OfficialEmployee, WarehouseWorker, MasterChef}

	CityDataTest = []*protocol.City{&Tehran, &Shiraz, &Tabriz, &Karaj, &Mahabad, &Noshahr, &LostHeaven, &Rasht}

	SexDataTest = []*protocol.Sex{Female, Male, Other}

	PersonDataTest = []*protocol.Person{&AlirezaPerson, &KianooshPerson, &OmidPerson, &KhorshidPerson, &RezaPerson, &TomPerson, &FrankPerson, &DelaramPerson}
	//----------- Test Vars -----------------

	// cities variable
	Tehran     = protocol.City{Name: "Tehran", AriaCode: "021"}
	Shiraz     = protocol.City{Name: "Shiraz", AriaCode: "061"}
	Tabriz     = protocol.City{Name: "Tabriz", AriaCode: "041"}
	Karaj      = protocol.City{Name: "Karaj", AriaCode: "026"}
	Mahabad    = protocol.City{Name: "Mahabad", AriaCode: "051"}
	Noshahr    = protocol.City{Name: "Noshahr", AriaCode: "011"}
	LostHeaven = protocol.City{Name: "LostHeaven", AriaCode: "+1"}
	Rasht      = protocol.City{Name: "Rasht", AriaCode: "013"}

	// Person variable
	AlirezaPerson  = protocol.Person{FirstName: "Alireza", LastName: "Mokhtari Garakani", GenderID: 2}
	KianooshPerson = protocol.Person{FirstName: "Kianoosh", LastName: "Ashayeri Zadeh", GenderID: 2}
	OmidPerson     = protocol.Person{FirstName: "Omid", LastName: "Hekayati", GenderID: 2}
	KhorshidPerson = protocol.Person{FirstName: "Khorshid", LastName: "Mokhtari Garakani", GenderID: 1}
	RezaPerson     = protocol.Person{FirstName: "Reza", LastName: "Aslejedian", GenderID: 2}
	TomPerson      = protocol.Person{FirstName: "Tommy", LastName: "Angelo", GenderID: 2}
	FrankPerson    = protocol.Person{FirstName: "Frank", LastName: "Colleti", GenderID: 2}
	DelaramPerson  = protocol.Person{FirstName: "Delaram", GenderID: 1}

	// job variable
	AccountManager   = &protocol.Job{Name: "AccountManager", LocationID: 2, Description: "", BasicPaymentPerHour: 0}
	Driver           = &protocol.Job{Name: "Driver", LocationID: 1, Description: "", BasicPaymentPerHour: 0}
	StoreManager     = &protocol.Job{Name: "StoreManager", LocationID: 2, Description: "", BasicPaymentPerHour: 0}
	Developer        = &protocol.Job{Name: "Developer", LocationID: 5, Description: "", BasicPaymentPerHour: 0}
	SecurityMan      = &protocol.Job{Name: "SecurityMan", LocationID: 4, Description: "", BasicPaymentPerHour: 0}
	OfficialEmployee = &protocol.Job{Name: "OfficialEmployee", LocationID: 6, Description: "", BasicPaymentPerHour: 0}
	WarehouseWorker  = &protocol.Job{Name: "WarehouseWorker ", LocationID: 3, Description: "", BasicPaymentPerHour: 0}
	MasterChef       = &protocol.Job{Name: "MasterChef", LocationID: 7, Description: "", BasicPaymentPerHour: 0}

	// sex variable
	Female = &protocol.Sex{Id: 1, Name: "Female"}
	Male   = &protocol.Sex{Id: 2, Name: "Male"}
	Other  = &protocol.Sex{Id: 3, Name: "Other"}
)
