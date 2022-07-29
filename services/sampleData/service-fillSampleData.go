package sampleData

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/config"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/services/city"
	"github.com/ar-mokhtari/tel-note/services/contact"
	"github.com/ar-mokhtari/tel-note/services/country"
	"github.com/ar-mokhtari/tel-note/services/customer"
	"github.com/ar-mokhtari/tel-note/services/job"
	"github.com/ar-mokhtari/tel-note/services/person"
	"github.com/ar-mokhtari/tel-note/services/sex"
	"net/http"
	"strconv"
)

type fillData struct{}
type cityStruct struct {
	cityInput
}
type cityInput struct {
	id          uint
	name        string  `json:"id"`
	englishName string  `json:"name"`
	ariaCode    string  `json:"english_name"`
	lat         float64 `json:"aria_code"`
	lng         float64 `json:"lat"`
}

func (c *cityInput) ID() uint            { return c.id }
func (c *cityInput) Name() string        { return c.name }
func (c *cityInput) EnglishName() string { return c.englishName }
func (c *cityInput) AriaCode() string    { return c.ariaCode }
func (c *cityInput) Lat() float64        { return c.lat }
func (c *cityInput) Lng() float64        { return c.lng }

var FillData fillData

func (fd *fillData) FillSimpleData() (result [][]string, err error) {
	for _, data := range env.SexDataTest {
		sex.AddSex.Do(sex.AddRequest(*data))
	}
	//for _, data := range env.CityDataTest {
	//	city.NewCity(*data)
	//}
	var cities [][]string
	cities, err = convertor.GetDataFromExcel(config.MainPath+"/env/IranCities.csv", true)
	for _, cityPack := range cities {
		lat, _ := strconv.ParseFloat(cityPack[6], 64)
		lng, _ := strconv.ParseFloat(cityPack[7], 64)
		var inputData *city.City
		inputData.id = 0
		inputData.name = cityPack[0]
		inputData.englishName = cityPack[1]
		inputData.ariaCode = cityPack[2]
		inputData.lat = lat
		inputData.lng = lng
		temp := city.SetCity(inputData)

		city.NewCity.Do(temp)
	}
	for _, data := range env.JobDataTest {
		_ = job.NewJob.Do(job.NewRequest(*data))
	}
	for _, data := range env.PersonDataTest {
		person.AddPerson.Do(person.AddRequest(*data))
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
