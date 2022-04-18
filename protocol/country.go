package protocol

import "time"

type (
	Country struct {
		ID           uint      `json:"id"`
		Name         string    `json:"country_name"`
		ShortName    string    `json:"country_short_name"`
		PrePhoneCode uint      `json:"country_phone_code"`
		CapitalID    uint      `json:"capital_id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	CountryServices interface {
		GetCountry() []*Country
		CallCountry() []*Country
		NewCountry(newCountry Country) error
		EditCountry(editedCountry Country) error
		DeleteCountry(IDS []uint) uint
		FindCountryByChar(insertChar string) []*Country
	}
)
