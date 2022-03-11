package city

import "tel-note/protocol"

func CallDistanceTimeTwoCities(cityNoOne, CityNoTwo protocol.City) ([]uint, protocol.ResponseStatus) {
	if data, status := storage.CallTimeDistanceTwoCities(cityNoOne, CityNoTwo); status {
		return data, protocol.ResponseStatus{State: status}
	}
	return []uint{}, protocol.ResponseStatus{State: false}
}
