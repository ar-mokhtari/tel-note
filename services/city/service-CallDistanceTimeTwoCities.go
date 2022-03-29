//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/SDK/neshan"
	"tel-note/lib/callAPI"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type distanceTimeService struct{}

var DistanceTimeService distanceTimeService

func (handler *distanceTimeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	firstCity := r.Header.Get("firstCity")
	secondCity := r.Header.Get("secondCity")
	_, uintFirstCity := convertor.StrToUint(firstCity)
	_, uintSecondCity := convertor.StrToUint(secondCity)
	firstStatus, dataFirstCity := FindCityByID(uintFirstCity)
	secondStatus, dataSecondCity := FindCityByID(uintSecondCity)
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

func (handler *distanceTimeService) Do(cityNoOne, cityNoTwo protocol.City) ([]uint, protocol.ResponseStatus) {
	MapParams := map[string]string{
		"type":         "car",
		"origins":      fmt.Sprintf("%f,%f", cityNoOne.Lat, cityNoOne.Lng),
		"destinations": fmt.Sprintf("%f,%f", cityNoTwo.Lat, cityNoTwo.Lng),
	}
	MapHeaders := map[string]string{
		"Api-Key": neshan.ApiKey,
	}
	responseData := callAPI.CallGetAPIs(
		neshan.DistanceMatrixURL,
		MapParams,
		MapHeaders,
	)
	var AllResult neshan.DistanceMatrix
	json.Unmarshal(responseData, &AllResult)
	switch AllResult.Rows == nil {
	case true:
		return []uint{}, protocol.ResponseStatus{State: false}
	case false:
		return []uint{
			AllResult.Rows[0].Elements[0].Duration.Value,
			AllResult.Rows[0].Elements[0].Distance.Value,
		}, protocol.ResponseStatus{State: true}
	}
	return []uint{}, protocol.ResponseStatus{State: false}
}
