package contact

import (
	"context"
	"tel-note/internal/protocol"
	"tel-note/internal/storage/memory"
)

type (
	storageContact struct {
		contacts *protocol.Contact
	}
)

func (m *storageContact) AddContact(ctx context.Context, inputContact protocol.Contact) bool {
	var (
		lastID uint
	)
	for _, data := range m.contacts {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//find job name by ID to return for end user
	inputContact.JobInfo.Name = MainData.FindJobById(inputContact.JobInfo.Id)
	//find gender name by ID to return for end user
	inputContact.Gender.Name = MainData.FindGenderNameById(inputContact.Gender.Id)
	//marge inputs to create a contact
	result := &Contact{
		Id: lastID,
		memory.Person: &memory.Person{
			FirstName: inputContact.FirstName,
			LastName:  inputContact.LastName,
			Gender: &memory.Sex{
				Id:   inputContact.Gender.Id,
				Name: inputContact.Gender.Name,
			},
		},
		Tel:         inputContact.Tel,
		Cellphone:   inputContact.Cellphone,
		Description: inputContact.Description,
		memory.JobInfo: &memory.JobInfo{
			Id:   inputContact.JobInfo.Id,
			Name: inputContact.JobInfo.Name,
		},
	}
	MainData.contacts = append(MainData.contacts, result)
	return true, *MainData
}

//func (MainData *storage) FindContactByID(id uint) (Contact, bool) {
//	var (
//		data Contact
//	)
//	for _, data := range MainData.ContactData {
//		if data.Id == id {
//			return *data, true
//			break
//		}
//	}
//	return data, false
//}
//
//func (MainData *storage) FindContactByChar(insertChar string) (memory.AllData, uint) {
//	var (
//		counter    uint
//		resultData memory.AllData
//	)
//	for _, data := range MainData.ContactData {
//		if strings.Contains(data.FirstName, insertChar) {
//			counter += 1
//			resultData.ContactData = append(resultData.ContactData, data)
//		}
//	}
//	return resultData, counter
//}
//
//func (MainData *storage) EditContactByID(newData Contact, ID uint) bool {
//	for index, data := range MainData.ContactData {
//		if data.Id == ID {
//			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
//			if newData.FirstName != "" {
//				(MainData.ContactData)[index].FirstName = newData.FirstName
//			}
//			if newData.LastName != "" {
//				(MainData.ContactData)[index].LastName = newData.LastName
//			}
//			if newData.Description != "" {
//				(MainData.ContactData)[index].Description = newData.Description
//			}
//			if newData.Tel != "" {
//				(MainData.ContactData)[index].Tel = newData.Tel
//			}
//			if newData.Cellphone != "" {
//				(MainData.ContactData)[index].Cellphone = newData.Cellphone
//			}
//			if newData.Gender.Id != 0 {
//				(MainData.ContactData)[index].Gender.Id = newData.Gender.Id
//				(MainData.ContactData)[index].Gender.Name = MainData.FindGenderNameById(newData.Gender.Id)
//			}
//			if newData.JobInfo.Id != 0 {
//				(MainData.ContactData)[index].JobInfo.Id = newData.JobInfo.Id
//				(MainData.ContactData)[index].JobInfo.Name = MainData.FindJobById(newData.JobInfo.Id)
//			}
//			return true
//		}
//	}
//	return false
//}
//
//func (MainData *storage) DeleteContactByID(ID uint) bool {
//	for index, data := range MainData.ContactData {
//		if data.Id == ID {
//			MainData.ContactData = append(MainData.ContactData[:index], MainData.ContactData[index+1:]...)
//			return true
//		}
//	}
//	return false
//}
//
//func (MainData *storage) DeleteAll() bool {
//	MainData.ContactData = MainData.ContactData[0:0]
//	return true
//}
