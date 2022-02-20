package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/storage/memory"
)

func NewContact(inputContact memory.Contact) (*config.ResponseStatus, memory.AllData) {
	storage.AddContact(inputContact)
	var resData memory.AllData
	if resStatus, resData := memory.AllDataTool.AddContact(&config.MainData, inputContact); resStatus {
		return &config.ResponseStatus{State: true}, resData
	}
	return &config.ResponseStatus{State: false}, resData
}

func FindContactByID(id uint) (*memory.Contact, bool) {
	if data, state := memory.AllDataTool.FindContactByID(&config.MainData, id); state == true {
		return &data, true
	}
	return &memory.Contact{}, false
}

func FindContactByChar(insertChar string) (*memory.AllData, uint) {
	if resultData, counter := memory.AllDataTool.FindContactByChar(&config.MainData, insertChar); counter > 1 {
		return &resultData, counter
	}
	return &memory.AllData{}, 0
}

func EditContactByID(newData memory.Contact, ID uint) *config.ResponseStatus {
	if memory.AllDataTool.EditContactByID(&config.MainData, newData, ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: true, String: "not found"}
}

func DeleteContactByID(ID uint) *config.ResponseStatus {
	if memory.AllDataTool.DeleteContactByID(&config.MainData, ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}

func DeleteAll() *config.ResponseStatus {
	if memory.AllDataTool.DeleteAll(&config.MainData) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}
