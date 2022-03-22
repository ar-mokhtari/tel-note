package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/fillSampleData"
	"tel-note/services/general"
	"tel-note/services/globalVars"
	"tel-note/services/identity"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
	"time"
)

var (
	separator  = strings.Repeat("-", 20)
	separator7 = strings.Repeat(separator, 7)
	scanner    = bufio.NewScanner(os.Stdin)
)

func ShowMenu() (output string) {
	output += fmt.Sprintln(separator7)
	output += fmt.Sprintln("*** Main menu: ***\n", "Please select:")
	for _, data := range GlobalMenu.Group {
		output += fmt.Sprintf("%-3s %s %3s \n", separator, data.GroupName, separator)
		for _, detail := range data.Row {
			output += fmt.Sprintf("%s			|	%s\n", detail.Title, detail.Description)
		}
	}
	output += fmt.Sprintln(separator7)
	return output
}

func RunMenu() {
	//TODO: CompleteAllEntityFeature(Person|Job|Sex|City|Contact)
	for {
		var userInput string
		fmt.Scanln(&userInput)
		//regulator actions only:
		userInput = strings.ToUpper(userInput)
		var notSelected = false
		switch identity.IsRegulator {
		case true:
			switch userInput {
			case InsertSomeSamplesData:
				fillSampleData.FillSimpleDataInMainData()
				fmt.Println(ShowMenuWarn)
			case PrintAllData:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range contact.GetPool.GetContacts() {
					_, person := person.FindPersonByID(data.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(data.JobID)
					_, city := city.FindCityByID(job.LocationID)
					if data.CellphoneCollection == nil {
						data.CellphoneCollection = append(data.CellphoneCollection, protocol.CellPhone{})
					}
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.CellphoneCollection[0].CellPhone, job.LocationID, city.Name, data.Description)
				}
				fmt.Println(separator7)
				fmt.Println("Top 10 City Data:")
				if cityCollection := city.GetCities(); cityCollection != nil {
					for _, data := range cityCollection[:10] {
						fmt.Printf("%3v | %-15v \n", data.Id, data.Name)
					}
					fmt.Println(separator7)
					fmt.Println("Top 10 Country Data:")
					fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
						"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
					fmt.Println()
				}
				if countryData := country.GetCountry(); countryData != nil {
					for _, data := range country.GetCountry()[:10] {
						fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
							data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
					}
				}
				fmt.Println(separator7)
				fmt.Println("Job Data:")
				for _, data := range job.GetJobs() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("sex Data:")
				for _, data := range sex.GetSex() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("person Data:")
				for _, data := range person.GetPersons() {
					fmt.Printf("%3v | %-15s \t %30s \n", data.Id, data.FirstName, data.LastName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer group Data:")
				fmt.Printf("%3v | %-15s   \n", "Id", "Groupname")
				fmt.Println()
				for _, group := range customer.GetCustomerGroup() {
					fmt.Printf("%3v | %-25s  \n",
						group.GroupID, group.GroupName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer group relation Data:")
				fmt.Printf("%3v | %-25s | %-13v | %-45s   \n", "Id", "Groupname", "CustomerID", "FullName")
				fmt.Println()
				for _, relation := range customer.GetCustomerGroupRelation() {
					customerObject := customer.FindCustomerByID(relation.CustomerID)
					_, personObject := person.FindPersonByID(customerObject.PersonID)
					groupObject := customer.FindGroupByID(relation.GroupID)
					fmt.Printf("%3v | %-25s | %-13v | %-45s \n",
						relation.ID, groupObject.GroupName, relation.CustomerID, personObject.FirstName+" "+personObject.LastName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"Id", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range globalVars.CustomerMapStore {
					_, person := person.FindPersonByID(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(ShowMenuWarn)
			case CheckIranNationalCode:
				var nationalCode string
				fmt.Println("insert new NationalCode")
				fmt.Scanln(&nationalCode)
				if !general.CheckIranNationalCode(nationalCode) {
					fmt.Println("Invalid national code")
					fmt.Println(ShowMenuWarn)
				} else {
					fmt.Printf("OK")
				}
				fmt.Println(ShowMenuWarn)
			default:
				notSelected = true
			}
			switch userInput {
			case ShowMenuList:
				fmt.Println(ShowMenu())
			case NewContactRecord:
				var (
					tel, cellphone  string
					jobID, personID uint
				)
				//input person
				fmt.Println("Please inter personID:")
				fmt.Scanln(&personID)
				//input job info
				fmt.Println("Please inter job ID:")
				fmt.Scanln(&jobID)
				//input tel
				fmt.Println("Please inter tel:")
				fmt.Scanln(&tel)
				//input cellphone
				fmt.Println("Please inter cellphone: ")
				fmt.Println("sample format: 'cellphone1,description1,cellphone2,description2,...'")
				fmt.Scanln(&cellphone)
				var cellPack []protocol.CellPhone
				cellPhoneCollection := strings.Split(cellphone, ",")
				if (len(cellPhoneCollection))%2 == 0 {
					var tempCell protocol.CellPhone
					for index, data := range cellPhoneCollection {
						if (index+1)%2 != 0 {
							tempCell.CellPhone = data
						} else {
							tempCell.Description = data
							cellPack = append(cellPack, tempCell)
							tempCell = protocol.CellPhone{}
						}
					}
				}
				//input description
				fmt.Println("Please inter description:")
				scanner.Scan()
				contact.PoolContact.NewContact(
					protocol.Contact{
						PersonID:            personID,
						JobID:               jobID,
						Tel:                 tel,
						CellphoneCollection: cellPack,
						Description:         scanner.Text(),
					})
				fmt.Println(">> New contact added done <<")
				fmt.Println(ShowMenuWarn)
			case ListOfContact:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range contact.GetPool.GetContacts() {
					_, person := person.FindPersonByID(data.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(data.JobID)
					_, city := city.FindCityByID(job.LocationID)
					if data.CellphoneCollection == nil {
						data.CellphoneCollection = append(data.CellphoneCollection, protocol.CellPhone{})
					}
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, (data.CellphoneCollection)[0].CellPhone, job.LocationID, city.Name, data.Description)
				}
				fmt.Println(ShowMenuWarn)
			case FindOneContactById:
				var id uint
				fmt.Println("Please insert your contact ID:")
				fmt.Scanln(&id)
				isFound, result := contact.FindByIDPool.FindContactByID(id)
				if isFound.State {
					fmt.Println(separator7)
					fmt.Println("Contact Data:")
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
						"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
					fmt.Println("")
					_, person := person.FindPersonByID(result.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(result.JobID)
					_, city := city.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.CellphoneCollection, job.LocationID, city.Name, result.Description)
				} else {
					fmt.Println("not found")
				}
				fmt.Println(ShowMenuWarn)
			case FindContactContainingSomeCharacter:
				var insertChar string
				fmt.Println("insert character(s):")
				fmt.Scanln(&insertChar)
				_, resultData, resultCount := contact.FindByCharPool.FindContactByChar(insertChar)
				if resultCount == 0 {
					fmt.Println("not found")
				} else {
					fmt.Println(resultCount, "record(s) found")
					fmt.Println(separator7)
					fmt.Println("Contact Data:")
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
						"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
					fmt.Println("")
					for _, data := range resultData {
						_, person := person.FindPersonByID(data.PersonID)
						genderID := person.GenderID
						gender, _ := sex.FindSexByID(uint8(genderID))
						_, job := job.FindJobByID(data.JobID)
						_, city := city.FindCityByID(job.LocationID)
						fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
							data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.CellphoneCollection, job.LocationID, city.Name, data.Description)
					}
				}
				fmt.Println(ShowMenuWarn)
			case FindAndEditContactByContactId:
				var (
					insertContactID, jobID, personID uint
					tel, cellphone                   string
				)
				fmt.Println("Please insert your contact id you want to edit:")
				fmt.Scanln(&insertContactID)
				isFound, result := contact.FindByIDPool.FindContactByID(insertContactID)
				if isFound.State {
					fmt.Println(separator7)
					fmt.Println("Contact Data:")
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
						"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
					fmt.Println("")
					_, person := person.FindPersonByID(result.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(result.JobID)
					_, city := city.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.CellphoneCollection, job.LocationID, city.Name, result.Description)
				} else {
					fmt.Println("not found")
					fmt.Println(ShowMenuWarn)
				}
				fmt.Println("new person ID:")
				fmt.Scanln(&personID)
				fmt.Println("new tel:")
				fmt.Scanln(&tel)
				fmt.Println("new cellphone:")
				fmt.Scanln(&cellphone)
				var cellPack []protocol.CellPhone
				if len(cellphone)%2 != 0 {
					cellPhoneCollection := strings.Split(cellphone, ",")
					var tempCell protocol.CellPhone
					for index, data := range cellPhoneCollection {
						if (index + 1%2) != 0 {
							tempCell.CellPhone = data
						} else {
							tempCell.Description = data
							cellPack = append(cellPack, tempCell)
							tempCell = protocol.CellPhone{}
						}
					}
				}
				fmt.Println("new job id:")
				fmt.Scanln(&jobID)
				fmt.Println("new description:")
				scanner.Scan()
				editedContact := protocol.Contact{
					PersonID:            personID,
					JobID:               jobID,
					Tel:                 tel,
					CellphoneCollection: cellPack,
					Description:         scanner.Text(),
				}
				state := contact.EditPool.EditContactByID(editedContact, insertContactID)
				if state.State {
					fmt.Printf("contact no %v updated", insertContactID)
				}
				fmt.Println(ShowMenuWarn)
			case DeleteContactById:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting a contact ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == YES {
					var deleteID uint
					fmt.Println("insert your contact id that you want to delete:")
					fmt.Scanln(&deleteID)
					status := contact.PoolDelContactID.DeleteContactByID(deleteID)
					switch status.State {
					case true:
						fmt.Printf("contact with id number:  %d deleted.", deleteID)
					case false:
						fmt.Println("something wrong")
					}
				}
				fmt.Println(ShowMenuWarn)
			case DeleteAllContacts:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting all of contacts ***")
				fmt.Println("are you sure? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == YES {
					resultStatus := contact.PoolDelAllContact.DeleteAll()
					fmt.Println(resultStatus.String)
				}
				fmt.Println(ShowMenuWarn)
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
							status = *contact.PoolDelContactID.DeleteContactByID(uint(uintDelID))
							fmt.Println(status.String)
						}
					}
				}
				fmt.Println(ShowMenuWarn)
			case NewCustomerGroup:
				var groupName string
				fmt.Println("insert group name:")
				scanner.Scan()
				groupName = scanner.Text()
				customer.NewGroup(groupName)
				fmt.Println(ShowMenuWarn)
			case NewCustomerGRelation:
				var customerID, groupID uint
				fmt.Println("insert customerID:")
				fmt.Scanln(&customerID)
				fmt.Println("insert groupID:")
				fmt.Scanln(&groupID)
				customer.NewRelation(customerID, groupID)
				fmt.Println(ShowMenuWarn)
			case CustomerGroupList:
				fmt.Println(separator7)
				fmt.Println("Customer group Data:")
				fmt.Printf("%3v | %-15s   \n", "Id", "Groupname")
				fmt.Println()
				for _, group := range customer.GetCustomerGroup() {
					fmt.Printf("%3v | %-25s  \n",
						group.GroupID, group.GroupName)
				}
				fmt.Println(ShowMenuWarn)
			case CustomerGroupRelationList:
				fmt.Println(separator7)
				fmt.Println("Customer group relation Data:")
				fmt.Printf("%3v | %-25s | %-13v | %-45s   \n", "Id", "Groupname", "CustomerID", "FullName")
				fmt.Println()
				for _, relation := range customer.GetCustomerGroupRelation() {
					customerObject := customer.FindCustomerByID(relation.CustomerID)
					_, personObject := person.FindPersonByID(customerObject.PersonID)
					groupObject := customer.FindGroupByID(relation.GroupID)
					fmt.Printf("%3v | %-25s | %-13v | %-45s \n",
						relation.ID, groupObject.GroupName, relation.CustomerID, personObject.FirstName+" "+personObject.LastName)
				}
				fmt.Println(ShowMenuWarn)
			case FindCustomerByGroupID:
				var groupID uint
				fmt.Println("insert group id:")
				fmt.Scanln(&groupID)
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Println(">>> ", customer.FindGroupByID(groupID).GroupName, " <<<")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"CID", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range customer.FindCustomerByGroupID(groupID).CustomerData {
					_, person := person.FindPersonByID(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(ShowMenuWarn)
			case NewCustomer:
				var (
					personID uint
				)
				fmt.Println("insert person id")
				fmt.Scanln(&personID)
				fmt.Println("insert customer description")
				scanner.Scan()
				desc := scanner.Text()
				customer.AddCustomer(protocol.Customer{
					PersonID:    personID,
					Description: desc,
				})
				fmt.Println(ShowMenuWarn)
			case ListOfCustomer:
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"Id", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range globalVars.CustomerMapStore {
					_, person := person.FindPersonByID(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(ShowMenuWarn)
			case EditCustomerByCustomerId:
				var customerID, personID uint
				var description string
				fmt.Println("insert customer id")
				fmt.Scanln(&customerID)
				fmt.Println("insert new person id")
				fmt.Scanln(&personID)
				fmt.Println("insert new description")
				fmt.Scanln(&description)
				customer.EditCustomer(customerID, protocol.Customer{
					PersonID:    personID,
					Description: description,
				})
				fmt.Println(ShowMenuWarn)
			case DeleteCustomerById:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting customer(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == YES {
					var deleteIDS string
					fmt.Println("insert your customer id(s) that you want to delete, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := strings.Split(deleteIDS, ",")
					var idPackInt []uint
					for _, i := range idPack {
						j := general.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
						customer.DeleteCustomerById(uint(j))
					}
					fmt.Printf("%v customer(s) has been deleted", idPackInt)
				}
				fmt.Println(ShowMenuWarn)
			case DeleteAllCustomers:
				globalVars.CustomerMapStore = nil
				fmt.Println("all customers deleted")
				fmt.Println(ShowMenuWarn)
			case NewPerson:
				var (
					FirstName       string
					LastName        string
					DOB             string
					BirthLocationID uint
					GenderID        uint
					NationalCode    string
					description     string
				)
				fmt.Println("insert new FirstName")
				scanner.Scan()
				FirstName = scanner.Text()
				fmt.Println("insert new LastName")
				scanner.Scan()
				LastName = scanner.Text()
				fmt.Println("insert new DOB (format:  20060102)")
				fmt.Scanln(&DOB)
				dboTime, _ := time.Parse("20060102", DOB)
				fmt.Println("insert new BirthLocationID")
				fmt.Scanln(&BirthLocationID)
				fmt.Println("insert new GenderID")
				fmt.Scanln(&GenderID)
				fmt.Println("insert new NationalCode")
				fmt.Scanln(&NationalCode)
				if !general.CheckIranNationalCode(NationalCode) {
					fmt.Println("Invalid national code")
					fmt.Println(ShowMenuWarn)
				}
				fmt.Println("insert new Description")
				scanner.Scan()
				description = scanner.Text()
				if person.NewPerson(protocol.Person{
					FirstName:       FirstName,
					LastName:        LastName,
					DOB:             dboTime,
					BirthLocationID: BirthLocationID,
					GenderID:        GenderID,
					NationalCode:    NationalCode,
					Description:     description,
				}).State {
					fmt.Println("New person added")
				}
				//some code
				fmt.Println(ShowMenuWarn)
			case ListOfPerson:
				fmt.Println(separator7)
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				//TODO::: decide to:
				//create a method called "GetPersons" to responds updated data live
				//use a global variable and after every event have to updated it
				for _, data := range person.GetPersons() /*person.GetPersons().PersonData*/ {
					genderID := data.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, city := city.FindCityByID(data.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
				}
				fmt.Println(ShowMenuWarn)
			case FindOnePersonById:
				var personID uint
				fmt.Println("insert person id")
				fmt.Scanln(&personID)
				if status, data := person.FindPersonByID(personID); status.State {
					fmt.Println(separator7)
					fmt.Println("Person Data:")
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
					fmt.Println("")
					genderID := data.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, city := city.FindCityByID(data.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
				} else {
					fmt.Println("not found")
				}
				fmt.Println(ShowMenuWarn)
			case FindPersonContainingSomeCharacter:
				fmt.Println("insert character")
				scanner.Scan()
				if state, data := person.FindPersonByChar(scanner.Text()); state.State {
					fmt.Println(separator7)
					fmt.Println(len(data), "record(s) found")
					fmt.Println("Person Data:")
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
					fmt.Println("")
					for _, personData := range data {
						genderID := personData.GenderID
						gender, _ := sex.FindSexByID(uint8(genderID))
						_, city := city.FindCityByID(personData.BirthLocationID)
						fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
							personData.Id, personData.FirstName, personData.LastName, gender.Name, personData.BirthLocationID, city.Name, (personData.DOB).String()[0:10], personData.Description)
					}
				}
				fmt.Println(ShowMenuWarn)
			case FindAndEditPersonByPersonId:
				fmt.Println("insert person id to (find and) edit")
				var personID uint
				fmt.Scanln(&personID)
				if status, data := person.FindPersonByID(personID); status.State {
					fmt.Println(separator7)
					fmt.Println("Person Data:")
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
					fmt.Println("")
					genderID := data.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, city := city.FindCityByID(data.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-4v | %-12v | %-13v |  %-3v  \n",
						data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
					fmt.Println("insert FirstName")
					scanner.Scan()
					FirstName := scanner.Text()
					fmt.Println("insert LastName")
					scanner.Scan()
					LastName := scanner.Text()
					fmt.Println("insert DOB")
					scanner.Scan()
					DOB := scanner.Text()
					dboTime, _ := time.Parse("20060102", DOB)
					fmt.Println("insert BirthLocationID")
					scanner.Scan()
					BirthLocationID := scanner.Text()
					BirthLocationIDUINT, _ := strconv.ParseUint(BirthLocationID, 10, 8)
					fmt.Println("insert GenderID")
					var confirmDel string
					fmt.Println("*** important, be careful, you are deleting person(s) ***")
					fmt.Println("do you want to continue? (yes or no)")
					fmt.Scanln(&confirmDel)
					if strings.ToLower(confirmDel) == YES {
						var deleteIDS string
						fmt.Println("insert your person id(s) that you want to delete, separate id's by ',':")
						fmt.Scanln(&deleteIDS)
						idPack := strings.Split(deleteIDS, ",")
						var idPackInt []uint
						for _, i := range idPack {
							j := general.StrToUint(i)
							idPackInt = append(idPackInt, uint(j))
						}
						resNums := person.DeletePerson(idPackInt)
						fmt.Printf("%v person(s) has been deleted", resNums)
					}
					fmt.Println(ShowMenuWarn)
					scanner.Scan()
					GenderID := scanner.Text()
					GenderIDUINT, _ := strconv.ParseUint(GenderID, 10, 8)
					fmt.Println("insert NationalCode")
					scanner.Scan()
					NationalCode := scanner.Text()
					fmt.Println("insert Description")
					scanner.Scan()
					Description := scanner.Text()
					newPerson := protocol.Person{
						Id:              personID,
						FirstName:       FirstName,
						LastName:        LastName,
						DOB:             dboTime,
						BirthLocationID: uint(BirthLocationIDUINT),
						GenderID:        uint(GenderIDUINT),
						NationalCode:    NationalCode,
						Description:     Description,
					}
					if person.EditPerson(newPerson.Id, newPerson).State {
						fmt.Printf("person no %v edited", personID)
					} else {
						fmt.Println("some thing wrong")
					}
					fmt.Println(ShowMenuWarn)
				}
				fmt.Println("not found")
				fmt.Println(ShowMenuWarn)
			case DeletePersonById:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting person(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == YES {
					var deleteIDS string
					fmt.Println("insert your person id(s) that you want to delete, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := strings.Split(deleteIDS, ",")
					var idPackInt []uint
					for _, i := range idPack {
						j := general.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := person.DeletePerson(idPackInt)
					fmt.Printf("%v person(s) has been deleted", resNums)
				}
				fmt.Println(ShowMenuWarn)
			case DeleteMultiPersonByIds:
				//some code
				fmt.Println(ShowMenuWarn)
			case DeleteAllPersons:
				//some code
				fmt.Println(ShowMenuWarn)
			case InsertNewCity:
				var inputCity, ariaCode string
				var lat, lng float64
				fmt.Println("insert city name:")
				scanner.Scan()
				inputCity = scanner.Text()
				fmt.Println("insert aria code:")
				fmt.Scanln(&ariaCode)
				fmt.Println("insert lat:")
				fmt.Scanln(&lat)
				fmt.Println("insert lng:")
				fmt.Scanln(&lng)
				if city.NewCity(protocol.City{
					Name:     inputCity,
					AriaCode: ariaCode,
					Lat:      lat,
					Lng:      lng,
				}).State {
					fmt.Println("New city added")
				}
				fmt.Println(ShowMenuWarn)
			case ListOfCities:
				dataJSON, _ := json.MarshalIndent(city.GetCities(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range city.GetCities() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(ShowMenuWarn)
			case EditCityById:
				var inputID uint
				var inputName, ariaCode string
				fmt.Println("insert city id:")
				fmt.Scanln(&inputID)
				fmt.Println("insert new city name:")
				scanner.Scan()
				inputName = scanner.Text()
				fmt.Println("insert new aria code:")
				fmt.Scanln(&ariaCode)
				if city.EditCityByID(inputID, protocol.City{
					Name:     inputName,
					AriaCode: ariaCode,
				}).State {
					fmt.Println("City changed ...")
				}
				fmt.Println(ShowMenuWarn)
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
						j := general.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := city.DeleteCity(idPackInt)
					fmt.Printf("%v city(ies) has been deleted", resNums)
				}
				fmt.Println(ShowMenuWarn)
			case CallDistanceTimeTwoCities:
				var firstCityCode, secondCityCode uint
				fmt.Println("insert first city id:")
				fmt.Scanln(&firstCityCode)
				fmt.Println("insert second city id:")
				fmt.Scanln(&secondCityCode)
				firstStatus, dataFirstCity := city.FindCityByID(firstCityCode)
				secondStatus, dataSecondCity := city.FindCityByID(secondCityCode)
				if firstStatus.State && secondStatus.State {
					result, state := city.ServiceHandler.CallDistanceTimeTwoCities(dataFirstCity, dataSecondCity)
					if state.State {
						fmt.Printf(
							"Time duration with online traffic is: %v hours ", result[0]/3600)
						if minutes := result[0] % 3600; minutes != 0 {
							fmt.Printf("and %v min", minutes/60)
						}
						fmt.Printf("\nDistance is: %v km\n", result[1]/1000)
					} else {
						fmt.Println("method send nothing")
					}
				} else {
					fmt.Println("not found")
				}
				fmt.Println(ShowMenuWarn)
			case AddCountry:
				var newCountryName string
				fmt.Println("insert new country name")
				scanner.Scan()
				newCountryName = scanner.Text()
				var newShortName string
				fmt.Println("insert new country newShortName")
				fmt.Scanln(&newShortName)
				var prePhoneCode uint
				fmt.Println("insert new country prePhoneCode")
				fmt.Scanln(&prePhoneCode)
				var capitalID uint
				fmt.Println("insert new country capitalID")
				fmt.Scanln(&capitalID)
				country.NewCountry(protocol.Country{
					Name:         newCountryName,
					ShortName:    newShortName,
					PrePhoneCode: prePhoneCode,
					CapitalID:    capitalID,
				})
				fmt.Println(ShowMenuWarn)
			case EditCountry:
				var countryID uint
				fmt.Println("insert country id")
				fmt.Scanln(&countryID)
				var newCountryName string
				fmt.Println("insert new country name")
				scanner.Scan()
				newCountryName = scanner.Text()
				var newShortName string
				fmt.Println("insert new country newShortName")
				fmt.Scanln(&newShortName)
				var prePhoneCode uint
				fmt.Println("insert new country prePhoneCode")
				fmt.Scanln(&prePhoneCode)
				var capitalID uint
				fmt.Println("insert new country capitalID")
				fmt.Scanln(&capitalID)
				country.EditCountry(protocol.Country{
					ID:           countryID,
					Name:         newCountryName,
					ShortName:    newShortName,
					PrePhoneCode: prePhoneCode,
					CapitalID:    capitalID,
				})
				fmt.Println(ShowMenuWarn)
			case DeleteCountry:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting country(ies) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == YES {
					var deleteIDS string
					fmt.Println("insert your country id(s) that you want to delete, for more than one, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := strings.Split(deleteIDS, ",")
					var idPackInt []uint
					for _, i := range idPack {
						j := general.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := country.DeleteCountry(idPackInt)
					fmt.Printf("%v country(ies) has been deleted", resNums)
				}
				fmt.Println(ShowMenuWarn)
			case FindCountryByChar:
				println("insert character")
				scanner.Scan()
				insertChar := scanner.Text()
				fmt.Println(separator7)
				fmt.Println("Country Data:")
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
					"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
				fmt.Println()
				for _, data := range country.FindCountryByChar(insertChar) {
					fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
						data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
				}
				fmt.Println(ShowMenuWarn)
			case CountriesList:
				fmt.Println(separator7)
				fmt.Println("Country Data:")
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
					"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
				fmt.Println()
				if country.GetCountry() != nil {
					for _, data := range country.GetCountry() {
						fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
							data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
					}
				}
				fmt.Println(ShowMenuWarn)
			case InsertNewJob:
				var (
					Name                string
					LocationID          uint64
					Description         string
					BasicPaymentPerHour uint64
				)
				fmt.Println("insert new Name")
				scanner.Scan()
				Name = scanner.Text()
				fmt.Println("insert new LocationID")
				scanner.Scan()
				LocationID, _ = strconv.ParseUint(scanner.Text(), 10, 8)
				fmt.Println("insert new Description")
				scanner.Scan()
				Description = scanner.Text()
				fmt.Println("insert new BasicPaymentPerHour")
				scanner.Scan()
				BasicPaymentPerHour, _ = strconv.ParseUint(scanner.Text(), 10, 8)
				if job.NewJob(protocol.Job{
					Name:                Name,
					LocationID:          uint(LocationID),
					Description:         Description,
					BasicPaymentPerHour: uint(BasicPaymentPerHour),
				}).State {
					fmt.Println("New job added")
				}
				fmt.Println(ShowMenuWarn)
			case EditJobById:
				var (
					inputID             uint
					Name                string
					LocationID          uint64
					Description         string
					BasicPaymentPerHour uint64
				)
				fmt.Println("insert job id:")
				fmt.Scanln(&inputID)
				fmt.Println("insert new Name")
				scanner.Scan()
				Name = scanner.Text()
				fmt.Println("insert new LocationID")
				scanner.Scan()
				LocationID, _ = strconv.ParseUint(scanner.Text(), 10, 8)
				fmt.Println("insert new Description")
				scanner.Scan()
				Description = scanner.Text()
				fmt.Println("insert new BasicPaymentPerHour")
				scanner.Scan()
				BasicPaymentPerHour, _ = strconv.ParseUint(scanner.Text(), 10, 8)
				if job.EditJobInfoByID(inputID, protocol.Job{
					Name:                Name,
					LocationID:          uint(LocationID),
					Description:         Description,
					BasicPaymentPerHour: uint(BasicPaymentPerHour),
				}).State {
					fmt.Println("Job changed ...")
				}
				fmt.Println(ShowMenuWarn)
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
						j := general.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := job.DeleteJobByID(idPackInt)
					fmt.Printf("%v job(s) has been deleted", resNums)
				}
				fmt.Println(ShowMenuWarn)
			case ListOfJob:
				dataJSON, _ := json.MarshalIndent(job.GetJobs(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range job.GetJobs() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(ShowMenuWarn)
			case InsertNewSex:
				var insertName string
				fmt.Println("Insert sex name")
				fmt.Scanln(&insertName)
				if sex.NewSex(
					protocol.Sex{
						Name: insertName,
					},
				).State {
					fmt.Println("New sex info added")
				}
				fmt.Println(ShowMenuWarn)
			case EditSex:
				var insertID uint8
				var insertName string
				fmt.Println("Insert sex id:")
				fmt.Scanln(&insertID)
				fmt.Println("Insert new name:")
				fmt.Scanln(&insertName)
				if sex.EditSexByID(protocol.Sex{
					Id:   insertID,
					Name: insertName,
				}).State {
					fmt.Println("sex info updated")
				} else {
					fmt.Println("some thing wrong")
				}
				fmt.Println(ShowMenuWarn)
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
				fmt.Println(ShowMenuWarn)
			case ListOfSex:
				dataJSON, _ := json.MarshalIndent(sex.GetSex(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range sex.GetSex() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(ShowMenuWarn)
			//something wrong:
			case RESET:
				globalVars.DeleteAllGlobalVars()
				RunCli()
			default:
				if !identity.IsRegulator || notSelected {
					fmt.Println("bad input, please insert character of menu list")
				}
			}
		}
	}
}

//TODO::: TalkToOmidHekayatiAboutSendDataInputValidationToClientForToCheckBeforeRequestToServer(LikeYII2)
