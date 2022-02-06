package storage

import "fmt"

type (
	Contact struct {
		Id          uint   `json:"id"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Tel         string `json:"tel,omitempty"`
		Cellphone   string `json:"cellphone,omitempty"`
		Description string `json:"description,omitempty"`
	}

	AllContact struct {
		ContactData []Contact
	}
)


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

	println("Please inter your first name:")
	fmt.Scanln(&firstName)

	//input and save lastname
	println("Please inter your last name:")
	fmt.Scanln(&lastName)

	//input and save tel
	println("Please inter your tel:")
	fmt.Scanln(&tel)

	//input and save cellphone
	println("Please inter your cellphone:")
	fmt.Scanln(&cellphone)

	//input and save description lastID +
	println("Please inter your description:")
	fmt.Scanln(&description)

	//marge inputs to create a contact
	result := Contact{Id: lastID, FirstName: firstName, LastName: lastName, Tel: tel, Cellphone: cellphone, Description: description}

	MainData.ContactData = append(MainData.ContactData, result)

	println(">> New record done <<")
}
