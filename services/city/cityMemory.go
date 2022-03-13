package city

import (
	"encoding/json"
	"fmt"
	"strings"
	"tel-note/SDK/neshan"
	"tel-note/protocol"
	"tel-note/services/general"
)

type storageMemory protocol.CityStorage

func (AllCity *storageMemory) GetCities() protocol.CityStorage {
	return protocol.CityStorage(*AllCity)
}

func (AllCity *storageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range AllCity.CityData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (AllCity *storageMemory) FindCityByID(inputID uint) (bool, protocol.City) {
	for _, data := range AllCity.CityData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.City{}
}

func (AllCity *storageMemory) NewCity(inputCity protocol.City) bool {
	var LastID uint
	for _, data := range AllCity.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.City{
		Id:          LastID,
		Name:        inputCity.Name,
		EnglishName: inputCity.EnglishName,
		AriaCode:    inputCity.AriaCode,
		Lat:         inputCity.Lat,
		Lng:         inputCity.Lng,
	}
	AllCity.CityData = append(AllCity.CityData, &result)
	return true
}

func (AllCity *storageMemory) EditCity(ID uint, newCity protocol.City) bool {
	for index, data := range AllCity.CityData {
		if data.Id == ID {
			//TODO:: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newCity.Name != "" {
				(AllCity.CityData)[index].Name = newCity.Name
			}
			if newCity.EnglishName != "" {
				(AllCity.CityData)[index].EnglishName = newCity.EnglishName
			}
			if newCity.AriaCode != "" {
				(AllCity.CityData)[index].AriaCode = newCity.AriaCode
			}
			if newCity.Lat != 0.0 {
				(AllCity.CityData)[index].Lat = newCity.Lat
			}
			if newCity.Lng != 0.0 {
				(AllCity.CityData)[index].Lng = newCity.Lng
			}
			return true
		}
	}
	return false
}

func (AllCity *storageMemory) DeleteCityByID(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range AllCity.CityData {
			if data.Id == id {
				AllCity.CityData = append((AllCity.CityData)[:index], (AllCity.CityData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}

func (AllCity *storageMemory) CallTimeDistanceTwoCities(cityNoOne, cityNoTwo protocol.City) ([]uint, bool) {
	MapParams := map[string]string{
		"type":         "car",
		"origins":      fmt.Sprintf("%f,%f", cityNoOne.Lat, cityNoOne.Lng),
		"destinations": fmt.Sprintf("%f,%f", cityNoTwo.Lat, cityNoTwo.Lng),
	}
	MapHeaders := map[string]string{
		"Api-Key": neshan.ApiKey,
	}
	responseData := general.CallGetAPIs(
		neshan.DistanceMatrixURL,
		MapParams,
		MapHeaders,
	)
	var AllResult neshan.DistanceMatrix
	json.Unmarshal(responseData, &AllResult)
	switch AllResult.Rows == nil {
	case true:
		return []uint{}, false
	case false:
		return []uint{
			AllResult.Rows[0].Elements[0].Duration.Value,
			AllResult.Rows[0].Elements[0].Distance.Value,
		}, true
	}
	return []uint{}, false
}
