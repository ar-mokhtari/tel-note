//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"strings"
	"tel-note/SDK/neshan"
	"tel-note/lib/callAPI"
	"tel-note/protocol"
)

type storageMemory struct {
	CityData []*protocol.City
	PoolByID map[uint]*protocol.City
}

var (
	storage storageMemory
	//TODO::: What does this do?
	_ protocol.CityServices = &storage
)

func (allCity *storageMemory) GetCities() []*protocol.City {
	return allCity.CityData
}

func (allCity *storageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range allCity.CityData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (allCity *storageMemory) FindCityByID(inputID uint) (bool, protocol.City) {
	for _, data := range allCity.CityData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.City{}
}

func (allCity *storageMemory) NewCity(inputCity protocol.City) bool {
	var LastID uint
	for _, data := range allCity.CityData {
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
	allCity.CityData = append(allCity.CityData, &result)
	return true
}

func (allCity *storageMemory) EditCity(ID uint, newCity protocol.City) bool {
	for index, data := range allCity.CityData {
		if data.Id == ID {
			//TODO:: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newCity.Name != "" {
				(allCity.CityData)[index].Name = newCity.Name
			}
			if newCity.EnglishName != "" {
				(allCity.CityData)[index].EnglishName = newCity.EnglishName
			}
			if newCity.AriaCode != "" {
				(allCity.CityData)[index].AriaCode = newCity.AriaCode
			}
			if newCity.Lat != 0.0 {
				(allCity.CityData)[index].Lat = newCity.Lat
			}
			if newCity.Lng != 0.0 {
				(allCity.CityData)[index].Lng = newCity.Lng
			}
			return true
		}
	}
	return false
}

func (allCity *storageMemory) DeleteCityByID(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range allCity.CityData {
			if data.Id == id {
				allCity.CityData = append((allCity.CityData)[:index], (allCity.CityData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}

func (allCity *storageMemory) CallTimeDistanceTwoCities(cityNoOne, cityNoTwo protocol.City) ([]uint, bool) {
	MapParams := map[string]string{
		"type":         "car",
		"origins":      fmt.Sprintf("%f,%f", cityNoOne.Lat, cityNoOne.Lng),
		"destinations": fmt.Sprintf("%f,%f", cityNoTwo.Lat, cityNoTwo.Lng),
	}
	MapHeaders := map[string]string{
		"Api-Key": neshan.ApiKey,
	}
	responseData := callAPI.CallGetAPIs(
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
