package country

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type callCountry struct{}

var CallCountry callCountry

func (cc *callCountry) Do() []*protocol.Country {
	return storage.CallCountry()
}
func (cc *callCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.PostMethod:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cc.Do())
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
