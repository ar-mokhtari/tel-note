package allData

import (
	"tel-note/internal/config"
	"tel-note/internal/service/basicInfo"
	"tel-note/internal/service/contact"
)

func FillSimpleDataInMainData() {
	for _, data := range config.MainDataTest.ContactData {
		contact.NewContact(*data)
	}
	for _, data := range config.MainDataTest.CityData {
		basicInfo.NewCity(data.Name)
	}
	//Todo: Person and Job also have to add
}
