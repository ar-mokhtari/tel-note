//go:build memory

package city

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findCityID struct{}

var FindCityID findCityID

func (fci *findCityID) FindCityByID(inputID uint) protocol.City {
	if _, data := storage.GetCity(inputID); data != nil {
		return data
	}
	return nil
}

func (fci *findCityID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case env.GetMethod:
		err, cityID := convertor.StrToUint(r.FormValue("cityID"))
		if err == nil {
			if city := fci.FindCityByID(cityID); city != nil {
				json.NewEncoder(w).Encode(struct {
					Status uint
					Data   protocol.City
				}{200, city})
			} else {
				json.NewEncoder(w).Encode(struct {
					State   uint
					Message string
				}{400, "not found"})
			}
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
