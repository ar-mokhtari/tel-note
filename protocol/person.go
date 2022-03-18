package protocol

import "time"

type (
	Person struct {
		Id              uint
		FirstName       string
		LastName        string
		DOB             time.Time
		BirthLocationID uint
		GenderID        uint
		NationalCode    string
		Description     string
		CreateAt        time.Time
		UpdateAt        time.Time
	}
	PersonServices interface {
		GetPersons() []*Person
		FindPersonByChar(inputChar string) (bool, []*Person)
		FindPersonByID(inputID uint) (status bool, res Person)
		NewPerson(person Person) bool
		EditPerson(ID uint, newPerson Person) bool
		DeletePerson(IDS []uint) (resDel []uint)
	}
)
