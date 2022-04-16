package person

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type getPerson struct{}

var GetPerson getPerson

func (gp *getPerson) Do() []*protocol.Person {
	return storage.GetPersons()
}

func (gp getPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var result = gp.Do()
	json.NewEncoder(w).Encode(struct {
		State uint
		Data  []*protocol.Person
	}{200, result})
}
