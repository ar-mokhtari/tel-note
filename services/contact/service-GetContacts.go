package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type contactGetPool struct{}

var GetPool *contactGetPool

func (allData *contactGetPool) GetContacts() []*protocol.Contact {
	return storage.GetContacts()
}

func (allData *contactGetPool) GetContactServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := allData.GetContacts
	for _, data := range result() {
		json.NewEncoder(w).Encode(data)
	}
}
