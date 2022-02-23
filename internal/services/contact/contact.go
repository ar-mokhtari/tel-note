package contact

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
	"tel-note/internal/services/globalVars"
	"unsafe"
)

func NewContact(inputContact protocol.Contact) (status config.ResponseStatus, resData protocol.ContactStorage) {
	if status.State, globalVars.AllContact = storage.AddContact(inputContact); status.State {
		return config.ResponseStatus{State: true}, resData
	}
	return config.ResponseStatus{State: false}, protocol.ContactStorage{}

}

func FindContactByID(id uint) (config.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return config.ResponseStatus{State: true}, data
	}
	return config.ResponseStatus{State: false}, protocol.Contact{}
}

func FindContactByChar(insertChar string) (protocol.ContactStorage, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return data, uint(unsafe.Sizeof(data))
	}
	return protocol.ContactStorage{}, 0
}

func EditContactByID(newData protocol.Contact, ID uint) *config.ResponseStatus {
	if storage.EditContactByID(newData, ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: true, String: "not found"}
}

func DeleteContactByID(ID uint) *config.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}

func DeleteAll() *config.ResponseStatus {
	if storage.DeleteAll() {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}
