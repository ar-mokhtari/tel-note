package storage

import (
	"fmt"
	"strings"
	"tel-note/internal/config"
)

type (
	AllDataTool interface {
		AddContact(inputContact Contact) (bool, AllData)
		FindContactByID(id uint) (Contact, bool)
		FindContactByChar(insertChar string) (AllData, uint)
		EditContactByID(newData Contact, ID uint) config.ResponseStatus
		DeleteContactByID(ID uint) config.ResponseStatus
		DeleteAll() config.ResponseStatus
	}

	AllData struct {
		ContactData []*Contact
		CityData    []*City
		jobData     []*jobInfo
	}
)

func (MainData *AllData) AddContact(inputContact Contact) (bool, AllData) {
	var (
		lastID uint
	)
	for _, data := range MainData.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//marge inputs to create a contact
	result := &Contact{Id: lastID, Person: &Person{FirstName: inputContact.FirstName, LastName: inputContact.LastName}, Tel: inputContact.Tel, Cellphone: inputContact.Cellphone, Description: inputContact.Description}

	MainData.ContactData = append(MainData.ContactData, result)

	fmt.Println(">> New record done <<")

	return true, *MainData
}

func (MainData *AllData) FindContactByID(id uint) (Contact, bool) {
	var (
		data Contact
	)
	for _, data := range MainData.ContactData {
		if data.Id == id {
			return *data, true
			break
		}
	}
	return data, false
}

func (MainData *AllData) FindContactByChar(insertChar string) (AllData, uint) {
	var (
		counter    uint
		resultData AllData
	)
	for _, data := range MainData.ContactData {
		if strings.Contains(data.FirstName, insertChar) {
			counter += 1
			resultData.ContactData = append(resultData.ContactData, data)
		}
	}
	return resultData, counter
}

func (MainData *AllData) EditContactByID(newData Contact, ID uint) config.ResponseStatus {
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

func (MainData *AllData) DeleteContactByID(ID uint) config.ResponseStatus {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			MainData.ContactData = append(MainData.ContactData[:index], MainData.ContactData[index+1:]...)
			return config.ResponseStatus{State: true}
		}
	}
	return config.ResponseStatus{State: false}
}

func (MainData *AllData) DeleteAll() config.ResponseStatus {
	MainData.ContactData = MainData.ContactData[0:0]
	return config.ResponseStatus{State: true}
}
