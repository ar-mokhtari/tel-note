package cli

// MetaData & Usage
// A metadata is a data that describes other data in a structured way.
// We will use content variables to describe our hardcoded string data.

const (
	//general meta data
	YES             = "yes"
	NO              = "no"
	OK              = "ok"
	Male            = "male"
	Female          = "female"
	RegulatorString = "reg"
	RESET           = "RESET"

	//TODO::: + string(rune(iota)) (for generate unique routes)

	// contact meta data
	NewContactRecord                   = "N"
	ListOfContact                      = "L"
	FindOneContactById                 = "F"
	FindContactContainingSomeCharacter = "FC"
	FindAndEditContactByContactId      = "E"
	DeleteContactById                  = "D"
	DeleteMultiContactByIds            = "DM"
	DeleteAllContacts                  = "DA"

	// insert sample data
	InsertSomeSamplesContacts = "DATA"

	//Customer meta data
	NewCustomer              = "NCU"
	ListOfCustomer           = "LCU"
	EditCustomerByCustomerId = "ECU"
	DeleteCustomerById       = "DCU"
	DeleteMultiCustomerByIds = "DMCU"
	DeleteAllCustomers       = "DACU"

	//person meta data
	NewPerson                         = "NP"
	ListOfPerson                      = "LP"
	FindOnePersonById                 = "FP"
	FindPersonContainingSomeCharacter = "FPC"
	FindAndEditPersonByPersonId       = "EP"
	DeletePersonById                  = "DP"
	DeleteMultiPersonByIds            = "DMP"
	DeleteAllPersons                  = "DAP"

	//city meta data
	InsertNewCity  = "NC"
	EditCityById   = "EC"
	DeleteCityById = "DC"
	ListOfCities   = "LC"

	//job meta data
	InsertNewJob  = "NJ"
	EditJobById   = "EJ"
	DeleteJobById = "DJ"
	ListOfJob     = "LJ"

	//sex meta data
	InsertNewSex = "NS"
	ListOfSex    = "LS"
	EditSex      = "ES"
	DeleteSex    = "DS"

	//print all
	PrintAllData = "ALL"
)
