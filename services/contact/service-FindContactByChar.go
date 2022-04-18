package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type findContactChar struct{}

var FindContactChar findContactChar

func (fcc *findContactChar) Do(insertChar string) (bool, []*protocol.Contact, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return true, data, uint(len(data))
	}
	return false, []*protocol.Contact{}, 0
}

func (fcc *findContactChar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		char := r.FormValue("char")
		if status, result, count := fcc.Do(char); status && char != "" {
			json.NewEncoder(w).Encode(struct {
				Status      int
				ResultCount uint
				Data        []*protocol.Contact
			}{200, count, result})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
