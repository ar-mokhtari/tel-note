package fillSampleData

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tel-note/env"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/general"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
)

type fillData struct{}
type AllDataCollection struct {
	contact               []*protocol.Contact
	customer              []*protocol.Customer
	customerGroup         []*protocol.CustomerGroup
	customerGroupRelation []*protocol.CustomerGroupRelation
	customerRelation      protocol.CustomerGRelationStorage
	person                []*protocol.Person
	countries             []*protocol.Country
	cities                []*protocol.City
}

var FillDataStruct fillData

func FillSimpleDataInMainData() bool {
	for _, data := range env.SexDataTest {
		sex.NewSex(*data)
	}
	//for _, data := range env.CityDataTest {
	//	city.NewCity(*data)
	//}
	fmt.Println("importing cities ...")
	cities, _ := general.GetDataFromExcel("././env/IranCities.csv", true)
	for _, cityPack := range cities {
		lat, _ := strconv.ParseFloat(cityPack[6], 64)
		lng, _ := strconv.ParseFloat(cityPack[7], 64)
		city.NewCity(protocol.City{
			Name:     cityPack[3],
			AriaCode: "",
			Lat:      lat,
			Lng:      lng,
		})
	}
	for _, data := range env.JobDataTest {
		job.NewJob(*data)
	}
	for _, data := range env.PersonDataTest {
		person.NewPerson(*data)
	}
	for _, data := range env.CustomerGroup {
		customer.NewGroup(data.GroupName)
	}
	for _, data := range env.CustomerRelation {
		customer.NewRelation(data.CustomerID, data.GroupID)
	}
	for _, data := range env.CustomersDataTest.CustomerData {
		customer.AddCustomer(*data)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range env.ContactDataTest {
		contact.PoolContact.NewContact(*data)
	}
	//call api test for fill countries
	fmt.Println("" +
		"importing countries ...")
	countries := country.CallCountry.Do()
	for _, data := range countries {
		country.NewCountry(*data)
	}
	return true
}

func (fillData *fillData) DoFillData() bool {
	return FillSimpleDataInMainData()
}

func (fillData *fillData) DoGetData() (result AllDataCollection) {
	result.contact = contact.GetPool.GetContacts()
	result.customerGroup = customer.GetCustomerGroup()
	result.customerGroupRelation = customer.GetCustomerGroupRelation()
	result.person = person.GetPersons()
	result.countries = country.GetCountry()
	result.cities = city.GetCities()
	return result
}

func (fillData *fillData) ServeGetDataHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fillData.DoGetData().contact)
	json.NewEncoder(w).Encode(fillData.DoGetData().customer)
	json.NewEncoder(w).Encode(fillData.DoGetData().customerGroup)
	json.NewEncoder(w).Encode(fillData.DoGetData().customerGroupRelation)
	json.NewEncoder(w).Encode(fillData.DoGetData().customerRelation)
	json.NewEncoder(w).Encode(fillData.DoGetData().person)
	json.NewEncoder(w).Encode(fillData.DoGetData().countries)
	json.NewEncoder(w).Encode(fillData.DoGetData().cities)
}

func (fillData *fillData) ServeFillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if fillData.DoFillData() {
		fmt.Fprintf(w, "data insert successfuly")
	} else {
		fmt.Fprintf(w, "some thing wrong")
	}
}
