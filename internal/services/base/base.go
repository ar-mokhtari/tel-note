package base

import (
	"tel-note/internal/config"
	"tel-note/internal/storage/memory"
)

func NewSex(SexName string) config.ResponseStatus {
	if memory.AllDataTool.NewSex(&config.MainData, SexName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func EditSexNameByID(ID uint8, NewName string) config.ResponseStatus {
	if memory.AllDataTool.EditSex(&config.MainData, ID, NewName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteSexByID(ID uint8) config.ResponseStatus {
	if memory.AllDataTool.DeleteSexByID(&config.MainData, ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
