package cli

// MetaData & Usage
// A metadata is a data that describes other data in a structured way.
// We will use content variables to describe our hardcoded string data.
//TODO::: + string(rune(iota)) (for generate unique routes)

const (
	// YES general meta data
	// NewContactRecord contact meta data
	// InsertSomeSamplesContacts insert sample data
	// NewCustomerGroup Customer meta data
	// NewPerson person meta data
	// InsertNewCity city meta data
	// AddCountry countries meta data
	// InsertNewJob job meta data
	// InsertNewSex sex meta data
	// PrintAllData

	YES                   = "yes"
	NO                    = "no"
	OK                    = "ok"
	Male                  = "male"
	Female                = "female"
	RegulatorString       = "reg"
	RESET                 = "RESET"
	CheckIranNationalCode = "ICNC"

	NewContactRecord                   = "N"
	ListOfContact                      = "L"
	FindOneContactById                 = "F"
	FindContactContainingSomeCharacter = "FC"
	FindAndEditContactByContactId      = "E"
	DeleteContactById                  = "D"
	DeleteMultiContactByIds            = "DM"
	DeleteAllContacts                  = "DA"

	InsertSomeSamplesContacts = "DATA"

	NewCustomerGroup          = "NCUG"
	NewCustomerGRelation      = "NCUGR"
	CustomerGroupList         = "CUGL"
	CustomerGroupRelationList = "CUGRL"
	FindCustomerByGroupID     = "FCUBG"
	NewCustomer               = "NCU"
	ListOfCustomer            = "LCU"
	EditCustomerByCustomerId  = "ECU"
	DeleteCustomerById        = "DCU"
	DeleteMultiCustomerByIds  = "DMCU"
	DeleteAllCustomers        = "DACU"

	NewPerson                         = "NP"
	ListOfPerson                      = "LP"
	FindOnePersonById                 = "FP"
	FindPersonContainingSomeCharacter = "FPC"
	FindAndEditPersonByPersonId       = "EP"
	DeletePersonById                  = "DP"
	DeleteMultiPersonByIds            = "DMP"
	DeleteAllPersons                  = "DAP"

	InsertNewCity             = "NC"
	EditCityById              = "EC"
	DeleteCityById            = "DC"
	ListOfCities              = "LC"
	CallDistanceTimeTwoCities = "CDTTC"

	AddCountry        = "ACOU"
	EditCountry       = "ECOU"
	DeleteCountry     = "DCOU"
	FindCountryByChar = "FCOUBC"
	CountriesList     = "COUL"

	InsertNewJob  = "NJ"
	EditJobById   = "EJ"
	DeleteJobById = "DJ"
	ListOfJob     = "LJ"

	InsertNewSex = "NS"
	ListOfSex    = "LS"
	EditSex      = "ES"
	DeleteSex    = "DS"

	PrintAllData = "ALL"
)
