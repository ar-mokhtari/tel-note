package protocol

type (
	Country struct {
		ID           uint
		Name         string `json:"country_name"`
		ShortName    string `json:"country_short_name"`
		PrePhoneCode uint   `json:"country_phone_code"`
		CapitalID    uint
	}
	CountryStorage  []*Country
	CountryServices interface {
		GetCountry() CountryStorage
		CallCountry()
		NewCountry(newCountry Country)
	}
)
