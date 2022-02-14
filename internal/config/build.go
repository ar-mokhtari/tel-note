package config

import "tel-note/internal/storage"

var (
	AllContact = make([]*storage.Contact, 0)
	AllCity    = make([]*storage.City, 0)
	MainData   = storage.AllData{ContactData: AllContact, CityData: AllCity}
)
