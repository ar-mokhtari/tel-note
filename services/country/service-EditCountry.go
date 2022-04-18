package country

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type (
	editCountry  struct{}
	EditRequest  protocol.Country
	editResponse struct {
		State   uint
		Message string
	}
)

var (
	EditCountry editCountry
)

func (ec *editCountry) Do(editedCountry EditRequest) error {
	return storage.EditCountry(protocol.Country(editedCountry))
}

func (ec *editCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		if err := ec.Do(req); err != nil {
			res.EncoderJson(w, editResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, editResponse{
				200,
				fmt.Sprintf("country #%v edited", req.ID),
			})
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

func (ecq *EditRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *editResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
