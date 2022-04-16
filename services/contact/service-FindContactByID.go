package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type findContactID struct{}

var FindContactID findContactID

func (fci *findContactID) FindContactByID(id uint) (protocol.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.Contact{}
}

func (fci *findContactID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.FormValue("id")
	_, uintID := convertor.StrToUint(id)
	if status, result := fci.FindContactByID(uintID); status.State {
		json.NewEncoder(w).Encode(struct {
			Status      int
			ContactData protocol.Contact
		}{200, result})
	}
}
