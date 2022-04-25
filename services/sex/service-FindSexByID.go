package sex

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
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
