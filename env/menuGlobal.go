package env

import (
	"tel-note/protocol"
)

var (
	GlobalMenu = protocol.Menu{Group: []protocol.MenuGroup{
		{"Contact", []protocol.MenuDetail{
			{NewContactRecord, NewContactRecordD, NewContactRecordR, DataEntryTypeForm},
			{ListOfContact, ListOfContactD, ListOfContactR, ListTypeForm},
			{FindOneContactById, FindOneContactByIdD, FindOneContactByIdR, ActionType},
			{FindContactContainingSomeCharacter, FindContactContainingSomeCharacterD, FindContactContainingSomeCharacterR, ActionType},
			{FindAndEditContactByContactId, FindAndEditContactByContactIdD, FindAndEditContactByContactIdR, DataEntryTypeForm},
			{DeleteContactById, DeleteContactByIdD, DeleteContactByIdR, ActionType},
			{DeleteMultiContactByIds, DeleteMultiContactByIdsD, DeleteMultiContactByIdsR, ActionType},
			{DeleteAllContacts, DeleteAllContactsD, DeleteAllContactsR, ActionType},
		}},
		{"Customer group", []protocol.MenuDetail{
			{NewCustomerGroup, NewCustomerGroupD, NewCustomerGroupR, DataEntryTypeForm},
			{NewCustomerGRelation, NewCustomerGRelationD, NewCustomerGRelationR, DataEntryTypeForm},
			{CustomerGroupList, CustomerGroupListD, CustomerGroupListR, ListTypeForm},
			{CustomerGroupRelationList, CustomerGroupRelationListD, CustomerGroupRelationListR, ListTypeForm},
			{FindCustomerByGroupID, FindCustomerByGroupIDD, FindCustomerByGroupIDR, ActionType},
		}},
		{"Customer", []protocol.MenuDetail{
			{NewCustomer, NewCustomerD, NewCustomerR, DataEntryTypeForm},
			{ListOfCustomer, ListOfCustomerD, ListOfCustomerR, ListTypeForm},
			{EditCustomerByCustomerId, EditCustomerByCustomerIdD, EditCustomerByCustomerIdR, DataEntryTypeForm},
			{DeleteCustomerById, DeleteCustomerByIdD, DeleteCustomerByIdR, ActionType},
			{DeleteMultiCustomerByIds, DeleteMultiCustomerByIdsD, DeleteMultiCustomerByIdsR, ActionType},
			{DeleteAllCustomers, DeleteAllCustomersD, DeleteAllCustomersR, ActionType},
		}},
		{"Person", []protocol.MenuDetail{
			{NewPerson, NewPersonD, NewPersonR, DataEntryTypeForm},
			{ListOfPerson, ListOfPersonD, ListOfPersonR, ListTypeForm},
			{FindOnePersonById, FindOnePersonByIdD, FindOnePersonByIdR, ActionType},
			{FindPersonContainingSomeCharacter, FindPersonContainingSomeCharacterD, FindPersonContainingSomeCharacterR, ActionType},
			{FindAndEditPersonByPersonId, FindAndEditPersonByPersonIdD, FindAndEditPersonByPersonIdR, DataEntryTypeForm},
			{DeletePersonById, DeletePersonByIdD, DeletePersonByIdR, ActionType},
			{DeleteMultiPersonByIds, DeleteMultiPersonByIdsD, DeleteMultiPersonByIdsR, ActionType},
			{DeleteAllPersons, DeleteAllPersonsD, DeleteAllPersonsR, ActionType},
		}},
		{"Country", []protocol.MenuDetail{
			{AddCountry, AddCountryD, AddCountryR, DataEntryTypeForm},
			{EditCountry, EditCountryD, EditCountryR, DataEntryTypeForm},
			{DeleteCountry, DeleteCountryD, DeleteCountryR, ActionType},
			{FindCountryByChar, FindCountryByCharD, FindCountryByCharR, ActionType},
			{CountriesList, CountriesListD, CountriesListR, ListTypeForm},
		}},
		{"City", []protocol.MenuDetail{
			{ListOfCities, ListOfCitiesD, ListOfCitiesR, ListTypeForm},
			{InsertNewCity, InsertNewCityD, InsertNewCityR, DataEntryTypeForm},
			{EditCityById, EditCityByIdD, EditCityByIdR, DataEntryTypeForm},
			{FindCityByChar, FindCityByCharD, FindCityByCharR, ActionType},
			{FindCityById, FindCityByIdD, FindCityByIdR, ActionType},
			{DeleteCityById, DeleteCityByIdD, DeleteCityByIdR, ActionType},
			{CallDistanceTimeTwoCities, CallDistanceTimeTwoCitiesD, CallDistanceTimeTwoCitiesR, ReportTypeForm},
		}},
		{"Job", []protocol.MenuDetail{
			{InsertNewJob, InsertNewJobD, InsertNewJobR, DataEntryTypeForm},
			{EditJobById, EditJobByIdD, EditJobByIdR, DataEntryTypeForm},
			{DeleteJobById, DeleteJobByIdD, DeleteJobByIdR, ActionType},
			{ListOfJob, ListOfJobD, ListOfJobR, ListTypeForm},
		}},
		{"Sex Menu", []protocol.MenuDetail{
			{InsertNewSex, InsertNewSexD, InsertNewSexR, DataEntryTypeForm},
			{ListOfSex, ListOfSexD, ListOfSexR, ListTypeForm},
			{EditSex, EditSexD, EditSexR, DataEntryTypeForm},
			{DeleteSex, DeleteSexD, DeleteSexR, ActionType},
		}},
		{"Sample data", []protocol.MenuDetail{
			{InsertSomeSamplesData, InsertSomeSamplesDataD, InsertSomeSamplesDataR, DataEntryTypeForm},
			{PrintAllData, PrintAllDataD, PrintAllDataR, ReportTypeForm},
		}},
		{"Reports", []protocol.MenuDetail{
			{ContactReport, ContactReportD, ContactReportR, ReportTypeForm},
		}},
		{"General", []protocol.MenuDetail{
			{CheckIranNationalCode, CheckIranNationalCodeD, CheckIranNationalCodeR, ListTypeForm},
			{RESET, RESETD, RESETR, ActionType},
		}},
	}}
)
