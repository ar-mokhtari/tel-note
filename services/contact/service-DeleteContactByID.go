package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
	"tel-note/services/general"
)

type deleteContactByID struct{}

var PoolDelContactID deleteContactByID

func (allData *deleteContactByID) DeleteContactByID(ID uint) *protocol.ResponseStatus {
	if storage.DeleteContactByID(ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}

func (allData *deleteContactByID) DeleteContactByIDServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Header.Get("id")
	if state := allData.DeleteContactByID(general.StrToUint(id)); state.State {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{200, "contact deleted"})
	}
}
