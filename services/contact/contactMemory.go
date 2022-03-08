package contact

import (
	"tel-note/protocol"
	"tel-note/services/person"
)

type (
	storageMemory protocol.ContactStorage
)

func (AllContact *storageMemory) GetContacts() protocol.ContactStorage {
	return protocol.ContactStorage(*AllContact)
}

func (AllContact *storageMemory) AddContact(inputContact protocol.Contact) bool {
	var (
		lastID uint
	)
	for _, data := range AllContact.Data {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	result := protocol.Contact{
		Id:          lastID,
		PersonID:    inputContact.PersonID,
		Tel:         inputContact.Tel,
		Cellphone:   inputContact.Cellphone,
		Description: inputContact.Description,
		JobID:       inputContact.JobID,
	}
	AllContact.Data = append(AllContact.Data, &result)
	return true
}

func (AllContact *storageMemory) FindContactByID(id uint) (bool, protocol.Contact) {
	for _, data := range AllContact.Data {
		if data.Id == id {
			return true, *data
		}
	}
	return false, protocol.Contact{}
}

func (AllContact *storageMemory) FindContactByChar(insertChar string) (status bool, result protocol.ContactStorage) {
	if state, res := person.FindPersonByChar(insertChar); state.State {
		for _, data := range AllContact.Data {
			if len(res.PersonData) == 0 {
				break
			}
			if ((res.PersonData)[0]).Id == data.PersonID {
				status = true
				result.Data = append(result.Data, data)
				res.PersonData = (res.PersonData)[1:]
			}
		}
	}
	return status, result
}

func (AllContact *storageMemory) EditContactByID(newData protocol.Contact, ID uint) bool {
	for index, data := range AllContact.Data {
		if data.Id == ID {
			//TODO:: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newData.PersonID != 0 {
				(AllContact.Data)[index].PersonID = newData.PersonID
			}
			if newData.Description != "" {
				(AllContact.Data)[index].Description = newData.Description
			}
			if newData.Tel != "" {
				(AllContact.Data)[index].Tel = newData.Tel
			}
			if newData.Cellphone != "" {
				(AllContact.Data)[index].Cellphone = newData.Cellphone
			}
			if newData.JobID != 0 {
				(AllContact.Data)[index].JobID = newData.JobID
			}
			return true
		}
	}
	return false
}

func (AllContact *storageMemory) DeleteContactByID(ID uint) bool {
	for index, data := range AllContact.Data {
		if data.Id == ID {
			AllContact.Data = append((AllContact.Data)[:index], (AllContact.Data)[index+1:]...)
			//TODO: BadSolution,HandelMultiDeleteInServiceOrStorageNotInClient
			return true
		}
	}
	return false
}

func (AllContact *storageMemory) DeleteAll() bool {
	AllContact.Data = (AllContact.Data)[0:0]
	return true
}
