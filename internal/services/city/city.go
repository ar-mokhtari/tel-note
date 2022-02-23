package city

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
)

func NewCity(city protocol.City) (status config.ResponseStatus) {
	if status.State, globalVars.AllCity = storage.NewCity(city); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditCityByID(ID uint, NewCity protocol.City) config.ResponseStatus {
	if storage.EditCity(ID, NewCity) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteCity(IDS []uint) []uint {
	return storage.DeleteCity(IDS)
}

func FindCityByID(inputID uint) (config.ResponseStatus, protocol.City) {
	if state, data := storage.FindCityByID(inputID); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.City{}
}
