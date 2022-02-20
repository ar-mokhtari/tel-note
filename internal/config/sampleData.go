package config

import "tel-note/internal/storage/memory"

var (
	//  contact variables:
	Alireza  = memory.Contact{Person: AlirezaPerson, Cellphone: "09121234567", Description: "none", JobInfo: &memory.JobInfo{Id: 5}}
	Kianoosh = memory.Contact{Person: KianooshPerson, Cellphone: "0916", Description: "", JobInfo: &memory.JobInfo{Id: 4}}
	Omid     = memory.Contact{Person: OmidPerson, Cellphone: "0916", Description: "", JobInfo: &memory.JobInfo{Id: 3}}
	Khorshid = memory.Contact{Person: KhorshidPerson, Cellphone: "0912", Description: "", JobInfo: &memory.JobInfo{Id: 2}}
	Reza     = memory.Contact{Person: RezaPerson, Cellphone: "0914", Description: "", JobInfo: &memory.JobInfo{Id: 1}}
	Tom      = memory.Contact{Person: TomPerson, Cellphone: "+39", Description: "", JobInfo: &memory.JobInfo{Id: 6}}
	Frank    = memory.Contact{Person: FrankPerson, Cellphone: "+39", Description: "", JobInfo: &memory.JobInfo{Id: 8}}
	Delaram  = memory.Contact{Person: DelaramPerson, JobInfo: &memory.JobInfo{Id: 4}}

	MainDataTest = memory.AllData{
		ContactData: []*memory.Contact{&Reza, &Tom, &Frank, &Khorshid, &Omid, &Kianoosh, &Alireza, &Delaram},
		CityData:    []*memory.City{&Tehran, &Shiraz, &Tabriz, &Karaj, &Mahabad, &Noshahr, &LostHeaven},
		JobData:     []*memory.JobInfo{AccountManager, Driver, StoreManager, Developer, SecurityMan, OfficialEmployee, WarehouseWorker, MasterChef},
		SexData:     []*memory.Sex{Female, Male, Other},
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
	AlirezaPerson  = &memory.Person{FirstName: "Alireza", LastName: "Mokhtari Garakani", Gender: Male}
	KianooshPerson = &memory.Person{FirstName: "Kianoosh", LastName: "Ashayeri Zadeh", Gender: Male}
	OmidPerson     = &memory.Person{FirstName: "Omid", LastName: "Hekayati", Gender: Male}
	KhorshidPerson = &memory.Person{FirstName: "Khorshid", LastName: "Mokhtari Garakani", Gender: Female}
	RezaPerson     = &memory.Person{FirstName: "Reza", LastName: "Aslejedian", Gender: Male}
	TomPerson      = &memory.Person{FirstName: "Tommy", LastName: "Angelo", Gender: Male}
	FrankPerson    = &memory.Person{FirstName: "Frank", LastName: "Colleti", Gender: Male}
	DelaramPerson  = &memory.Person{FirstName: "Delaram", Gender: Female}

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
	Female = &memory.Sex{Id: 1, Name: "Female"}
	Male   = &memory.Sex{Id: 2, Name: "Male"}
	Other  = &memory.Sex{Id: 3, Name: "Other"}
)
