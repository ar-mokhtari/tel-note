package country

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type callCountry struct{}

var CallCountry callCountry

func (allCountry *callCountry) Do() []*protocol.Country {
	return storageService.CallCountry()
}
func (allCountry *callCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allCountry.Do())
}
