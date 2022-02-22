package city

import (
	"strings"
	"tel-note/internal/protocol"
)

type StorageMemory protocol.CityStorage

func (AllCity *StorageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range AllCity.CityData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (AllCity *StorageMemory) FindCityByID(inputID uint) (bool, protocol.City) {
	for _, data := range AllCity.CityData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.City{}
}

func (AllCity *StorageMemory) NewCity(inputCity protocol.City) bool {
	var LastID uint
	for _, data := range AllCity.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID = +1
	result := protocol.City{
		Id:       LastID,
		Name:     inputCity.Name,
		AriaCode: inputCity.AriaCode,
	}
	AllCity.CityData = append(AllCity.CityData, &result)
	return true
}

func (AllCity *StorageMemory) EditCity(ID uint, newCity protocol.City) bool {
	for index, data := range AllCity.CityData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newCity.Name != "" {
				(AllCity.CityData)[index].Name = newCity.Name
			}
			if newCity.Name != "" {
				(AllCity.CityData)[index].Name = newCity.Name
			}
			if newCity.AriaCode != "" {
				(AllCity.CityData)[index].AriaCode = newCity.AriaCode
			}
			return true
		}
	}
	return false
}

func (AllCity *StorageMemory) DeleteCity(IDS []uint) (resDel []uint) {
	for index, id := range IDS {
		for _, data := range AllCity.CityData {
			if data.Id == id {
				AllCity.CityData = append((AllCity.CityData)[:index], (AllCity.CityData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
