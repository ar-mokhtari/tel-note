package fillSampleData

import (
	"tel-note/internal/config"
	"tel-note/internal/services/city"
	"tel-note/internal/services/contact"
	"tel-note/internal/services/job"
	"tel-note/internal/services/person"
	"tel-note/internal/services/sex"
)

func FillSimpleDataInMainData() {
	for _, data := range config.SexDataTest {
		sex.NewSex(data.Name)
	}
	for _, data := range config.PersonDataTest {
		person.NewPerson(*data)
	}
	for _, data := range config.CityDataTest {
		city.NewCity(*data)
	}
	for _, data := range config.JobDataTest {
		job.NewJob(*data)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range (config.ContactDataTest).Data {
		contact.NewContact(*data)
	}
}
