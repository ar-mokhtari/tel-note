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
	AllContact struct {
		ContactData []*Contact
	}
)

// AddContact todo: merge with func the same ("NewContact")
func (MainData *AllContact) AddContact() {
	var (
		lastID uint
	)
	//define input variable for save to struct
	var firstName, lastName, tel, cellphone, description string
	//input and save firstname
	for _, data := range MainData.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//append lastID

	fmt.Println("Please inter your first name:")
	fmt.Scanln(&firstName)

	//input and save lastname
	fmt.Println("Please inter your last name:")
	fmt.Scanln(&lastName)

	//input and save tel
	fmt.Println("Please inter your tel:")
	fmt.Scanln(&tel)

	//input and save cellphone
	fmt.Println("Please inter your cellphone:")
	fmt.Scanln(&cellphone)

	//input and save description lastID +
	fmt.Println("Please inter your description:")
	fmt.Scanln(&description)

	//marge inputs to create a contact
	result := &Contact{Id: lastID, Person: &Person{FirstName: firstName, LastName: lastName}, Tel: tel, Cellphone: cellphone, Description: description}

	MainData.ContactData = append(MainData.ContactData, result)

	fmt.Println(">> New record done <<")
}
