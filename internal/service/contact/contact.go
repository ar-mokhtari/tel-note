package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

func NewContact(inputContact storage.Contact) (config.ResponseStatus, storage.AllData) {
	var resData storage.AllData
	if resStatus, resData := storage.AllDataTool.AddContact(&config.MainData, inputContact); resStatus {
		return config.ResponseStatus{State: true}, resData
	}
	return config.ResponseStatus{State: false}, resData
}

func FindContactByID(id uint) (storage.Contact, bool) {
	if data, state := storage.AllDataTool.FindContactByID(&config.MainData, id); state == true {
		return data, true
	}
	return storage.Contact{}, false
}

func FindContactByChar(insertChar string) (storage.AllData, uint) {
	if resultData, counter := storage.AllDataTool.FindContactByChar(&config.MainData, insertChar); counter > 1 {
		return resultData, counter
	}
	return storage.AllData{}, 0
}

func EditContactByID(newData storage.Contact, ID uint) config.ResponseStatus {
	if storage.AllDataTool.EditContactByID(&config.MainData, newData, ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: true, String: "not found"}
}

func DeleteContactByID(ID uint) config.ResponseStatus {
	if storage.AllDataTool.DeleteContactByID(&config.MainData, ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteAll() config.ResponseStatus {
	if storage.AllDataTool.DeleteAll(&config.MainData) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
