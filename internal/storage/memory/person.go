package memory

import "time"

type (
	Person struct {
		Id            uint   `json:"id"`
		FirstName     string `json:"first_name,omitempty"`
		LastName      string `json:"last_name,omitempty"`
		DOB           time.Time
		BirthLocation *City
		Gender        *sex
		NationalCode  string
		Description   string `json:"description,omitempty"`
	}
	AllPerson struct {
		PersonData *[]Person
	}
)
