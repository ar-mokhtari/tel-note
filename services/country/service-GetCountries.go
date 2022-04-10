package country

import "tel-note/protocol"

func GetCountry() []*protocol.Country {
	return storageService.GetCountry()
}
