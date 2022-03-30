package cli

import "tel-note/env"

var (
	GlobalMenu = Menu{[]MenuGroup{
		{"Contact Menu", []MenuDetail{
			{env.NewContactRecord, "new record"},
			{env.ListOfContact, "list of contact"},
			{env.FindOneContactById, "list of contact"},
			{env.FindContactContainingSomeCharacter, "find one contact by id"},
			{env.FindAndEditContactByContactId, "find contact, contain some character"},
			{env.DeleteContactById, "find and edit contact by contact id"},
			{env.DeleteMultiContactByIds, "delete contact by id"},
			{env.DeleteAllContacts, "delete multi contact by id(s)"},
		}},
		{"Customer group Menu", []MenuDetail{
			{env.NewCustomerGroup, "new Customer group"},
			{env.NewCustomerGRelation, "new customer group relation"},
			{env.CustomerGroupList, "customer group list"},
			{env.CustomerGroupRelationList, "customer relation group list"},
			{env.FindCustomerByGroupID, "find customer(s) by group id"},
		}},
		{"Customer Menu", []MenuDetail{
			{env.NewCustomer, "new Customer"},
			{env.EditCustomerByCustomerId, "edit customer by id"},
			{env.ListOfCustomer, "list of Customer(s)"},
			{env.DeleteCustomerById, "delete Customer by id"},
			{env.DeleteMultiCustomerByIds, "delete multi Customer by id(s)"},
			{env.DeleteAllCustomers, "delete all Customer"},
		}},
		{"Person Menu", []MenuDetail{
			{env.NewPerson, "new person"},
			{env.ListOfPerson, "list of person(s)"},
			{env.FindOnePersonById, "find one person by id"},
			{env.FindPersonContainingSomeCharacter, "find person, contain some character"},
			{env.FindAndEditPersonByPersonId, "find and edit person by person id"},
			{env.DeletePersonById, "delete person by id"},
			{env.DeleteMultiPersonByIds, "delete multi person by id(s)"},
			{env.DeleteAllPersons, "delete all person"},
		}},
		{"Country menu", []MenuDetail{
			{env.AddCountry, "new country"},
			{env.EditCountry, "edit country"},
			{env.DeleteCountry, "delete country(ies)"},
			{env.FindCountryByChar, "find country by char"},
			{env.CountriesList, "list of countries data"},
		}},
		{"City Menu", []MenuDetail{
			{env.InsertNewCity, "insert new city"},
			{env.ListOfCities, "list of city(ies)"},
			{env.EditCityById, "edit city by id"},
			{env.DeleteCityById, "delete city by id"},
			{env.CallDistanceTimeTwoCities, "find distance/time online traffic two cities"},
		}},
		{"Job Menu", []MenuDetail{
			{env.InsertNewJob, "insert new job"},
			{env.ListOfJob, "list of job(s)"},
			{env.EditJobById, "edit job by id"},
			{env.DeleteJobById, "delete job by id"},
		}},
		{"Sex Menu", []MenuDetail{
			{env.InsertNewSex, "insert new sex"},
			{env.EditSex, "edit sex by id"},
			{env.DeleteSex, "delete sex by id"},
			{env.ListOfSex, "list of sex"},
		}},
		{"General Menu", []MenuDetail{
			{env.RESET, "reset program (as new user role)"},
			{env.InsertSomeSamplesData, "insert some sample's data"},
			{env.PrintAllData, "print all data"},
			{env.CheckIranNationalCode, "check iran national code"},
		}},
	}}
)
