package fillSampleData

import (
	"tel-note/env"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
)

func FillSimpleDataInMainData() {
	for _, data := range env.SexDataTest {
		sex.NewSex(*data)
	}
	for _, data := range env.CityDataTest {
		city.NewCity(*data)
	}
	for _, data := range env.JobDataTest {
		job.NewJob(*data)
	}
	for _, data := range env.PersonDataTest {
		person.NewPerson(*data)
	}
	for _, data := range env.CustomerGroup {
		customer.NewGroup(data.GroupName)
	}
	for _, data := range env.CustomerRelation {
		customer.NewRelation(data.CustomerID, data.GroupID)
	}
	for _, data := range env.CustomersDataTest.Data {
		customer.AddCustomer(*data)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range (env.ContactDataTest).Data {
		contact.NewContact(*data)
	}
	//call api test for fill countries
	country.CallCountry()
}
