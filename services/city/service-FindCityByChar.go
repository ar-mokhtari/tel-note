//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type findByCharService struct{}

var FindByCharService findByCharService

func (allData *findByCharService) FindCityByChar(inputChar string) (protocol.ResponseStatus, []uint) {
	if state, data := storage.FindCityByChar(inputChar); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, []uint{}
}

func (allData *findByCharService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var char struct {
		Char string `json:"char"`
	}
	if err := json.NewDecoder(r.Body).Decode(&char); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if status, data := allData.FindCityByChar(char.Char); status.State {
		type cityResult struct {
			CityCode uint
			CityName string
		}
		var dataResult []cityResult
		for _, result := range data {
			_, city := storage.FindCityByID(result)
			dataResult = append(dataResult, cityResult{city.Id, city.Name})
		}
		json.NewEncoder(w).Encode(struct {
			Status      uint
			ResultCount uint
			Data        []cityResult
		}{200, uint(len(data)), dataResult})
	} else {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "not found"})
	}
}
