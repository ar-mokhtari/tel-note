package fillSampleData

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tel-note/config"
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

func (fillData *fillData) FillSimpleDataInMainData() (result [][]string, err error) {
	for _, data := range env.SexDataTest {
		sex.NewSex(*data)
	}
	//for _, data := range env.CityDataTest {
	//	city.NewCity(*data)
	//}
	var cities [][]string
	cities, err = general.GetDataFromExcel(config.MainPath+"env/IranCities.csv", true)
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
	countries := country.CallCountry.Do()
	for _, data := range countries {
		country.NewCountry(*data)
	}
	return cities, err
}

func (fillData *fillData) DoFillData() error {
	_, err := fillData.FillSimpleDataInMainData()
	return err
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
	result := fillData.DoGetData()
	json.NewEncoder(w).Encode(result.cities)
	json.NewEncoder(w).Encode(result.contact)
	json.NewEncoder(w).Encode(result.customer)
	json.NewEncoder(w).Encode(result.customerGroup)
	json.NewEncoder(w).Encode(result.customerGroupRelation)
	json.NewEncoder(w).Encode(result.customerRelation)
	json.NewEncoder(w).Encode(result.person)
	json.NewEncoder(w).Encode(result.countries)
}

func (fillData *fillData) ServeFillData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := fillData.DoFillData(); err != nil {
		json.NewEncoder(w).Encode(struct {
			Err string
		}{fmt.Sprintf("error is: %v", err)})

	} else {
		json.NewEncoder(w).Encode(struct {
			State   int
			Message string
		}{200, "Data insert successfully"})
	}
}
