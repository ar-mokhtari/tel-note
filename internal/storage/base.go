package storage

import (
	"fmt"
)

type (
	AllDataTool interface {
		GetANewCodeID()
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

//func (MainData *AllData) FindContactByID(id uint) (Contact, bool) {
//	var (
//		data Contact
//	)
//	for _, data := range (config.MainData).ContactData {
//		if data.Id == id {
//			return *data, true
//			break
//		}
//	}
//	return data, false
//}
