package protocol

import "time"

type (
	Country struct {
		ID           uint
		Name         string `json:"country_name"`
		ShortName    string `json:"country_short_name"`
		PrePhoneCode uint   `json:"country_phone_code"`
		CapitalID    uint
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	CountryServices interface {
		GetCountry() []*Country
		CallCountry() []*Country
		NewCountry(newCountry Country)
		EditCountry(editedCountry Country)
		DeleteCountry(IDS []uint) uint
		FindCountryByChar(insertChar string) []*Country
	}
)
