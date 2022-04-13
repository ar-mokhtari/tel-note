//go:build memory

package city

import (
	"encoding/json"
	"errors"
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
	_       protocol.CityServices = &storage
)

func (sm *storageMemory) GetCities() []*protocol.City {
	return sm.CityData
}

func (sm *storageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range sm.CityData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (sm *storageMemory) FindCityByID(inputID uint) (bool, protocol.City) {
	for _, data := range sm.CityData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.City{}
}

func (sm *storageMemory) NewCity(inputCity protocol.City) bool {
	var LastID uint
	for _, data := range sm.CityData {
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
	sm.CityData = append(sm.CityData, &result)
	return true
}

func (sm *storageMemory) EditCity(newCity protocol.City) (err error) {
	var index, exitingCity = sm.GetCity(newCity.Id)
	// check for found city
	if (protocol.City{}) == exitingCity {
		return errors.New("not found")
	}
	// try to edit
	if newCity.Name != "" {
		(sm.CityData)[index].Name = newCity.Name
	}
	if newCity.EnglishName != "" {
		(sm.CityData)[index].EnglishName = newCity.EnglishName
	}
	if newCity.AriaCode != "" {
		(sm.CityData)[index].AriaCode = newCity.AriaCode
	}
	if newCity.Lat != 0.0 {
		(sm.CityData)[index].Lat = newCity.Lat
	}
	if newCity.Lng != 0.0 {
		(sm.CityData)[index].Lng = newCity.Lng
	}
	return nil
}

func (sm *storageMemory) GetCity(ID uint) (cityIndex int, city protocol.City) {
	for index, data := range sm.CityData {
		if data.Id == ID {
			return index, *data
		}
	}
	return -1, protocol.City{}
}

func (sm *storageMemory) DeleteCityByID(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range sm.CityData {
			if data.Id == id {
				sm.CityData = append((sm.CityData)[:index], (sm.CityData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}

func (sm *storageMemory) CallTimeDistanceTwoCities(cityNoOne, cityNoTwo protocol.City) ([]uint, bool) {
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
