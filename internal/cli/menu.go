package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/service/contact"
	"tel-note/internal/storage"
)

func ShowMenu(MainData *storage.AllContact) {
	fmt.Println("------------------")
	fmt.Println("*** Main menu: ***", "Please select:") //todo: Why fmt.Println don't split contents line
	fmt.Print("'N|n' 	 	new record\n")
	fmt.Print("'L|n' 	 	list of contact\n")
	fmt.Print("'F|f' 	 	find one contact by id\n")
	fmt.Print("'FF|ff' 	find contact, contain some character\n")
	fmt.Print("'e|E' 		find and edit contact by contact id\n")
	fmt.Print("'d|D' 		delete contact by id\n")
	fmt.Println("------------------")
	runMenu(MainData)
}

func runMenu(MainData *storage.AllContact) {
	for {
		var userInput string
		fmt.Scanln(&userInput)
		//fmt.Println(userInput) todo: when negative value inserted, maybe because of define uint for id; The positive value is entered again, why?
		switch userInput {
		case "N", "n":
			//todo: “To use function, or to use method, that is the question”
			var firstName, lastName, tel, cellphone, description string
			fmt.Println("Please inter your first name:")
			fmt.Scanln(&firstName)
			//input lastname
			fmt.Println("Please inter your last name:")
			fmt.Scanln(&lastName)
			//input tel
			fmt.Println("Please inter your tel:")
			fmt.Scanln(&tel)
			//input cellphone
			fmt.Println("Please inter your cellphone:")
			fmt.Scanln(&cellphone)
			//input description lastID +
			fmt.Println("Please inter your description:")
			fmt.Scanln(&description)
			contact.NewContact(MainData, firstName, lastName, tel, cellphone, description)
			//(*storage.AllContact).AddContact(MainData)
			fmt.Println(">> New record done <<")
			ShowMenu(MainData)
		case "L", "l":
			dataJSON, _ := json.MarshalIndent(MainData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(MainData)
			ShowMenu(MainData)
		case "F", "f":
			var id uint
			fmt.Println("Please insert your contact ID:")
			fmt.Scanln(&id)
			result, isFound := contact.FindContactByID(MainData, id)
			if isFound {
				fmt.Println(result)
			} else {
				fmt.Println("not found")
			}
			ShowMenu(MainData)
		case "FF", "ff":
			var insertChar string
			fmt.Println("insert character(s):")
			fmt.Scanln(&insertChar)
			resultData, resultCount := contact.FindContactByChar(MainData, insertChar)
			fmt.Println(resultData)
			if resultCount == 0 {
				fmt.Println("not found")
			} else {
				fmt.Println(resultCount, "record(s) found")
			}
			ShowMenu(MainData)
		case "E", "e":
			var (
				insertContactID                                  uint
				firstName, lastName, tel, cellphone, description string
			)
			fmt.Println("Please insert your contact id you want to edit:")
			fmt.Scanln(&insertContactID)
			result, isFound := contact.FindContactByID(MainData, insertContactID)
			if isFound {
				fmt.Println(result)
			} else {
				fmt.Println("not found")
				ShowMenu(MainData)
			}
			fmt.Println("new first name:")
			fmt.Scanln(&firstName)
			fmt.Println("new last name:")
			fmt.Scanln(&lastName)
			fmt.Println("new description:")
			fmt.Scanln(&tel)
			fmt.Println("new tel:")
			fmt.Scanln(&cellphone)
			fmt.Println("new cellphone:")
			fmt.Scanln(&description)
			editedContact := storage.Contact{FirstName: firstName, LastName: lastName, Tel: tel, Cellphone: cellphone, Description: description}
			if state := contact.EditContactByID(MainData, editedContact, insertContactID); state.State != false {
				fmt.Println("contact edited successful")
			} else {
				println("contact not changed")
			}
			ShowMenu(MainData)
		case "D", "d":
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting a contact ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == config.OkStatus {
				var deleteID uint
				fmt.Println("insert your contact id that you want to delete:")
				fmt.Scanln(&deleteID)
				status := contact.DeleteContactByID(MainData, deleteID)
				if status.State {
					fmt.Printf("contact with id number:  %d deleted.", deleteID)
				}
			}
			ShowMenu(MainData)
		//something wrong:
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}
