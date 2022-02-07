package cli

import (
	"encoding/json"
	"fmt"
	"tel-note/internal/service/menu/contact"
	"tel-note/internal/storage"
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
		//something wrong:
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}
