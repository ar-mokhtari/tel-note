package city

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindCityByID(inputID uint) (config.ResponseStatus, protocol.City) {
	if state, data := storage.FindCityByID(inputID); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.City{}
}
