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

func ShowMenu() {
	fmt.Println(separator7)
	fmt.Println("*** Main menu: ***\n", "Please select:")
	//Contact Menu
	fmt.Printf("%-3s %s %3s \n", separator, "Contact Menu", separator)
	fmt.Print(NewContactRecord, "			|	new record\n")
	fmt.Print(ListOfContact, "			|	list of contact\n")
	fmt.Print(FindOneContactById, "			|	find one contact by id\n")
	fmt.Print(FindContactContainingSomeCharacter, "			|	find contact, contain some character\n")
	fmt.Print(FindAndEditContactByContactId, "			|	find and edit contact by contact id\n")
	fmt.Print(DeleteContactById, "			|	delete contact by id\n")
	fmt.Print(DeleteMultiContactByIds, "			|	delete multi contact by id(s)\n")
	fmt.Print(DeleteAllContacts, "			|	delete all contacts\n")
	//Customer Menu
	fmt.Printf("%-3s %s %3s \n", separator, "Customer group Menu", separator)
	fmt.Print(NewCustomerGroup, "			|	new Customer group\n")
	fmt.Print(NewCustomerGRelation, "			|	new customer group relation\n")
	fmt.Print(CustomerGroupList, "			|	customer group list\n")
	fmt.Print(CustomerGroupRelationList, "			|	customer relation group list\n")
	fmt.Print(FindCustomerByGroupID, "			|	find customer(s) by group id\n")
	fmt.Printf("%-3s %s %3s \n", separator, "Customer Menu", separator)
	fmt.Print(NewCustomer, "			|	new Customer\n")
	fmt.Print(EditCustomerByCustomerId, "			|	edit customer by id\n")
	fmt.Print(ListOfCustomer, "			|	list of Customer(s)\n")
	fmt.Print(DeleteCustomerById, "			|	delete Customer by id\n")
	fmt.Print(DeleteMultiCustomerByIds, "			|	delete multi Customer by id(s)\n")
	fmt.Print(DeleteAllCustomers, "			|	delete all Customer\n")
	//Person Menu
	fmt.Printf("%-3s %s %3s \n", separator, "Person Menu", separator)
	fmt.Print(NewPerson, "			|	new person\n")
	fmt.Print(ListOfPerson, "			|	list of person(s)\n")
	fmt.Print(FindOnePersonById, "			|	find one person by id\n")
	fmt.Print(FindPersonContainingSomeCharacter, "			|	find person, contain some character\n")
	fmt.Print(FindAndEditPersonByPersonId, "			|	find and edit person by person id\n")
	fmt.Print(DeletePersonById, "			|	delete person by id\n")
	fmt.Print(DeleteMultiPersonByIds, "			|	delete multi person by id(s)\n")
	fmt.Print(DeleteAllPersons, "			|	delete all person\n")
	//country menu
	fmt.Printf("%-3s %s %3s \n", separator, "Country menu", separator)
	fmt.Print(AddCountry, "			|	new country\n")
	fmt.Print(EditCountry, "			|	edit country\n")
	fmt.Print(DeleteCountry, "			|	delete country(ies)\n")
	fmt.Print(FindCountryByChar, "			|	find country by char\n")
	fmt.Print(CountriesList, "			|	list of countries data\n")
	//City Menu
	fmt.Printf("%-3s %s %3s \n", separator, "City Menu", separator)
	fmt.Print(InsertNewCity, "			|	insert new city\n")
	fmt.Print(ListOfCities, "			|	list of city(ies)\n")
	fmt.Print(EditCityById, "			|	edit city by id\n")
	fmt.Print(DeleteCityById, "			|	delete city by id\n")
	//Job Menu
	fmt.Printf("%-3s %s %3s \n", separator, "Job Menu", separator)
	fmt.Print(InsertNewJob, "			|	insert new job\n")
	fmt.Print(ListOfJob, "			|	list of job(s)\n")
	fmt.Print(EditJobById, "			|	edit job by id\n")
	fmt.Print(DeleteJobById, "			|	delete job by id\n")
	//Sex Menu
	fmt.Printf("%-3s %s %3s \n", separator, "Sex Menu", separator)
	fmt.Print(InsertNewSex, "			|	insert new sex\n")
	fmt.Print(EditSex, "			|	edit sex by id\n")
	fmt.Print(DeleteSex, "			|	delete sex by id\n")
	fmt.Print(ListOfSex, "			|	list of sex\n")
	fmt.Println(separator7)
	fmt.Print(RESET, "			|	reset program (as new user role)\n")
	fmt.Println(separator7)
	//regulator menu
	if identity.IsRegulator {
		fmt.Printf("%-3s %s %3s \n", separator, "Fill sample data", separator)
		fmt.Print(InsertSomeSamplesContacts, "			|	insert some sample's contacts\n")
		fmt.Printf("%-3s %s %3s \n", separator, "Print All", separator)
		fmt.Print(PrintAllData, "			|	print all data\n")
		fmt.Println(separator7)
	}
	runMenu()
}

func runMenu() {
	//TODO: CompleteAllEntityFeature(Person|Job|Sex|City|Contact)
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
				for _, data := range contact.GetContacts().Data {
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
				for _, data := range city.GetCities().CityData {
					fmt.Printf("%3v | %-15v \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("Top 10 Country Data:")
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
					"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
				fmt.Println()
				for _, data := range country.GetCountry()[:10] {
					fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v | %-19v \n",
						data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
				}
				fmt.Println(separator7)
				fmt.Println("Job Data:")
				for _, data := range job.GetJobs().JobData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("sex Data:")
				for _, data := range sex.GetSex().SexData {
					fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
				}
				fmt.Println(separator7)
				fmt.Println("person Data:")
				for _, data := range person.GetPersons().PersonData {
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
				ShowMenu()
			}
		}
		switch userInput {
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
			fmt.Println("Please inter cellphone:")
			fmt.Scanln(&cellphone)
			//input description
			fmt.Println("Please inter description:")
			scanner.Scan()
			contact.NewContact(
				protocol.Contact{
					PersonID:    personID,
					JobID:       jobID,
					Tel:         tel,
					Cellphone:   cellphone,
					Description: scanner.Text(),
				})
			fmt.Println(">> New contact added done <<")
			ShowMenu()
		case ListOfContact:
			fmt.Println(separator7)
			fmt.Println("Contact Data:")
			fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
				"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
			fmt.Println("")
			for _, data := range contact.GetContacts().Data {
				_, person := person.FindPersonByID(data.PersonID)
				genderID := person.GenderID
				gender, _ := sex.FindSexByID(uint8(genderID))
				_, job := job.FindJobByID(data.JobID)
				_, city := city.FindCityByID(job.LocationID)
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
					data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.Cellphone, job.LocationID, city.Name, data.Description)
			}
			ShowMenu()
		case FindOneContactById:
			var id uint
			fmt.Println("Please insert your contact ID:")
			fmt.Scanln(&id)
			isFound, result := contact.FindContactByID(id)
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
					result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.Cellphone, job.LocationID, city.Name, result.Description)
			} else {
				fmt.Println("not found")
			}
			ShowMenu()
		case FindContactContainingSomeCharacter:
			var insertChar string
			fmt.Println("insert character(s):")
			fmt.Scanln(&insertChar)
			resultData, resultCount := contact.FindContactByChar(insertChar)
			if resultCount == 0 {
				fmt.Println("not found")
			} else {
				fmt.Println(resultCount, "record(s) found")
				fmt.Println(separator7)
				fmt.Println("Contact Data:")
				fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-v | %-12v | %-3v  \n",
					"Id", "PID", "PersonName", "PersonFamily", "JID", "JobName", "Gender", "Cellphone", "LoID", "jobCity", "Desc")
				fmt.Println("")
				for _, data := range resultData.Data {
					_, person := person.FindPersonByID(data.PersonID)
					genderID := person.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, job := job.FindJobByID(data.JobID)
					_, city := city.FindCityByID(job.LocationID)
					fmt.Printf("%3v | %-3v | %-15s | %-20v | %-3v | %-20s | %-8v | %-12v | %-4v | %-12v | %-3v  \n",
						data.Id, data.PersonID, person.FirstName, person.LastName, data.JobID, job.Name, gender.Name, data.Cellphone, job.LocationID, city.Name, data.Description)
				}
			}
			ShowMenu()
		case FindAndEditContactByContactId:
			var (
				insertContactID, jobID, personID uint
				tel, cellphone                   string
			)
			fmt.Println("Please insert your contact id you want to edit:")
			fmt.Scanln(&insertContactID)
			isFound, result := contact.FindContactByID(insertContactID)
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
					result.Id, result.PersonID, person.FirstName, person.LastName, result.JobID, job.Name, gender.Name, result.Cellphone, job.LocationID, city.Name, result.Description)
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
			scanner.Scan()
			editedContact := protocol.Contact{
				PersonID:    personID,
				JobID:       jobID,
				Tel:         tel,
				Cellphone:   cellphone,
				Description: scanner.Text(),
			}
			state := contact.EditContactByID(editedContact, insertContactID)
			if state.State {
				fmt.Printf("contact no %v updated", insertContactID)
			}
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
		case NewCustomerGroup:
			var groupName string
			fmt.Println("insert group name:")
			scanner.Scan()
			groupName = scanner.Text()
			customer.NewGroup(groupName)
			ShowMenu()
		case NewCustomerGRelation:
			var customerID, groupID uint
			fmt.Println("insert customerID:")
			fmt.Scanln(&customerID)
			fmt.Println("insert groupID:")
			fmt.Scanln(&groupID)
			customer.NewRelation(customerID, groupID)
			ShowMenu()
		case CustomerGroupList:
			fmt.Println(separator7)
			fmt.Println("Customer group Data:")
			fmt.Printf("%3v | %-15s   \n", "Id", "Groupname")
			fmt.Println()
			for _, group := range customer.GetCustomerGroup() {
				fmt.Printf("%3v | %-25s  \n",
					group.GroupID, group.GroupName)
			}
			ShowMenu()
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
			ShowMenu()
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
			for ID, data := range customer.FindCustomerByGroupID(groupID).Data {
				_, person := person.FindPersonByID(data.PersonID)
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-25v | %-25v | %-23v \n",
					ID, person.FirstName, person.LastName, data.PersonID, (data.CreateAt).String()[0:19], (data.UpdatedAt).String()[0:19], data.Description)
			}
			ShowMenu()
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
			ShowMenu()
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
			ShowMenu()
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
			ShowMenu()
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
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					idPackInt = append(idPackInt, uint(j))
					customer.DeleteCustomerById(uint(j))
				}
				fmt.Printf("%v customer(s) has been deleted", idPackInt)
			}
			ShowMenu()
		case DeleteAllCustomers:
			globalVars.CustomerMapStore = nil
			fmt.Println("all customers deleted")
			ShowMenu()
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
			ShowMenu()
		case ListOfPerson:
			fmt.Println(separator7)
			fmt.Println("Person Data:")
			fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
				"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
			fmt.Println("")
			//TODO::: decide to:
			//create a method called "GetPersons" to responds updated data live
			//use a global variable and after every event have to updated it
			for _, data := range person.GetPersons().PersonData /*person.GetPersons().PersonData*/ {
				genderID := data.GenderID
				gender, _ := sex.FindSexByID(uint8(genderID))
				_, city := city.FindCityByID(data.BirthLocationID)
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					data.Id, data.FirstName, data.LastName, gender.Name, data.BirthLocationID, city.Name, (data.DOB).String()[0:10], data.Description)
			}
			ShowMenu()
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
			ShowMenu()
		case FindPersonContainingSomeCharacter:
			fmt.Println("insert character")
			scanner.Scan()
			if state, data := person.FindPersonByChar(scanner.Text()); state.State {
				fmt.Println(separator7)
				fmt.Println(len(data.PersonData), "record(s) found")
				fmt.Println("Person Data:")
				fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
					"Id", "PersonName", "PersonFamily", "Gender", "BLoID", "jobCity", "DOB", "Desc")
				fmt.Println("")
				for _, personData := range data.PersonData {
					genderID := personData.GenderID
					gender, _ := sex.FindSexByID(uint8(genderID))
					_, city := city.FindCityByID(personData.BirthLocationID)
					fmt.Printf("%3v | %-15s | %-20v | %-8v | %-5v | %-12v | %-13v | %-3v  \n",
						personData.Id, personData.FirstName, personData.LastName, gender.Name, personData.BirthLocationID, city.Name, (personData.DOB).String()[0:10], personData.Description)
				}
			}
			ShowMenu()
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
						j, err := strconv.Atoi(i)
						if err != nil {
							panic(err)
						}
						idPackInt = append(idPackInt, uint(j))
					}
					resNums := person.DeletePerson(idPackInt)
					fmt.Printf("%v person(s) has been deleted", resNums)
				}
				ShowMenu()
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
				ShowMenu()
			}
			fmt.Println("not found")
			ShowMenu()
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
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					idPackInt = append(idPackInt, uint(j))
				}
				resNums := person.DeletePerson(idPackInt)
				fmt.Printf("%v person(s) has been deleted", resNums)
			}
			ShowMenu()
		case DeleteMultiPersonByIds:
			//some code
			ShowMenu()
		case DeleteAllPersons:
			//some code
			ShowMenu()
		case InsertNewCity:
			var inputCity, ariaCode string
			fmt.Println("insert city name:")
			scanner.Scan()
			inputCity = scanner.Text()
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
			dataJSON, _ := json.MarshalIndent(city.GetCities().CityData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range city.GetCities().CityData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
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
			ShowMenu()
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
			ShowMenu()
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
					j, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					idPackInt = append(idPackInt, uint(j))
				}
				resNums := country.DeleteCountry(idPackInt)
				fmt.Printf("%v country(ies) has been deleted", resNums)
			}
			ShowMenu()
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
			ShowMenu()
		case CountriesList:
			fmt.Println(separator7)
			fmt.Println("Country Data:")
			fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
				"Id", "CountryName", "ShName", "PrePhone", "CapID", "create@", "updated@")
			fmt.Println()
			for _, data := range country.GetCountry() {
				fmt.Printf("%3v | %-45s | %-8v | %-10v | %-12v | %-19v  | %-19v   \n",
					data.ID, data.Name, data.ShortName, data.PrePhoneCode, data.CapitalID, (data.CreatedAt).String()[0:16], (data.UpdatedAt).String()[0:16])
			}
			ShowMenu()
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
			ShowMenu()
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
				fmt.Printf("%v job(s) has been deleted", resNums)
			}
			ShowMenu()
		case ListOfJob:
			dataJSON, _ := json.MarshalIndent(job.GetJobs().JobData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range job.GetJobs().JobData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
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
			ShowMenu()
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
			dataJSON, _ := json.MarshalIndent(sex.GetSex().SexData, "", "  ")
			fmt.Println(string(dataJSON))
			fmt.Println(separator7)
			for _, data := range sex.GetSex().SexData {
				fmt.Printf("%3v | %-15s \n", data.Id, data.Name)
			}
			ShowMenu()
		//something wrong:
		case RESET:
			globalVars.DeleteAllGlobalVars()
			RunApp()
		default:
			fmt.Println("bad input, please insert character of menu list")
		}
	}
}

//TODO::: TalkToOmidHekayatiAboutSendDataInputValidationToClientForToCheckBeforeRequestToServer(LikeYII2)
