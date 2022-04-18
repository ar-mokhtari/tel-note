//go:build !mysql || !mongodb

package country

import (
	"encoding/json"
	"strings"
	"tel-note/SDK/Universal"
	"tel-note/lib/callAPI"
	"tel-note/protocol"
	"time"
)

type (
	storageMemory struct {
		CountryData []*protocol.Country
	}
)

var (
	storage storageMemory
	_       protocol.CountryServices = &storage
)

func (allCountries *storageMemory) GetCountry() []*protocol.Country {
	return allCountries.CountryData
}

func (allCountries *storageMemory) CallCountry() []*protocol.Country {
	//generate new token
	MapTokenHeaders := map[string]string{
		"Accept":     "application/json",
		"api-token":  Universal.UniversaltutorialToken,
		"user-email": Universal.Email,
	}
	responseTokenData := callAPI.CallGetAPIs(
		Universal.GetTokenURL,
		map[string]string{},
		MapTokenHeaders,
	)
	//var token struct {
	//	authToken string `json:"auth_token"`
	//}
	token := map[string]string{
		"auth_token": "",
	}
	json.Unmarshal(responseTokenData, &token)
	//call api with new token
	MapHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + string(token["auth_token"]),
	}
	responseData := callAPI.CallGetAPIs(
		Universal.GetCountryURL,
		map[string]string{},
		MapHeaders,
	)
	var AllResult []*protocol.Country
	json.Unmarshal(responseData, &AllResult)
	return AllResult
}

func (allCountries *storageMemory) NewCountry(newCountry protocol.Country) {
	var LastID uint
	for _, country := range allCountries.CountryData {
		if country.ID > LastID {
			LastID = country.ID
		}
	}
	LastID += 1
	allCountries.CountryData = append(allCountries.CountryData, &protocol.Country{
		ID:           LastID,
		Name:         newCountry.Name,
		CapitalID:    newCountry.CapitalID,
		ShortName:    newCountry.ShortName,
		PrePhoneCode: newCountry.PrePhoneCode,
		CreatedAt:    time.Now(),
	})
}

func (allCountries *storageMemory) EditCountry(editedCountry protocol.Country) {
	for index, country := range allCountries.CountryData {
		if country.ID == editedCountry.ID {
			if editedCountry.Name != "" {
				(allCountries.CountryData)[index].Name = editedCountry.Name
			}
			if editedCountry.ShortName != "" {
				(allCountries.CountryData)[index].ShortName = editedCountry.ShortName
			}
			if editedCountry.PrePhoneCode != 0 {
				(allCountries.CountryData)[index].PrePhoneCode = editedCountry.PrePhoneCode
			}
			if editedCountry.CapitalID != 0 {
				(allCountries.CountryData)[index].CapitalID = editedCountry.CapitalID
			}
			if editedCountry.Name != "" ||
				editedCountry.ShortName != "" ||
				editedCountry.PrePhoneCode != 0 ||
				editedCountry.CapitalID != 0 {
				(allCountries.CountryData)[index].UpdatedAt = time.Now()
			}
		}
	}
}
func (allCountries *storageMemory) DeleteCountry(IDS []uint) uint {
	//TODO::: when first input 4,5,13,120 then input 2,1,3,4,5,13,120 there was bug
	var counter uint
	for index, country := range allCountries.CountryData {
		for _, ID := range IDS {
			if country.ID == uint(ID) {
				allCountries.CountryData = append((allCountries.CountryData)[:index], (allCountries.CountryData)[index+1:]...)
				counter += 1
			}
		}
	}
	return counter
}
func (allCountries *storageMemory) FindCountryByChar(insertChar string) []*protocol.Country {
	var result []*protocol.Country
	for _, country := range allCountries.CountryData {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(insertChar)) {
			result = append(result, country)
		}
	}
	return result
}
