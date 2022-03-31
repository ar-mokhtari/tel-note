package sampleData

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/person"
)

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
type getData struct{}

var GetDataStruct getData

func (fillData *getData) DoGetData() (result AllDataCollection) {
	result.contact = contact.GetPool.GetContacts()
	result.customerGroup = customer.GetCustomerGroup()
	result.customerGroupRelation = customer.GetCustomerGroupRelation()
	result.person = person.GetPersons()
	result.countries = country.GetCountry()
	result.cities = city.GetCityPool.GetCities()
	return result
}

func (fillData *getData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	result := fillData.DoGetData()
	json.NewEncoder(w).Encode(struct {
		State    uint
		CityData []*protocol.City
	}{200, result.cities})
	//json.NewEncoder(w).Encode(result.contact)
	//json.NewEncoder(w).Encode(result.customer)
	//json.NewEncoder(w).Encode(result.customerGroup)
	//json.NewEncoder(w).Encode(result.customerGroupRelation)
	//json.NewEncoder(w).Encode(result.customerRelation)
	//json.NewEncoder(w).Encode(result.person)
	//json.NewEncoder(w).Encode(result.countries)
}
