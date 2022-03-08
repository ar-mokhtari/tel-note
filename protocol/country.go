package protocol

type (
	Country struct {
		Name         string
		CapitalID    uint
		ShortName    string
		PrePhoneCode string
	}
	CountryStorage  []*Country
	CountryServices interface {
		GetCountry() CountryStorage
	}
)
