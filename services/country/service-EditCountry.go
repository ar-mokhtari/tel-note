package country

import "tel-note/protocol"

func EditCountry(editedCountry protocol.Country) {
	storageService.EditCountry(editedCountry)
}
