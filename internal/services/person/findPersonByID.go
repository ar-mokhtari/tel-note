package person

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindPersonByID(ID uint) (state config.ResponseStatus, result protocol.Person) {
	if status, res := storage.FindPersonByID(ID); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, protocol.Person{}
}
