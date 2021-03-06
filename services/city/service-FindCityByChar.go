//go:build memory

package city

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findByCharService struct{}

var FindByCharService findByCharService

func (fcs *findByCharService) FindCityByChar(inputChar string) (protocol.ResponseStatus, []uint) {
	if state, data := storage.FindCityByChar(inputChar); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, []uint{}
}

func (fcs *findByCharService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case env.GetMethod:
		var char struct {
			Char string `json:"char"`
		}
		if err := json.NewDecoder(r.Body).Decode(&char); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if status, data := fcs.FindCityByChar(char.Char); status.State {
			var dataResult []protocol.City
			for _, result := range data {
				_, city := storage.GetCity(result)
				dataResult = append(dataResult, city)
			}
			json.NewEncoder(w).Encode(struct {
				Status      uint
				ResultCount uint
				Data        []protocol.City
			}{200, uint(len(data)), dataResult})
		} else {
			json.NewEncoder(w).Encode(struct {
				State   uint
				Message string
			}{400, "not found"})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
