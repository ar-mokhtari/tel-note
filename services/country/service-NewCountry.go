package country

import "tel-note/protocol"

func NewCountry(newCountry protocol.Country) {
	storage.NewCountry(newCountry)
}
