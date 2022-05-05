package person

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getPerson struct{}

var GetPerson getPerson

func (gp *getPerson) Do() []*protocol.Person {
	return storage.GetPersons()
}

func (gp *getPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var result = gp.Do()
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  []*protocol.Person
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			Status int
			Data   string
		}{400, "method not support"})
	}
}
