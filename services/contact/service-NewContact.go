package contact

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	newContact         struct{}
	NewContactRequest  protocol.Contact
	newContactResponse struct {
		State   uint
		Message string
	}
)

var (
	NewContact newContact
)

func (ec *newContact) Do(inputContact NewContactRequest) (err error) {
	if err = storage.AddContact(protocol.Contact(inputContact)); err != nil {
		return err
	}
	return nil
}

func (ec *newContact) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req NewContactRequest
		res newContactResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ec.Do(req); err != nil {
			res.EncoderJson(w, newContactResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, newContactResponse{
				200,
				fmt.Sprintf("Contact added"),
			})
		}
	default:
		res.EncoderJson(w, newContactResponse{400, "method not support"})

	}
}

func (ecq *NewContactRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *newContactResponse) EncoderJson(w http.ResponseWriter, output newContactResponse) {
	json.NewEncoder(w).Encode(output)
}
