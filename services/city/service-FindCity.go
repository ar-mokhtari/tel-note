package city

import (
	"tel-note/protocol"
)

func FindCityByID(inputID uint) (protocol.ResponseStatus, protocol.City) {
	if state, data := storage.FindCityByID(inputID); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.City{}
}
