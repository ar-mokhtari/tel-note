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
	}
	PersonStorage struct {
		PersonData *[]Person
	}
	PersonServices interface {
		FindPersonByChar(inputChar string) (bool, []uint)
		FindPersonByID(inputID uint) (status bool, res Person)
		NewPerson(person Person) bool
		EditPerson(ID uint, newPerson Person) bool
		DeletePerson(IDS []uint) (resDel []uint)
	}
)
