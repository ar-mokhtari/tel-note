package contact

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
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
