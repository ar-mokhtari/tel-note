//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type newCityPool struct{}

var NewCityPool newCityPool

func (allData *newCityPool) NewCity(city protocol.City) (status protocol.ResponseStatus) {
	if status.State = storage.NewCity(city); status.State {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}

func (allData *newCityPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != env.PostMethod {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	} else {
		var inputCity struct {
			Name        string
			EnglishName string
			AriaCode    string
			Lat         string
			Lng         string
		}
		if err := json.NewDecoder(r.Body).Decode(&inputCity); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCity := new(protocol.City)
		newCity.Name = inputCity.Name
		newCity.EnglishName = inputCity.EnglishName
		newCity.AriaCode = inputCity.AriaCode
		_, newCity.Lat = convertor.StrToFloat64(inputCity.Lat)
		_, newCity.Lng = convertor.StrToFloat64(inputCity.Lng)
		if status := allData.NewCity(*newCity); status.State {
			json.NewEncoder(w).Encode(struct {
				Status  uint
				Message string
			}{200, "new city added"})
		} else {
			json.NewEncoder(w).Encode(struct {
				State   uint
				Message string
			}{400, "could not add"})
		}
	}
}
