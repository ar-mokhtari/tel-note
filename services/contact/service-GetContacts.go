package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type contactGetPool struct{}

var GetPool contactGetPool

func (allData *contactGetPool) GetContacts() []*protocol.Contact {
	return storage.GetContacts()
}

func (allData *contactGetPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := allData.GetContacts()
	json.NewEncoder(w).Encode(struct {
		State       uint
		ResultCount uint
		Data        []*protocol.Contact
	}{200, uint(len(result)), result})
}
