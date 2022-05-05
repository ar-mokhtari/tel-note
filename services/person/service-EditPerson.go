package person

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	editPerson   struct{}
	EditRequest  protocol.Person
	editResponse struct {
		State   uint
		Message string
	}
)

var EditPerson editPerson

func (ep *editPerson) Do(ID uint, NewPerson EditRequest) error {
	if err := storage.EditPerson(ID, protocol.Person(NewPerson)); err != nil {
		return err
	}
	return nil
}

func (ep *editPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditRequest
		res editResponse
	)
	switch r.Method {
	case env.PatchMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err, id := convertor.StrToUint(r.FormValue("personID")); err != nil {
			res.EncoderJson(w, editResponse{400, err.Error()})
		} else {
			if err := ep.Do(id, req); err != nil {
				res.EncoderJson(w, editResponse{400, err.Error()})
			} else {
				res.EncoderJson(w, editResponse{
					200,
					fmt.Sprintf("person edited"),
				})
			}
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

//TODO::: handle time input when json decode unmarshal (like DOB)

func (epq *EditRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&epq)
}

func (eps *editResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
