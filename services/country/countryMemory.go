package country

import (
	"encoding/json"
	"strings"
	"tel-note/SDK/Universal"
	"tel-note/protocol"
	"tel-note/services/general"
	"time"
)

type storageCountry protocol.CountryStorage

func (AllCountries *storageCountry) GetCountry() protocol.CountryStorage {
	return protocol.CountryStorage(*AllCountries)
}

func (AllCountries *storageCountry) CallCountry() protocol.CountryStorage {
	//generate new token
	MapTokenHeaders := map[string]string{
		"Accept":     "application/json",
		"api-token":  Universal.UniversaltutorialToken,
		"user-email": Universal.Email,
	}
	responseTokenData := general.CallGetAPIs(
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
	responseData := general.CallGetAPIs(
		Universal.GetCountryURL,
		map[string]string{},
		MapHeaders,
	)
	var AllResult protocol.CountryStorage
	json.Unmarshal(responseData, &AllResult)
	return AllResult
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
