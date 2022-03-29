//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type getCityPool struct{}

var GetCityPool getCityPool

func (allData *getCityPool) GetCities() []*protocol.City {
	return storage.GetCities()
}

func (allData *getCityPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, city := range allData.GetCities() {
		json.NewEncoder(w).Encode(city)
	}
}
