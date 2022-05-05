package contact

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getContact struct{}

var GetContact getContact

func (gc *getContact) Do() []*protocol.Contact {
	return storage.GetContacts()
}

func (gc *getContact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		result := gc.Do()
		json.NewEncoder(w).Encode(struct {
			State       uint
			ResultCount uint
			Data        []*protocol.Contact
		}{200, uint(len(result)), result})
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
