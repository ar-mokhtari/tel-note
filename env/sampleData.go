package env

import (
	"github.com/ar-mokhtari/tel-note/protocol"
	"time"
)

type cityStruct struct {
	city
}
type city struct {
	id          uint
	name        string
	englishName string
	ariaCode    string
	lat         float64
	lng         float64
}

func (c city) ID() uint            { return c.id }
func (c city) Name() string        { return c.name }
func (c city) EnglishName() string { return c.englishName }
func (c city) AriaCode() string    { return c.ariaCode }
func (c city) Lat() float64        { return c.lat }
func (c city) Lng() float64        { return c.lng }

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

	CityDataTest = []protocol.City{
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
	Tehran     = city{name: "Tehran", englishName: "THR", ariaCode: "021", lat: 0, lng: 0}
	Shiraz     = city{name: "Shiraz", ariaCode: "061", lat: 29.6655016, lng: 52.3929325}
	Tabriz     = city{name: "Tabriz", ariaCode: "041", lat: 38.0802852, lng: 46.1536431}
	Karaj      = city{name: "Karaj", ariaCode: "026", lat: 35.8109689, lng: 50.8772975}
	Mahabad    = city{name: "Mahabad", ariaCode: "051", lat: 36.7659726, lng: 45.6948439}
	Noshahr    = city{name: "Noshahr", ariaCode: "011", lat: 36.6502418, lng: 51.4693375}
	LostHeaven = city{name: "LostHeaven", ariaCode: "+1"}
	Rasht      = city{name: "Rasht", ariaCode: "013", lat: 37.244096, lng: 49.5163231}
	Abadan     = city{name: "Abadan", ariaCode: "0631", lat: 30.3902492, lng: 48.1193191}
	Kish       = city{name: "Kish", ariaCode: "0764442", lat: 26.5360317, lng: 53.9043458}

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
