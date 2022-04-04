package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type deleteContactByID struct{}

var DelContactIDPool deleteContactByID

func (allData *deleteContactByID) DeleteContactByID(ID uint) *protocol.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}

func (allData *deleteContactByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Header.Get("id")
	_, uintID := convertor.StrToUint(id)
	if state := allData.DeleteContactByID(uintID); state.State {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{200, "contact deleted"})
	}
}
