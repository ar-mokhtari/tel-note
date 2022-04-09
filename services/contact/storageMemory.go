package contact

import (
	"tel-note/protocol"
	"tel-note/services/person"
)

type (
	storageMemory struct {
		ContactData []*protocol.Contact
	}
)

var (
	storage storageMemory
	_       protocol.ContactServices = &storage
)

func (allContact *storageMemory) GetContacts() []*protocol.Contact {
	return allContact.ContactData
}

func (allContact *storageMemory) AddContact(inputContact protocol.Contact) bool {
	var (
		lastID uint
	)
	for _, data := range allContact.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	result := protocol.Contact{
		Id:                  lastID,
		PersonID:            inputContact.PersonID,
		JobID:               inputContact.JobID,
		Tel:                 inputContact.Tel,
		CellphoneCollection: inputContact.CellphoneCollection,
		Address:             inputContact.Address,
		Description:         inputContact.Description,
	}
	allContact.ContactData = append(allContact.ContactData, &result)
	return true
}

func (allContact *storageMemory) FindContactByID(id uint) (bool, protocol.Contact) {
	for _, data := range allContact.ContactData {
		if data.Id == id {
			return true, *data
		}
	}
	return false, protocol.Contact{}
}

func (allContact *storageMemory) FindContactByChar(insertChar string) (status bool, result []*protocol.Contact) {
	if state, res := person.FindPersonByChar(insertChar); state.State {
		for _, data := range allContact.ContactData {
			if len(res) == 0 {
				break
			}
			if ((res)[0]).Id == data.PersonID {
				status = true
				result = append(result, data)
				res = (res)[1:]
			}
		}
	}
	return status, result
}

func (allContact *storageMemory) EditContactByID(newData protocol.Contact, ID uint) bool {
	for index, data := range allContact.ContactData {
		if data.Id == ID {
			if newData.PersonID != 0 {
				(allContact.ContactData)[index].PersonID = newData.PersonID
			}
			if newData.JobID != 0 {
				(allContact.ContactData)[index].JobID = newData.JobID
			}
			if newData.Tel != "" {
				(allContact.ContactData)[index].Tel = newData.Tel
			}
			if newData.CellphoneCollection != nil {
				(allContact.ContactData)[index].CellphoneCollection = newData.CellphoneCollection
			}
			if newData.Address != "" {
				(allContact.ContactData)[index].Address = newData.Address
			}
			if newData.Description != "" {
				(allContact.ContactData)[index].Description = newData.Description
			}
			return true
		}
	}
	return false
}

func (allContact *storageMemory) DeleteContactByID(ID uint) bool {
	for index, data := range allContact.ContactData {
		if data.Id == ID {
			allContact.ContactData = append((allContact.ContactData)[:index], (allContact.ContactData)[index+1:]...)
			//TODO: BadSolution,HandelMultiDeleteInServiceOrStorageNotInClient
			return true
		}
	}
	return false
}

func (allContact *storageMemory) DeleteAll() bool {
	allContact.ContactData = (allContact.ContactData)[0:0]
	return true
}
