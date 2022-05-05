package country

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findCountryChar struct{}

var FindCountryChar findCountryChar

func (fci *findCountryChar) Do(insertChar string) []*protocol.Country {
	return storage.FindCountryByChar(insertChar)
}

func (fci *findCountryChar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		inputChar := r.FormValue("inputChar")
		json.NewEncoder(w).Encode(fci.Do(inputChar))
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
