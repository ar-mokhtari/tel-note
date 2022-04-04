package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type deleteAllContact struct{}

var DelAllContactPool deleteAllContact

func (allData *deleteAllContact) DeleteAll() *protocol.ResponseStatus {
	if storage.DeleteAll() {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}

func (allData *deleteAllContact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if status := allData.DeleteAll(); status.State {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Massage string
		}{200, "All contact data deleted"})
	}
}
