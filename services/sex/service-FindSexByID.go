package sex

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findSexID struct{}

var FindSexID findSexID

func (fpi *findSexID) Do(ID uint) protocol.Sex {
	return storage.FindSexByID(byte(ID))
}

func (fpi *findSexID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		sexID := r.FormValue("sexID")
		_, sexUID := convertor.StrToUint(sexID)
		data := fpi.Do(sexUID)
		json.NewEncoder(w).Encode(data)
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
