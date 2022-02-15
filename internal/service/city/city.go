package city

import (
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

func NewCity(CityName string) config.ResponseStatus {
	if storage.AllDataTool.NewCity(&config.MainData, CityName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditCityByID(ID uint, NewCityName string) config.ResponseStatus {
	if storage.AllDataTool.EditCityByID(&config.MainData, ID, NewCityName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
