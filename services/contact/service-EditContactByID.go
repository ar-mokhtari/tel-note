package contact

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	editContact         struct{}
	EditContactRequest  protocol.Contact
	editContactResponse struct {
		State   uint
		Message string
	}
)

var (
	EditContact editContact
)

func (ec *editContact) Do(inputContact EditContactRequest) (err error) {
	if err = storage.EditContactByID(protocol.Contact(inputContact)); err != nil {
		return err
	}
	return nil
}

func (ec *editContact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditContactRequest
		res editContactResponse
	)
	switch r.Method {
	case env.PatchMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ec.Do(req); err != nil {
			res.EncoderJson(w, editContactResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, editContactResponse{
				200,
				fmt.Sprintf("Contact #%v edited", req.Id),
			})
		}
	default:
		res.EncoderJson(w, editContactResponse{400, "method not support"})

	}
}

func (ecq *EditContactRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *editContactResponse) EncoderJson(w http.ResponseWriter, output editContactResponse) {
	json.NewEncoder(w).Encode(output)
}
