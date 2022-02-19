package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/service/allData"
	"tel-note/internal/service/city"
	"tel-note/internal/service/contact"
	"tel-note/internal/service/identity"
	"tel-note/internal/storage/memory"
)

func ShowMenu() {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("*** Main menu: ***\n", "Please select:")
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("'N'			|	new record\n")
	fmt.Print("'L'			|	list of contact\n")
	fmt.Print("'F'			|	find one contact by id\n")
	fmt.Print("'FC'			|	find contact, contain some character\n")
	fmt.Print("'E'			|	find and edit contact by contact id\n")
	fmt.Print("'D'			|	delete contact by id\n")
	fmt.Print("'DM'			|	delete multi contact by id(s)\n")
	fmt.Print("'DA'			|	delete all contacts\n")
	fmt.Println("-------------------------------------------------------------")
	if identity.IsAdmin {
		fmt.Print("'DATA'			|	insert some sample's contacts\n")
		fmt.Print("'NC'			|	insert new city\n")
		fmt.Print("'LC'			|	list of cities\n")
		fmt.Print("'EC'			|	edit city by id\n")
		fmt.Println("-------------------------------------------------------------")
		fmt.Print("'ALL'			|	print all data\n")
		fmt.Println("-------------------------------------------------------------")
	}
	runMenu()
}

func runMenu() {
	for {
		var userInput string
		fmt.Scanln(&userInput)
		//admin actions only:
		userInput = strings.ToUpper(userInput)
		if identity.IsAdmin {
			switch userInput {
			case InsertSomeSamplesContacts:
				allData.FillSimpleDataInMainData()
				ShowMenu()
			case InsertNewCity:
				var inputCity string
				fmt.Println("insert city name:")
				fmt.Scanln(&inputCity)
				city.NewCity(inputCity)
				ShowMenu()
			case ListOfCities:
				dataJSON, _ := json.MarshalIndent(config.MainData, "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println("-------------------------------------------------------------")
				for _, data := range (config.MainData).CityData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				ShowMenu()
			case EditCityById:
				var inputID uint
				var inputName string
				fmt.Println("insert city id:")
				fmt.Scanln(&inputID)
				fmt.Println("insert new city name:")
				fmt.Scanln(&inputName)
				if city.EditCityByID(inputID, inputName).State {
					fmt.Println("City changed ...")
				}
				ShowMenu()
			case PrintAllData:
				fmt.Println("-------------------------------------------------------------")
				fmt.Println("Contact Data:")
				//TODO::: Find a way to loop a struct that contain other struct(s)
				//e := reflect.ValueOf(config.MainData.ContactData).Elem()
				//for i := 0; i < e.NumField(); i++ {
				//	varName := e.Type().Field(i).Name
				//	varType := e.Type().Field(i).Type
				//	varValue := e.Field(i).Interface()
				//	fmt.Printf("%v %v %v\n", varName, varType, varValue)
				//}
				for _, data := range config.MainData.ContactData {
					fmt.Printf("%3v | %-10s | %-20v | %-2v | %-15v | %-10v | %-15v | %-10v | %-5v \n", data.Id, data.FirstName, data.LastName, data.Tel, data.Cellphone, data.Description, data.JobInfo.Name, data.JobInfo.Id, data.Person.Gender.Id)
				}
				fmt.Println("-------------------------------------------------------------")
				fmt.Println("City Data:")
				for _, data := range config.MainData.CityData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println("-------------------------------------------------------------")
				fmt.Println("Job Data:")
				for _, data := range config.MainData.JobData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				ShowMenu()
			}
		}
		switch userInput {
		case NewRecord:
			var (
				firstName, lastName, tel, cellphone, description, gender string
				genderID                                                 uint8
				jobID                                                    uint
			)
			fmt.Println("Please inter first name:")
			fmt.Scanln(&firstName)
			//input lastname
			fmt.Println("Please inter last name:")
			fmt.Scanln(&lastName)
			//input gender
			fmt.Println("Please inter gender: (m/f)")
			fmt.Scanln(&gender)
			switch strings.ToLower(gender) {
			case "f":
				genderID = 1
			case "m":
				genderID = 2
			}
			//input tel
			fmt.Println("Please inter tel:")
			fmt.Scanln(&tel)
			//input cellphone
			fmt.Println("Please inter cellphone:")
			fmt.Scanln(&cellphone)
			//input description
			fmt.Println("Please inter description:")
			fmt.Scanln(&description)
			//input job info
			fmt.Println("Please inter job ID:")
			fmt.Scanln(&jobID)
			contact.NewContact(
				memory.Contact{
					Person: &memory.Person{
						FirstName: firstName,
						LastName:  lastName,
						Gender:    &memory.Sex{Id: genderID},
					},
					Tel:         tel,
					Cellphone:   cellphone,
					Description: description,
					JobInfo:     &memory.JobInfo{Id: jobID},
				})
			fmt.Println(">> New record done <<")
			ShowMenu()
		case ListOfContact:
			dataJSON, _ := json.MarshalIndent(config.MainData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println("-------------------------------------------------------------")
			for _, data := range (config.MainData).ContactData {
				fmt.Printf("%3v | %-15s | %-35v | %-5v | %-15v | %-5v\n", data.Id, data.FirstName, data.LastName, data.Tel, data.Cellphone, data.Description)
			}
			ShowMenu()
		case FindOneContactById:
			var id uint
			fmt.Println("Please insert your contact ID:")
			fmt.Scanln(&id)
			result, isFound := contact.FindContactByID(id)
			if isFound {
				fmt.Println(result)
			} else {
				fmt.Println("not found")
			}
			ShowMenu()
		case FindContactContainingSomeCharacter:
			var insertChar string
			fmt.Println("insert character(s):")
			fmt.Scanln(&insertChar)
			resultData, resultCount := contact.FindContactByChar(insertChar)
			fmt.Println(resultData)
			if resultCount == 0 {
				fmt.Println("not found")
			} else {
				fmt.Println(resultCount, "record(s) found")
			}
			ShowMenu()
		case FindAndEditContactByContactId:
			var (
				insertContactID                                  uint
				firstName, lastName, tel, cellphone, description string
			)
			fmt.Println("Please insert your contact id you want to edit:")
			fmt.Scanln(&insertContactID)
			result, isFound := contact.FindContactByID(insertContactID)
			if isFound {
				fmt.Println(result)
			} else {
				fmt.Println("not found")
				ShowMenu()
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
			editedContact := memory.Contact{Person: &memory.Person{FirstName: firstName, LastName: lastName}, Tel: tel, Cellphone: cellphone, Description: description}
			state := contact.EditContactByID(editedContact, insertContactID)
			fmt.Println(state.String)
			ShowMenu()
		case DeleteContactById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting a contact ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == OK {
				var deleteID uint
				fmt.Println("insert your contact id that you want to delete:")
				fmt.Scanln(&deleteID)
				status := contact.DeleteContactByID(deleteID)
				switch status.State {
				case true:
					fmt.Printf("contact with id number:  %d deleted.", deleteID)
				case false:
					fmt.Println("something wrong")
				}
			}
			ShowMenu()
		case DeleteAllContacts:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting all of contacts ***")
			fmt.Println("are you sure? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == OK {
				resultStatus := contact.DeleteAll()
				fmt.Println(resultStatus.String)
			}
			ShowMenu()
		case DeleteMultiContactByIds:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting contact(s) ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == OK {
				var deleteIDS string
				var status config.ResponseStatus
				fmt.Println("insert your contact id(s) that you want to delete, separate id's by ',':")
				fmt.Scanln(&deleteIDS)
				idPack := strings.Split(deleteIDS, ",")
				for _, delID := range idPack {
					uintDelID, _ := strconv.ParseUint(delID, 10, 8)
					if uint(uintDelID) <= 0 {
						fmt.Println("not equal kind")
						break
					} else {
						status = *contact.DeleteContactByID(uint(uintDelID))
						fmt.Println(status.String)
					}
				}
			}
			ShowMenu()
		//something wrong:
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}
