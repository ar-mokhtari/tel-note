package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/fillSampleData"
	"tel-note/services/globalVars"
	"tel-note/services/identity"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
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
	fmt.Printf("%-3s %s %3s \n", separator, "sex Menu", separator)
	fmt.Print("'NS'			|	insert new sex\n")
	fmt.Print("'ES'			|	edit sex by id\n")
	fmt.Print("'DS'			|	delete sex by id\n")
	fmt.Print("'LS'			|	list of sex\n")
	fmt.Println(separator7)
	fmt.Print("'RESET'			|	reset program (as new user role)\n")
	fmt.Println(separator7)

	if identity.IsRegulator {
		fmt.Printf("%-3s %s %3s \n", separator, "Fill sample data", separator)
		fmt.Print("'DATA'			|	insert some sample's contacts\n")
		fmt.Printf("%-3s %s %3s \n", separator, "Print All", separator)
		fmt.Print("'ALL'			|	print all data\n")
		fmt.Println(separator7)
	}
	runMenu()
}

func runMenu() {
	//TODO::: CompleteAllEntityFeature(Person|Job|Sex|City|Contact)
	for {
		var userInput string
		fmt.Scanln(&userInput)
		//regulator actions only:
		userInput = strings.ToUpper(userInput)
		if identity.IsRegulator {
			switch userInput {
			case InsertSomeSamplesContacts:
				fillSampleData.FillSimpleDataInMainData()
				ShowMenu()
			case PrintAllData:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range globalVars.AllContact.Data {
					_, person := person.FindPersonByID(data.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(data.JobID)
					_, city := city.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.Cellphone, job.LocationID, city.Name, data.Description)
				}
				fmt.Println(separator7)
				fmt.Println("City Data:")
				for _, data := range globalVars.AllCity.CityData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("Job Data:")
				for _, data := range globalVars.AllJob.JobData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("sex Data:")
				for _, data := range globalVars.AllSex.SexData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("person Data:")
				for _, data := range globalVars.AllPerson.PersonData {
					fmt.Printf("%3v | %-15s \t %30s \n", data.Id, data.FirstName, data.LastName)
				}
				ShowMenu()
			}
		}
		switch userInput {
		case NewContactRecord:
			var (
				tel, cellphone, description string
				jobID, personID             uint
			)
			//input person
			fmt.Println("Please inter personID:")
			fmt.Scanln(&personID)
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
				protocol.Contact{
					PersonID:    personID,
					Tel:         tel,
					Cellphone:   cellphone,
					Description: description,
					JobID:       jobID,
				})
			fmt.Println(">> New contact added done <<")
			ShowMenu()
			//TODO::: ArrangeOutputTable
		case ListOfContact:
			dataJSON, _ := json.MarshalIndent(globalVars.AllContact, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			fmt.Printf("%3v | %-10v | %-20v | %-15v | %-15v | %-10v | %-5v | %-20v \n", "Id", "personID", "Tl", "Cellphone", "Desc", "JobID", "JobName", "Gender")
			fmt.Println("")
			for _, data := range globalVars.AllContact.Data {
				fmt.Printf("%3v | %-10v | %-20v | %-15v | %-15v | %-10v | %-5v | %-20v \n", data.Id, data.PersonID, data.Tel, data.Cellphone, data.Description, data.JobID, "JobName", "Gender.Name")
			}
			ShowMenu()
		case FindOneContactById:
			var id uint
			fmt.Println("Please insert your contact ID:")
			fmt.Scanln(&id)
			isFound, result := contact.FindContactByID(id)
			if isFound.State {
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
				insertContactID, jobID, personID uint
				tel, cellphone, description      string
			)
			fmt.Println("Please insert your contact id you want to edit:")
			fmt.Scanln(&insertContactID)
			isFound, result := contact.FindContactByID(insertContactID)
			if isFound.State {
				//TODO::: CleanPrintingFoundedResult
				fmt.Println(result)
			} else {
				fmt.Println("not found")
				ShowMenu()
			}
			fmt.Println("new person ID:")
			fmt.Scanln(&personID)
			fmt.Println("new tel:")
			fmt.Scanln(&tel)
			fmt.Println("new cellphone:")
			fmt.Scanln(&cellphone)
			fmt.Println("new job id:")
			fmt.Scanln(&jobID)
			fmt.Println("new description:")
			fmt.Scanln(&description)
			editedContact := protocol.Contact{
				PersonID:    personID,
				Tel:         tel,
				Cellphone:   cellphone,
				JobID:       jobID,
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
				var status protocol.ResponseStatus
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
			var inputCity, ariaCode string
			fmt.Println("insert city name:")
			fmt.Scanln(&inputCity)
			fmt.Println("insert aria code:")
			fmt.Scanln(&ariaCode)
			if city.NewCity(protocol.City{
				Name:     inputCity,
				AriaCode: ariaCode,
			}).State {
				fmt.Println("New city added")
			}
			ShowMenu()
		case ListOfCities:
			dataJSON, _ := json.MarshalIndent(globalVars.AllCity.CityData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range globalVars.AllCity.CityData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		case EditCityById:
			var inputID uint
			var inputName, ariaCode string
			fmt.Println("insert city id:")
			fmt.Scanln(&inputID)
			fmt.Println("insert new city name:")
			fmt.Scanln(&inputName)
			fmt.Println("insert new aria code:")
			fmt.Scanln(&inputName)
			if city.EditCityByID(inputID, protocol.City{
				Name:     inputName,
				AriaCode: ariaCode,
			}).State {
				fmt.Println("City changed ...")
			}
			ShowMenu()
		case DeleteCityById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting city(ies) ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
				var deleteIDS string
				fmt.Println("insert your city id(s) that you want to delete, for more than one, separate id's by ',':")
				fmt.Scanln(&deleteIDS)
				idPack := strings.Split(deleteIDS, ",")
				var idPackInt []uint
				for _, i := range idPack {
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					idPackInt = append(idPackInt, uint(j))
				}
				resNums := city.DeleteCity(idPackInt)
				fmt.Printf("%v city(ies) has been deleted", resNums)
			}
			ShowMenu()
		case InsertNewJob:
			var inputJob string
			fmt.Println("insert job name:")
			fmt.Scanln(&inputJob)
			if job.NewJob(protocol.Job{
				Name:                inputJob,
				LocationID:          0,
				Description:         "",
				BasicPaymentPerHour: 0,
			}).State {
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
			if job.EditJobInfoByID(inputID, protocol.Job{
				Name:                inputName,
				LocationID:          0,
				Description:         "",
				BasicPaymentPerHour: 0,
			}).State {
				fmt.Println("Job changed ...")
			}
			ShowMenu()
		case DeleteJobById:
			var confirmDel string
			fmt.Println("*** important, be careful, you are deleting job(s) ***")
			fmt.Println("do you want to continue? (yes or no)")
			fmt.Scanln(&confirmDel)
			if strings.ToLower(confirmDel) == YES {
				var deleteIDS string
				fmt.Println("insert your job id(s) that you want to delete, for more than one, separate id's by ',':")
				fmt.Scanln(&deleteIDS)
				idPack := strings.Split(deleteIDS, ",")
				var idPackInt []uint
				for _, i := range idPack {
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					idPackInt = append(idPackInt, uint(j))
				}
				resNums := job.DeleteJobByID(idPackInt)
				fmt.Printf("%v city(ies) has been deleted", resNums)
			}
			ShowMenu()
		case ListOfJob:
			dataJSON, _ := json.MarshalIndent(globalVars.AllJob.JobData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range globalVars.AllJob.JobData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		case InsertNewSex:
			var insertName string
			fmt.Println("Insert sex name")
			fmt.Scanln(&insertName)
			if sex.NewSex(insertName).State {
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
			if sex.EditSexByID(insertID, insertName).State {
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
				if sex.DeleteSexByID(insertID).State {
					fmt.Println("sex info deleted")
				} else {
					fmt.Println("some thing wrong")
				}
			}
			ShowMenu()
		case ListOfSex:
			dataJSON, _ := json.MarshalIndent(globalVars.AllSex.SexData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range globalVars.AllSex.SexData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		//something wrong:
		case RESET:
			RunApp()
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}

//TODO::: TalkToOmidHekayatiAboutSendDataInputValidationToClientForToCheckBeforeRequestToServer(LikeYII2)
