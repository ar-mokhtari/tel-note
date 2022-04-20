package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/lib/validator"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/globalVars"
	"tel-note/services/identity"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sampleData"
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
	for _, data := range env.GlobalMenu.Group {
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
			case env.InsertData:
				sampleData.FillData.FillSimpleData()
				fmt.Println(env.ShowMenuWarn)
			case env.GetData:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range contact.GetContact.Do() {
					person := person.FindPersonID.Do(data.PersonID)
					genderID := person.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, job := job.FindJobID.Do(data.JobID)
					_, city := city.FindCityID.FindCityByID(job.LocationID)
					if data.CellphoneCollection == nil {
						data.CellphoneCollection = append(data.CellphoneCollection, protocol.CellPhone{})
					}
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.CellphoneCollection[0].CellPhone, job.LocationID, city.Name, data.Description)
				}
				fmt.Println(separator7)
				fmt.Println("Top 10 City Data:")
				if cityCollection := city.GetCity.Do(); cityCollection != nil {
					for _, data := range cityCollection[:10] {
						fmt.Printf("%3v | %-15v \n", data.Id, data.Name)
					}
					fmt.Println(separator7)
					fmt.Println("Top 10 Country Data:")
					fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
						"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
					fmt.Println()
				}
				if countryData := country.GetCountry.Do(); countryData != nil {
					for _, data := range country.GetCountry.Do()[:10] {
						fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
							data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
					}
				}
				fmt.Println(separator7)
				fmt.Println("Job Data:")
				for _, data := range job.GetJob.Do() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("sex Data:")
				for _, data := range sex.GetSex.Do() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("person Data:")
				for _, data := range person.GetPerson.Do() {
					fmt.Printf("%3v | %-15s \t %30s \n", data.Id, data.FirstName, data.LastName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer group Data:")
				fmt.Printf("%3v | %-15s   \n", "Id", "Groupname")
				fmt.Println()
				for _, group := range customer.GetCustomerGroup.Do() {
					fmt.Printf("%3v | %-25s  \n",
						group.GroupID, group.GroupName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer group relation Data:")
				fmt.Printf("%3v | %-25s | %-13v | %-45s   \n", "Id", "Groupname", "CustomerID", "FullName")
				fmt.Println()
				for _, relation := range customer.GetCustomerGroupRelation.Do() {
					customerObject := customer.FindCustomerByID.Do(relation.CustomerID)
					personObject := person.FindPersonID.Do(customerObject.PersonID)
					groupObject := customer.FindGroupByID.Do(relation.GroupID)
					fmt.Printf("%3v | %-25s | %-13v | %-45s \n",
						relation.ID, groupObject.GroupName, relation.CustomerID, personObject.FirstName+" "+personObject.LastName)
				}
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"Id", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range globalVars.CustomerMapStore {
					person := person.FindPersonID.Do(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.CheckIranNationalCode:
				var nationalCode string
				fmt.Println("insert new NationalCode")
				fmt.Scanln(&nationalCode)
				if !validator.CheckNationalID(nationalCode) {
					fmt.Println("Invalid national code")
					fmt.Println(env.ShowMenuWarn)
				} else {
					fmt.Printf("OK")
				}
				fmt.Println(env.ShowMenuWarn)
			default:
				notSelected = true
			}
			switch userInput {
			case env.ShowMenuList:
				fmt.Println(ShowMenu())
			case env.ContactNew:
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
				contact.NewContact.Do(
					contact.NewContactRequest{
						PersonID:            personID,
						JobID:               jobID,
						Tel:                 tel,
						CellphoneCollection: cellPack,
						Description:         scanner.Text(),
					})
				fmt.Println(">> New contact added done <<")
				fmt.Println(env.ShowMenuWarn)
			case env.ContactList:
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range contact.GetContact.Do() {
					person := person.FindPersonID.Do(data.PersonID)
					genderID := person.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, job := job.FindJobID.Do(data.JobID)
					_, city := city.FindCityID.FindCityByID(job.LocationID)
					if data.CellphoneCollection == nil {
						data.CellphoneCollection = append(data.CellphoneCollection, protocol.CellPhone{})
					}
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, (data.CellphoneCollection)[0].CellPhone, job.LocationID, city.Name, data.Description)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.ContactFindId:
				var id uint
				fmt.Println("Please insert your contact ID:")
				fmt.Scanln(&id)
				isFound, result := contact.FindContactID.Do(id)
				if isFound.State {
					fmt.Println(separator7)
					fmt.Println("Contact Data:")
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
						"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
					fmt.Println("")
					person := person.FindPersonID.Do(result.PersonID)
					genderID := person.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, job := job.FindJobID.Do(result.JobID)
					_, city := city.FindCityID.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.CellphoneCollection, job.LocationID, city.Name, result.Description)
				} else {
					fmt.Println("not found")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.FindContactChar:
				var insertChar string
				fmt.Println("insert character(s):")
				fmt.Scanln(&insertChar)
				_, resultData, resultCount := contact.FindContactChar.Do(insertChar)
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
						person := person.FindPersonID.Do(data.PersonID)
						genderID := person.GenderID
						gender := sex.FindSexID.Do(genderID)
						_, job := job.FindJobID.Do(data.JobID)
						_, city := city.FindCityID.FindCityByID(job.LocationID)
						fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
							data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.CellphoneCollection, job.LocationID, city.Name, data.Description)
					}
				}
				fmt.Println(env.ShowMenuWarn)
			case env.EditContactId:
				var (
					insertContactID, jobID, personID uint
					tel, cellphone                   string
				)
				fmt.Println("Please insert your contact id you want to edit:")
				fmt.Scanln(&insertContactID)
				isFound, result := contact.FindContactID.Do(insertContactID)
				if isFound.State {
					fmt.Println(separator7)
					fmt.Println("Contact Data:")
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
						"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
					fmt.Println("")
					person := person.FindPersonID.Do(result.PersonID)
					genderID := person.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, job := job.FindJobID.Do(result.JobID)
					_, city := city.FindCityID.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.CellphoneCollection, job.LocationID, city.Name, result.Description)
				} else {
					fmt.Println("not found")
					fmt.Println(env.ShowMenuWarn)
				}
				fmt.Println("new person ID:")
				fmt.Scanln(&personID)
				fmt.Println("new tel:")
				fmt.Scanln(&tel)
				fmt.Println("new cellphone:")
				fmt.Scanln(&cellphone)
				var cellPack []protocol.CellPhone
				if len(cellphone)%2 != 0 {
					cellPhoneCollection := convertor.StrToSlice(cellphone)
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
				editedContact := contact.NewContactRequest{
					Id:                  insertContactID,
					PersonID:            personID,
					JobID:               jobID,
					Tel:                 tel,
					CellphoneCollection: cellPack,
					Address:             "",
					Description:         scanner.Text(),
				}
				err := contact.NewContact.Do(editedContact)
				if err != nil {
					fmt.Printf("contact no %v updated", insertContactID)
				} else {
					fmt.Println(err.Error())
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteContactId:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting a contact ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteID uint
					fmt.Println("insert your contact id that you want to delete:")
					fmt.Scanln(&deleteID)
					status := contact.DelContactID.Do(deleteID)
					switch status.State {
					case true:
						fmt.Printf("contact with id number:  %d deleted.", deleteID)
					case false:
						fmt.Println("something wrong")
					}
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteAllContacts:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting all of contacts ***")
				fmt.Println("are you sure? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					resultStatus := contact.DelAllContact.Do()
					fmt.Println(resultStatus.String)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteMultiContact:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting contact(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					var status protocol.ResponseStatus
					fmt.Println("insert your contact id(s) that you want to delete, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					for _, delID := range idPack {
						uintDelID, _ := strconv.ParseUint(delID, 10, 8)
						if uint(uintDelID) <= 0 {
							fmt.Println("not equal kind")
							break
						} else {
							status = *contact.DelContactID.Do(uint(uintDelID))
							fmt.Println(status.String)
						}
					}
				}
				fmt.Println(env.ShowMenuWarn)
			case env.NewCustomerGroup:
				var groupName string
				fmt.Println("insert group name:")
				scanner.Scan()
				groupName = scanner.Text()
				customer.NewGroup.Do(groupName)
				fmt.Println(env.ShowMenuWarn)
			case env.NewCustGRelation:
				var customerID, groupID uint
				fmt.Println("insert customerID:")
				fmt.Scanln(&customerID)
				fmt.Println("insert groupID:")
				fmt.Scanln(&groupID)
				customer.NewGrpRelation.Do(customerID, groupID)
				fmt.Println(env.ShowMenuWarn)
			case env.CustomerGroupList:
				fmt.Println(separator7)
				fmt.Println("Customer group Data:")
				fmt.Printf("%3v | %-15s   \n", "Id", "Groupname")
				fmt.Println()
				for _, group := range customer.GetCustomerGroup.Do() {
					fmt.Printf("%3v | %-25s  \n",
						group.GroupID, group.GroupName)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.CustGRelationList:
				fmt.Println(separator7)
				fmt.Println("Customer group relation Data:")
				fmt.Printf("%3v | %-25s | %-13v | %-45s   \n", "Id", "Groupname", "CustomerID", "FullName")
				fmt.Println()
				for _, relation := range customer.GetCustomerGroupRelation.Do() {
					customerObject := customer.FindCustomerByID.Do(relation.CustomerID)
					personObject := person.FindPersonID.Do(customerObject.PersonID)
					groupObject := customer.FindGroupByID.Do(relation.GroupID)
					fmt.Printf("%3v | %-25s | %-13v | %-45s \n",
						relation.ID, groupObject.GroupName, relation.CustomerID, personObject.FirstName+" "+personObject.LastName)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.FindCustGrpRelationByGrpID:
				var groupID uint
				fmt.Println("insert group id:")
				fmt.Scanln(&groupID)
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Println(">>> ", customer.FindGroupByID.Do(groupID).GroupName, " <<<")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"CID", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range customer.FindCustGrpRelationByGrpID.Do(groupID).CustomerData {
					person := person.FindPersonID.Do(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.NewCustomer:
				var (
					personID uint
				)
				fmt.Println("insert person id")
				fmt.Scanln(&personID)
				fmt.Println("insert customer description")
				scanner.Scan()
				desc := scanner.Text()
				readyToFill := customer.NewRequest{
					PersonID:    personID,
					Description: desc,
				}
				customer.AddCustomer.Do(readyToFill)
				fmt.Println(env.ShowMenuWarn)
			case env.CustomerList:
				fmt.Println(separator7)
				fmt.Println("Customer Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v |  %-23v  \n",
					"Id", "Customername", "CustomerFamily", "PID", "Creat@", "Update@", "Desc")
				fmt.Println()
				for ID, data := range globalVars.CustomerMapStore {
					person := person.FindPersonID.Do(data.PersonID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
						ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.EditCustomer:
				var customerID, personID uint
				var description string
				fmt.Println("insert customer id")
				fmt.Scanln(&customerID)
				fmt.Println("insert new person id")
				fmt.Scanln(&personID)
				fmt.Println("insert new description")
				fmt.Scanln(&description)
				err := customer.EditCustomer.Do(customerID, customer.EditRequest{
					PersonID:    personID,
					Description: description,
				})
				if err != nil {
					fmt.Println("not edited")
				} else {
					fmt.Println("edit successful")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteCustomer:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting customer(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					fmt.Println("insert your customer id(s) that you want to delete, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					var idPackInt []uint
					for _, i := range idPack {
						_, j := convertor.StrToUint(i)
						idPackInt = append(idPackInt, j)
						customer.DeleteCustomer.Do(idPackInt)
					}
					fmt.Printf("%v customer(s) has been deleted", idPackInt)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteAllCust:
				globalVars.CustomerMapStore = nil
				fmt.Println("all customers deleted")
				fmt.Println(env.ShowMenuWarn)
			case env.NewPerson:
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
				if !validator.CheckNationalID(NationalCode) {
					fmt.Println("Invalid national code")
					fmt.Println(env.ShowMenuWarn)
				}
				fmt.Println("insert new Description")
				scanner.Scan()
				description = scanner.Text()
				person.AddPerson.Do(person.AddRequest{
					FirstName:       FirstName,
					LastName:        LastName,
					DOB:             dboTime,
					BirthLocationID: BirthLocationID,
					GenderID:        GenderID,
					NationalCode:    NationalCode,
					Description:     description,
				})
				//some code
				fmt.Println(env.ShowMenuWarn)
			case env.ListOfPerson:
				fmt.Println(separator7)
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				for _, data := range person.GetPerson.Do() {
					genderID := data.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, city := city.FindCityID.FindCityByID(data.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.FindOnePersonById:
				var personID uint
				fmt.Println("insert person id")
				fmt.Scanln(&personID)
				data := person.FindPersonID.Do(personID)
				fmt.Println(separator7)
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				genderID := data.GenderID
				gender := sex.FindSexID.Do(genderID)
				_, city := city.FindCityID.FindCityByID(data.BirthLocationID)
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
				fmt.Println(env.ShowMenuWarn)
			case env.FindPersonChar:
				fmt.Println("insert character")
				scanner.Scan()
				data := person.FindCharPerson.Do(scanner.Text())
				fmt.Println(separator7)
				fmt.Println(len(data), "record(s) found")
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				for _, personData := range data {
					genderID := personData.GenderID
					gender := sex.FindSexID.Do(genderID)
					_, city := city.FindCityID.FindCityByID(personData.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						personData.Id, personData.FirstName, personData.LastName, gender.Name, personData.BirthLocationID, city.Name, (personData.DOB).String()[0:10], personData.Description)

				}
				fmt.Println(env.ShowMenuWarn)
			case env.EditPerson:
				fmt.Println("insert person id to (find and) edit")
				var personID uint
				fmt.Scanln(&personID)
				data := person.FindPersonID.Do(personID)
				fmt.Println(separator7)
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				genderID := data.GenderID
				gender := sex.FindSexID.Do(genderID)
				_, city := city.FindCityID.FindCityByID(data.BirthLocationID)
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
				if err := person.EditPerson.Do(newPerson.Id, person.EditRequest(newPerson)); err == nil {
					fmt.Printf("person no %v edited", personID)
				} else {
					fmt.Println("some thing wrong")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeletePerson:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting person(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					fmt.Println("insert your person id(s) that you want to delete, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					var idPackInt []uint
					for _, i := range idPack {
						_, j := convertor.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := person.DeletePerson.Do(idPackInt)
					fmt.Printf("%v person(s) has been deleted", resNums)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteMultiPerson:
				//some code
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteAllPersons:
				//some code
				fmt.Println(env.ShowMenuWarn)
			case env.InsertNewCity:
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
				err := city.NewCity.Do(city.NewCityRequest{
					Name:        inputCity,
					EnglishName: "",
					AriaCode:    ariaCode,
					Lat:         lat,
					Lng:         lng,
				})
				if err != nil {
					fmt.Println("New city added")
				} else {
					fmt.Println(err.Error())
				}
				fmt.Println(env.ShowMenuWarn)
			case env.ListOfCities:
				dataJSON, _ := json.MarshalIndent(city.GetCity.Do(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range city.GetCity.Do() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.EditCityById:
				var inputID uint
				var inputName, ariaCode string
				fmt.Println("insert city id:")
				fmt.Scanln(&inputID)
				fmt.Println("insert new city name:")
				scanner.Scan()
				inputName = scanner.Text()
				fmt.Println("insert new aria code:")
				fmt.Scanln(&ariaCode)
				var result city.EditCityRequest
				result.Id, result.Name, result.AriaCode =
					inputID, inputName, ariaCode
				if err := city.EditCity.Do(result); err == nil {
					fmt.Println("City changed ...")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteCityById:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting city(ies) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					fmt.Println("insert your city id(s) that you want to delete, for more than one, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					var idPackInt []uint
					for _, i := range idPack {
						_, j := convertor.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := city.DeleteCity.Do(idPackInt)
					fmt.Printf("%v city(ies) has been deleted", resNums)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.CallDistanceTimeTwoCities:
				var firstCityCode, secondCityCode uint
				fmt.Println("insert first city id:")
				fmt.Scanln(&firstCityCode)
				fmt.Println("insert second city id:")
				fmt.Scanln(&secondCityCode)
				firstStatus, dataFirstCity := city.FindCityID.FindCityByID(firstCityCode)
				secondStatus, dataSecondCity := city.FindCityID.FindCityByID(secondCityCode)
				if firstStatus.State && secondStatus.State {
					result, state := city.DistanceTimeService.Do(dataFirstCity, dataSecondCity)
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
				fmt.Println(env.ShowMenuWarn)
			case env.AddCountry:
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
				country.NewCountry.Do(country.NewRequest{
					Name:         newCountryName,
					ShortName:    newShortName,
					PrePhoneCode: prePhoneCode,
					CapitalID:    capitalID,
				})
				fmt.Println(env.ShowMenuWarn)
			case env.EditCountry:
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
				country.EditCountry.Do(country.EditRequest{
					ID:           countryID,
					Name:         newCountryName,
					ShortName:    newShortName,
					PrePhoneCode: prePhoneCode,
					CapitalID:    capitalID,
				})
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteCountry:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting country(ies) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					fmt.Println("insert your country id(s) that you want to delete, for more than one, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					var idPackInt []uint
					for _, i := range idPack {
						_, j := convertor.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums, _ := country.DeleteCountry.Do(idPackInt)
					fmt.Printf("%v country(ies) has been deleted", resNums)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.FindCountryByChar:
				println("insert character")
				scanner.Scan()
				insertChar := scanner.Text()
				fmt.Println(separator7)
				fmt.Println("Country Data:")
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
					"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
				fmt.Println()
				for _, data := range country.FindCountryChar.Do(insertChar) {
					fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
						data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
				}
				fmt.Println(env.ShowMenuWarn)
			case env.CountriesList:
				fmt.Println(separator7)
				fmt.Println("Country Data:")
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
					"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
				fmt.Println()
				if country.GetCountry.Do() != nil {
					for _, data := range country.GetCountry.Do() {
						fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
							data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
					}
				}
				fmt.Println(env.ShowMenuWarn)
			case env.InsertNewJob:
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
				job.NewJob.Do(job.NewRequest{
					Name:                Name,
					LocationID:          uint(LocationID),
					Description:         Description,
					BasicPaymentPerHour: uint(BasicPaymentPerHour),
				})
				fmt.Println("New job added")
				fmt.Println(env.ShowMenuWarn)
			case env.EditJobById:
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
				job.EditJob.Do(inputID, job.EditRequest{
					Name:                Name,
					LocationID:          uint(LocationID),
					Description:         Description,
					BasicPaymentPerHour: uint(BasicPaymentPerHour),
				})
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteJobById:
				var confirmDel string
				fmt.Println("*** important, be careful, you are deleting job(s) ***")
				fmt.Println("do you want to continue? (yes or no)")
				fmt.Scanln(&confirmDel)
				if strings.ToLower(confirmDel) == env.YES {
					var deleteIDS string
					fmt.Println("insert your job id(s) that you want to delete, for more than one, separate id's by ',':")
					fmt.Scanln(&deleteIDS)
					idPack := convertor.StrToSlice(deleteIDS)
					var idPackInt []uint
					for _, i := range idPack {
						_, j := convertor.StrToUint(i)
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := job.DeleteJob.Do(idPackInt)
					fmt.Printf("%v job(s) has been deleted", resNums)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.ListOfJob:
				dataJSON, _ := json.MarshalIndent(job.GetJob.Do(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range job.GetJob.Do() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(env.ShowMenuWarn)
			case env.InsertNewSex:
				var insertName string
				fmt.Println("Insert sex name")
				fmt.Scanln(&insertName)
				if sex.AddSex.Do(
					sex.AddRequest{
						Name: insertName,
					},
				) == nil {
					fmt.Println("New sex info added")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.EditSex:
				var insertID uint8
				var insertName string
				fmt.Println("Insert sex id:")
				fmt.Scanln(&insertID)
				fmt.Println("Insert new name:")
				fmt.Scanln(&insertName)
				if sex.EditSex.Do(sex.EditRequest{
					Id:   insertID,
					Name: insertName,
				}) == nil {
					fmt.Println("sex info updated")
				} else {
					fmt.Println("some thing wrong")
				}
				fmt.Println(env.ShowMenuWarn)
			case env.DeleteSex:
				fmt.Println("you are deleting some data, are you sure? (yes/no)")
				var confirmInsert string
				fmt.Scanln(&confirmInsert)
				if strings.ToLower(confirmInsert) == env.YES {
					var insertID uint8
					fmt.Println("insert sex id to delete:")
					fmt.Scanln(&insertID)
					if sex.DeleteSexByID.Do(insertID) == nil {
						fmt.Println("sex info deleted")
					} else {
						fmt.Println("some thing wrong")
					}
				}
				fmt.Println(env.ShowMenuWarn)
			case env.ListOfSex:
				dataJSON, _ := json.MarshalIndent(sex.GetSex.Do(), "", "  ")
				fmt.Println(string(dataJSON))
				fmt.Println(separator7)
				for _, data := range sex.GetSex.Do() {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(env.ShowMenuWarn)
			//something wrong:
			case env.RESET:
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
