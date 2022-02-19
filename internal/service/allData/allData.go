package allData

import (
	"tel-note/internal/config"
	"tel-note/internal/service/city"
	"tel-note/internal/service/contact"
	"tel-note/internal/service/jobInfo"
)

func FillSimpleDataInMainData() {
	for _, data := range config.MainDataTest.ContactData {
		contact.NewContact(*data)
	}
	for _, data := range config.MainDataTest.CityData {
		city.NewCity(data.Name)
	}
	for _, data := range config.MainDataTest.JobData {
		jobInfo.NewJob(data.Name)
	}

}
