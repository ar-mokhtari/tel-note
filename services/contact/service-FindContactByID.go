package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type contactFindIDContactPool struct{}

var FindByIDPool contactFindIDContactPool

func (allData *contactFindIDContactPool) FindContactByID(id uint) (protocol.ResponseStatus, protocol.Contact) {
	if state, data := storage.FindContactByID(id); state {
		return protocol.ResponseStatus{State: true}, data
	}
	return protocol.ResponseStatus{State: false}, protocol.Contact{}
}

func (allData *contactFindIDContactPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.FormValue("id")
	if status, result := allData.FindContactByID(convertor.StrToUint(id)); status.State {
		json.NewEncoder(w).Encode(struct {
			Status      int
			ContactData protocol.Contact
		}{200, result})
	}
}
