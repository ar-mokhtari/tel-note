package sampleData

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/globalVars"
	"tel-note/services/person"
)

type AllDataCollection struct {
	contact               []*protocol.Contact
	customer              map[uint]*protocol.Customer
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
	result.customer = globalVars.CustomerMapStore
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
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	result := fillData.DoGetData()
	json.NewEncoder(w).Encode(struct {
		State                     uint
		CityData                  []*protocol.City
		ContactData               []*protocol.Contact
		CustomerData              map[uint]*protocol.Customer
		CustomerGroupData         []*protocol.CustomerGroup
		CustomerGroupRelationData []*protocol.CustomerGroupRelation
		CustomerRelationData      protocol.CustomerGRelationStorage
		PersonData                []*protocol.Person
		CountriesData             []*protocol.Country
	}{200,
		result.cities,
		result.contact,
		result.customer,
		result.customerGroup,
		result.customerGroupRelation,
		result.customerRelation,
		result.person,
		result.countries,
	})

}
