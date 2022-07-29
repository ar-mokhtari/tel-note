//go:build !mysql || !mongodb

package city

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ar-mokhtari/tel-note/lib/callAPI"
	"github.com/ar-mokhtari/tel-note/protocol"
	"github.com/ar-mokhtari/tel-note/sdk/neshan"
	"strings"
)

type storageMemory struct {
	CityData []DTO
	//PoolByID map[uint]*protocol.City
}

var (
	storage storageMemory
	_       protocol.CityStorageServices = &storage
)

func (sm *storageMemory) GetCities() []protocol.City {
	return sm.GetCities()
}

func (sm *storageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range sm.CityData {
		if strings.Contains(data.NameF, inputChar) {
			res = append(res, data.IDF)
			status = true
		}
	}
	return status, res
}

func (sm *storageMemory) GetCity(ID uint) (cityIndex int, city protocol.City) {
	for index, data := range sm.CityData {
		if data.IDF == ID {
			return index, &data
		}
	}
	return -1, nil
}

func (sm *storageMemory) NewCity(inputCity protocol.City) (err error) {
	var LastID uint
	for _, data := range sm.CityData {
		if data.IDF > LastID {
			LastID = data.IDF
		}
	}
	LastID += 1
	result := DTO{
		IDF:          LastID,
		NameF:        inputCity.Name(),
		EnglishNameF: inputCity.EnglishName(),
		AriaCodeF:    inputCity.AriaCode(),
		LatF:         inputCity.Lat(),
		LngF:         inputCity.Lng(),
	}
	sm.CityData = append(sm.CityData, result)
	return nil
}

func (sm *storageMemory) EditCity(newCity protocol.City) error {
	index, city := sm.GetCity(newCity.ID())
	// check for found city
	if city == nil {
		return errors.New("not found")
	}
	// try to edit
	if newCity.Name() != "" {
		sm.CityData[index].NameF = newCity.Name()
	}
	if newCity.EnglishName() != "" {
		(sm.CityData)[index].EnglishNameF = newCity.EnglishName()
	}
	if newCity.AriaCode() != "" {
		(sm.CityData)[index].AriaCodeF = newCity.AriaCode()
	}
	if newCity.Lat() != 0.0 {
		(sm.CityData)[index].LatF = newCity.Lat()
	}
	if newCity.Lng() != 0.0 {
		(sm.CityData)[index].LngF = newCity.Lng()
	}
	return nil
}

func (sm *storageMemory) DeleteCityByID(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range sm.CityData {
			if data.IDF == id {
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
