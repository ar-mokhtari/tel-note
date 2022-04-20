package protocol

import "time"

type (
	Person struct {
		Id              uint      `json:"id"`
		FirstName       string    `json:"first_name"`
		LastName        string    `json:"last_name"`
		DOB             time.Time `json:"dob"`
		BirthLocationID uint      `json:"birth_location_id"`
		GenderID        uint      `json:"gender_id"`
		NationalCode    string    `json:"national_code"`
		Description     string    `json:"description"`
		CreateAt        time.Time `json:"create_at"`
		UpdateAt        time.Time `json:"update_at"`
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
