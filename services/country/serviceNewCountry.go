package country

import "tel-note/protocol"

func NewCountry(newCountry protocol.Country) {
	storageService.NewCountry(newCountry)
}
