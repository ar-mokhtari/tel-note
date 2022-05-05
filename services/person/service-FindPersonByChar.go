package person

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	findCharPerson struct{}
)

var FindCharPerson findCharPerson

func (fpc *findCharPerson) Do(inputChar string) []*protocol.Person {
	res := storage.FindPersonByChar(inputChar)
	return res
}

func (fpc *findCharPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		inputChar := r.FormValue("char")
		result := fpc.Do(inputChar)
		json.NewEncoder(w).Encode(struct {
			Status int
			Data   []*protocol.Person
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			Status int
			Data   string
		}{400, "method not support"})
	}
}
