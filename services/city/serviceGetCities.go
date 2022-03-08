package city

import "tel-note/protocol"

func GetCities() protocol.CityStorage {
	return storage.GetCities()
}
