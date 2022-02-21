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
	}
)
