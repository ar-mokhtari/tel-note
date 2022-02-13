package storage

import (
	"fmt"
)

type (
	Contact struct {
		Id uint `json:"id"`
		*Person
		*jobInfo
		Tel         string `json:"tel,omitempty"`
		Cellphone   string `json:"cellphone,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

// AddContact todo: merge with func the same ("NewContact")
func (MainData *AllData) AddContact(firstName, lastName, tel, cellphone, description string) AllData {
	var (
		lastID uint
	)
	//input and save firstname
	for _, data := range MainData.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//marge inputs to create a contact
	result := &Contact{Id: lastID, Person: &Person{FirstName: firstName, LastName: lastName}, Tel: tel, Cellphone: cellphone, Description: description}

	MainData.ContactData = append(MainData.ContactData, result)

	fmt.Println(">> New record done <<")
	return *MainData
}
