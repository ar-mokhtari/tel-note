package contact

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type deleteAllContact struct{}

var DelAllContact deleteAllContact

func (dac *deleteAllContact) Do() *protocol.ResponseStatus {
	if storage.DeleteAll() {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: false}
}

func (dac *deleteAllContact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.DeleteMethod:
		w.Header().Set("Content-Type", "application/json")
		if status := dac.Do(); status.State {
			json.NewEncoder(w).Encode(struct {
				State   uint
				Massage string
			}{200, "All contact data deleted"})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
