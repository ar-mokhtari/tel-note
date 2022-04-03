package env

// MetaData & Usage
// A metadata is a data that describes other data in a structured way.
// We will use content variables to describe our hardcoded string data.
//TODO::: + string(rune(iota)) (for generate unique routes)

const (
	//contact
	NewContactRecord  = "N"
	NewContactRecordD = "new contact"
	NewContactRecordR = "/new-contact"

	ListOfContact  = "L"
	ListOfContactD = "list of contact"
	ListOfContactR = "/get-contact"

	FindOneContactById  = "F"
	FindOneContactByIdD = "find contact by id"
	FindOneContactByIdR = "/find-contact-id"

	FindContactContainingSomeCharacter  = "FC"
	FindContactContainingSomeCharacterD = "find contact by char"
	FindContactContainingSomeCharacterR = "/find-contact-char"

	FindAndEditContactByContactId  = "E"
	FindAndEditContactByContactIdD = "edit contact"
	FindAndEditContactByContactIdR = "/edit-contact"

	DeleteContactById  = "D"
	DeleteContactByIdD = "delete contact by id"
	DeleteContactByIdR = "/delete-contact-id"

	DeleteMultiContactByIds  = "DM"
	DeleteMultiContactByIdsD = "delete multi contact"
	DeleteMultiContactByIdsR = ""

	DeleteAllContacts  = "DA"
	DeleteAllContactsD = "delete all contact"
	DeleteAllContactsR = "/delete-contact-all"

	//sample data
	InsertSomeSamplesData  = "DATA"
	InsertSomeSamplesDataD = "insert some sample data "
	InsertSomeSamplesDataR = "/fill-data"

	PrintAllData  = "ALL"
	PrintAllDataD = "get all data"
	PrintAllDataR = "/get-data"

	//customer
	NewCustomerGroup  = "NCUG"
	NewCustomerGroupD = "new customer group"
	NewCustomerGroupR = ""

	NewCustomerGRelation  = "NCUGR"
	NewCustomerGRelationD = "new customer group relation"
	NewCustomerGRelationR = ""

	CustomerGroupList  = "CUGL"
	CustomerGroupListD = "customer group list"
	CustomerGroupListR = ""

	CustomerGroupRelationList  = "CUGRL"
	CustomerGroupRelationListD = "customer group relation"
	CustomerGroupRelationListR = ""

	FindCustomerByGroupID  = "FCUBG"
	FindCustomerByGroupIDD = "find customer by id"
	FindCustomerByGroupIDR = ""

	NewCustomer  = "NCU"
	NewCustomerD = "new customer"
	NewCustomerR = ""

	ListOfCustomer  = "LCU"
	ListOfCustomerD = "list of customers"
	ListOfCustomerR = ""

	EditCustomerByCustomerId  = "ECU"
	EditCustomerByCustomerIdD = "edit customer"
	EditCustomerByCustomerIdR = ""

	DeleteCustomerById  = "DCU"
	DeleteCustomerByIdD = "delete customer"
	DeleteCustomerByIdR = ""

	DeleteMultiCustomerByIds  = "DMCU"
	DeleteMultiCustomerByIdsD = "delete multi customer"
	DeleteMultiCustomerByIdsR = ""

	DeleteAllCustomers  = "DACU"
	DeleteAllCustomersD = "delete all customer"
	DeleteAllCustomersR = ""

	//person
	NewPerson  = "NP"
	NewPersonD = "new person"
	NewPersonR = ""

	ListOfPerson  = "LP"
	ListOfPersonD = "list of person(s)"
	ListOfPersonR = ""

	FindOnePersonById  = "FP"
	FindOnePersonByIdD = "find person by id"
	FindOnePersonByIdR = "find-person-id"

	FindPersonContainingSomeCharacter  = "FPC"
	FindPersonContainingSomeCharacterD = "finf person by char"
	FindPersonContainingSomeCharacterR = ""

	FindAndEditPersonByPersonId  = "EP"
	FindAndEditPersonByPersonIdD = "edit person"
	FindAndEditPersonByPersonIdR = ""

	DeletePersonById  = "DP"
	DeletePersonByIdD = "delete person"
	DeletePersonByIdR = ""

	DeleteMultiPersonByIds  = "DMP"
	DeleteMultiPersonByIdsD = "delete multi person(s)"
	DeleteMultiPersonByIdsR = ""

	DeleteAllPersons  = "DAP"
	DeleteAllPersonsD = "delete all person(s)"
	DeleteAllPersonsR = ""

	//city
	InsertNewCity  = "NC"
	InsertNewCityR = "/add-city"
	InsertNewCityD = "insert new city"

	EditCityById  = "EC"
	EditCityByIdR = "/edit-city"
	EditCityByIdD = "edit city by id"

	FindCityById  = "FCBI"
	FindCityByIdR = "/find-city-id"
	FindCityByIdD = "find city by id"

	FindCityByChar  = "FCBC"
	FindCityByCharR = "/find-city-char"
	FindCityByCharD = "find city by char"

	DeleteCityById  = "DC"
	DeleteCityByIdR = "/delete-city"
	DeleteCityByIdD = "delete city by id"

	ListOfCities  = "LC"
	ListOfCitiesR = "/get-city"
	ListOfCitiesD = "list of city(ies)"

	CallDistanceTimeTwoCities  = "CDTTC"
	CallDistanceTimeTwoCitiesR = "distance-time-between-two-city"
	CallDistanceTimeTwoCitiesD = "find distance/time online traffic two cities"

	//country
	AddCountry  = "ACOU"
	AddCountryD = "new country"
	AddCountryR = ""

	EditCountry  = "ECOU"
	EditCountryD = "edit country"
	EditCountryR = ""

	DeleteCountry  = "DCOU"
	DeleteCountryD = "delete country"
	DeleteCountryR = ""

	FindCountryByChar  = "FCOUBC"
	FindCountryByCharD = "find country by char"
	FindCountryByCharR = ""

	CountriesList  = "COUL"
	CountriesListD = "list of country(ies)"
	CountriesListR = "/call-country-list"

	//job
	InsertNewJob  = "NJ"
	InsertNewJobD = "new job"
	InsertNewJobR = ""

	EditJobById  = "EJ"
	EditJobByIdD = "edit job"
	EditJobByIdR = ""

	DeleteJobById  = "DJ"
	DeleteJobByIdD = "delete job"
	DeleteJobByIdR = ""

	ListOfJob  = "LJ"
	ListOfJobD = "list of job(s)"
	ListOfJobR = ""

	//sex
	InsertNewSex  = "NS"
	InsertNewSexD = "new sex"
	InsertNewSexR = ""

	ListOfSex  = "LS"
	ListOfSexD = "list of sex(s)"
	ListOfSexR = ""

	EditSex  = "ES"
	EditSexD = "edit sex"
	EditSexR = ""

	DeleteSex  = "DS"
	DeleteSexD = "delete sex"
	DeleteSexR = ""

	//general
	CheckIranNationalCode  = "ICNC"
	CheckIranNationalCodeD = "check Iran national code"
	CheckIranNationalCodeR = "/check-national-id"

	RESET  = "RESET"
	RESETD = ""
	RESETR = ""
)
