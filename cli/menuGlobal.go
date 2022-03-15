package cli

var (
	GlobalMenu = Menu{[]MenuGroup{
		{"Contact Menu", []MenuDetail{
			{NewContactRecord, "new record"},
			{ListOfContact, "list of contact"},
			{FindOneContactById, "list of contact"},
			{FindContactContainingSomeCharacter, "find one contact by id"},
			{FindAndEditContactByContactId, "find contact, contain some character"},
			{DeleteContactById, "find and edit contact by contact id"},
			{DeleteMultiContactByIds, "delete contact by id"},
			{DeleteAllContacts, "delete multi contact by id(s)"},
		}},
		{"Customer group Menu", []MenuDetail{
			{NewCustomerGroup, "new Customer group"},
			{NewCustomerGRelation, "new customer group relation"},
			{CustomerGroupList, "customer group list"},
			{CustomerGroupRelationList, "customer relation group list"},
			{FindCustomerByGroupID, "find customer(s) by group id"},
		}},
		{"Customer Menu", []MenuDetail{
			{NewCustomer, "new Customer"},
			{EditCustomerByCustomerId, "edit customer by id"},
			{ListOfCustomer, "list of Customer(s)"},
			{DeleteCustomerById, "delete Customer by id"},
			{DeleteMultiCustomerByIds, "delete multi Customer by id(s)"},
			{DeleteAllCustomers, "delete all Customer"},
		}},
		{"Person Menu", []MenuDetail{
			{NewPerson, "new person"},
			{ListOfPerson, "list of person(s)"},
			{FindOnePersonById, "find one person by id"},
			{FindPersonContainingSomeCharacter, "find person, contain some character"},
			{FindAndEditPersonByPersonId, "find and edit person by person id"},
			{DeletePersonById, "delete person by id"},
			{DeleteMultiPersonByIds, "delete multi person by id(s)"},
			{DeleteAllPersons, "delete all person"},
		}},
		{"Country menu", []MenuDetail{
			{AddCountry, "new country"},
			{EditCountry, "edit country"},
			{DeleteCountry, "delete country(ies)"},
			{FindCountryByChar, "find country by char"},
			{CountriesList, "list of countries data"},
		}},
		{"City Menu", []MenuDetail{
			{InsertNewCity, "insert new city"},
			{ListOfCities, "list of city(ies)"},
			{EditCityById, "edit city by id"},
			{DeleteCityById, "delete city by id"},
			{CallDistanceTimeTwoCities, "find distance/time online traffic two cities"},
		}},
		{"Job Menu", []MenuDetail{
			{InsertNewJob, "insert new job"},
			{ListOfJob, "list of job(s)"},
			{EditJobById, "edit job by id"},
			{DeleteJobById, "delete job by id"},
		}},
		{"Sex Menu", []MenuDetail{
			{InsertNewSex, "insert new sex"},
			{EditSex, "edit sex by id"},
			{DeleteSex, "delete sex by id"},
			{ListOfSex, "list of sex"},
		}},
		{"General Menu", []MenuDetail{
			{RESET, "reset program (as new user role)"},
			{InsertSomeSamplesData, "insert some sample's data"},
			{PrintAllData, "print all data"},
			{CheckIranNationalCode, "check iran national code"},
		}},
	}}
)
