package country

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type (
	newCountry  struct{}
	NewRequest  protocol.Country
	newResponse struct {
		State   uint
		Message string
	}
)

var (
	NewCountry newCountry
)

func (nc *newCountry) Do(newCountry NewRequest) error {
	return storage.NewCountry(protocol.Country(newCountry))
}

func (nc *newCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req NewRequest
		res newResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := nc.Do(req); err != nil {
			res.EncoderJson(w, newResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, newResponse{
				200,
				fmt.Sprintf("country added"),
			})
		}
	default:
		res.EncoderJson(w, newResponse{400, "method not support"})
	}
}

func (ecq *NewRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *newResponse) EncoderJson(w http.ResponseWriter, output newResponse) {
	json.NewEncoder(w).Encode(output)
}
