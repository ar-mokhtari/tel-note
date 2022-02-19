package memory

import (
	"fmt"
	"strings"
)

type (
	AllDataTool interface {
		// Contact CRUD basic info
		AddContact(inputContact Contact) (bool, AllData)
		FindContactByID(id uint) (Contact, bool)
		FindContactByChar(insertChar string) (AllData, uint)
		EditContactByID(newData Contact, ID uint) bool
		DeleteContactByID(ID uint) bool
		DeleteAll() bool
		// City CRUD basic info
		NewCity(CityName string) bool
		EditCityByID(ID uint, NewCityName string) bool
		// JobInfo CRUD basic info
		NewJob(jobName string) bool
	}

	AllData struct {
		ContactData []*Contact
		JobData     []*JobInfo
		PersonData  []*Person
		CityData    []*City
		SexData     []*Sex
	}
)

// Contact CRUD basic info

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
	result := &Contact{Id: lastID, Person: &Person{FirstName: inputContact.FirstName, LastName: inputContact.LastName, Gender: &Sex{Id: inputContact.Gender.Id}}, Tel: inputContact.Tel, Cellphone: inputContact.Cellphone, Description: inputContact.Description, JobInfo: &JobInfo{Id: inputContact.JobInfo.Id}}

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

func (MainData *AllData) EditContactByID(newData Contact, ID uint) bool {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
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
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteContactByID(ID uint) bool {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			MainData.ContactData = append(MainData.ContactData[:index], MainData.ContactData[index+1:]...)
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteAll() bool {
	MainData.ContactData = MainData.ContactData[0:0]
	return true
}

// City CRUD basic info

func (MainData *AllData) NewCity(CityName string) bool {
	var LastID uint
	for _, data := range MainData.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := City{
		Id:   uint(LastID) + 1,
		Name: CityName,
	}
	MainData.CityData = append(MainData.CityData, &result)
	return true
}

func (MainData *AllData) EditCityByID(ID uint, NewCityName string) bool {
	for _, data := range MainData.CityData {
		if data.Id == ID {
			data.Name = NewCityName
			return true
		}
	}
	return false
}

// Job CRUD basic info

func (MainData *AllData) NewJob(jobName string) bool {
	var LastID uint
	for _, data := range MainData.JobData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := JobInfo{
		Id:   uint(LastID) + 1,
		Name: jobName,
	}
	MainData.JobData = append(MainData.JobData, &result)
	return true
}
