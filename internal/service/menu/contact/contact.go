package contact

import (
	"tel-note/internal/storage"
	"strings"
)

func NewContact(MainData *storage.AllContact, firstName, lastName, tel, cellphone, description string) {
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

	MainData.ContactData = append(MainData.ContactData, *result)
}

func FindContactByID(MainData *storage.AllContact, id uint) storage.Contact {
	var (
		data storage.Contact
	)
	for _, data := range MainData.ContactData {
		if data.Id == id {
			return data
			break
		}
	}
	return data
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
