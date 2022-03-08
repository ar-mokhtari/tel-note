package country

import "tel-note/protocol"

func FindCountryByChar(insertChar string) protocol.CountryStorage {
	return storageService.FindCountryByChar(insertChar)
}
