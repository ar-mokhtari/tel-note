package country

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"tel-note/env/apis"
	"tel-note/protocol"
	"time"
)

type storageCountry protocol.CountryStorage

func (AllCountries *storageCountry) GetCountry() protocol.CountryStorage {
	return protocol.CountryStorage(*AllCountries)
}

func (AllCountries *storageCountry) CallCountry() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", apis.UniversalTutorialURL, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", apis.UniversalTutorialAPIKey)
	response, callErr := client.Do(req)
	if callErr != nil {
		fmt.Println(callErr.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
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
		CreatedAt:    time.Now(),
	})
}

func (AllCountries *storageCountry) EditCountry(editedCountry protocol.Country) {
	for index, country := range *AllCountries {
		if country.ID == editedCountry.ID {
			if editedCountry.Name != "" {
				(*AllCountries)[index].Name = editedCountry.Name
			}
			if editedCountry.ShortName != "" {
				(*AllCountries)[index].ShortName = editedCountry.ShortName
			}
			if editedCountry.PrePhoneCode != 0 {
				(*AllCountries)[index].PrePhoneCode = editedCountry.PrePhoneCode
			}
			if editedCountry.CapitalID != 0 {
				(*AllCountries)[index].CapitalID = editedCountry.CapitalID
			}
			if editedCountry.Name != "" ||
				editedCountry.ShortName != "" ||
				editedCountry.PrePhoneCode != 0 ||
				editedCountry.CapitalID != 0 {
				(*AllCountries)[index].UpdatedAt = time.Now()
			}
		}
	}
}
func (AllCountries *storageCountry) DeleteCountry(IDS []uint) uint {
	var counter uint
	for index, country := range *AllCountries {
		for _, ID := range IDS {
			if country.ID == ID {
				*AllCountries = append((*AllCountries)[:index], (*AllCountries)[index+1:]...)
				counter += 1
			}
		}
	}
	return counter
}
func (AllCountries *storageCountry) FindCountryByChar(insertChar string) protocol.CountryStorage {
	var result protocol.CountryStorage
	for _, country := range *AllCountries {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(insertChar)) {
			result = append(result, country)
		}
	}
	return result
}
