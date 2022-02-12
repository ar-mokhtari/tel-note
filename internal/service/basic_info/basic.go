package basic_info

import (
	"fmt"
	"github.com/fatih/structs"
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

// NewCity CRUD basic info
func NewCity(MainData *storage.AllCities, CityName string) {

	names := structs.Names(&storage.Contact{})
	fmt.Println(names) // ["Foo", "Bar"]

	var LastID uint
	for _, data := range MainData.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := storage.City{
		Id:   uint(LastID) + 1,
		Name: CityName,
	}
	MainData.CityData = append(MainData.CityData, &result)
}

func EditCityByID(MainCity *storage.AllCities, ID uint, NewCityName string) config.ResponseStatus {
	for _, data := range MainCity.CityData {
		if data.Id == ID {
			data.Name = NewCityName
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{State: false}
}
