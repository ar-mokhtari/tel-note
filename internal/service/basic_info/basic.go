package basic_info

import (
	"tel-note/internal/storage"
)

// NewCity CRUD basic info
func NewCity(MainData *storage.AllCities, CityName string) {
	var LastID uint
	for _, data := range MainData.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := storage.City{
		Id:   uint(LastID),
		Name: CityName,
	}
	MainData.CityData = append(MainData.CityData, &result)
}
