package config

import (
	"tel-note/internal/protocol"
	"tel-note/internal/storage/memory"
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

	ContactDataTest = protocol.ContactStorage{
		Data: []*protocol.Contact{&Alireza, &Kianoosh, &Omid, &Khorshid, &Reza, &Tom, &Frank, &Delaram},
	}
	MainDataTest = memory.AllData{
		CityData: []*memory.City{&Tehran, &Shiraz, &Tabriz, &Karaj, &Mahabad, &Noshahr, &LostHeaven},
		JobData:  []*memory.JobInfo{AccountManager, Driver, StoreManager, Developer, SecurityMan, OfficialEmployee, WarehouseWorker, MasterChef},
		SexData:  []*memory.Sex{Female, Male, Other},
	}

	// cities variable
	Tehran     = memory.City{Name: "Tehran", AriaCode: "021"}
	Shiraz     = memory.City{Name: "Shiraz", AriaCode: "061"}
	Tabriz     = memory.City{Name: "Tabriz", AriaCode: "041"}
	Karaj      = memory.City{Name: "Karaj", AriaCode: "026"}
	Mahabad    = memory.City{Name: "Mahabad", AriaCode: "051"}
	Noshahr    = memory.City{Name: "Noshahr", AriaCode: "011"}
	LostHeaven = memory.City{Name: "LostHeaven", AriaCode: "+1"}
	Rasht      = memory.City{Name: "Rasht", AriaCode: "013"}

	// Person variable
	AlirezaPerson  = protocol.Person{FirstName: "Alireza", LastName: "Mokhtari Garakani", GenderID: 0}
	KianooshPerson = protocol.Person{FirstName: "Kianoosh", LastName: "Ashayeri Zadeh", GenderID: 0}
	OmidPerson     = protocol.Person{FirstName: "Omid", LastName: "Hekayati", GenderID: 0}
	KhorshidPerson = protocol.Person{FirstName: "Khorshid", LastName: "Mokhtari Garakani", GenderID: 1}
	RezaPerson     = protocol.Person{FirstName: "Reza", LastName: "Aslejedian", GenderID: 0}
	TomPerson      = protocol.Person{FirstName: "Tommy", LastName: "Angelo", GenderID: 0}
	FrankPerson    = protocol.Person{FirstName: "Frank", LastName: "Colleti", GenderID: 0}
	DelaramPerson  = protocol.Person{FirstName: "Delaram", GenderID: 1}

	// job variable
	AccountManager   = &memory.JobInfo{Name: "AccountManager", Location: &LostHeaven, Description: "", BasicPaymentPerHour: 0}
	Driver           = &memory.JobInfo{Name: "Driver", Location: &LostHeaven, Description: "", BasicPaymentPerHour: 0}
	StoreManager     = &memory.JobInfo{Name: "StoreManager", Location: &Tabriz, Description: "", BasicPaymentPerHour: 0}
	Developer        = &memory.JobInfo{Name: "Developer", Location: &Karaj, Description: "", BasicPaymentPerHour: 0}
	SecurityMan      = &memory.JobInfo{Name: "SecurityMan", Location: &Mahabad, Description: "", BasicPaymentPerHour: 0}
	OfficialEmployee = &memory.JobInfo{Name: "OfficialEmployee", Location: &Shiraz, Description: "", BasicPaymentPerHour: 0}
	WarehouseWorker  = &memory.JobInfo{Name: "WarehouseWorker ", Location: &Noshahr, Description: "", BasicPaymentPerHour: 0}
	MasterChef       = &memory.JobInfo{Name: "MasterChef", Location: &Rasht, Description: "", BasicPaymentPerHour: 0}

	// sex variable
	Male   = &memory.Sex{Id: 1, Name: "Male"}
	Female = &memory.Sex{Id: 2, Name: "Female"}
	Other  = &memory.Sex{Id: 3, Name: "Other"}
)
