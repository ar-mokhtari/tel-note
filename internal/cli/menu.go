package cli

import (
	"TelNote/internal/service/menu/contact"
	"TelNote/internal/storage"
	"encoding/json"
	"fmt"
)

func ShowMenu(MainData *storage.AllContact) {
	fmt.Println("------------------")
	fmt.Println("*** Main menu: ***", "Please select:") //Todo: Why fmt.Println don't split contents line
	fmt.Print("'N|n' 	 	new record\n")
	fmt.Print("'L|n' 	 	list of contact\n")
	fmt.Print("'F|f' 	 	find one contact\n")
	fmt.Print("'FF|ff' 	find contact, contain some character\n")
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

			contact.NewContact(MainData, firstName, lastName, tel, cellphone, description)
			//(*storage.AllContact).AddContact(MainData)

			println(">> New record done <<")

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
			fmt.Println(contact.FindContactByID(MainData, id))
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
				fmt.Println(resultCount, "record found")
			}
			ShowMenu(MainData)
		//something wrong:
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}
