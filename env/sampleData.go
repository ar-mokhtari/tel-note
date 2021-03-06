package env

import (
	"github.com/ar-mokhtari/tel-note/protocol"
	"time"
)

var (
	//  contact variables:
	Alireza  = protocol.Contact{PersonID: 1, CellphoneCollection: []protocol.CellPhone{{CellPhone: "09121234567", Description: ""}, {CellPhone: "09121", Description: "Home"}}, Description: "none", JobID: 5}
	Kianoosh = protocol.Contact{PersonID: 2, CellphoneCollection: []protocol.CellPhone{{CellPhone: "0916", Description: "manager"}, {CellPhone: "09161", Description: "office"}}, Description: "", JobID: 4}
	Omid     = protocol.Contact{PersonID: 3, CellphoneCollection: []protocol.CellPhone{{CellPhone: "0916", Description: ""}}, Description: "", JobID: 3}
	Khorshid = protocol.Contact{PersonID: 4, CellphoneCollection: []protocol.CellPhone{{CellPhone: "0912", Description: ""}}, Description: "", JobID: 2}
	Reza     = protocol.Contact{PersonID: 5, CellphoneCollection: []protocol.CellPhone{{CellPhone: "0914", Description: ""}}, Description: "", JobID: 1}
	Tom      = protocol.Contact{PersonID: 6, CellphoneCollection: []protocol.CellPhone{{CellPhone: "+1", Description: ""}}, Description: "", JobID: 6}
	Frank    = protocol.Contact{PersonID: 7, CellphoneCollection: []protocol.CellPhone{{CellPhone: "+1", Description: ""}}, Description: "", JobID: 8}
	Delaram  = protocol.Contact{PersonID: 8, JobID: 4}

	//----------- Test Vars -----------------
	ContactDataTest = []*protocol.Contact{
		&Alireza,
		&Kianoosh,
		&Omid,
		&Khorshid,
		&Reza,
		&Tom,
		&Frank,
		&Delaram,
	}
	JobDataTest = []*protocol.Job{
		AccountManager,
		Driver,
		StoreManager,
		Developer,
		SecurityMan,
		OfficialEmployee,
		WarehouseWorker,
		MasterChef,
	}

	CityDataTest = []*protocol.City{
		&Tehran,
		&Shiraz,
		&Tabriz,
		&Karaj,
		&Mahabad,
		&Noshahr,
		&LostHeaven,
		&Rasht,
		&Abadan,
		&Kish,
	}

	SexDataTest = []*protocol.Sex{
		FemaleGender,
		MaleGender,
		Other,
	}

	PersonDataTest = []*protocol.Person{
		&AlirezaPerson,
		&KianooshPerson,
		&OmidPerson,
		&KhorshidPerson,
		&RezaPerson,
		&TomPerson,
		&FrankPerson,
		&DelaramPerson,
	}

	CustomersDataTest = protocol.CustomerStorage{CustomerData: map[uint]*protocol.Customer{
		0: &AlirezaCust,
		1: &KianooshCust,
		2: &KhorshidCust,
		3: &RezaCust,
		4: &FrankCust,
	}}

	//----------- Test Vars -----------------

	// cities variable
	//Tehran     = {1, "", ",", 0.0, 0.0}
	Shiraz     = protocol.City{Name: "Shiraz", AriaCode: "061", Lat: 29.6655016, Lng: 52.3929325}
	Tabriz     = protocol.City{Name: "Tabriz", AriaCode: "041", Lat: 38.0802852, Lng: 46.1536431}
	Karaj      = protocol.City{Name: "Karaj", AriaCode: "026", Lat: 35.8109689, Lng: 50.8772975}
	Mahabad    = protocol.City{Name: "Mahabad", AriaCode: "051", Lat: 36.7659726, Lng: 45.6948439}
	Noshahr    = protocol.City{Name: "Noshahr", AriaCode: "011", Lat: 36.6502418, Lng: 51.4693375}
	LostHeaven = protocol.City{Name: "LostHeaven", AriaCode: "+1"}
	Rasht      = protocol.City{Name: "Rasht", AriaCode: "013", Lat: 37.244096, Lng: 49.5163231}
	Abadan     = protocol.City{Name: "Abadan", AriaCode: "0631", Lat: 30.3902492, Lng: 48.1193191}
	Kish       = protocol.City{Name: "Kish", AriaCode: "0764442", Lat: 26.5360317, Lng: 53.9043458}

	// Person variable
	AlirezaPerson = protocol.Person{
		FirstName: "Alireza",
		LastName:  "Mokhtari Garakani",
		GenderID:  2, BirthLocationID: 5, DOB: time.Unix(295345600, 0)}
	KianooshPerson = protocol.Person{
		FirstName: "Kianoosh",
		LastName:  "Ashayeri Zadeh",
		GenderID:  2, BirthLocationID: 2, DOB: time.Unix(922000000, 0)}
	OmidPerson = protocol.Person{
		FirstName: "Omid",
		LastName:  "Hekayati",
		GenderID:  2, BirthLocationID: 1, DOB: time.Unix(332000000, 0)}
	KhorshidPerson = protocol.Person{
		FirstName: "Khorshid",
		LastName:  "Mokhtari Garakani",
		GenderID:  1, BirthLocationID: 7, DOB: time.Unix(1430592000, 0)}
	RezaPerson = protocol.Person{
		FirstName: "Reza",
		LastName:  "Aslejedian",
		GenderID:  2, BirthLocationID: 8, DOB: time.Unix(995000000, 0)}
	TomPerson = protocol.Person{
		FirstName: "Tommy",
		LastName:  "Angelo",
		GenderID:  2, BirthLocationID: 4, DOB: time.Unix(230000000, 0)}
	FrankPerson = protocol.Person{
		FirstName: "Frank",
		LastName:  "Colleti",
		GenderID:  2, BirthLocationID: 3, DOB: time.Unix(180000000, 0)}
	DelaramPerson = protocol.Person{
		FirstName: "Delaram",
		GenderID:  1, BirthLocationID: 1}

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
	FemaleGender = &protocol.Sex{Id: 1, Name: "Female"}
	MaleGender   = &protocol.Sex{Id: 2, Name: "Male"}
	Other        = &protocol.Sex{Id: 3, Name: "Other"}

	// customer variable
	AlirezaCust  = protocol.Customer{PersonID: 1, Description: "first"}
	KianooshCust = protocol.Customer{PersonID: 2, Description: "second"}
	OmidCust     = protocol.Customer{PersonID: 3, Description: "important"}
	KhorshidCust = protocol.Customer{PersonID: 4, Description: "power"}
	RezaCust     = protocol.Customer{PersonID: 5, Description: "none"}
	FrankCust    = protocol.Customer{PersonID: 7, Description: "---"}

	CustomerGroupExm = protocol.CustomerGroupStorage{
		&protocol.CustomerGroup{GroupID: 1, GroupName: "Real Customers"},
		&protocol.CustomerGroup{GroupID: 2, GroupName: "Legal Customers"},
	}

	CustomerRelation = protocol.CustomerGRelationStorage{
		&protocol.CustomerGroupRelation{CustomerID: 1, GroupID: 2},
		&protocol.CustomerGroupRelation{CustomerID: 3, GroupID: 2},
		&protocol.CustomerGroupRelation{CustomerID: 5, GroupID: 2},
		&protocol.CustomerGroupRelation{CustomerID: 2, GroupID: 1},
		&protocol.CustomerGroupRelation{CustomerID: 4, GroupID: 1},
	}
)
