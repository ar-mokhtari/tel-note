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
	var char struct{ char string }
	err := json.NewDecoder(r.Body).Decode(&char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if status, data := allData.FindCityByChar(char.char); status.State {
		for _, result := range data {
			_, city := storage.FindCityByID(result)
			json.NewEncoder(w).Encode(struct {
				CityCode uint
				CityName string
			}{city.Id, city.Name})
		}
	}
}
