//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type editCity struct{}

var EditCity editCity //Pool

//TODO::: change method name to "DO"
func (allData *editCity) Do(req *editCityRequest) (err error) {
	if err := storage.EditCity(req); err != nil {
		return nil
	}
	return err
}

func (allData *editCity) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != env.PatchMethod {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "don't support request"})
	} else {
		if err := json.NewDecoder(r.Body).Decode(&editCityRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCity := new(protocol.City)
		newCity.Id = uint(editCityRequest.Id)
		newCity.Name = editCityRequest.Name
		newCity.EnglishName = editCityRequest.EnglishName
		newCity.AriaCode = editCityRequest.AriaCode
		_, newCity.Lat = convertor.StrToFloat64(editCityRequest.Lat)
		_, newCity.Lng = convertor.StrToFloat64(editCityRequest.Lng)
		if status := allData.Do(*newCity); status.State {
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
}

type editCityRequest struct {
	Id          int
	Name        string
	EnglishName string
	AriaCode    string
	Lat         string
	Lng         string
}

func (ec *editCityRequest) MarshalJson() {

}

func (ec *editCityRequest) UnMarshalJson() {

}
