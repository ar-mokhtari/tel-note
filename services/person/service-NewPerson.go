package person

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	addPerson   struct{}
	AddRequest  protocol.Person
	addResponse struct {
		State   uint
		Message string
	}
)

var AddPerson addPerson

func (ep *addPerson) Do(NewPerson AddRequest) error {
	if err := storage.NewPerson(protocol.Person(NewPerson)); err != nil {
		return err
	}
	return nil
}

func (ep *addPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req AddRequest
		res addResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ep.Do(req); err != nil {
			res.EncoderJson(w, addResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, addResponse{
				200,
				fmt.Sprintf("person added"),
			})
		}
	default:
		res.EncoderJson(w, addResponse{400, "method not support"})
	}
}

//TODO::: handle time input when json decode unmarshal (like DOB)

func (epq *AddRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&epq)
}

func (eps *addResponse) EncoderJson(w http.ResponseWriter, output addResponse) {
	json.NewEncoder(w).Encode(output)
}
