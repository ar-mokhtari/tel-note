package config

import (
	"tel-note/internal/storage/memory"
)

type (
	Greeting struct {
		General     string `json:"general,omitempty"`
		Description string `json:"description,omitempty"`
	}
	Role           int
	ResponseStatus struct {
		State  bool
		String string
	}
)

const (
	Admin Role = iota + 10
	Manager
	Basic
	AdminString = "admin"
	OkStatus    = "yes"
)

var (
	//  contact variables:
	Alireza  = memory.Contact{Person: &memory.Person{FirstName: "Alireza", LastName: "Mokhtari Garakani"}, Cellphone: "09121234567", Description: "none"}
	Kianoosh = memory.Contact{Person: &memory.Person{FirstName: "Kianoosh", LastName: "Ashayeri Zadeh"}, Cellphone: "0916", Description: ""}
	Omid     = memory.Contact{Person: &memory.Person{FirstName: "Omid", LastName: "Hekayati"}, Cellphone: "0916", Description: ""}
	Khorshid = memory.Contact{Person: &memory.Person{FirstName: "Khorshid", LastName: "Mokhtari Garakani"}, Cellphone: "0912", Description: ""}
	Reza     = memory.Contact{Person: &memory.Person{FirstName: "Reza", LastName: "Aslejedian"}, Cellphone: "0914", Description: ""}
	Tom      = memory.Contact{Person: &memory.Person{FirstName: "Tommy", LastName: "Angelo"}, Cellphone: "+39", Description: ""}
	Frank    = memory.Contact{Person: &memory.Person{FirstName: "Frank", LastName: "Colleti"}, Cellphone: "+39", Description: ""}

	MainDataTest = memory.AllData{ContactData: []*memory.Contact{&Reza, &Tom, &Frank, &Khorshid, &Omid, &Kianoosh, &Alireza}, CityData: []*memory.City{&Tehran, &Shiraz, &Tabriz, &Karaj, &Mahabad, &Noshahr, &LostHeaven}}

	// cities variable
	Tehran     = memory.City{Name: "Tehran"}
	Shiraz     = memory.City{Name: "Shiraz"}
	Tabriz     = memory.City{Name: "Tabriz"}
	Karaj      = memory.City{Name: "Karaj"}
	Mahabad    = memory.City{Name: "Mahabad"}
	Noshahr    = memory.City{Name: "Noshahr"}
	LostHeaven = memory.City{Name: "LostHeaven"}

	//Todo: Person and Job also have to add
	/*	// persons variable

	 */ // job variable
)
