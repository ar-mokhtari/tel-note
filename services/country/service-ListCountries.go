package country

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getCountry struct{}

var GetCountry getCountry

func (gc *getCountry) Do() (res []*protocol.Country) {
	for _, data := range storage.ListCountries() {
		if country := storage.FindCountryByID(data); country != nil {
			res = append(res, country)
		}
	}
	return res
}

func (gc *getCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		result := gc.Do()
		json.NewEncoder(w).Encode(struct {
			State       uint
			ResultCount uint
			Data        []*protocol.Country
		}{200, uint(len(result)), result})
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
