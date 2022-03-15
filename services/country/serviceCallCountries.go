package country

import "tel-note/protocol"

func CallCountry() protocol.CountryStorage {
	return storageService.CallCountry()
}
