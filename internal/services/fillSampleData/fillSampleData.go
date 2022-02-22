package fillSampleData

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/city"
	"tel-note/internal/services/contact"
	"tel-note/internal/services/job"
	"tel-note/internal/services/sex"
)

func FillSimpleDataInMainData() {
	for _, data := range config.SexDataTest {
		sex.NewSex(data.Name)
	}
	for _, data := range config.CityDataTest {
		city.NewCity(protocol.City{
			Name:     data.Name,
			AriaCode: data.AriaCode,
		})
	}
	for _, data := range config.JobDataTest {
		job.NewJob(protocol.Job{
			Name:                data.Name,
			LocationID:          data.LocationID,
			Description:         data.Description,
			BasicPaymentPerHour: data.BasicPaymentPerHour,
		})
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range (config.ContactDataTest).Data {
		contact.NewContact(*data)
	}
}
