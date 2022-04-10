package env

// MetaData & Usage
// A metadata is a data that describes other data in a structured way.
// We will use content variables to describe our hardcoded string data.
const (
	//methods:
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	PatchMethod  = "PATCH"
	DeleteMethod = "DELETE"
	//---
	YES = "yes"
	NO  = "no"
	OK  = "ok"
	//---
	Male   = "male"
	Female = "female"
	//---
	ShowMenuList    = "/"
	ShowMenuWarn    = "To see menu, insert ('/') then enter"
	RegulatorString = "reg"

	//template route types
	//front-end use this type for manage menu
	DataEntryTypeForm = 100   //(RW) get and return data to create entry form type in front-end
	ListTypeForm      = 200   //(RW) get and return data to create list form type in front-end
	ReportTypeForm    = 500   //(R) return data to create report form type in front-end
	ActionType        = 10000 //get and return data for do some action(s)
)
