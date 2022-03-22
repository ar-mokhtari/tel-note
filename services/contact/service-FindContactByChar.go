package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type contactFindCharContactPool struct{}

var FindByCharPool *contactFindCharContactPool

func (allData *contactFindCharContactPool) FindContactByChar(insertChar string) (bool, []*protocol.Contact, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return true, data, uint(len(data))
	}
	return false, []*protocol.Contact{}, 0
}

func (allData *contactFindCharContactPool) FindContactByCharServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	char := r.Header.Get("char")
	if status, result, count := allData.FindContactByChar(char); status {
		json.NewEncoder(w).Encode(struct {
			Status      int
			ResultCount uint
		}{200, count})
		for _, data := range result {
			json.NewEncoder(w).Encode(data)
		}
	}
}
