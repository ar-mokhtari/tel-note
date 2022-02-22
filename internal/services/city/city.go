package city

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func NewCity(city protocol.City) config.ResponseStatus {
	if Storage.NewCity(city) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditCityByID(ID uint, NewCity protocol.City) config.ResponseStatus {
	if Storage.EditCity(ID, NewCity) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteCity(IDS []uint) []uint {
	return Storage.DeleteCity(IDS)
}
