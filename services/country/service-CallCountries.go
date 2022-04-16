package country

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type callCountry struct{}

var CallCountry callCountry

func (cc *callCountry) Do() []*protocol.Country {
	return storage.CallCountry()
}
func (cc *callCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cc.Do())
}
