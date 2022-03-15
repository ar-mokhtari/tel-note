package fillSampleData

import (
	"fmt"
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
	for _, data := range env.CustomersDataTest.Data {
		customer.AddCustomer(*data)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range (env.ContactDataTest).Data {
		contact.NewContact(*data)
	}
	//call api test for fill countries
	fmt.Println("" +
		"importing countries ...")
	countries := country.CallCountry()
	for _, data := range countries {
		country.NewCountry(*data)
	}
	return true
}
