//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type editCityPool struct{}

var EditCityPool editCityPool

func (allData *editCityPool) EditCityByID(NewCity protocol.City) protocol.ResponseStatus {
	if storage.EditCity(NewCity) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}

func (allData *editCityPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inputCity struct {
		Id          int
		Name        string
		EnglishName string
		AriaCode    string
		Lat         string
		Lng         string
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &inputCity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newCity := new(protocol.City)
	newCity.Id = uint(inputCity.Id)
	newCity.Name = inputCity.Name
	newCity.EnglishName = inputCity.EnglishName
	newCity.AriaCode = inputCity.AriaCode
	err, newCity.Lat = convertor.StrToFloat64(inputCity.Lat)
	err, newCity.Lng = convertor.StrToFloat64(inputCity.Lng)
	if status := allData.EditCityByID(*newCity); status.State {
		json.NewEncoder(w).Encode(struct {
			Status  uint
			Message string
		}{200, fmt.Sprintf("City #%v edited", newCity.Id)})
	} else {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "could not be edited"})
	}
}
