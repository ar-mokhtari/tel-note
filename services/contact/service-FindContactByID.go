package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type findContactID struct{}

var FindContactID findContactID

func (fci *findContactID) Do(id uint) (protocol.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.Contact{}
}

func (fci *findContactID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		_, uintID := convertor.StrToUint(id)
		if status, result := fci.Do(uintID); status.State {
			json.NewEncoder(w).Encode(struct {
				Status      int
				ContactData protocol.Contact
			}{200, result})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
