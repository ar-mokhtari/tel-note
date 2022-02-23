package sex

import (
	"tel-note/internal/config"
	"tel-note/internal/services/globalVars"
)

func NewSex(SexName string) (status config.ResponseStatus) {
	if status.State, globalVars.AllSex = storage.NewSex(SexName); status.State {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
