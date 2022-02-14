package contact

import (
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

func NewContact(inputContact storage.Contact) (*storage.AllData, config.ResponseStatus) {
	var (
		lastID uint
	)
	//define input variable for save to struct
	//input and save firstname

	//lastID = MainData.GetANewCodeID()
	for _, data := range (config.MainData).ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//marge inputs to create a contact
	result := &storage.Contact{Id: lastID, Person: &storage.Person{FirstName: inputContact.FirstName, LastName: inputContact.LastName}, Tel: inputContact.Tel, Cellphone: inputContact.Cellphone, Description: inputContact.Description}
	//append to pointed storage
	(config.MainData).ContactData = append((config.MainData).ContactData, result)
	return &config.MainData, config.ResponseStatus{State: true}
}

func FindContactByID(id uint) (storage.Contact, bool) {
	var (
		data storage.Contact
	)
	for _, data := range (config.MainData).ContactData {
		if data.Id == id {
			return *data, true
			break
		}
	}
	return data, false
}

func FindContactByChar(insertChar string) (storage.AllData, uint) {
	var (
		counter    uint
		resultData storage.AllData
	)
	for _, data := range (config.MainData).ContactData {
		if strings.Contains(data.FirstName, insertChar) {
			counter += 1
			resultData.ContactData = append(resultData.ContactData, data)
		}
	}
	return resultData, counter
}

func EditContactByID(newData storage.Contact, ID uint) config.ResponseStatus {
	//todo: is there any function to detect index of an element's slice?
	for index, data := range (config.MainData).ContactData {
		if data.Id == ID {
			//todo: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newData.FirstName != "" {
				((config.MainData).ContactData)[index].FirstName = newData.FirstName
			}
			if newData.LastName != "" {
				((config.MainData).ContactData)[index].LastName = newData.LastName
			}
			if newData.Description != "" {
				((config.MainData).ContactData)[index].Description = newData.Description
			}
			if newData.Tel != "" {
				((config.MainData).ContactData)[index].Tel = newData.Tel
			}
			if newData.Cellphone != "" {
				((config.MainData).ContactData)[index].Cellphone = newData.Cellphone
			}
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{String: "not found"}
}

func DeleteContactByID(ID uint) config.ResponseStatus {
	for index, data := range (config.MainData).ContactData {
		if data.Id == ID {
			(config.MainData).ContactData = append((config.MainData).ContactData[:index], (config.MainData).ContactData[index+1:]...)
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{State: false}
}

func DeleteAll() config.ResponseStatus {
	(config.MainData).ContactData = (config.MainData).ContactData[0:0]
	return config.ResponseStatus{State: true}
}
