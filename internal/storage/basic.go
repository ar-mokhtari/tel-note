package storage

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
	City struct {
		Id   uint
		Name string
	}
	AllCities struct {
		CityData []*City
	}
	sex struct {
		Id   uint8
		Name string
	}
	jobInfo struct {
		Id                  uint
		Name                string
		Location            *City
		Description         string
		BasicPaymentPerHour uint
	}
)
