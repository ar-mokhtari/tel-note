package basic_info

import (
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

// NewCity CRUD basic info
func NewCity(CityName string) {
	var LastID uint
	for _, data := range (config.MainData).CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := storage.City{
		Id:   uint(LastID) + 1,
		Name: CityName,
	}
	(config.MainData).CityData = append((config.MainData).CityData, &result)
}

func EditCityByID(ID uint, NewCityName string) config.ResponseStatus {
	for _, data := range (config.MainData).CityData {
		if data.Id == ID {
			data.Name = NewCityName
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{State: false}
}
