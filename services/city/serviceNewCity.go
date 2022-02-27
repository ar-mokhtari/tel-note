package city

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
)

func NewCity(city protocol.City) (status protocol.ResponseStatus) {
	if status.State, globalVars.CityStore = storage.NewCity(city); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}

}
