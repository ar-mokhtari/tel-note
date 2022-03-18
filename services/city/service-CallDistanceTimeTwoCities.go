package city

import (
	"fmt"
	"net/http"
	"tel-note/protocol"
	"tel-note/services/general"
)

type serviceHandler struct{}

var ServiceHandler serviceHandler

func (handler *serviceHandler) CallDistanceTimeTwoCities(cityNoOne, CityNoTwo protocol.City) ([]uint, protocol.ResponseStatus) {
	if data, status := storage.CallTimeDistanceTwoCities(cityNoOne, CityNoTwo); status {
		return data, protocol.ResponseStatus{State: status}
	}
	return []uint{}, protocol.ResponseStatus{State: false}
}

func (handler *serviceHandler) DistanceTimeServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	firstCity := r.Header.Get("firstCity")
	secondCity := r.Header.Get("secondCity")
	firstStatus, dataFirstCity := FindCityByID(general.StrToUint(firstCity))
	secondStatus, dataSecondCity := FindCityByID(general.StrToUint(secondCity))
	if firstStatus.State && secondStatus.State {
		result, state := handler.CallDistanceTimeTwoCities(dataFirstCity, dataSecondCity)
		if state.State {
			fmt.Fprintf(w,
				"First city: %v Second city: %v \n", dataFirstCity.Name, dataSecondCity.Name)
			fmt.Fprintf(w,
				"Time duration with online traffic is: %v hours ", result[0]/3600)
			if minutes := result[0] % 3600; minutes != 0 {
				fmt.Fprintf(w, "and %v min", minutes/60)
			}
			fmt.Fprintf(w, "\nDistance is: %v km\n", result[1]/1000)
		} else {
			fmt.Fprintf(w, "method send nothing")
		}
	} else {
		fmt.Fprintf(w, "not found")
	}
}
