//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type getCity struct{}

var GetCity getCity

func (gc *getCity) Do() []*protocol.City {
	return storage.GetCities()
}

func (gc *getCity) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	switch r.Method {
	case env.GetMethod:
		result := gc.Do()
		json.NewEncoder(w).Encode(struct {
			State       uint
			CountResult uint
			Data        []*protocol.City
		}{200, uint(len(result)), result})
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
