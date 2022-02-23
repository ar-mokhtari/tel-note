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
