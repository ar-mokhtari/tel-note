package country

import "tel-note/protocol"

func GetCountry() protocol.CountryStorage {
	return storageService.GetCountry()
}
