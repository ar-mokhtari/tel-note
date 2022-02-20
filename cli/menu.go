package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/service/base"
	"tel-note/internal/service/city"
	"tel-note/internal/service/contact"
	"tel-note/internal/service/fillSampleData"
	"tel-note/internal/service/identity"
	"tel-note/internal/service/jobInfo"
	"tel-note/internal/storage/memory"
)

var (
	separator  = strings.Repeat("-", 20)
	separator7 = strings.Repeat(separator, 7)
)

func ShowMenu() {
	fmt.Println(separator7)
	fmt.Println("*** Main menu: ***\n", "Please select:")
	fmt.Printf("%-3s %s %3s \n", separator, "Contact Menu", separator)
	fmt.Print("'N'			|	new record\n")
	fmt.Print("'L'			|	list of contact\n")
	fmt.Print("'F'			|	find one contact by id\n")
	fmt.Print("'FC'			|	find contact, contain some character\n")
	fmt.Print("'E'			|	find and edit contact by contact id\n")
	fmt.Print("'D'			|	delete contact by id\n")
	fmt.Print("'DM'			|	delete multi contact by id(s)\n")
	fmt.Print("'DA'			|	delete all contacts\n")
	fmt.Printf("%-3s %s %3s \n", separator, "City Menu", separator)
	fmt.Print("'NC'			|	insert new city\n")
	fmt.Print("'LC'			|	list of city(ies)\n")
	fmt.Print("'EC'			|	edit city by id\n")
	fmt.Print("'DC'			|	delete city by id\n")
	fmt.Printf("%-3s %s %3s \n", separator, "Job Menu", separator)
	fmt.Print("'NJ'			|	insert new job\n")
	fmt.Print("'LJ'			|	list of job(s)\n")
	fmt.Print("'EJ'			|	edit job by id\n")
	fmt.Print("'DJ'			|	delete job by id\n")
	fmt.Printf("%-3s %s %3s \n", separator, "Base Menu", separator)
	fmt.Print("'NS'			|	insert new sex\n")
	fmt.Print("'ES'			|	edit sex by id\n")
	fmt.Print("'DS'			|	delete sex by id\n")
	fmt.Print("'LS'			|	list of sex\n")
	fmt.Println(separator7)

	if identity.IsAdmin {
		fmt.Printf("%-3s %s %3s \n", separator, "Fill sample data", separator)
		fmt.Print("'DATA'			|	insert some sample's contacts\n")
		fmt.Printf("%-3s %s %3s \n", separator, "Print All", separator)
		fmt.Print("'ALL'			|	print all data\n")
		fmt.Println(separator7)
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
				fillSampleData.FillSimpleDataInMainData()
				ShowMenu()
			case PrintAllData:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				//TODO::: Find a way to loop a struct that contain other struct(s)
				//e := reflect.ValueOf(config.MainData.ContactData).Elem()
				//for i := 0; i < e.NumField(); i++ {
				//	varName := e.Type().Field(i).Name
				//	varType := e.Type().Field(i).Type
				//	varValue := e.Field(i).Interface()
				//	fmt.Printf("%v %v %v\n", varName, varType, varValue)
				//}
				fmt.Printf("%3v | %-10s | %-20v | %-8v | %-15v | %-10v | %-5v | %-20v | %-5v \n", "Id", "FName", "LName", "Tl", "Cellphone", "Desc", "JobID", "JobName", "Gender")
				fmt.Println("")
				for _, data := range config.MainData.ContactData {
					fmt.Printf("%3v | %-10s | %-20v | %-8v | %-15v | %-10v | %-5v | %-20v | %-5v \n", data.Id, data.FirstName, data.LastName, data.Tel, data.Cellphone, data.Description, data.JobInfo.Id, data.JobInfo.Name, data.Person.Gender.Name)
				}
				fmt.Println(separator7)
				fmt.Println("City Data:")
				for _, data := range config.MainData.CityData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("Job Data:")
				for _, data := range config.MainData.JobData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				ShowMenu()
			}
		}
		switch userInput {
		case NewContactRecord:
			var (
				firstName, lastName, tel, cellphone, description string
				genderID                                         uint8
				jobID                                            uint
			)
			fmt.Println("Please inter first name:")
			fmt.Scanln(&firstName)
			//input lastname
			fmt.Println("Please inter last name:")
			fmt.Scanln(&lastName)
			//input gender
			fmt.Println("Please inter genderID:")
			fmt.Scanln(&genderID)
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
			fmt.Println(">> New contact added done <<")
			ShowMenu()
		case ListOfContact:
			dataJSON, _ := json.MarshalIndent(config.MainData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			fmt.Printf("%3v | %-10s | %-20v | %-8v | %-15v | %-10v | %-5v | %-20v | %-5v \n", "Id", "FName", "LName", "Tl", "Cellphone", "Desc", "JobID", "JobName", "Gender")
			fmt.Println("")
			for _, data := range config.MainData.ContactData {
				fmt.Printf("%3v | %-10s | %-20v | %-8v | %-15v | %-10v | %-5v | %-20v | %-5v \n", data.Id, data.FirstName, data.LastName, data.Tel, data.Cellphone, data.Description, data.JobInfo.Id, data.JobInfo.Name, data.Person.Gender.Name)
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
				insertContactID, jobID                           uint
				firstName, lastName, tel, cellphone, description string
				genderID                                         uint8
			)
			fmt.Println("Please insert your contact id you want to edit:")
			fmt.Scanln(&insertContactID)
			result, isFound := contact.FindContactByID(insertContactID)
			if isFound {
				fmt.Println(*result)
			} else {
				fmt.Println("not found")
				ShowMenu()
			}
			fmt.Println("new first name:")
			fmt.Scanln(&firstName)
			fmt.Println("new last name:")
			fmt.Scanln(&lastName)
			fmt.Println("new genderID:")
			fmt.Scanln(&genderID)
			fmt.Println("new tel:")
			fmt.Scanln(&tel)
			fmt.Println("new cellphone:")
			fmt.Scanln(&cellphone)
			fmt.Println("new job id:")
			fmt.Scanln(&jobID)
			fmt.Println("new description:")
			fmt.Scanln(&description)
			editedContact := memory.Contact{
				Person: &memory.Person{
					FirstName: firstName,
					LastName:  lastName,
					Gender:    &memory.Sex{Id: genderID},
				},
				Tel:         tel,
				Cellphone:   cellphone,
				JobInfo:     &memory.JobInfo{Id: jobID},
				Description: description,
			}
			state := contact.EditContactByID(editedContact, insertContactID)
			fmt.Println(state.String)
			ShowMenu()
		case DeleteContactById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting a contact ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
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
			if strings.ToLower(confirmDel) == YES {
				resultStatus := contact.DeleteAll()
				fmt.Println(resultStatus.String)
			}
			ShowMenu()
		case DeleteMultiContactByIds:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting contact(s) ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
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
		case InsertNewCity:
			var inputCity string
			fmt.Println("insert city name:")
			fmt.Scanln(&inputCity)
			if city.NewCity(inputCity).State {
				fmt.Println("New city added")
			}
			ShowMenu()
		case ListOfCities:
			dataJSON, _ := json.MarshalIndent(config.MainData.CityData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
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
		case DeleteCityById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting a city ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
				var deleteID uint
				fmt.Println("insert your city id that you want to delete:")
				fmt.Scanln(&deleteID)
				status := city.DeleteCityByID(deleteID)
				switch status.State {
				case true:
					fmt.Printf("city with id number:  %d deleted.", deleteID)
				case false:
					fmt.Println("something wrong")
				}
			}
			ShowMenu()
		case InsertNewJob:
			var inputJob string
			fmt.Println("insert job name:")
			fmt.Scanln(&inputJob)
			if jobInfo.NewJob(inputJob).State {
				fmt.Println("New job added")
			}
			ShowMenu()
		case EditJobById:
			var inputID uint
			var inputName string
			fmt.Println("insert job id:")
			fmt.Scanln(&inputID)
			fmt.Println("insert new job name:")
			fmt.Scanln(&inputName)
			if jobInfo.EditJobInfoByID(inputID, inputName).State {
				fmt.Println("Job changed ...")
			}
			ShowMenu()
		case DeleteJobById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting a job ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
				var deleteID uint
				fmt.Println("insert your job id that you want to delete:")
				fmt.Scanln(&deleteID)
				status := jobInfo.DeleteJobByID(deleteID)
				switch status.State {
				case true:
					fmt.Printf("job with id number:  %d deleted.", deleteID)
				case false:
					fmt.Println("something wrong")
				}
			}
			ShowMenu()
		case ListOfJob:
			dataJSON, _ := json.MarshalIndent(config.MainData.JobData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range (config.MainData).JobData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		case InsertNewSex:
			var insertName string
			fmt.Println("Insert sex name")
			fmt.Scanln(&insertName)
			if base.NewSex(insertName).State {
				fmt.Println("New sex info added")
			}
			ShowMenu()
		case EditSex:
			var insertID uint8
			var insertName string
			fmt.Println("Insert sex id:")
			fmt.Scanln(&insertID)
			fmt.Println("Insert new name:")
			fmt.Scanln(&insertName)
			if base.EditSexNameByID(insertID, insertName).State {
				fmt.Println("sex info updated")
			} else {
				fmt.Println("some thing wrong")
			}
			ShowMenu()
		case DeleteSex:
			fmt.Println("you are deleting some data, are you sure? (yes/no)")
			var confirmInsert string
			fmt.Scanln(&confirmInsert)
			if strings.ToLower(confirmInsert) == YES {
				var insertID uint8
				fmt.Println("insert sex id to delete:")
				fmt.Scanln(&insertID)
				if base.DeleteSexByID(insertID).State {
					fmt.Println("sex info deleted")
				} else {
					fmt.Println("some thing wrong")
				}
			}
			ShowMenu()
		case ListOfSex:
			dataJSON, _ := json.MarshalIndent(config.MainData.SexData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range (config.MainData).SexData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		//something wrong:
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}
