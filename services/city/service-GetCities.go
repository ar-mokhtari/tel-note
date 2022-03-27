//go:build memory

package city

import "tel-note/protocol"

func GetCities() []*protocol.City {
	return storage.GetCities()
}
