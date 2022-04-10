package cli

import "tel-note/env"

var (
	GlobalMenu = Menu{[]MenuGroup{
		{"Contact", []MenuDetail{
			{env.NewContactRecord, env.NewContactRecordD, env.NewContactRecordR},
			{env.ListOfContact, env.ListOfContactD, env.ListOfContactR},
			{env.FindOneContactById, env.FindOneContactByIdD, env.FindOneContactByIdR},
			{env.FindContactContainingSomeCharacter, env.FindContactContainingSomeCharacterD, env.FindContactContainingSomeCharacterR},
			{env.FindAndEditContactByContactId, env.FindAndEditContactByContactIdD, env.FindAndEditContactByContactIdR},
			{env.DeleteContactById, env.DeleteContactByIdD, env.DeleteContactByIdR},
			{env.DeleteMultiContactByIds, env.DeleteMultiContactByIdsD, env.DeleteMultiContactByIdsR},
			{env.DeleteAllContacts, env.DeleteAllContactsD, env.DeleteAllContactsR},
		}},
		{"Customer group", []MenuDetail{
			{env.NewCustomerGroup, env.NewCustomerGroupD, env.NewCustomerGroupR},
			{env.NewCustomerGRelation, env.NewCustomerGRelationD, env.NewCustomerGRelationR},
			{env.CustomerGroupList, env.CustomerGroupListD, env.CustomerGroupListR},
			{env.CustomerGroupRelationList, env.CustomerGroupRelationListD, env.CustomerGroupRelationListR},
			{env.FindCustomerByGroupID, env.FindCustomerByGroupIDD, env.FindCustomerByGroupIDR},
		}},
		{"Customer", []MenuDetail{
			{env.NewCustomer, env.NewCustomerD, env.NewCustomerR},
			{env.ListOfCustomer, env.ListOfCustomerD, env.ListOfCustomerR},
			{env.EditCustomerByCustomerId, env.EditCustomerByCustomerIdD, env.EditCustomerByCustomerIdR},
			{env.DeleteCustomerById, env.DeleteCustomerByIdD, env.DeleteCustomerByIdR},
			{env.DeleteMultiCustomerByIds, env.DeleteMultiCustomerByIdsD, env.DeleteMultiCustomerByIdsR},
			{env.DeleteAllCustomers, env.DeleteAllCustomersD, env.DeleteAllCustomersR},
		}},
		{"Person", []MenuDetail{
			{env.NewPerson, env.NewPersonD, env.NewPersonR},
			{env.ListOfPerson, env.ListOfPersonD, env.ListOfPersonR},
			{env.FindOnePersonById, env.FindOnePersonByIdD, env.FindOnePersonByIdR},
			{env.FindPersonContainingSomeCharacter, env.FindPersonContainingSomeCharacterD, env.FindPersonContainingSomeCharacterR},
			{env.FindAndEditPersonByPersonId, env.FindAndEditPersonByPersonIdD, env.FindAndEditPersonByPersonIdR},
			{env.DeletePersonById, env.DeletePersonByIdD, env.DeletePersonByIdR},
			{env.DeleteMultiPersonByIds, env.DeleteMultiPersonByIdsD, env.DeleteMultiPersonByIdsR},
			{env.DeleteAllPersons, env.DeleteAllPersonsD, env.DeleteAllPersonsR},
		}},
		{"Country", []MenuDetail{
			{env.AddCountry, env.AddCountryD, env.AddCountryR},
			{env.EditCountry, env.EditCountryD, env.EditCountryR},
			{env.DeleteCountry, env.DeleteCountryD, env.DeleteCountryR},
			{env.FindCountryByChar, env.FindCountryByCharD, env.FindCountryByCharR},
			{env.CountriesList, env.CountriesListD, env.CountriesListR},
		}},
		{"City", []MenuDetail{
			{env.ListOfCities, env.ListOfCitiesD, env.ListOfCitiesR},
			{env.InsertNewCity, env.InsertNewCityD, env.InsertNewCityR},
			{env.EditCityById, env.EditCityByIdD, env.EditCityByIdR},
			{env.FindCityByChar, env.FindCityByCharD, env.FindCityByCharR},
			{env.FindCityById, env.FindCityByIdD, env.FindCityByIdR},
			{env.DeleteCityById, env.DeleteCityByIdD, env.DeleteCityByIdR},
			{env.CallDistanceTimeTwoCities, env.CallDistanceTimeTwoCitiesD, env.CallDistanceTimeTwoCitiesR},
		}},
		{"Job", []MenuDetail{
			{env.InsertNewJob, env.InsertNewJobD, env.InsertNewJobR},
			{env.EditJobById, env.EditJobByIdD, env.EditJobByIdR},
			{env.DeleteJobById, env.DeleteJobByIdD, env.DeleteJobByIdR},
			{env.ListOfJob, env.ListOfJobD, env.ListOfJobR},
		}},
		{"Sex Menu", []MenuDetail{
			{env.InsertNewSex, env.InsertNewSexD, env.InsertNewSexR},
			{env.ListOfSex, env.ListOfSexD, env.ListOfSexR},
			{env.EditSex, env.EditSexD, env.EditSexR},
			{env.DeleteSex, env.DeleteSexD, env.DeleteSexR},
		}},
		{"Sample data", []MenuDetail{
			{env.InsertSomeSamplesData, env.InsertSomeSamplesDataD, env.InsertSomeSamplesDataR},
			{env.PrintAllData, env.PrintAllDataD, env.PrintAllDataR},
		}},
		{"Reports", []MenuDetail{
			{env.ContactReport, env.ContactReportD, env.ContactReportR},
		}},
		{"General", []MenuDetail{
			{env.CheckIranNationalCode, env.CheckIranNationalCodeD, env.CheckIranNationalCodeR},
			{env.RESET, env.RESETD, env.RESETR},
		}},
	}}
)
