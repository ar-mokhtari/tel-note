package country

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type countryGetPool struct{}

var GetCountry countryGetPool

func (allData *countryGetPool) GetCountry() []*protocol.Country {
	return storage.GetCountry()
}

func (allData *countryGetPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	result := allData.GetCountry()
	json.NewEncoder(w).Encode(struct {
		State       uint
		ResultCount uint
		Data        []*protocol.Country
	}{200, uint(len(result)), result})
}
