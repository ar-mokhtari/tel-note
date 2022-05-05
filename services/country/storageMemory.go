//go:build !mysql || !mongodb

package country

import (
	"encoding/json"
	"errors"
	"github.com/ar-mokhtari/tel-note/lib/callAPI"
	"github.com/ar-mokhtari/tel-note/protocol"
	"github.com/ar-mokhtari/tel-note/sdk/universal"
	"strings"
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

func (sm *storageMemory) GetCountry() []*protocol.Country {
	return sm.CountryData
}

func (sm *storageMemory) CallCountry() []*protocol.Country {
	//generate new token
	MapTokenHeaders := map[string]string{
		"Accept":     "application/json",
		"api-token":  universal.UniversaltutorialToken,
		"user-email": universal.Email,
	}
	responseTokenData := callAPI.CallGetAPIs(
		universal.GetTokenURL,
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
		universal.GetCountryURL,
		map[string]string{},
		MapHeaders,
	)
	var AllResult []*protocol.Country
	json.Unmarshal(responseData, &AllResult)
	return AllResult
}

func (sm *storageMemory) NewCountry(newCountry protocol.Country) error {
	var LastID uint
	for _, country := range sm.CountryData {
		if country.ID > LastID {
			LastID = country.ID
		}
	}
	LastID += 1
	sm.CountryData = append(sm.CountryData, &protocol.Country{
		ID:           LastID,
		Name:         newCountry.Name,
		CapitalID:    newCountry.CapitalID,
		ShortName:    newCountry.ShortName,
		PrePhoneCode: newCountry.PrePhoneCode,
		CreatedAt:    time.Now(),
	})
	return nil
}

func (sm *storageMemory) EditCountry(editedCountry protocol.Country) error {
	for index, country := range sm.CountryData {
		if country.ID == editedCountry.ID {
			if editedCountry.Name != "" {
				(sm.CountryData)[index].Name = editedCountry.Name
			}
			if editedCountry.ShortName != "" {
				(sm.CountryData)[index].ShortName = editedCountry.ShortName
			}
			if editedCountry.PrePhoneCode != 0 {
				(sm.CountryData)[index].PrePhoneCode = editedCountry.PrePhoneCode
			}
			if editedCountry.CapitalID != 0 {
				(sm.CountryData)[index].CapitalID = editedCountry.CapitalID
			}
			if editedCountry.Name != "" ||
				editedCountry.ShortName != "" ||
				editedCountry.PrePhoneCode != 0 ||
				editedCountry.CapitalID != 0 {
				(sm.CountryData)[index].UpdatedAt = time.Now()
			}
			return nil
		}
	}
	return errors.New("country not found")
}
func (sm *storageMemory) DeleteCountry(IDS []uint) uint {
	//TODO::: when first input 4,5,13,120 then input 2,1,3,4,5,13,120 there was bug
	var counter uint
	for index, country := range sm.CountryData {
		for _, ID := range IDS {
			if country.ID == uint(ID) {
				sm.CountryData = append((sm.CountryData)[:index], (sm.CountryData)[index+1:]...)
				counter += 1
			}
		}
	}
	return counter
}
func (sm *storageMemory) FindCountryByChar(insertChar string) []*protocol.Country {
	var result []*protocol.Country
	for _, country := range sm.CountryData {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(insertChar)) {
			result = append(result, country)
		}
	}
	return result
}
