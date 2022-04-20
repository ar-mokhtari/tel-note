package env

import (
	"tel-note/protocol"
)

var (
	GlobalMenu = protocol.Menu{Group: []protocol.MenuGroup{
		{Contact, []protocol.MenuDetail{
			{ContactNew, ContactNewD, ContactNewR, DataEntryTypeForm},
			{ContactList, ContactListD, ContactListR, ListTypeForm},
			{ContactFindId, ContactFindIdD, ContactFindIdR, ActionType},
			{FindContactChar, FindContactCharD, FindContactCharR, ActionType},
			{EditContactId, EditContactIdD, EditContactIdR, DataEntryTypeForm},
			{DeleteContactId, DeleteContactIdD, DeleteContactIdR, ActionType},
			{DeleteMultiContact, DeleteMultiContactD, DeleteMultiContactR, ActionType},
			{DeleteAllContacts, DeleteAllContactsD, DeleteAllContactsR, ActionType},
		}},
		{CustomerGroup, []protocol.MenuDetail{
			{NewCustomerGroup, NewCustomerGroupD, NewCustomerGroupR, DataEntryTypeForm},
			{NewCustGRelation, NewCustGRelationD, NewCustGRelationR, DataEntryTypeForm},
			{CustomerGroupList, CustomerGroupListD, CustomerGroupListR, ListTypeForm},
			{CustGRelationList, CustGRelationListD, CustGRelationListR, ListTypeForm},
			{FindCustGrpID, FindCustGrpIDD, FindCustGrpIDR, ActionType},
			{FindCustGrpRelationByGrpID, FindCustGrpRelationByGrpIDD, FindCustGrpRelationByGrpIDR, ActionType},
		}},
		{Customer, []protocol.MenuDetail{
			{NewCustomer, NewCustomerD, NewCustomerR, DataEntryTypeForm},
			{CustomerList, CustomerListD, CustomerListR, ListTypeForm},
			{EditCustomer, EditCustomerD, EditCustomerR, DataEntryTypeForm},
			{FindCustID, FindCustIDD, FindCustIDR, ActionType},
			{DeleteCustomer, DeleteCustomerD, DeleteCustomerR, ActionType},
			{DeleteMultiCust, DeleteMultiCustD, DeleteMultiCustR, ActionType},
			{DeleteAllCust, DeleteAllCustD, DeleteAllCustR, ActionType},
		}},
		{Person, []protocol.MenuDetail{
			{NewPerson, NewPersonD, NewPersonR, DataEntryTypeForm},
			{ListOfPerson, ListOfPersonD, ListOfPersonR, ListTypeForm},
			{FindOnePersonById, FindOnePersonByIdD, FindOnePersonByIdR, ActionType},
			{FindPersonChar, FindPersonCharD, FindPersonCharR, ActionType},
			{EditPerson, EditPersonD, EditPersonR, DataEntryTypeForm},
			{DeletePerson, DeletePersonD, DeletePersonR, ActionType},
			{DeleteMultiPerson, DeleteMultiPersonD, DeleteMultiPersonR, ActionType},
			{DeleteAllPersons, DeleteAllPersonsD, DeleteAllPersonsR, ActionType},
		}},
		{Country, []protocol.MenuDetail{
			{AddCountry, AddCountryD, AddCountryR, DataEntryTypeForm},
			{EditCountry, EditCountryD, EditCountryR, DataEntryTypeForm},
			{DeleteCountry, DeleteCountryD, DeleteCountryR, ActionType},
			{FindCountryByChar, FindCountryByCharD, FindCountryByCharR, ActionType},
			{CountriesList, CountriesListD, CountriesListR, ListTypeForm},
			{CountriesCall, CountriesCallD, CountriesCallR, ActionType},
		}},
		{City, []protocol.MenuDetail{
			{ListOfCities, ListOfCitiesD, ListOfCitiesR, ListTypeForm},
			{InsertNewCity, InsertNewCityD, InsertNewCityR, DataEntryTypeForm},
			{EditCityById, EditCityByIdD, EditCityByIdR, DataEntryTypeForm},
			{FindCityByChar, FindCityByCharD, FindCityByCharR, ActionType},
			{FindCityById, FindCityByIdD, FindCityByIdR, ActionType},
			{DeleteCityById, DeleteCityByIdD, DeleteCityByIdR, ActionType},
			{CallDistanceTimeTwoCities, CallDistanceTimeTwoCitiesD, CallDistanceTimeTwoCitiesR, ReportTypeForm},
		}},
		{Job, []protocol.MenuDetail{
			{InsertNewJob, InsertNewJobD, InsertNewJobR, DataEntryTypeForm},
			{EditJobById, EditJobByIdD, EditJobByIdR, DataEntryTypeForm},
			{DeleteJobById, DeleteJobByIdD, DeleteJobByIdR, ActionType},
			{ListOfJob, ListOfJobD, ListOfJobR, ListTypeForm},
		}},
		{Sex, []protocol.MenuDetail{
			{InsertNewSex, InsertNewSexD, InsertNewSexR, DataEntryTypeForm},
			{ListOfSex, ListOfSexD, ListOfSexR, ListTypeForm},
			{EditSex, EditSexD, EditSexR, DataEntryTypeForm},
			{DeleteSex, DeleteSexD, DeleteSexR, ActionType},
		}},
		{SampleData, []protocol.MenuDetail{
			{InsertData, InsertDataD, InsertDataR, DataEntryTypeForm},
			{GetData, GetDataD, GetDataR, ReportTypeForm},
		}},
		{Report, []protocol.MenuDetail{
			{ContactReport, ContactReportD, ContactReportR, ReportTypeForm},
		}},
		{General, []protocol.MenuDetail{
			{CheckIranNationalCode, CheckIranNationalCodeD, CheckIranNationalCodeR, ActionType},
			{RESET, RESETD, RESETR, ActionType},
		}},
	}}
)
