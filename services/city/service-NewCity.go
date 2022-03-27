//go:build memory

package city

import (
	"tel-note/protocol"
)

func NewCity(city protocol.City) (status protocol.ResponseStatus) {
	if status.State = storage.NewCity(city); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}

}
