package person

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
)

type (
	deletePerson   struct{}
	DeleteRequest  []uint
	deleteResponse struct {
		State   uint
		Message string
	}
)

var DeletePerson deletePerson

func (dp *deletePerson) Do(IDs []uint) []uint {
	return storage.DeletePerson(IDs)
}

func (dp *deletePerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req DeleteRequest
		res deleteResponse
	)
	switch r.Method {
	case env.DeleteMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		dp.Do(req)
		res.EncoderJson(w, deleteResponse{
			200,
			fmt.Sprintf("person(s) deleted"),
		})

	default:
		res.EncoderJson(w, deleteResponse{400, "method not support"})
	}
}

func (ecq *DeleteRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *deleteResponse) EncoderJson(w http.ResponseWriter, output deleteResponse) {
	json.NewEncoder(w).Encode(output)
}
