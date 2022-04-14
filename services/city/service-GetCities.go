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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	result := allData.GetCities()
	json.NewEncoder(w).Encode(struct {
		State       uint
		CountResult uint
		Data        []*protocol.City
	}{200, uint(len(result)), result})
}
