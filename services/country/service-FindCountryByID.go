package country

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findCountryID struct{}

var FindCountryID findCountryID

func (fci *findCountryID) Do(insertChar string) []*protocol.Country {
	return storage.FindCountryByChar(insertChar)
}

func (fci *findCountryID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		inputID := r.FormValue("inputID")
		json.NewEncoder(w).Encode(fci.Do(inputID))
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
