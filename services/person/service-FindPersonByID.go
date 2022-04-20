package person

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type findPersonID struct{}

var FindPersonID findPersonID

func (fpi *findPersonID) Do(ID uint) protocol.Person {
	return storage.FindPersonByID(ID)
}

func (fpi *findPersonID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		personID := r.FormValue("pid")
		_, personUID := convertor.StrToUint(personID)
		data := fpi.Do(uint(personUID))
		json.NewEncoder(w).Encode(data)
	default:
		json.NewEncoder(w).Encode(struct {
			Status int
			Data   string
		}{400, "method not support"})
	}
}
