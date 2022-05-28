//go:build !mysql || !mongodb

package country

import (
	"errors"
	"github.com/ar-mokhtari/tel-note/protocol"
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

func (sm *storageMemory) ListCountries() (res []uint) {
	for _, data := range sm.CountryData {
		res = append(res, data.ID)
	}
	return res
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
func (sm *storageMemory) FindCountryByID(insertID uint) *protocol.Country {
	for _, country := range sm.CountryData {
		if country.ID == insertID {
			return country
		}
	}
	return nil
}
