package person

import (
	"tel-note/internal/config"
)

func FindPersonByChar(inputChar string) (state config.ResponseStatus, result []uint) {
	if status, res := storage.FindPersonByChar(inputChar); status {
		return config.ResponseStatus{State: true}, res
	}
	return config.ResponseStatus{State: false}, []uint{}
}
