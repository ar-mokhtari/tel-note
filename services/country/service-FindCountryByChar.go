package country

import "tel-note/protocol"

func FindCountryByChar(insertChar string) []*protocol.Country {
	return storageService.FindCountryByChar(insertChar)
}
