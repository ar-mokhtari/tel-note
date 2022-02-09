package contact

import (
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/storage"
)

func NewContact(MainData *storage.AllContact, firstName, lastName, tel, cellphone, description string) *storage.AllContact {
	var (
		lastID uint
	)
	//define input variable for save to struct
	//input and save firstname
	for _, data := range MainData.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//marge inputs to create a contact
	result := &storage.Contact{Id: lastID, FirstName: firstName, LastName: lastName, Tel: tel, Cellphone: cellphone, Description: description}
	//append to pointed storage
	MainData.ContactData = append(MainData.ContactData, *result)
	return MainData
}

func FindContactByID(MainData *storage.AllContact, id uint) (storage.Contact, bool) {
	var (
		data storage.Contact
	)
	for _, data := range MainData.ContactData {
		if data.Id == id {
			return data, true
			break
		}
	}
	return data, false
}

func FindContactByChar(MainData *storage.AllContact, insertChar string) (storage.AllContact, uint) {
	var (
		counter    uint
		resultData storage.AllContact
	)
	for _, data := range MainData.ContactData {
		if strings.Contains(data.FirstName, insertChar) {
			counter += 1
			resultData.ContactData = append(resultData.ContactData, data)
		}
	}
	return resultData, counter
}

func EditContactByID(MainData *storage.AllContact, newData storage.Contact, ID uint) config.ResponseStatus {
	//todo: is there any function to detect index of an element's slice?
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			//todo: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newData.FirstName != "" {
				(MainData.ContactData)[index].FirstName = newData.FirstName
			}
			if newData.LastName != "" {
				(MainData.ContactData)[index].LastName = newData.LastName
			}
			if newData.Description != "" {
				(MainData.ContactData)[index].Description = newData.Description
			}
			if newData.Tel != "" {
				(MainData.ContactData)[index].Tel = newData.Tel
			}
			if newData.Cellphone != "" {
				(MainData.ContactData)[index].Cellphone = newData.Cellphone
			}
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{String: "not found"}
}

func DeleteContactByID(MainData *storage.AllContact, ID uint) config.ResponseStatus {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			MainData.ContactData = append(MainData.ContactData[:index], MainData.ContactData[index+1:]...)
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{State: false}
}

func DeleteAll(MainData *storage.AllContact) config.ResponseStatus {
	MainData.ContactData = MainData.ContactData[0:0]
	return config.ResponseStatus{State: true}
}

func FillSimpleDataInMainData(MainData *storage.AllContact) {
	for index, data := range config.MainDataTest.ContactData {
		data.Id = uint(index) //uint(index)
		MainData.ContactData = append(MainData.ContactData, data)
	}
}
