package config

import "tel-note/internal/storage"

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
	Alireza  = storage.Contact{Person: &storage.Person{FirstName: "Alireza", LastName: "Mokhtari Garakani"}, Cellphone: "09121234567", Description: "none"}
	Kianoosh = storage.Contact{Person: &storage.Person{FirstName: "Kianoosh", LastName: "Ashayeri Zadeh"}, Cellphone: "0916", Description: ""}
	Omid     = storage.Contact{Person: &storage.Person{FirstName: "Omid", LastName: "Hekayati"}, Cellphone: "0916", Description: ""}
	Khorshid = storage.Contact{Person: &storage.Person{FirstName: "Khorshid", LastName: "Mokhtari Garakani"}, Cellphone: "0912", Description: ""}
	Reza     = storage.Contact{Person: &storage.Person{FirstName: "Reza", LastName: "Aslejedian"}, Cellphone: "0914", Description: ""}
	Tom      = storage.Contact{Person: &storage.Person{FirstName: "Tommy", LastName: "Angelo"}, Cellphone: "+39", Description: ""}
	Frank    = storage.Contact{Person: &storage.Person{FirstName: "Frank", LastName: "Colleti"}, Cellphone: "+39", Description: ""}

	MainDataTest = storage.AllData{ContactData: []*storage.Contact{&Reza, &Tom, &Frank, &Khorshid, &Omid, &Kianoosh, &Alireza}}
)
