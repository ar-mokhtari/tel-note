package cli

import "tel-note/env"

var (
	GlobalMenu = Menu{[]MenuGroup{
		{"Contact", []MenuDetail{
			{env.NewContactRecord, env.NewContactRecordD, env.NewContactRecordR, env.DataEntryTypeForm},
			{env.ListOfContact, env.ListOfContactD, env.ListOfContactR, env.ListTypeForm},
			{env.FindOneContactById, env.FindOneContactByIdD, env.FindOneContactByIdR, env.ActionType},
			{env.FindContactContainingSomeCharacter, env.FindContactContainingSomeCharacterD, env.FindContactContainingSomeCharacterR, env.ActionType},
			{env.FindAndEditContactByContactId, env.FindAndEditContactByContactIdD, env.FindAndEditContactByContactIdR, env.DataEntryTypeForm},
			{env.DeleteContactById, env.DeleteContactByIdD, env.DeleteContactByIdR, env.ActionType},
			{env.DeleteMultiContactByIds, env.DeleteMultiContactByIdsD, env.DeleteMultiContactByIdsR, env.ActionType},
			{env.DeleteAllContacts, env.DeleteAllContactsD, env.DeleteAllContactsR, env.ActionType},
		}},
		{"Customer group", []MenuDetail{
			{env.NewCustomerGroup, env.NewCustomerGroupD, env.NewCustomerGroupR, env.DataEntryTypeForm},
			{env.NewCustomerGRelation, env.NewCustomerGRelationD, env.NewCustomerGRelationR, env.DataEntryTypeForm},
			{env.CustomerGroupList, env.CustomerGroupListD, env.CustomerGroupListR, env.ListTypeForm},
			{env.CustomerGroupRelationList, env.CustomerGroupRelationListD, env.CustomerGroupRelationListR, env.ListTypeForm},
			{env.FindCustomerByGroupID, env.FindCustomerByGroupIDD, env.FindCustomerByGroupIDR, env.ActionType},
		}},
		{"Customer", []MenuDetail{
			{env.NewCustomer, env.NewCustomerD, env.NewCustomerR, env.DataEntryTypeForm},
			{env.ListOfCustomer, env.ListOfCustomerD, env.ListOfCustomerR, env.ListTypeForm},
			{env.EditCustomerByCustomerId, env.EditCustomerByCustomerIdD, env.EditCustomerByCustomerIdR, env.DataEntryTypeForm},
			{env.DeleteCustomerById, env.DeleteCustomerByIdD, env.DeleteCustomerByIdR, env.ActionType},
			{env.DeleteMultiCustomerByIds, env.DeleteMultiCustomerByIdsD, env.DeleteMultiCustomerByIdsR, env.ActionType},
			{env.DeleteAllCustomers, env.DeleteAllCustomersD, env.DeleteAllCustomersR, env.ActionType},
		}},
		{"Person", []MenuDetail{
			{env.NewPerson, env.NewPersonD, env.NewPersonR, env.DataEntryTypeForm},
			{env.ListOfPerson, env.ListOfPersonD, env.ListOfPersonR, env.ListTypeForm},
			{env.FindOnePersonById, env.FindOnePersonByIdD, env.FindOnePersonByIdR, env.ActionType},
			{env.FindPersonContainingSomeCharacter, env.FindPersonContainingSomeCharacterD, env.FindPersonContainingSomeCharacterR, env.ActionType},
			{env.FindAndEditPersonByPersonId, env.FindAndEditPersonByPersonIdD, env.FindAndEditPersonByPersonIdR, env.DataEntryTypeForm},
			{env.DeletePersonById, env.DeletePersonByIdD, env.DeletePersonByIdR, env.ActionType},
			{env.DeleteMultiPersonByIds, env.DeleteMultiPersonByIdsD, env.DeleteMultiPersonByIdsR, env.ActionType},
			{env.DeleteAllPersons, env.DeleteAllPersonsD, env.DeleteAllPersonsR, env.ActionType},
		}},
		{"Country", []MenuDetail{
			{env.AddCountry, env.AddCountryD, env.AddCountryR, env.DataEntryTypeForm},
			{env.EditCountry, env.EditCountryD, env.EditCountryR, env.DataEntryTypeForm},
			{env.DeleteCountry, env.DeleteCountryD, env.DeleteCountryR, env.ActionType},
			{env.FindCountryByChar, env.FindCountryByCharD, env.FindCountryByCharR, env.ActionType},
			{env.CountriesList, env.CountriesListD, env.CountriesListR, env.ListTypeForm},
		}},
		{"City", []MenuDetail{
			{env.ListOfCities, env.ListOfCitiesD, env.ListOfCitiesR, env.ListTypeForm},
			{env.InsertNewCity, env.InsertNewCityD, env.InsertNewCityR, env.DataEntryTypeForm},
			{env.EditCityById, env.EditCityByIdD, env.EditCityByIdR, env.DataEntryTypeForm},
			{env.FindCityByChar, env.FindCityByCharD, env.FindCityByCharR, env.ActionType},
			{env.FindCityById, env.FindCityByIdD, env.FindCityByIdR, env.ActionType},
			{env.DeleteCityById, env.DeleteCityByIdD, env.DeleteCityByIdR, env.ActionType},
			{env.CallDistanceTimeTwoCities, env.CallDistanceTimeTwoCitiesD, env.CallDistanceTimeTwoCitiesR, env.ReportTypeForm},
		}},
		{"Job", []MenuDetail{
			{env.InsertNewJob, env.InsertNewJobD, env.InsertNewJobR, env.DataEntryTypeForm},
			{env.EditJobById, env.EditJobByIdD, env.EditJobByIdR, env.DataEntryTypeForm},
			{env.DeleteJobById, env.DeleteJobByIdD, env.DeleteJobByIdR, env.ActionType},
			{env.ListOfJob, env.ListOfJobD, env.ListOfJobR, env.ListTypeForm},
		}},
		{"Sex Menu", []MenuDetail{
			{env.InsertNewSex, env.InsertNewSexD, env.InsertNewSexR, env.DataEntryTypeForm},
			{env.ListOfSex, env.ListOfSexD, env.ListOfSexR, env.ListTypeForm},
			{env.EditSex, env.EditSexD, env.EditSexR, env.DataEntryTypeForm},
			{env.DeleteSex, env.DeleteSexD, env.DeleteSexR, env.ActionType},
		}},
		{"Sample data", []MenuDetail{
			{env.InsertSomeSamplesData, env.InsertSomeSamplesDataD, env.InsertSomeSamplesDataR, env.DataEntryTypeForm},
			{env.PrintAllData, env.PrintAllDataD, env.PrintAllDataR, env.ReportTypeForm},
		}},
		{"Reports", []MenuDetail{
			{env.ContactReport, env.ContactReportD, env.ContactReportR, env.ReportTypeForm},
		}},
		{"General", []MenuDetail{
			{env.CheckIranNationalCode, env.CheckIranNationalCodeD, env.CheckIranNationalCodeR, env.ListTypeForm},
			{env.RESET, env.RESETD, env.RESETR, env.ActionType},
		}},
	}}
)
