package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type deleteContactByID struct{}

var DelContactID deleteContactByID

func (dci *deleteContactByID) DeleteContactByID(ID uint) *protocol.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}

func (dci *deleteContactByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Header.Get("id")
	_, uintID := convertor.StrToUint(id)
	if state := dci.DeleteContactByID(uintID); state.State {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{200, "contact deleted"})
	}
}
