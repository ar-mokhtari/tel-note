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

	//base meta data
	InsertNewSex = "NS"
	ListOfSex    = "LS"
	EditSex      = "ES"
	DeleteSex    = "DS"

	//print all
	PrintAllData = "ALL"
)
