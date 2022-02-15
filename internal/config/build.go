package config

import (
	"tel-note/internal/storage/memory"
)

var (
	AllContact = make([]*memory.Contact, 0)
	AllCity    = make([]*memory.City, 0)
	MainData   = memory.AllData{ContactData: AllContact, CityData: AllCity}
)
