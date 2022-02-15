package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

func NewContact(inputContact storage.Contact) (config.ResponseStatus, storage.AllData) {
	var resData storage.AllData
	if resStatus, resData := (*storage.AllData).AddContact(&config.MainData, inputContact); resStatus == true {
		return config.ResponseStatus{State: true}, resData
	}
	return config.ResponseStatus{State: false}, resData
}

func FindContactByID(id uint) (storage.Contact, bool) {
	if data, state := (*storage.AllData).FindContactByID(&config.MainData, id); state == true {
		return data, true
	}
	return storage.Contact{}, false
}

func FindContactByChar(insertChar string) (storage.AllData, uint) {
	if resultData, counter := (*storage.AllData).FindContactByChar(&config.MainData, insertChar); counter > 1 {
		return resultData, counter
	}
	return storage.AllData{}, 0
}

func EditContactByID(newData storage.Contact, ID uint) config.ResponseStatus {
	if (*storage.AllData).EditContactByID(&config.MainData, newData, ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: true, String: "not found"}
}

func DeleteContactByID(ID uint) config.ResponseStatus {
	if (*storage.AllData).DeleteContactByID(&config.MainData, ID) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteAll() config.ResponseStatus {
	if (*storage.AllData).DeleteAll(&config.MainData) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}
