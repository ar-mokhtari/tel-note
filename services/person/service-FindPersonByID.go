package person

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
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
