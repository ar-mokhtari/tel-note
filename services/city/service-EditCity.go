package city

import (
	"tel-note/protocol"
)

func EditCityByID(ID uint, NewCity protocol.City) protocol.ResponseStatus {
	if storage.EditCity(ID, NewCity) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
