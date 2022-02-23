package city

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func EditCityByID(ID uint, NewCity protocol.City) config.ResponseStatus {
	if storage.EditCity(ID, NewCity) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
