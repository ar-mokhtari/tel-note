package country

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tel-note/env/apis"
	"tel-note/protocol"
)

type storageCountry protocol.CountryStorage

func (AllCountries *storageCountry) GetCountry() protocol.CountryStorage {
	return protocol.CountryStorage(*AllCountries)
}

func (AllCountries *storageCountry) CallCountry() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apis.UniversalTutorialURL, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", apis.UniversalTutorialAPIKey)
	response, callErr := client.Do(req)
	if callErr != nil {
		fmt.Println(callErr.Error())
		os.Exit(1)
	}
	responseData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalln(readErr)
	}
	var AllResult protocol.CountryStorage
	json.Unmarshal(responseData, &AllResult)
	for _, country := range AllResult {
		storageService.NewCountry(*country)
	}
}

func (AllCountries *storageCountry) NewCountry(newCountry protocol.Country) {
	var LastID uint
	for _, country := range *AllCountries {
		if country.ID > LastID {
			LastID = country.ID
		}
	}
	LastID += 1
	*AllCountries = append(*AllCountries, &protocol.Country{
		ID:           LastID,
		Name:         newCountry.Name,
		CapitalID:    newCountry.CapitalID,
		ShortName:    newCountry.ShortName,
		PrePhoneCode: newCountry.PrePhoneCode,
	})
}
