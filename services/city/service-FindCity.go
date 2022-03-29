//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type findCityIDPool struct{}

var FindCityIDPool findCityIDPool

func (allData *findCityIDPool) FindCityByID(inputID uint) (protocol.ResponseStatus, protocol.City) {
	if state, data := storage.FindCityByID(inputID); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.City{}
}

func (allData *findCityIDPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err, cityID := convertor.StrToUint(r.Header.Get("cityID"))
	if err != nil {
		if status, city := allData.FindCityByID(cityID); status.State {
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
}
