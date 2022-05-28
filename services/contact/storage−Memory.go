//go:build !mysql || !mongodb

package contact

import (
	"errors"
	"github.com/ar-mokhtari/tel-note/protocol"
	"github.com/ar-mokhtari/tel-note/services/person"
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

func (allContact *storageMemory) AddContact(inputContact protocol.Contact) error {
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
	return nil
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
	res := person.FindCharPerson.Do(insertChar)
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
	return status, result
}

func (allContact *storageMemory) EditContactByID(newData protocol.Contact) error {
	for index, data := range allContact.ContactData {
		if data.Id == newData.Id {
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
			return nil
		}
	}
	return errors.New("contact not found")
}

func (allContact *storageMemory) DeleteContactByID(ID uint) bool {
	for index, data := range allContact.ContactData {
		if data.Id == ID {
			allContact.ContactData = append((allContact.ContactData)[:index], (allContact.ContactData)[index+1:]...)
			return true
		}
	}
	return false
}

func (allContact *storageMemory) DeleteAll() bool {
	allContact.ContactData = (allContact.ContactData)[0:0]
	return true
}
