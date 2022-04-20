package env

// MetaData & Usage
// A metadata is a data that describes other data in a structured way.
// We will use content variables to describe our hardcoded string data.
//TODO::: + string(rune(iota)) (for generate unique routes)

const (
	//menu
	MenuR = "/menu-list"

	//contact
	Contact = "contact"

	ContactNew  = "N"
	ContactNewD = "new contact"
	ContactNewR = "/new-contact"

	ContactList  = "L"
	ContactListD = "list of contact"
	ContactListR = "/get-contact"

	ContactFindId  = "F"
	ContactFindIdD = "find contact by id"
	ContactFindIdR = "/find-contact-id"

	FindContactChar  = "FC"
	FindContactCharD = "find contact by char"
	FindContactCharR = "/find-contact-char"

	EditContactId  = "E"
	EditContactIdD = "edit contact"
	EditContactIdR = "/edit-contact"

	DeleteContactId  = "D"
	DeleteContactIdD = "delete contact by id"
	DeleteContactIdR = "/delete-contact-id"

	DeleteMultiContact  = "DM"
	DeleteMultiContactD = "delete multi contact"
	DeleteMultiContactR = ""

	DeleteAllContacts  = "DA"
	DeleteAllContactsD = "delete all contact"
	DeleteAllContactsR = "/delete-contact-all"

	//sample data
	SampleData  = "sample data"
	InsertData  = "DATA"
	InsertDataD = "insert some sample data "
	InsertDataR = "/fill-data"

	GetData  = "ALL"
	GetDataD = "get all data"
	GetDataR = "/get-data"

	//customer
	Customer      = "Customer"
	CustomerGroup = "Customer group"

	NewCustomerGroup  = "NCUG"
	NewCustomerGroupD = "new customer group"
	NewCustomerGroupR = "/new-cust-group"

	NewCustGRelation  = "NCUGR"
	NewCustGRelationD = "new customer group relation"
	NewCustGRelationR = "/new-cust-grp-relation"

	CustomerGroupList  = "CUGL"
	CustomerGroupListD = "customer group list"
	CustomerGroupListR = "/get-customer-group"

	CustGRelationList  = "CUGRL"
	CustGRelationListD = "customer group relation"
	CustGRelationListR = "/get-cust-g-relation"

	FindCustID  = "FCI"
	FindCustIDD = "find customer by id"
	FindCustIDR = "/find-cust-id"

	FindCustGrpID  = "FCGI"
	FindCustGrpIDD = "find customer group by id"
	FindCustGrpIDR = "/find-cust-grp-id"

	FindCustGrpRelationByGrpID  = "FCUBG"
	FindCustGrpRelationByGrpIDD = "find customer group by id"
	FindCustGrpRelationByGrpIDR = "/find-cust-grp-relation-id"

	NewCustomer  = "NCU"
	NewCustomerD = "new customer"
	NewCustomerR = "/new-customer"

	CustomerList  = "LCU"
	CustomerListD = "list of customers"
	CustomerListR = "/get-customer"

	EditCustomer  = "ECU"
	EditCustomerD = "edit customer"
	EditCustomerR = "/edit-customer"

	DeleteCustomer  = "DCU"
	DeleteCustomerD = "delete customer"
	DeleteCustomerR = "/delete-customer"

	DeleteMultiCust  = "DMCU"
	DeleteMultiCustD = "delete multi customer"
	DeleteMultiCustR = ""

	DeleteAllCust  = "DACU"
	DeleteAllCustD = "delete all customer"
	DeleteAllCustR = ""

	//person
	Person     = "Person"
	NewPerson  = "NP"
	NewPersonD = "new person"
	NewPersonR = ""

	ListOfPerson  = "LP"
	ListOfPersonD = "list of person(s)"
	ListOfPersonR = "/get-person"

	FindOnePersonById  = "FP"
	FindOnePersonByIdD = "find person by id"
	FindOnePersonByIdR = "find-person-id"

	FindPersonChar  = "FPC"
	FindPersonCharD = "find person by char"
	FindPersonCharR = ""

	EditPerson  = "EP"
	EditPersonD = "edit person"
	EditPersonR = ""

	DeletePerson  = "DP"
	DeletePersonD = "delete person"
	DeletePersonR = ""

	DeleteMultiPerson  = "DMP"
	DeleteMultiPersonD = "delete multi person(s)"
	DeleteMultiPersonR = ""

	DeleteAllPersons  = "DAP"
	DeleteAllPersonsD = "delete all person(s)"
	DeleteAllPersonsR = ""

	//city
	City           = "city"
	InsertNewCity  = "NC"
	InsertNewCityD = "insert new city"
	InsertNewCityR = "/add-city"

	EditCityById  = "EC"
	EditCityByIdD = "edit city by id"
	EditCityByIdR = "/edit-city"

	FindCityById  = "FCBI"
	FindCityByIdD = "find city by id"
	FindCityByIdR = "/find-city-id"

	FindCityByChar  = "FCBC"
	FindCityByCharD = "find city by char"
	FindCityByCharR = "/find-city-char"

	DeleteCityById  = "DC"
	DeleteCityByIdD = "delete city by id"
	DeleteCityByIdR = "/delete-city"

	ListOfCities  = "LC"
	ListOfCitiesD = "list of city(ies)"
	ListOfCitiesR = "/get-city"

	CallDistanceTimeTwoCities  = "CDTTC"
	CallDistanceTimeTwoCitiesD = "find distance/time online traffic two cities"
	CallDistanceTimeTwoCitiesR = "/distance-time-between-two-city"

	//country
	Country = "country"

	AddCountry  = "ACOU"
	AddCountryD = "new country"
	AddCountryR = "/new-country"

	EditCountry  = "ECOU"
	EditCountryD = "edit country"
	EditCountryR = "/edit-country"

	DeleteCountry  = "DCOU"
	DeleteCountryD = "delete country"
	DeleteCountryR = "/delete-country"

	FindCountryByChar  = "FCOUBC"
	FindCountryByCharD = "find country by char"
	FindCountryByCharR = "/find-country-char"

	CountriesList  = "COUL"
	CountriesListD = "list of country(ies)"
	CountriesListR = "/get-country"

	CountriesCall  = "COUC"
	CountriesCallD = "call for get country(ies)"
	CountriesCallR = "/call-country-list"

	//job
	Job = "Job"

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
	ListOfJobR = "/get-job"

	//sex
	Sex           = "sex"
	InsertNewSex  = "NS"
	InsertNewSexD = "new sex"
	InsertNewSexR = ""

	ListOfSex  = "LS"
	ListOfSexD = "list of sex(s)"
	ListOfSexR = "/get-sex"

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

	//reports
	Report = "reports"

	ContactReport  = ""
	ContactReportD = "contact general report"
	ContactReportR = "/report-contact"

	General = "general"
	RESET   = "RESET"
	RESETD  = ""
	RESETR  = ""
)
