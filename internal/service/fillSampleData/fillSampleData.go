package fillSampleData

import (
	"tel-note/internal/config"
	"tel-note/internal/service/base"
	"tel-note/internal/service/city"
	"tel-note/internal/service/contact"
	"tel-note/internal/service/jobInfo"
)

func FillSimpleDataInMainData() {
	for _, data := range config.MainDataTest.SexData {
		base.NewSex(data.Name)
	}
	for _, data := range config.MainDataTest.CityData {
		city.NewCity(data.Name)
	}
	for _, data := range config.MainDataTest.JobData {
		jobInfo.NewJob(data.Name)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range config.MainDataTest.ContactData {
		contact.NewContact(*data)
	}
}
