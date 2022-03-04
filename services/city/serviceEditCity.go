package city

import (
	"tel-note/protocol"
)

func EditCityByID(ID uint, NewCity protocol.City) protocol.ResponseStatus {
	if stare, _ := storage.EditCity(ID, NewCity); stare {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
