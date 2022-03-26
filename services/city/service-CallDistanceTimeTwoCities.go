package city

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
	"tel-note/services/general"
)

type distanceTimeService struct{}

var DistanceTimeService distanceTimeService

func (handler *distanceTimeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	firstCity := r.Header.Get("firstCity")
	secondCity := r.Header.Get("secondCity")
	firstStatus, dataFirstCity := FindCityByID(general.StrToUint(firstCity))
	secondStatus, dataSecondCity := FindCityByID(general.StrToUint(secondCity))
	if firstStatus.State && secondStatus.State {
		result, state := handler.Do(dataFirstCity, dataSecondCity)
		if state.State {
			json.NewEncoder(w).Encode(
				struct {
					Status     uint
					FirstCity  string
					SecondCity string
					Duration   struct {
						Hour   float32
						Minute float32
					}
					Distance uint
				}{200, dataFirstCity.Name, dataSecondCity.Name,
					struct {
						Hour   float32
						Minute float32
					}{float32(result[0] / 3600), float32(result[0] % 3600 / 60)},
					result[1] / 1000},
			)
		} else {
			json.NewEncoder(w).Encode(
				struct {
					Status  uint
					Massage string
				}{400, "method don't response any"})
		}
	} else {
		json.NewEncoder(w).Encode(
			struct {
				Status           uint
				Massage          string
				FirstCityStatus  bool
				SecondCityStatus bool
			}{
				401,
				"city(ies) not found",
				firstStatus.State,
				secondStatus.State,
			})
	}
}

func (handler *distanceTimeService) Do(cityNoOne, CityNoTwo protocol.City) ([]uint, protocol.ResponseStatus) {
	if data, status := storage.CallTimeDistanceTwoCities(cityNoOne, CityNoTwo); status {
		return data, protocol.ResponseStatus{State: status}
	}
	return []uint{}, protocol.ResponseStatus{State: false}
}
