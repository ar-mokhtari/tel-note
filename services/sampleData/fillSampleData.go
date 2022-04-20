package sampleData

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tel-note/config"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
)

type fillData struct{}

var FillData fillData

func (fd *fillData) FillSimpleData() (result [][]string, err error) {
	for _, data := range env.SexDataTest {
		sex.NewSex(*data)
	}
	//for _, data := range env.CityDataTest {
	//	city.NewCity(*data)
	//}
	var cities [][]string
	cities, err = convertor.GetDataFromExcel(config.MainPath+"/env/IranCities.csv", true)
	for _, cityPack := range cities {
		lat, _ := strconv.ParseFloat(cityPack[6], 64)
		lng, _ := strconv.ParseFloat(cityPack[7], 64)
		city.NewCity.Do(city.NewCityRequest{
			Name:        cityPack[3],
			EnglishName: "",
			AriaCode:    "",
			Lat:         lat,
			Lng:         lng,
		})
	}
	for _, data := range env.JobDataTest {
		_ = job.NewJob.Do(job.NewRequest(*data))
	}
	for _, data := range env.PersonDataTest {
		person.NewPerson(*data)
	}
	for _, data := range env.CustomerGroupExm {
		customer.NewGroup.Do(data.GroupName)
	}
	for _, data := range env.CustomerRelation {
		customer.NewGrpRelation.Do(data.CustomerID, data.GroupID)
	}
	for _, data := range env.CustomersDataTest.CustomerData {
		readyToFill := customer.NewRequest{
			PersonID:    data.PersonID,
			Description: data.Description,
		}
		customer.AddCustomer.Do(readyToFill)
	}
	//contact have to locate in end list, because it's elements has dependent to upper steps (city/job/...)
	for _, data := range env.ContactDataTest {
		contact.NewContact.Do(contact.NewContactRequest(*data))
	}
	//call api test for fill countries
	countries := country.CallCountry.Do()
	for _, data := range countries {
		country.NewCountry.Do(country.NewRequest(*data))
	}
	return cities, err
}

func (fd *fillData) Do() error {
	_, err := fd.FillSimpleData()
	return err
}

func (fd *fillData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method != env.PostMethod {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	} else {
		if err := fd.Do(); err != nil {
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
}
