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

func (fci *findCityID) FindCityByID(inputID uint) (protocol.ResponseStatus, protocol.City) {
	if state, data := storage.FindCityByID(inputID); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.City{}
}

func (fci *findCityID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case env.GetMethod:
		err, cityID := convertor.StrToUint(r.FormValue("cityID"))
		if err == nil {
			if status, city := fci.FindCityByID(cityID); status.State {
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
