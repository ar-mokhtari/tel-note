package city

import "tel-note/protocol"

func CallDistanceTimeTwoCities(cityNoOne, CityNoTwo protocol.City) []string {
	return storage.CallTimeDistanceTwoCities(cityNoOne, CityNoTwo)
}
